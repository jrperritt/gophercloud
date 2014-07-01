package v1

import (
	"fmt"
	"github.com/rackspace/gophercloud/openstack/identity"
)

type Client struct {
	endpoint  string
	authority identity.AuthResults
	options   identity.AuthOptions
	token     *identity.Token
}

func NewClient(e string, a identity.AuthResults, o identity.AuthOptions) *Client {
	return &Client{
		endpoint:  e,
		authority: a,
		options:   o,
	}
}

func (c *Client) GetListURL() string {
	return fmt.Sprintf("%s/v1/images/detail", c.endpoint)
}

func (c *Client) GetHeaders() (map[string]string, error) {
	t, err := c.getAuthToken()
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{
		"X-Auth-Token": t,
	}, nil
}

// getAuthToken is a function that tries to retrieve an authentication token from a client's endpoint.
func (c *Client) getAuthToken() (string, error) {
	var err error

	if c.token == nil {
		c.token, err = identity.GetToken(c.authority)
		if err != nil {
			return "", err
		}
	}

	return c.token.Id, err
}
