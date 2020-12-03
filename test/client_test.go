package test

import (
	"github.com/davidji99/force-go/force"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewClient_WithAccessToken(t *testing.T) {
	client, err := force.New(
		force.AccessToken("SOME_TOKEN"),
		force.InstanceURL("https://instance.my.salesforce.com"))
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestNewClient_WithAccessTokenEmptyString(t *testing.T) {
	client, err := force.New(
		force.AccessToken(""),
		force.InstanceURL("https://instance.my.salesforce.com"))
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestNewClient_OAuth(t *testing.T) {
	client, err := force.New(force.OAuthCred(
		os.Getenv("OAUTH_USER"),
		os.Getenv("OAUTH_PASSWORD"),
		os.Getenv("OAUTH_CLIENT_ID"),
		os.Getenv("OAUTH_CLIENT_SECRET"),
	))
	assert.Nil(t, err)
	assert.NotNil(t, client)
}
