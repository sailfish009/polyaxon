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
)

// NewGetRunSettingsParams creates a new GetRunSettingsParams object
// with the default values initialized.
func NewGetRunSettingsParams() *GetRunSettingsParams {
	var ()
	return &GetRunSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunSettingsParamsWithTimeout creates a new GetRunSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunSettingsParamsWithTimeout(timeout time.Duration) *GetRunSettingsParams {
	var ()
	return &GetRunSettingsParams{

		timeout: timeout,
	}
}

// NewGetRunSettingsParamsWithContext creates a new GetRunSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunSettingsParamsWithContext(ctx context.Context) *GetRunSettingsParams {
	var ()
	return &GetRunSettingsParams{

		Context: ctx,
	}
}

// NewGetRunSettingsParamsWithHTTPClient creates a new GetRunSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunSettingsParamsWithHTTPClient(client *http.Client) *GetRunSettingsParams {
	var ()
	return &GetRunSettingsParams{
		HTTPClient: client,
	}
}

/*GetRunSettingsParams contains all the parameters to send to the API endpoint
for the get run settings operation typically these are written to a http.Request
*/
type GetRunSettingsParams struct {

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

// WithTimeout adds the timeout to the get run settings params
func (o *GetRunSettingsParams) WithTimeout(timeout time.Duration) *GetRunSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run settings params
func (o *GetRunSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run settings params
func (o *GetRunSettingsParams) WithContext(ctx context.Context) *GetRunSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run settings params
func (o *GetRunSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run settings params
func (o *GetRunSettingsParams) WithHTTPClient(client *http.Client) *GetRunSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run settings params
func (o *GetRunSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get run settings params
func (o *GetRunSettingsParams) WithOwner(owner string) *GetRunSettingsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get run settings params
func (o *GetRunSettingsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the get run settings params
func (o *GetRunSettingsParams) WithProject(project string) *GetRunSettingsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get run settings params
func (o *GetRunSettingsParams) SetProject(project string) {
	o.Project = project
}

// WithUUID adds the uuid to the get run settings params
func (o *GetRunSettingsParams) WithUUID(uuid string) *GetRunSettingsParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get run settings params
func (o *GetRunSettingsParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
