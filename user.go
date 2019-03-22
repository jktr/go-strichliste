package strichliste

import (
	"fmt"
	"github.com/jktr/go-strichliste/schema"
	"net/http"
)

type UserClient struct {
	client *Client
}

func (c *UserClient) Create(user *schema.UserCreateRequest) (*schema.User, *Response, error) {
	req, err := c.client.NewRequest(http.MethodPost, schema.EndpointUser, user)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleUserResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return &body.User, resp, err
}

func (c *UserClient) getByX(x string) (*schema.User, *Response, error) {
	path := fmt.Sprintf("%s/%s", schema.EndpointUser, x)

	req, err := c.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleUserResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}
	return &body.User, resp, nil
}

func (c *UserClient) Get(id schema.ID) (*schema.User, *Response, error) {
	return c.getByX(fmt.Sprintf("%d", id))
}

func (c *UserClient) GetByName(name string) (*schema.User, *Response, error) {
	return c.getByX(name)
}

func (c *UserClient) List(opt *ListOpts) ([]schema.User, *Response, error) {
	path := fmt.Sprintf("%s?%s", schema.EndpointUser, opt.values().Encode())

	req, err := c.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.MultiUserResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return body.Users, resp, nil
}

func (c *UserClient) Lookup(query string, opt *ListOpts) ([]schema.User, *Response, error) {

	v := opt.values()
	v.Add("query", query)

	path := fmt.Sprintf("%s?%s", schema.EndpointUserSearch, v.Encode())

	req, err := c.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.MultiUserResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return body.Users, resp, nil
}

func (c *UserClient) updateByX(x string, user *schema.UserUpdateRequest) (*schema.User, *Response, error) {
	path := fmt.Sprintf("%s/%s", schema.EndpointUser, x)

	req, err := c.client.NewRequest(http.MethodPost, path, user)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleUserResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}
	return &body.User, resp, nil
}

func (c *UserClient) Update(id schema.ID, user *schema.UserUpdateRequest) (*schema.User, *Response, error) {
	return c.updateByX(fmt.Sprintf("%d", id), user)
}

func (c *UserClient) UpdateByName(name string, user *schema.UserUpdateRequest) (*schema.User, *Response, error) {
	return c.updateByX(name, user)
}

func (c *UserClient) Deactivate(id schema.ID) (*schema.User, *Response, error) {
	return c.Update(id, &schema.UserUpdateRequest{
		SetActive: new(bool), // false
	})
}
