// Code generated by goa v3.4.3, DO NOT EDIT.
//
// todo endpoints
//
// Command:
// $ goa gen todo/design

package todo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "todo" service endpoints.
type Endpoints struct {
	Hello  goa.Endpoint
	Show   goa.Endpoint
	Create goa.Endpoint
	Update goa.Endpoint
	Delete goa.Endpoint
}

// NewEndpoints wraps the methods of the "todo" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Hello:  NewHelloEndpoint(s),
		Show:   NewShowEndpoint(s),
		Create: NewCreateEndpoint(s),
		Update: NewUpdateEndpoint(s),
		Delete: NewDeleteEndpoint(s),
	}
}

// Use applies the given middleware to all the "todo" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Hello = m(e.Hello)
	e.Show = m(e.Show)
	e.Create = m(e.Create)
	e.Update = m(e.Update)
	e.Delete = m(e.Delete)
}

// NewHelloEndpoint returns an endpoint function that calls the method "hello"
// of service "todo".
func NewHelloEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*HelloPayload)
		return s.Hello(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "todo".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		res, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedTodo(res, "default")
		return vres, nil
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "todo".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreatePayload)
		return s.Create(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "todo".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePayload)
		return s.Update(ctx, p)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "todo".
func NewDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeletePayload)
		return s.Delete(ctx, p)
	}
}
