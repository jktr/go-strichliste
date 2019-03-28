package schema

const (
	EndpointArticle       = "/article"
	EndpointArticleSearch = "/article/search"
)

type Article struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Value       int       `json:"amount"`
	Barcode     *string   `json:"barcode"`
	IsActive    bool      `json:"active"`
	Precursor   *Article  `json:"precursor"`
	TimeCreated Timestamp `json:"created"`

	/*
	  usageCount is omitted because it has insufficiently specified semantics
	    - unclear meaning when reversing transactions
	    - unclear accounting period
	    - unclear meaning if "precursor" Articles exist
	*/
	usageCount int `json:"usageCount"`
}

type ArticleCreateRequest struct {
	Name    string `json:"name"`
	Value   int    `json:"amount"`
	Barcode string `json:"barcode,omitempty"`
}

type ArticleUpdateRequest struct {
	Name    string `json:"name"`
	Value   int    `json:"amount"`
	Barcode string `json:"barcode,omitempty"`
}

type SingleArticleResponse struct {
	Article Article `json:"article"`
}

type MultiArticleResponse struct {
	Articles []Article `json:"articles"`
}
