package strichliste

import (
	"fmt"
	"github.com/jktr/go-strichliste/schema"
	"net/http"
)

// An ArticleClient carries the necessary context to interact
// with the /article endpoint
type ArticleClient struct {
	client *Client
}

// POST /article
//   - ErrorParameterMissing
//   - ErrorParameterInvalid
//
// Creates a new article and returns it.
func (s *ArticleClient) Create(article *schema.ArticleCreateRequest) (*schema.Article, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, schema.EndpointArticle, article)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleArticleResponse
	resp, err := s.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}
	return &body.Article, resp, nil
}

// GET /article/{articleId}
//
// Retrieves an article by ID.
func (s *ArticleClient) Get(id int) (*schema.Article, *Response, error) {
	path := fmt.Sprintf("%s/%d", schema.EndpointArticle, id)

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleArticleResponse
	resp, err := s.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}
	return &body.Article, resp, nil
}

// GET /article
//
// Retrieves a list of articles (both active and inactive).
// Pagination is possible via ListOpts, which can be nil.
func (s *ArticleClient) List(opt *ListOpts) ([]schema.Article, *Response, error) {
	path := fmt.Sprintf("%s?%s", schema.EndpointArticle, opt.values().Encode())

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.MultiArticleResponse
	resp, err := s.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return body.Articles, resp, nil
}

func (s *ArticleClient) searchByX(x, query string, opt *ListOpts) ([]schema.Article, *Response, error) {

	v := opt.values()
	v.Add(x, query)

	path := fmt.Sprintf("%s?%s", schema.EndpointArticleSearch, v.Encode())

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.MultiArticleResponse
	resp, err := s.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}

	return body.Articles, resp, nil
}

// GET /article/search
//
// Retrieves a list of articles whose names match, or include, the
// passed name. Pagination is possible via ListOpts, which can be nil.
func (s *ArticleClient) SearchByName(name string, opt *ListOpts) ([]schema.Article, *Response, error) {
	return s.searchByX("query", name, opt)
}

// GET /article/search
//
// Retrieves a list of articles whose barcodes match, or include, the
// passed barcode. Pagination is possible via ListOpts, which can be nil.
func (s *ArticleClient) SearchByBarcode(barcode string, opt *ListOpts) ([]schema.Article, *Response, error) {
	return s.searchByX("barcode", string(barcode), opt)
}

// POST /article/{articleId}
//   - ErrorArticleNotFound
//   - ErrorParameterMissing
//   - ErrorParameterInvalid
//
// Updates an article by ID. Note that this operation checks for
// referential integrity and may not update the article, but instead
// create a new one, referencing and deactivating the old version.
// The returned article is always new version — either replaced or
// updated.
func (s *ArticleClient) Update(id int, article *schema.ArticleUpdateRequest) (*schema.Article, *Response, error) {
	path := fmt.Sprintf("%s/%d", schema.EndpointArticle, id)

	req, err := s.client.NewRequest(http.MethodPost, path, article)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleArticleResponse
	resp, err := s.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}
	return &body.Article, resp, nil
}

// DELETE /article/{articleId}
//   - ErrorArticleNotFound
//
// Deactivates an article by ID; returns the deactivated article.
// Note that actual deletion is not possible.
func (s *ArticleClient) Deactivate(id int) (*schema.Article, *Response, error) {
	path := fmt.Sprintf("%s/%d", schema.EndpointArticle, id)

	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.SingleArticleResponse
	resp, err := s.client.Do(req, &body)
	if err != nil {
		return nil, resp, err
	}
	return &body.Article, resp, nil
}
