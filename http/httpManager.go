package http

import (
	"errors"
	"net/url"
	"strconv"
	"time"

	gjson "github.com/adimax2953/go-tool/json"

	"github.com/valyala/fasthttp"
)

type Option struct {
	URL                       string
	Name                      string
	MaxConnsPerHost           int
	MaxIdemponentCallAttempts int
	ReadTimeout               time.Duration
	WriteTimeout              time.Duration
}

// Service -
type Service struct {
	url    string
	name   string
	Client *fasthttp.Client
}

type Response struct {
}

// NewClient -發起一個http client
func NewClient(option *Option) (*Service, error) {
	c := &Service{
		url:  option.URL,
		name: option.Name,
		Client: &fasthttp.Client{
			// // MaxConnsPerHost  default is 512, increase to 16384
			MaxConnsPerHost: option.MaxConnsPerHost,
			ReadTimeout:     option.ReadTimeout * time.Second,
			WriteTimeout:    option.WriteTimeout * time.Second,
			// // retry 次數
			MaxIdemponentCallAttempts: option.MaxIdemponentCallAttempts,
		}}
	return c, nil
}

// GetByte - 取得回應格式為[]byte
func (c *Service) GetByte(path, query string) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	u, err := url.Parse(c.url + path)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query
	req.SetRequestURI(u.String())
	req.Header.Add("X-ServerID", c.name)
	timestamp := time.Now().Unix()
	req.Header.Add("X-API-Timestamp",
		strconv.FormatInt(timestamp, 10))
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod(GET)

	if err := c.Client.Do(req, resp); err != nil {
		return nil, err
	}

	bodyBytes := resp.Body()
	if resp.StatusCode() != 200 {
		return nil, errors.New(strconv.Itoa(resp.StatusCode()) + "," + string(bodyBytes))
	}
	return bodyBytes, nil
}

// GetString - 取得回應格式為String
func (c *Service) GetString(path, query string) (body string, err error) {

	b, err := c.GetByte(path, query)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// DelByte - 取得回應格式為[]byte
func (c *Service) DelByte(path, query string) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	u, err := url.Parse(c.url + path)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query
	req.SetRequestURI(u.String())
	req.Header.Add("X-ServerID", c.name)
	timestamp := time.Now().Unix()
	req.Header.Add("X-API-Timestamp",
		strconv.FormatInt(timestamp, 10))
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod(DELETE)

	if err := c.Client.Do(req, resp); err != nil {
		return nil, err
	}

	bodyBytes := resp.Body()
	if resp.StatusCode() != 200 {
		return nil, errors.New(strconv.Itoa(resp.StatusCode()) + "," + string(bodyBytes))
	}
	return bodyBytes, nil
}

// DelString - 取得回應格式為String
func (c *Service) DelString(path, query string) (body string, err error) {

	b, err := c.DelByte(path, query)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// PostByte - 取得回應格式為[]byte
func (c *Service) PostByte(path string, data interface{}) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(c.url + path)
	req.Header.Add("X-ServerID", c.name)
	timestamp := time.Now().Unix()
	req.Header.Add("X-API-Timestamp",
		strconv.FormatInt(timestamp, 10))
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod(POST)

	body, err := gjson.JsonMarshal(data)
	if err != nil {
		return nil, err
	}

	req.SetBody(body)
	if err := c.Client.Do(req, resp); err != nil {
		return nil, err
	}

	bodyBytes := resp.Body()
	if resp.StatusCode() != 200 && resp.StatusCode() != 204 {
		return nil, errors.New(strconv.Itoa(resp.StatusCode()) + "," + string(bodyBytes))
	}

	return bodyBytes, nil
}

// PostString - 取得回應格式為String
func (c *Service) PostString(path string, data interface{}) (body string, err error) {

	b, err := c.PostByte(path, data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// PutByte - 取得回應格式為[]byte
func (c *Service) PutByte(path string, data interface{}) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(c.url + path)
	req.Header.Add("X-ServerID", c.name)
	timestamp := time.Now().Unix()
	req.Header.Add("X-API-Timestamp",
		strconv.FormatInt(timestamp, 10))
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod(PUT)

	body, err := gjson.JsonMarshal(data)
	if err != nil {
		return nil, err
	}

	req.SetBody(body)
	if err := c.Client.Do(req, resp); err != nil {
		return nil, err
	}

	bodyBytes := resp.Body()
	if resp.StatusCode() != 200 && resp.StatusCode() != 204 {
		return nil, errors.New(strconv.Itoa(resp.StatusCode()) + "," + string(bodyBytes))
	}

	return bodyBytes, nil
}

// PutString - 取得回應格式為String
func (c *Service) PutString(path string, data interface{}) (body string, err error) {

	b, err := c.PutByte(path, data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// PatchByte - 取得回應格式為[]byte
func (c *Service) PatchByte(path string, data interface{}) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(c.url + path)
	req.Header.Add("X-ServerID", c.name)
	timestamp := time.Now().Unix()
	req.Header.Add("X-API-Timestamp",
		strconv.FormatInt(timestamp, 10))
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod(PATCH)

	body, err := gjson.JsonMarshal(data)
	if err != nil {
		return nil, err
	}

	req.SetBody(body)
	if err := c.Client.Do(req, resp); err != nil {
		return nil, err
	}

	bodyBytes := resp.Body()
	if resp.StatusCode() != 200 && resp.StatusCode() != 204 {
		return nil, errors.New(strconv.Itoa(resp.StatusCode()) + "," + string(bodyBytes))
	}

	return bodyBytes, nil
}

// PatchString - 取得回應格式為String
func (c *Service) PatchString(path string, data interface{}) (body string, err error) {

	b, err := c.PatchByte(path, data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
