// Code generated by go-swagger; DO NOT EDIT.

package todo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateTodoListHandlerFunc turns a function with the right signature into a update todo list handler
type UpdateTodoListHandlerFunc func(UpdateTodoListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateTodoListHandlerFunc) Handle(params UpdateTodoListParams) middleware.Responder {
	return fn(params)
}

// UpdateTodoListHandler interface for that can handle valid update todo list params
type UpdateTodoListHandler interface {
	Handle(UpdateTodoListParams) middleware.Responder
}

// NewUpdateTodoList creates a new http.Handler for the update todo list operation
func NewUpdateTodoList(ctx *middleware.Context, handler UpdateTodoListHandler) *UpdateTodoList {
	return &UpdateTodoList{Context: ctx, Handler: handler}
}

/* UpdateTodoList swagger:route PUT /update-todo todo updateTodoList

todo更新

*/
type UpdateTodoList struct {
	Context *middleware.Context
	Handler UpdateTodoListHandler
}

func (o *UpdateTodoList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateTodoListParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
