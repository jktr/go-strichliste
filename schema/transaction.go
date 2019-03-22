package schema

const EndpointTransaction = "/transaction"

type Transaction struct {
	ID          ID        `json:"id"`
	Issuer      User      `json:"user"`
	Value       Currency  `json:"amount"`
	Comment     string    `json:"comment"`
	TimeCreated Timestamp `json:"created"`

	IsReversed   bool `json:"deleted"`
	IsReversible bool `json:"isDeletable"`

	// only when sending money
	From *User `json:"sender"`
	To   *User `json:"recipient"`

	// only when buying articles
	Quantity *uint    `json:"quantity"`
	Article  *Article `json:"articleId"`
}

type TransactionCreateRequest struct {
	// "amount" overwrites Article's value if ArticleID is present
	Amount    Currency `json:"amount"`
	Comment   string   `json:"comment,omitempty"`
	Recipient *ID      `json:"recipient,omitempty"`

	// only when sending articles
	Quantity  *uint `json:"quantity,omitempty"`
	ArticleID *ID `json:"articleId,omitempty"`
}

type SingleTransactionResponse struct {
	Transaction Transaction `json:"transaction"`
}

type MultiTransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
}
