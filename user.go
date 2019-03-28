package strichliste

import (
	"fmt"
	"github.com/jktr/go-strichliste/schema"
	"net/http"
)

// A UserClient carries the necessary context to
// interact with the /user endpoint
type UserClient struct {
	client *Client
}

// POST /user
//
// Creates a new user and returns it.
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

// GET /user/{userId}
//
// Retrieves a user by ID.
func (c *UserClient) Get(id int) (*schema.User, *Response, error) {
	return c.getByX(fmt.Sprintf("%d", id))
}

// GET /user/{userID}
//
// Retrieves a user by name.
func (c *UserClient) GetByName(name string) (*schema.User, *Response, error) {
	return c.getByX(name)
}

// GET /user
//
// Retrieves the list of users (both active and inactive).
// Pagination is possible via ListOpts, which can be nil.
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

// GET /user/search
//
// Retrieves a list of users whose names match, or include, the passed
// name. Pagination is possible via ListOpts, which can be nil.
func (c *UserClient) Search(query string, opt *ListOpts) ([]schema.User, *Response, error) {

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

// POST /user/{userId}
//
// Updates a user by ID and returns it.
func (c *UserClient) Update(id int, user *schema.UserUpdateRequest) (*schema.User, *Response, error) {
	path := fmt.Sprintf("%s/%d", schema.EndpointUser, id)

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

// Deactivates a user by ID; returns the deactivated user.
// Note that actual deletion is not possible.
func (c *UserClient) Deactivate(id int) (*schema.User, *Response, error) {
	return c.Update(id, &schema.UserUpdateRequest{
		SetActive: new(bool), // false
	})
}
