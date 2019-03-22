package schema

const (
	EndpointArticle       = "/article"
	EndpointArticleSearch = "/article/search"
)

type Article struct {
	ID          ID        `json:"id"`
	Name        string    `json:"name"`
	Value       Currency  `json:"amount"`
	Barcode     *string   `json:"barcode"`
	IsActive    bool      `json:"active"`
	Precursor   *Article  `json:"precursor"`
	TimeCreated Timestamp `json:"created"`

	/*
	   Note that omit usageCount has insufficiently specified semantics.
	    - unclear meaning when reversing transactions
	    - unclear accounting period
	    - unclear interaction with "precursor" Articles
	*/
	UsageCount uint `json:"usageCount"`
}

type ArticleCreateRequest struct {
	Name    string   `json:"name"`
	Value   Currency `json:"amount"`
	Barcode Barcode  `json:"barcode,omitempty"`
}

type ArticleUpdateRequest struct {
	Name    string   `json:"name"`
	Value   Currency `json:"amount"`
	Barcode Barcode  `json:"barcode,omitempty"`
}

type SingleArticleResponse struct {
	Article Article `json:"article"`
}

type MultiArticleResponse struct {
	Articles []Article `json:"articles"`
}
