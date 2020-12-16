package force

// TokenResponse represents the response returned from making an OAuth request.
type TokenResponse struct {
	ID          *string `json:"id"`
	IssuedAt    *string `json:"issued_at"`
	InstanceURL *string `json:"instance_url"`
	AccessToken *string `json:"access_token"`
	Signature   *string `json:"signature"`
	TokenType   *string `json:"token_type"`
}

// TokenErrorResponse is the error response returned from an OAuth request.
type TokenErrorResponse struct {
	ErrorCode   TokenErrorCode `json:"error"`
	Description string         `json:"error_description"`
}

// TokenErrorCode represents the string error code returned in a TokenErrorResponse.
type TokenErrorCode string

// TokenErrorCodes represents a list of error codes returned by a TokenErrorResponse.
var TokenErrorCodes = struct {
	InvalidClient   TokenErrorCode
	InvalidClientID TokenErrorCode
	InvalidGrant    TokenErrorCode
}{
	InvalidClient:   "invalid_client",
	InvalidClientID: "invalid_client_id",
	InvalidGrant:    "invalid_grant",
}

// ToString is a helper method to return the string of a TokenErrorCode.
func (s TokenErrorCode) ToString() string {
	return string(s)
}
