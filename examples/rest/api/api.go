package api

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Headers map[string]any
	Body    any
}

type Response struct {
	Status int
	Body   map[string]any
}

func (r Response) Bind(strukt any) error {
	jsonBody, err := json.Marshal(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonBody, &strukt)
}

type RestClient interface {
	Get(url string) (*Response, error)
	Post(url string, payload *Request) (*Response, error)
	Put(url string, payload *Request) (*Response, error)
	Delete(url string) (*Response, error)
}

// API :
type ApiClient struct {
	rest    RestClient
	baseUrl string
}

func NewApiClient(restClient RestClient, baseUrl string) *ApiClient {
	return &ApiClient{restClient, baseUrl}
}

func (c *ApiClient) CreateAccount(req *CreateAccountRequest) (*CreateAccountResponse, error) {
	url := fmt.Sprintf("%s/signup", c.baseUrl)
	response, err := c.rest.Post(url, &Request{nil, req})
	if err != nil {
		return nil, err
	}

	createAccountResponse := &CreateAccountResponse{}
	if err := response.Bind(createAccountResponse); err != nil {
		return nil, err
	}
	return createAccountResponse, nil
}

func (c *ApiClient) Login(req *LoginRequest) (*LoginResponse, error) {
	url := fmt.Sprintf("%s/login", c.baseUrl)
	response, err := c.rest.Post(url, &Request{nil, req})
	if err != nil {
		return nil, err
	}

	loginResponse := &LoginResponse{}
	if err := response.Bind(loginResponse); err != nil {
		return nil, err
	}
	return loginResponse, nil
}

func (c *ApiClient) CreatePost(req *CreatePostRequest) (*CreatePostResponse, error) {
	url := fmt.Sprintf("%s/post", c.baseUrl)
	response, err := c.rest.Post(url, &Request{nil, req})
	if err != nil {
		return nil, err
	}

	createPostResponse := &CreatePostResponse{}
	if err := response.Bind(createPostResponse); err != nil {
		return nil, err
	}
	return createPostResponse, nil
}

func (c *ApiClient) CreateComment(req *CreateCommentRequest) (*CreateCommentResponse, error) {
	url := fmt.Sprintf("%s/post/:id/comment", c.baseUrl)
	response, err := c.rest.Post(url, &Request{nil, req})
	if err != nil {
		return nil, err
	}

	createCommentResponse := &CreateCommentResponse{}
	if err := response.Bind(createCommentResponse); err != nil {
		return nil, err
	}
	return createCommentResponse, nil
}

func (c *ApiClient) EditComment(req *EditCommentRequest) (*EditCommentResponse, error) {
	url := fmt.Sprintf("%s/post/:id/comment", c.baseUrl)
	response, err := c.rest.Put(url, &Request{nil, req})
	if err != nil {
		return nil, err
	}

	editCommentResponse := &EditCommentResponse{}
	if err := response.Bind(editCommentResponse); err != nil {
		return nil, err
	}
	return editCommentResponse, nil
}

func (c *ApiClient) DeleteComment(req *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	url := fmt.Sprintf("%s/post/:id/comment", c.baseUrl)
	response, err := c.rest.Delete(url)
	if err != nil {
		return nil, err
	}

	deleteCommentResponse := &DeleteCommentResponse{}
	if err := response.Bind(deleteCommentResponse); err != nil {
		return nil, err
	}
	return deleteCommentResponse, nil
}
