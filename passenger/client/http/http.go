package http

import (
	"bytes"
	"context"
	"encoding/json"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	http1 "net/http"
	"net/url"
	endpoint1 "passenger/pkg/endpoint"
	http2 "passenger/pkg/http"
	service "passenger/pkg/service"
	"strings"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]http.ClientOption) (service.PassengerService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var getPassengerInfoEndpoint endpoint.Endpoint
	{
		getPassengerInfoEndpoint = http.NewClient("POST", copyURL(u, "/get-passenger-info"), encodeHTTPGenericRequest, decodeGetPassengerInfoResponse, options["GetPassengerInfo"]...).Endpoint()
	}

	var publishOrderEndpoint endpoint.Endpoint
	{
		publishOrderEndpoint = http.NewClient("POST", copyURL(u, "/publish-order"), encodeHTTPGenericRequest, decodePublishOrderResponse, options["PublishOrder"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		GetPassengerInfoEndpoint: getPassengerInfoEndpoint,
		PublishOrderEndpoint:     publishOrderEndpoint,
	}, nil
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http1.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeGetPassengerInfoResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetPassengerInfoResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetPassengerInfoResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodePublishOrderResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodePublishOrderResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.PublishOrderResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
