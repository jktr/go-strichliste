package strichliste

import (
	"fmt"
	"github.com/jktr/go-strichliste/schema"
	"net/http"
)

type (
	// An ArticleClient carries the necessary context to
	// interact with the /transaction endpoint.
	TransactionClient struct {
		client *Client
	}

	// TransactionContext extends the TransactionClient
	// with additional user-specific context.
	TransactionContext struct {
		TransactionClient
		issuer  int    // user context
		comment string // stored here to allow easy reuse for many tx's
	}
)

// Create a user-specific context by user ID, which allows issuing
// transactions as that user.
func (c *TransactionClient) Context(user int) *TransactionContext {
	ctx := &TransactionContext{
		issuer: user,
	}
	ctx.client = c.client
	return ctx
}

// Set the comment to use for all new transactions that use the returned context.
func (c *TransactionContext) WithComment(comment string) *TransactionContext {
	ctx := *c
	ctx.comment = comment
	return &ctx
}

// Utility wrapper for Create; see Create for possible errors.
// Deposit or withdraw funds for the current User; returns the created transaction.
func (c *TransactionContext) Delta(amount int) (*schema.Transaction, *Response, error) {
	tcr := &schema.TransactionCreateRequest{
		Amount: amount,
	}
	return c.Create(tcr)
}

// Utility wrapper for Create; see Create for possible errors.
// Purchase a number of articles by ID with the current user; returns the created transaction.
func (c *TransactionContext) Purchase(article int, count int) (*schema.Transaction, *Response, error) {

	// XXX custom article price is not optional

	a, resp, err := c.client.Article.Get(article)
	if err != nil {
		return nil, resp, err
	}

	tcr := &schema.TransactionCreateRequest{
		Amount:    (-a.Value * count),
		ArticleID: &a.ID,
		Quantity:  &count,
	}
	return c.Create(tcr)
}

// Utility wrapper for Create; see Create for possible errors.
// Transfer an amount of funds from the current user to another by ID; returns the created transaction.
func (c *TransactionContext) TransferFunds(recipient int, amount int) (*schema.Transaction, *Response, error) {
	tcr := &schema.TransactionCreateRequest{
		Amount:    amount,
		Recipient: &recipient,
	}
	return c.Create(tcr)
}

// POST /user/{userId}/transaction
//   - ErrorUserNotFound
//   - ErrorParameterMissing
//   - ErrorParameterInvalid
//   - ErrorAccountBalanceBoundary
//   - ErrorTransactionBoundary
//   - TODO ErrorArticleNotFound?
//
// Creates a raw transaction and returns it.
// Consider using these wrappers for specific use-cases:
//   - Delta
//   - Purchase
//   - TransferFunds
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

// GET /user/{userId}/transaction/{txId}
//   - ErrorUserNotFound
//   - ErrorTransactionNotFound
//
// Retrieves a transaction by ID.
func (c *TransactionContext) Get(id int) (*schema.Transaction, *Response, error) {
	path := fmt.Sprintf("%s/%d%s/%d",
		schema.EndpointUser, c.issuer, schema.EndpointTransaction, id)

	req, err := c.client.NewRequest(http.MethodGet, path, nil)
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

// GET /transaction
//
// Retrieves a list of recent transactions.
// Pagination is possible via ListOpts, which can be nil.
func (c *TransactionClient) List(opt *ListOpts) ([]schema.Transaction, *Response, error) {
	path := fmt.Sprintf("%s?%s", schema.EndpointTransaction, opt.values().Encode())

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

// GET /user/{userId}/transaction
//   - ErrorUserNotFound
//
// Retrieves a list of transactions issued by this user.
// Pagination is possible via ListOpts, which can be nil.
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

// DELETE /user/{userId}/transaction
//   - ErrorUserNotFound
//   - ErrorTransactionNotFound
//   - ErrorTransactionNotDeletable
//
// Revert a transaction by ID; returns the reversed transaction.
// Not all transactions are reversible; check Transaction.IsReversible.
// Note that actual deletion is not possible.
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
