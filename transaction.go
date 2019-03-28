package strichliste

import (
	"fmt"
	"github.com/jktr/go-strichliste/schema"
	"net/http"
)

type (
	TransactionClient struct {
		client *Client
	}
	TransactionContext struct {
		TransactionClient
		issuer  int
		comment string
	}
)

func (c *TransactionClient) Context(user int) *TransactionContext {
	ctx := &TransactionContext{
		issuer: user,
	}
	ctx.client = c.client
	return ctx
}

func (c *TransactionContext) WithComment(comment string) *TransactionContext {
	ctx := *c
	ctx.comment = comment
	return &ctx
}

func (c *TransactionContext) Delta(amount int) (*schema.Transaction, *Response, error) {
	tcr := &schema.TransactionCreateRequest{
		Amount: amount,
	}
	return c.Create(tcr)
}

func (c *TransactionContext) Deposit(amount int) (*schema.Transaction, *Response, error) {
	return c.Delta(amount)
}

func (c *TransactionContext) Withdraw(amount int) (*schema.Transaction, *Response, error) {
	return c.Delta(-amount)
}

func (c *TransactionContext) Buy(article int, count int) (*schema.Transaction, *Response, error) {
	tcr := &schema.TransactionCreateRequest{
		ArticleID: &article,
		Quantity:  &count,
	}
	return c.Create(tcr)
}

func (c *TransactionContext) TransferFunds(recipient int, amount int) (*schema.Transaction, *Response, error) {
	tcr := &schema.TransactionCreateRequest{
		Amount:    amount,
		Recipient: &recipient,
	}
	return c.Create(tcr)
}

func (c *TransactionContext) Create(trc *schema.TransactionCreateRequest) (*schema.Transaction, *Response, error) {
	path := fmt.Sprintf("%s/%d%s",
		schema.EndpointUser, c.issuer, schema.EndpointTransaction)

	if trc.Comment == "" {
		trc.Comment = c.comment
	}

	req, err := c.client.NewRequest(http.MethodPost, path, trc)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleTransactionResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}
	return &body.Transaction, resp, nil
}

func (c *TransactionContext) Get(id int) (*schema.Transaction, *Response, error) {
	path := fmt.Sprintf("%s/%d%s/%d",
		schema.EndpointUser, c.issuer, schema.EndpointTransaction, id)

	req, err := c.client.NewRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleTransactionResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}
	return &body.Transaction, resp, nil
}

func (c *TransactionContext) List(opt *ListOpts) ([]schema.Transaction, *Response, error) {
	path := fmt.Sprintf("%s/%d%s?%s", schema.EndpointUser, c.issuer,
		schema.EndpointTransaction, opt.values().Encode())

	req, err := c.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.MultiTransactionResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return body.Transactions, resp, nil
}

func (c *TransactionContext) Revert(id int) (*schema.Transaction, *Response, error) {
	path := fmt.Sprintf("%s/%d%s/%d", schema.EndpointUser,
		c.issuer, schema.EndpointTransaction, id)

	req, err := c.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleTransactionResponse
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}
	return &body.Transaction, resp, nil
}
