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

package agents_v1

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

// NewPatchAgentParams creates a new PatchAgentParams object
// with the default values initialized.
func NewPatchAgentParams() *PatchAgentParams {
	var ()
	return &PatchAgentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAgentParamsWithTimeout creates a new PatchAgentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAgentParamsWithTimeout(timeout time.Duration) *PatchAgentParams {
	var ()
	return &PatchAgentParams{

		timeout: timeout,
	}
}

// NewPatchAgentParamsWithContext creates a new PatchAgentParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAgentParamsWithContext(ctx context.Context) *PatchAgentParams {
	var ()
	return &PatchAgentParams{

		Context: ctx,
	}
}

// NewPatchAgentParamsWithHTTPClient creates a new PatchAgentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAgentParamsWithHTTPClient(client *http.Client) *PatchAgentParams {
	var ()
	return &PatchAgentParams{
		HTTPClient: client,
	}
}

/*PatchAgentParams contains all the parameters to send to the API endpoint
for the patch agent operation typically these are written to a http.Request
*/
type PatchAgentParams struct {

	/*AgentUUID
	  UUID

	*/
	AgentUUID string
	/*Body
	  Agent body

	*/
	Body *service_model.V1Agent
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch agent params
func (o *PatchAgentParams) WithTimeout(timeout time.Duration) *PatchAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch agent params
func (o *PatchAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch agent params
func (o *PatchAgentParams) WithContext(ctx context.Context) *PatchAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch agent params
func (o *PatchAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch agent params
func (o *PatchAgentParams) WithHTTPClient(client *http.Client) *PatchAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch agent params
func (o *PatchAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentUUID adds the agentUUID to the patch agent params
func (o *PatchAgentParams) WithAgentUUID(agentUUID string) *PatchAgentParams {
	o.SetAgentUUID(agentUUID)
	return o
}

// SetAgentUUID adds the agentUuid to the patch agent params
func (o *PatchAgentParams) SetAgentUUID(agentUUID string) {
	o.AgentUUID = agentUUID
}

// WithBody adds the body to the patch agent params
func (o *PatchAgentParams) WithBody(body *service_model.V1Agent) *PatchAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch agent params
func (o *PatchAgentParams) SetBody(body *service_model.V1Agent) {
	o.Body = body
}

// WithOwner adds the owner to the patch agent params
func (o *PatchAgentParams) WithOwner(owner string) *PatchAgentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch agent params
func (o *PatchAgentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agent.uuid
	if err := r.SetPathParam("agent.uuid", o.AgentUUID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
