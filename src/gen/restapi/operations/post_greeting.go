// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostGreetingHandlerFunc turns a function with the right signature into a post greeting handler
type PostGreetingHandlerFunc func(PostGreetingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostGreetingHandlerFunc) Handle(params PostGreetingParams) middleware.Responder {
	return fn(params)
}

// PostGreetingHandler interface for that can handle valid post greeting params
type PostGreetingHandler interface {
	Handle(PostGreetingParams) middleware.Responder
}

// NewPostGreeting creates a new http.Handler for the post greeting operation
func NewPostGreeting(ctx *middleware.Context, handler PostGreetingHandler) *PostGreeting {
	return &PostGreeting{Context: ctx, Handler: handler}
}

/*
	PostGreeting swagger:route POST /hello postGreeting

PostGreeting post greeting API
*/
type PostGreeting struct {
	Context *middleware.Context
	Handler PostGreetingHandler
}

func (o *PostGreeting) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostGreetingParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
