package strichliste

import (
	"fmt"
	"github.com/jktr/go-strichliste/schema"
	"net/http"
)

type ArticleClient struct {
	client *Client
}

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

func (s *ArticleClient) SearchByName(name string, opt *ListOpts) ([]schema.Article, *Response, error) {
	return s.searchByX("query", name, opt)
}

func (s *ArticleClient) SearchByBarcode(barcode string, opt *ListOpts) ([]schema.Article, *Response, error) {
	return s.searchByX("barcode", string(barcode), opt)
}

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
