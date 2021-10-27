package force

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"regexp"
	"strings"
)

// Option is a functional option for configuring the API client.
type Option func(*Client) error

// HTTP allows for a custom HTTP.
func HTTP(http *simpleresty.Client) Option {
	return func(c *Client) error {
		c.http = http
		return nil
	}
}

// UserAgent allows overriding of the default User Agent.
func UserAgent(userAgent string) Option {
	return func(c *Client) error {
		c.userAgent = userAgent
		return nil
	}
}

// LoginURL sets the login url. URL cannot contain a trailing slash
func LoginURL(url string) Option {
	return func(c *Client) error {
		// Validate that there is no trailing slashes before setting the custom url
		if url[len(url)-1:] == "/" {
			return fmt.Errorf("custom login URL cannot contain a trailing slash")
		}

		c.loginURL = url
		return nil
	}
}

// InstanceURL sets the instance URL.
func InstanceURL(url string) Option {
	return func(c *Client) error {
		c.instanceURL = url
		return nil
	}
}

// APIVersion sets the API version. Must be in this format: `v46.0`.
func APIVersion(version string) Option {
	return func(c *Client) error {
		// Validate that version is 'vXX.X'
		regex := regexp.MustCompile(`v\d{2}.0`)
		isValid := regex.MatchString(version)

		if !isValid {
			return fmt.Errorf("invalid API version - format should be 'vXX.X'")
		}

		c.apiVersion = version
		return nil
	}
}

// CustomHTTPHeaders sets additional HTTPHeaders.
func CustomHTTPHeaders(headers map[string]string) Option {
	return func(c *Client) error {
		c.customHTTPHeaders = headers
		return nil
	}
}

// AccessToken sets the short-lived access token.
func AccessToken(t string) Option {
	return func(c *Client) error {
		if t == "" {
			return fmt.Errorf("access token cannot be empty string")
		}
		c.accessToken = t
		return nil
	}
}

// OAuthCred sets the credentials needed for OAuth.
func OAuthCred(username, password, clientID, clientSecret string) Option {
	return func(c *Client) error {
		// Make sure supplied arguments are not empty string aside from the security token
		args := []string{username, password, clientID, clientSecret}
		for _, a := range args {
			if strings.TrimSpace(a) == "" {
				return fmt.Errorf("cannot be empty string")
			}
		}

		c.oauthCred = &OAuthCredentials{
			ClientID:      clientID,
			ClientSecret:  clientSecret,
			Username:      username,
			Password:      password,
		}
		return nil
	}
}

// parseOptions parses the supplied options functions and returns a configured *Client instance.
func (c *Client) parseOptions(opts ...Option) error {
	// Range over each options function and apply it to our API type to
	// configure it. Options functions are applied in order, with any
	// conflicting options overriding earlier calls.
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}
