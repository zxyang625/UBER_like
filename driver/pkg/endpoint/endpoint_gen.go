// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	service "driver/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetDriverInfoEndpoint endpoint.Endpoint
	TakeOrderEndpoint     endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.DriverService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		GetDriverInfoEndpoint: MakeGetDriverInfoEndpoint(s),
		TakeOrderEndpoint:     MakeTakeOrderEndpoint(s),
	}
	for _, m := range mdw["GetDriverInfo"] {
		eps.GetDriverInfoEndpoint = m(eps.GetDriverInfoEndpoint)
	}
	for _, m := range mdw["TakeOrder"] {
		eps.TakeOrderEndpoint = m(eps.TakeOrderEndpoint)
	}
	return eps
}