// Code generated by goa v3.4.3, DO NOT EDIT.
//
// todo client
//
// Command:
// $ goa gen todo/design

package todo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "todo" service client.
type Client struct {
	HelloEndpoint  goa.Endpoint
	ShowEndpoint   goa.Endpoint
	CreateEndpoint goa.Endpoint
	UpdateEndpoint goa.Endpoint
	DeleteEndpoint goa.Endpoint
}

// NewClient initializes a "todo" service client given the endpoints.
func NewClient(hello, show, create, update, delete_ goa.Endpoint) *Client {
	return &Client{
		HelloEndpoint:  hello,
		ShowEndpoint:   show,
		CreateEndpoint: create,
		UpdateEndpoint: update,
		DeleteEndpoint: delete_,
	}
}

// Hello calls the "hello" endpoint of the "todo" service.
func (c *Client) Hello(ctx context.Context, p *HelloPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.HelloEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Show calls the "show" endpoint of the "todo" service.
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *Todo, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Todo), nil
}

// Create calls the "create" endpoint of the "todo" service.
func (c *Client) Create(ctx context.Context, p *CreatePayload) (res string, err error) {
	var ires interface{}
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Update calls the "update" endpoint of the "todo" service.
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (res string, err error) {
	var ires interface{}
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Delete calls the "delete" endpoint of the "todo" service.
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (res string, err error) {
	var ires interface{}
	ires, err = c.DeleteEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
