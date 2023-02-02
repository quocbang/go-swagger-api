// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"quocbang/go-swagger-api/server/swagger/models"
)

// ListAuthorizedAccountHandlerFunc turns a function with the right signature into a list authorized account handler
type ListAuthorizedAccountHandlerFunc func(ListAuthorizedAccountParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAuthorizedAccountHandlerFunc) Handle(params ListAuthorizedAccountParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ListAuthorizedAccountHandler interface for that can handle valid list authorized account params
type ListAuthorizedAccountHandler interface {
	Handle(ListAuthorizedAccountParams, *models.Principal) middleware.Responder
}

// NewListAuthorizedAccount creates a new http.Handler for the list authorized account operation
func NewListAuthorizedAccount(ctx *middleware.Context, handler ListAuthorizedAccountHandler) *ListAuthorizedAccount {
	return &ListAuthorizedAccount{Context: ctx, Handler: handler}
}

/* ListAuthorizedAccount swagger:route GET /account/authorized/department-oid/{departmentOID} account listAuthorizedAccount

查詢帳號權限清單

*/
type ListAuthorizedAccount struct {
	Context *middleware.Context
	Handler ListAuthorizedAccountHandler
}

func (o *ListAuthorizedAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListAuthorizedAccountParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ListAuthorizedAccountOKBody list authorized account o k body
//
// swagger:model ListAuthorizedAccountOKBody
type ListAuthorizedAccountOKBody struct {

	// data
	Data models.AccountData `json:"data"`
}

// Validate validates this list authorized account o k body
func (o *ListAuthorizedAccountOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListAuthorizedAccountOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("listAuthorizedAccountOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this list authorized account o k body based on the context it is used
func (o *ListAuthorizedAccountOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListAuthorizedAccountOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("listAuthorizedAccountOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListAuthorizedAccountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAuthorizedAccountOKBody) UnmarshalBinary(b []byte) error {
	var res ListAuthorizedAccountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
