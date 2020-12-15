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
	ErrorCode   string `json:"error"`
	Description string `json:"error_description"`
}
