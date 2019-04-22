package schema

const EndpointTransaction = "/transaction"

type Transaction struct {
	ID          int       `json:"id"`
	Issuer      User      `json:"user"`
	Value       int       `json:"amount"`
	Comment     string    `json:"comment"`
	TimeCreated Timestamp `json:"created"`

	IsReversed   bool `json:"deleted"`
	IsReversible bool `json:"isDeletable"`

	// only when sending money
	From *User `json:"sender"` // TODO always empty?
	To   *User `json:"recipient"`

	// only when buying articles
	Quantity *int     `json:"quantity"`
	Article  *Article `json:"articleId"`
}

type TransactionCreateRequest struct {
	// "amount" overwrites Article's value if ArticleID is present
	Amount    int    `json:"amount"`
	Comment   string `json:"comment,omitempty"`
	Recipient *int   `json:"recipientId,omitempty"`

	// only when sending articles
	Quantity  *int `json:"quantity,omitempty"`
	ArticleID *int `json:"articleId,omitempty"`
}

type SingleTransactionResponse struct {
	Transaction Transaction `json:"transaction"`
}

type MultiTransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
}
