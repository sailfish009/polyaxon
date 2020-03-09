// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewUpdateOrganizationMemberParams creates a new UpdateOrganizationMemberParams object
// with the default values initialized.
func NewUpdateOrganizationMemberParams() *UpdateOrganizationMemberParams {
	var ()
	return &UpdateOrganizationMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationMemberParamsWithTimeout creates a new UpdateOrganizationMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOrganizationMemberParamsWithTimeout(timeout time.Duration) *UpdateOrganizationMemberParams {
	var ()
	return &UpdateOrganizationMemberParams{

		timeout: timeout,
	}
}

// NewUpdateOrganizationMemberParamsWithContext creates a new UpdateOrganizationMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOrganizationMemberParamsWithContext(ctx context.Context) *UpdateOrganizationMemberParams {
	var ()
	return &UpdateOrganizationMemberParams{

		Context: ctx,
	}
}

// NewUpdateOrganizationMemberParamsWithHTTPClient creates a new UpdateOrganizationMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOrganizationMemberParamsWithHTTPClient(client *http.Client) *UpdateOrganizationMemberParams {
	var ()
	return &UpdateOrganizationMemberParams{
		HTTPClient: client,
	}
}

/*UpdateOrganizationMemberParams contains all the parameters to send to the API endpoint
for the update organization member operation typically these are written to a http.Request
*/
type UpdateOrganizationMemberParams struct {

	/*Body
	  Organization body

	*/
	Body *service_model.V1OrganizationMember
	/*MemberUser
	  User

	*/
	MemberUser string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update organization member params
func (o *UpdateOrganizationMemberParams) WithTimeout(timeout time.Duration) *UpdateOrganizationMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization member params
func (o *UpdateOrganizationMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization member params
func (o *UpdateOrganizationMemberParams) WithContext(ctx context.Context) *UpdateOrganizationMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization member params
func (o *UpdateOrganizationMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization member params
func (o *UpdateOrganizationMemberParams) WithHTTPClient(client *http.Client) *UpdateOrganizationMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization member params
func (o *UpdateOrganizationMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update organization member params
func (o *UpdateOrganizationMemberParams) WithBody(body *service_model.V1OrganizationMember) *UpdateOrganizationMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update organization member params
func (o *UpdateOrganizationMemberParams) SetBody(body *service_model.V1OrganizationMember) {
	o.Body = body
}

// WithMemberUser adds the memberUser to the update organization member params
func (o *UpdateOrganizationMemberParams) WithMemberUser(memberUser string) *UpdateOrganizationMemberParams {
	o.SetMemberUser(memberUser)
	return o
}

// SetMemberUser adds the memberUser to the update organization member params
func (o *UpdateOrganizationMemberParams) SetMemberUser(memberUser string) {
	o.MemberUser = memberUser
}

// WithOwner adds the owner to the update organization member params
func (o *UpdateOrganizationMemberParams) WithOwner(owner string) *UpdateOrganizationMemberParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update organization member params
func (o *UpdateOrganizationMemberParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param member.user
	if err := r.SetPathParam("member.user", o.MemberUser); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
