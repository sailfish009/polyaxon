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

package projects_v1

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
	"github.com/go-openapi/swag"
)

// NewUploadProjectArtifactParams creates a new UploadProjectArtifactParams object
// with the default values initialized.
func NewUploadProjectArtifactParams() *UploadProjectArtifactParams {
	var ()
	return &UploadProjectArtifactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadProjectArtifactParamsWithTimeout creates a new UploadProjectArtifactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadProjectArtifactParamsWithTimeout(timeout time.Duration) *UploadProjectArtifactParams {
	var ()
	return &UploadProjectArtifactParams{

		timeout: timeout,
	}
}

// NewUploadProjectArtifactParamsWithContext creates a new UploadProjectArtifactParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadProjectArtifactParamsWithContext(ctx context.Context) *UploadProjectArtifactParams {
	var ()
	return &UploadProjectArtifactParams{

		Context: ctx,
	}
}

// NewUploadProjectArtifactParamsWithHTTPClient creates a new UploadProjectArtifactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadProjectArtifactParamsWithHTTPClient(client *http.Client) *UploadProjectArtifactParams {
	var ()
	return &UploadProjectArtifactParams{
		HTTPClient: client,
	}
}

/*UploadProjectArtifactParams contains all the parameters to send to the API endpoint
for the upload project artifact operation typically these are written to a http.Request
*/
type UploadProjectArtifactParams struct {

	/*Overwrite
	  File path query params.

	*/
	Overwrite *bool
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Path
	  File path query params.

	*/
	Path *string
	/*Project
	  Project having access to the store

	*/
	Project string
	/*Uploadfile
	  The file to upload.

	*/
	Uploadfile runtime.NamedReadCloser
	/*UUID
	  Unique integer identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upload project artifact params
func (o *UploadProjectArtifactParams) WithTimeout(timeout time.Duration) *UploadProjectArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload project artifact params
func (o *UploadProjectArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload project artifact params
func (o *UploadProjectArtifactParams) WithContext(ctx context.Context) *UploadProjectArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload project artifact params
func (o *UploadProjectArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload project artifact params
func (o *UploadProjectArtifactParams) WithHTTPClient(client *http.Client) *UploadProjectArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload project artifact params
func (o *UploadProjectArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOverwrite adds the overwrite to the upload project artifact params
func (o *UploadProjectArtifactParams) WithOverwrite(overwrite *bool) *UploadProjectArtifactParams {
	o.SetOverwrite(overwrite)
	return o
}

// SetOverwrite adds the overwrite to the upload project artifact params
func (o *UploadProjectArtifactParams) SetOverwrite(overwrite *bool) {
	o.Overwrite = overwrite
}

// WithOwner adds the owner to the upload project artifact params
func (o *UploadProjectArtifactParams) WithOwner(owner string) *UploadProjectArtifactParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the upload project artifact params
func (o *UploadProjectArtifactParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPath adds the path to the upload project artifact params
func (o *UploadProjectArtifactParams) WithPath(path *string) *UploadProjectArtifactParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the upload project artifact params
func (o *UploadProjectArtifactParams) SetPath(path *string) {
	o.Path = path
}

// WithProject adds the project to the upload project artifact params
func (o *UploadProjectArtifactParams) WithProject(project string) *UploadProjectArtifactParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the upload project artifact params
func (o *UploadProjectArtifactParams) SetProject(project string) {
	o.Project = project
}

// WithUploadfile adds the uploadfile to the upload project artifact params
func (o *UploadProjectArtifactParams) WithUploadfile(uploadfile runtime.NamedReadCloser) *UploadProjectArtifactParams {
	o.SetUploadfile(uploadfile)
	return o
}

// SetUploadfile adds the uploadfile to the upload project artifact params
func (o *UploadProjectArtifactParams) SetUploadfile(uploadfile runtime.NamedReadCloser) {
	o.Uploadfile = uploadfile
}

// WithUUID adds the uuid to the upload project artifact params
func (o *UploadProjectArtifactParams) WithUUID(uuid string) *UploadProjectArtifactParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the upload project artifact params
func (o *UploadProjectArtifactParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *UploadProjectArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Overwrite != nil {

		// query param overwrite
		var qrOverwrite bool
		if o.Overwrite != nil {
			qrOverwrite = *o.Overwrite
		}
		qOverwrite := swag.FormatBool(qrOverwrite)
		if qOverwrite != "" {
			if err := r.SetQueryParam("overwrite", qOverwrite); err != nil {
				return err
			}
		}

	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Path != nil {

		// query param path
		var qrPath string
		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {
			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}

	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// form file param uploadfile
	if err := r.SetFileParam("uploadfile", o.Uploadfile); err != nil {
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
