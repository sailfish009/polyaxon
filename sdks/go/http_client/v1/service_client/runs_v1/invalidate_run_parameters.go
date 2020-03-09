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

package runs_v1

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

// NewInvalidateRunParams creates a new InvalidateRunParams object
// with the default values initialized.
func NewInvalidateRunParams() *InvalidateRunParams {
	var ()
	return &InvalidateRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInvalidateRunParamsWithTimeout creates a new InvalidateRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInvalidateRunParamsWithTimeout(timeout time.Duration) *InvalidateRunParams {
	var ()
	return &InvalidateRunParams{

		timeout: timeout,
	}
}

// NewInvalidateRunParamsWithContext creates a new InvalidateRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewInvalidateRunParamsWithContext(ctx context.Context) *InvalidateRunParams {
	var ()
	return &InvalidateRunParams{

		Context: ctx,
	}
}

// NewInvalidateRunParamsWithHTTPClient creates a new InvalidateRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInvalidateRunParamsWithHTTPClient(client *http.Client) *InvalidateRunParams {
	var ()
	return &InvalidateRunParams{
		HTTPClient: client,
	}
}

/*InvalidateRunParams contains all the parameters to send to the API endpoint
for the invalidate run operation typically these are written to a http.Request
*/
type InvalidateRunParams struct {

	/*Body*/
	Body *service_model.V1ProjectEntityResourceRequest
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project where the notification will be assigned

	*/
	Project string
	/*UUID
	  Uuid identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the invalidate run params
func (o *InvalidateRunParams) WithTimeout(timeout time.Duration) *InvalidateRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invalidate run params
func (o *InvalidateRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invalidate run params
func (o *InvalidateRunParams) WithContext(ctx context.Context) *InvalidateRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invalidate run params
func (o *InvalidateRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invalidate run params
func (o *InvalidateRunParams) WithHTTPClient(client *http.Client) *InvalidateRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invalidate run params
func (o *InvalidateRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the invalidate run params
func (o *InvalidateRunParams) WithBody(body *service_model.V1ProjectEntityResourceRequest) *InvalidateRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the invalidate run params
func (o *InvalidateRunParams) SetBody(body *service_model.V1ProjectEntityResourceRequest) {
	o.Body = body
}

// WithOwner adds the owner to the invalidate run params
func (o *InvalidateRunParams) WithOwner(owner string) *InvalidateRunParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the invalidate run params
func (o *InvalidateRunParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the invalidate run params
func (o *InvalidateRunParams) WithProject(project string) *InvalidateRunParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the invalidate run params
func (o *InvalidateRunParams) SetProject(project string) {
	o.Project = project
}

// WithUUID adds the uuid to the invalidate run params
func (o *InvalidateRunParams) WithUUID(uuid string) *InvalidateRunParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the invalidate run params
func (o *InvalidateRunParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *InvalidateRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
