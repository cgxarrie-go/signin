package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	errRequestUrlHasNotBeenSet        string = "url has not been set"
	errRequestHTTPMethodHasNotBeenSet string = "http method has not been set"
)

type request struct {
	request            interface{}
	url                string
	context            context.Context
	headers            map[string]string
	response           interface{}
	responseHeaders    map[string][]string
	expectedStatusCode int
	httpClient         http.Client
	httpMethod         string

	shouldVerifySatusCode       bool
	shouldUnmarshallResponse    bool
	shouldReturnResponseHeaders bool
	client                      client
	isCustomHTTPClientSet       bool
	isContextSet                bool
}

type client struct {
	maxIddleConns        int
	maxConnsPerHost      int
	maxIddleConnsPerHost int
	timeout              time.Duration
}

// Request instantiate a new fluent http request
func Request(req interface{}) *request {
	return &request{
		shouldVerifySatusCode:       false,
		isCustomHTTPClientSet:       false,
		shouldUnmarshallResponse:    false,
		shouldReturnResponseHeaders: false,
		isContextSet:                false,
		request:                     req,
		expectedStatusCode:          -1,
		headers:                     map[string]string{},
		client: client{
			maxIddleConns:        0,
			maxConnsPerHost:      0,
			maxIddleConnsPerHost: 0,
			timeout:              time.Minute,
		},
	}
}

func (r *request) WithSigninHeders(bearer string) *request {
	return r.WithHeader("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) "+
			"Gecko/20100101 Firefox/113.0").
		WithHeader("Accept", "application/json").
		WithHeader("Accept-Language", "en-US,en;q=0.5").
		WithHeader("Content-Type", "application/json").
		WithHeader("Referer", "https://companion.signin.app/").
		WithHeader("X-App-Version", "Web companion app/1.0.3").
		WithHeader("Origin", "https://companion.signin.app").
		WithHeader("Sec-Fetch-Dest", "empty").
		WithHeader("Sec-Fetch-Mode", "cors").
		WithHeader("Sec-Fetch-Site", "cross-site").
		WithHeader("Connection", "keep-alive").
		WithHeader("TE", "trailers").
		WithHeader("Authorization", fmt.Sprintf("Bearer %s", bearer))

}

// WithHTTPMethod sets the httpmethod to be used in the request
func (r *request) WithHTTPMethod(method string) *request {
	r.httpMethod = method
	return r
}

// WithHTTPClient sets the HTTP client to be used in for the request
// if not set, a new http client instance is used.
func (r *request) WithHTTPClient(client http.Client) *request {
	r.httpClient = client
	r.isCustomHTTPClientSet = true
	return r
}

// WithClientTimeout sets the timeout for the HTTP client
// Default timeout is 1 minute
func (r *request) WithClientTimeout(value time.Duration) *request {
	r.client.timeout = value
	return r
}

// WithClientMaxIddleConnections sets the maxIddleConns for the HTTP client
// Default value is 0
func (r *request) WithClientMaxIddleConnections(value int) *request {
	r.client.maxIddleConns = value
	return r
}

// WithClientMaxConnectionsPerHost sets the maxConnsPerHost for the HTTP client
// Default value is 0
func (r *request) WithClientMaxConnectionsPerHost(value int) *request {
	r.client.maxConnsPerHost = value
	return r
}

// WithClientMaxIddleConnectionsPerHost sets the maxIddleConnsPerHost
// for the HTTP client
// Default value is 0
func (r *request) WithClientMaxIddleConnectionsPerHost(value int) *request {
	r.client.maxIddleConnsPerHost = value
	return r
}

// WithExpectedStatusCode sets the expected status code for the response
// If not set, status code is not validated
func (r *request) WithExpectedStatusCode(code int) *request {
	r.expectedStatusCode = code
	r.shouldVerifySatusCode = true
	return r
}

// WithURL sets the url to be called to perform the request
func (r *request) WithURL(url string) *request {
	r.url = url
	return r
}

// WithContext sets the context for the http request
func (r *request) WithContext(ctx context.Context) *request {
	r.context = ctx
	r.isContextSet = true
	return r
}

// WithHeader adds a header to the http request
func (r *request) WithHeader(key string, value string) *request {
	r.headers[key] = value
	return r
}

// WithResponseBody sets the pointer to store the response in
func (r *request) WithResponseBody(resp interface{}) *request {
	r.response = resp
	r.shouldUnmarshallResponse = true
	return r
}

// WithResponseHeaders sets the map to store the response headers in
func (r *request) WithResponseHeaders(respHead map[string][]string) *request {
	r.responseHeaders = respHead
	r.shouldReturnResponseHeaders = true
	return r
}

// Do executes the request
func (r *request) Do() error {

	if err := r.validate(); err != nil {
		return fmt.Errorf("validating request : %w", err)
	}

	httpReq, err := r.prepareRequest()
	if err != nil {
		return err
	}

	httpClient := r.getHTTPClient()

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("error executing request %s", err)
	}

	defer resp.Body.Close() //nolint:errcheck

	err = r.decodeResponse(resp)

	return err
}

func (r request) prepareRequest() (httpReq *http.Request, err error) {
	jsonBody, err := json.Marshal(r.request)
	if err != nil {
		return httpReq, fmt.Errorf("error marshalling request %s", err)
	}
	bodyReader := bytes.NewReader(jsonBody)

	httpReq, err = r.getHTTPRequest(r.httpMethod, bodyReader)
	if err != nil {
		return httpReq, fmt.Errorf("error creating hhtpRequest request %s", err)
	}

	for headKey, headValue := range r.headers {
		httpReq.Header.Set(headKey, headValue)
	}

	return httpReq, nil
}

func (r *request) decodeResponse(resp *http.Response) (err error) {

	if r.shouldVerifySatusCode &&
		resp.StatusCode != r.expectedStatusCode {
		return fmt.Errorf("%s expected code %d, got %d", r.url,
			r.expectedStatusCode, resp.StatusCode)
	}

	if r.shouldReturnResponseHeaders {
		for k, v := range resp.Header {
			r.responseHeaders[k] = v
		}
	}

	if !r.shouldUnmarshallResponse {
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(r.response)
	if err != nil {
		return fmt.Errorf("cannot read response body: %s", err)
	}

	return nil
}

func (r request) validate() error {

	if r.url == "" {
		return fmt.Errorf(errRequestUrlHasNotBeenSet)

	}

	if r.httpMethod == "" {
		return fmt.Errorf(errRequestHTTPMethodHasNotBeenSet)
	}
	return nil
}

func (r request) getHTTPClient() http.Client {

	if r.isCustomHTTPClientSet {
		return r.httpClient
	}

	transport := &http.Transport{
		MaxIdleConns:        r.client.maxIddleConns,
		MaxConnsPerHost:     r.client.maxConnsPerHost,
		MaxIdleConnsPerHost: r.client.maxIddleConnsPerHost,
	}
	return http.Client{
		Transport: transport,
		Timeout:   r.client.timeout,
	}

}

func (r request) getHTTPRequest(httpMethod string, bodyReader *bytes.Reader) (*http.Request, error) {

	if r.isContextSet {
		return http.NewRequestWithContext(r.context, httpMethod, r.url,
			bodyReader)
	}

	return http.NewRequest(httpMethod, r.url, bodyReader)

}
