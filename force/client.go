package force

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"sync"
	"time"
)

const (
	// DefaultLoginURL is the default Salesforce login URL
	DefaultLoginURL = "https://login.salesforce.com"

	// DefaultAPIVersion is the default Force API version.
	DefaultAPIVersion = "v50.0"

	// DefaultUserAgent is the default user agent
	DefaultUserAgent = "force-go"

	// MediaTypeJSON
	MediaTypeJSON = "application/json"

	// FormURLEncodedHeader
	FormURLEncodedHeader = "application/x-www-form-urlencoded"
)

// A Client manages communication with the Salesforce Force API.
type Client struct {
	// clientMu protects the client during calls that modify the CheckRedirect func.
	clientMu sync.Mutex

	// HTTP client used to communicate with the API.
	http *simpleresty.Client

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// Additional HTTP headers
	customHTTPHeaders map[string]string

	// userAgent
	userAgent string

	// instanceURL
	instanceURL string

	// loginURL
	loginURL string

	// apiVersion
	apiVersion string

	// oauthCred
	oauthCred *oauthCred

	// accessToken
	accessToken string
}

// service represents the http
type service struct {
	client *Client
}

// New initializes a new Client.
func New(opts ...Option) (*Client, error) {
	c := &Client{
		http:              simpleresty.New(),
		customHTTPHeaders: map[string]string{},
		userAgent:         DefaultUserAgent,
		loginURL:          DefaultLoginURL,
		apiVersion:        DefaultAPIVersion,
		instanceURL:       "",
		accessToken:       "",
		oauthCred:         nil,
	}

	// Define any user custom Client settings
	if optErr := c.parseOptions(opts...); optErr != nil {
		return nil, optErr
	}

	// Authenticate
	authErr := c.authenticate()
	if authErr != nil {
		return nil, authErr
	}

	// Setup http
	c.setupClient(c.accessToken)

	return c, nil
}

func (c *Client) authenticate() error {
	// If access token is set, use the token. Otherwise, execute OAuth authorization request.
	if c.accessToken != "" {
		if c.instanceURL == "" {
			return fmt.Errorf("instance url must be set if authenticating via an already generated access token")
		}
		return nil
	}

	// Validate to make sure oauthCred is defined
	if c.oauthCred == nil {
		return fmt.Errorf("no oauth credentials defined")
	}

	// Execute oauth
	r, oauthErr := c.oauth()
	if oauthErr != nil {
		return oauthErr
	}

	c.instanceURL = r.GetInstanceURL()
	c.accessToken = r.GetAccessToken()

	return nil
}

func (c *Client) oauth() (*TokenResponse, error) {
	var tokenResponse *TokenResponse
	oClient := simpleresty.NewWithBaseURL(c.loginURL)
	url := oClient.RequestURL("/services/oauth2/token")

	_, err := oClient.R().
		SetFormData(map[string]string{
			"grant_type":    "password",
			"client_id":     c.oauthCred.ClientID,
			"client_secret": c.oauthCred.ClientSecret,
			"username":      c.oauthCred.Username,
			"password":      c.oauthCred.Password,
		}).
		SetHeaders(map[string]string{"Accept": MediaTypeJSON, "Content-Type": FormURLEncodedHeader}).
		SetResult(&tokenResponse).
		Post(url)

	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}

func (c *Client) setupClient(accessToken string) {
	c.http.SetBaseURL(c.instanceURL)

	c.http.SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("Content-type", MediaTypeJSON).
		SetHeader("Accept", MediaTypeJSON).
		SetHeader("User-Agent", c.userAgent).
		SetHeaders(c.customHTTPHeaders).
		SetTimeout(5 * time.Minute).
		SetAllowGetMethodPayload(true)
}

// Describe gets the metadata regarding a SObject.
func (c *Client) Describe(apiName string) (*SObjectMetadata, *simpleresty.Response, error) {
	var result *SObjectMetadata
	urlStr := c.http.RequestURL(fmt.Sprintf("/services/data/%s/sobjects/%s/describe", c.apiVersion, apiName))

	response, getErr := c.http.Get(urlStr, &result, nil)
	if getErr != nil {
		return nil, nil, getErr
	}

	return result, response, getErr
}

// Create a new SObject.
//
// This request does not return the newly created object regardless of status.
// Rather it returns a JSON result of SObjectCreateResult.
func (c *Client) Create(objectName string, opts interface{}) (*SObjectCreateResult, *simpleresty.Response, error) {
	var result *SObjectCreateResult
	urlStr := c.http.RequestURL(fmt.Sprintf("/services/data/%s/sobjects/%s", c.apiVersion, objectName))

	response, err := c.http.Post(urlStr, &result, opts)
	return result, response, err
}

// Update an existing SObject.
//
// The request does not return any body if the PATCH is successful.
func (c *Client) Update(objectName, objectId string, opts interface{}) (*simpleresty.Response, error) {
	urlStr := c.http.RequestURL(fmt.Sprintf("/services/data/%s/sobjects/%s/%s", c.apiVersion, objectName, objectId))

	response, err := c.http.Patch(urlStr, nil, opts)
	return response, err
}

// Destroy deletes an existing SObject.
func (c *Client) Destroy(objectName, objectId string) (*simpleresty.Response, error) {
	urlStr := c.http.RequestURL(fmt.Sprintf("/services/data/%s/sobjects/%s/%s", c.apiVersion, objectName, objectId))

	response, err := c.http.Delete(urlStr, nil, nil)
	return response, err
}

// QueryRequest represents a SOQL query.
type QueryRequest struct {
	SOQL string `url:"q,omitempty"`
}

// QueryResult represents the response when executing a SOQL query.
type QueryResult struct {
	Done           bool      `json:"done,omitempty"`
	TotalSize      int       `json:"totalSize,omitempty"`
	Records        []SObject `json:"records,omitempty"`
	NextRecordsURL string    `json:"nextRecordsUrl,omitempty"`
}

// Query executes a request to find SObjects.
func (c *Client) Query(q *QueryRequest) (*QueryResult, *simpleresty.Response, error) {
	var result QueryResult
	urlStr, urlStrErr := c.http.RequestURLWithQueryParams(fmt.Sprintf("/services/data/%s/query", c.apiVersion), q)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, getErr := c.http.Get(urlStr, &result, nil)
	if getErr != nil {
		return nil, nil, getErr
	}

	return &result, response, nil
}

// GetBaseSObjectQuery describes the object and then constructs the base select query.
func (c *Client) GetBaseSObjectQuery(objectName string) (string, error) {
	// Describe the SObject
	sobject, _, describeErr := c.Describe(objectName)
	if describeErr != nil {
		return "", describeErr
	}

	// The space after the second '%s' is required. Do not remove!
	queryString := fmt.Sprintf("select %s from %s ", sobject.GetFieldNamesString(), objectName)
	return queryString, nil
}
