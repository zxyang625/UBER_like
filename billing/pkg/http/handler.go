package http

import (
	endpoint "billing/pkg/endpoint"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	http1 "github.com/go-kit/kit/transport/http"
)

// makeGenBillHandler creates the handler logic
func makeGenBillHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/gen-bill", http1.NewServer(endpoints.GenBillEndpoint, decodeGenBillRequest, encodeGenBillResponse, options...))
}

// decodeGenBillRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGenBillRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GenBillRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGenBillResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGenBillResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetBillListHandler creates the handler logic
func makeGetBillListHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-bill-list", http1.NewServer(endpoints.GetBillListEndpoint, decodeGetBillListRequest, encodeGetBillListResponse, options...))
}

// decodeGetBillListRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetBillListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetBillListRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetBillListResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetBillListResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeGetBillHandler creates the handler logic
func makeGetBillHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-bill", http1.NewServer(endpoints.GetBillEndpoint, decodeGetBillRequest, encodeGetBillResponse, options...))
}

// decodeGetBillRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetBillRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetBillRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetBillResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetBillResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
