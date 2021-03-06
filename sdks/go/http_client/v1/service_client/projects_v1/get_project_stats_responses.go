// Copyright 2018-2021 Polyaxon, Inc.
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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetProjectStatsReader is a Reader for the GetProjectStats structure.
type GetProjectStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetProjectStatsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetProjectStatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectStatsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetProjectStatsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProjectStatsOK creates a GetProjectStatsOK with default headers values
func NewGetProjectStatsOK() *GetProjectStatsOK {
	return &GetProjectStatsOK{}
}

/* GetProjectStatsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetProjectStatsOK struct {
	Payload interface{}
}

func (o *GetProjectStatsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/stats][%d] getProjectStatsOK  %+v", 200, o.Payload)
}
func (o *GetProjectStatsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectStatsNoContent creates a GetProjectStatsNoContent with default headers values
func NewGetProjectStatsNoContent() *GetProjectStatsNoContent {
	return &GetProjectStatsNoContent{}
}

/* GetProjectStatsNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetProjectStatsNoContent struct {
	Payload interface{}
}

func (o *GetProjectStatsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/stats][%d] getProjectStatsNoContent  %+v", 204, o.Payload)
}
func (o *GetProjectStatsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectStatsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectStatsForbidden creates a GetProjectStatsForbidden with default headers values
func NewGetProjectStatsForbidden() *GetProjectStatsForbidden {
	return &GetProjectStatsForbidden{}
}

/* GetProjectStatsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetProjectStatsForbidden struct {
	Payload interface{}
}

func (o *GetProjectStatsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/stats][%d] getProjectStatsForbidden  %+v", 403, o.Payload)
}
func (o *GetProjectStatsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectStatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectStatsNotFound creates a GetProjectStatsNotFound with default headers values
func NewGetProjectStatsNotFound() *GetProjectStatsNotFound {
	return &GetProjectStatsNotFound{}
}

/* GetProjectStatsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetProjectStatsNotFound struct {
	Payload interface{}
}

func (o *GetProjectStatsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/stats][%d] getProjectStatsNotFound  %+v", 404, o.Payload)
}
func (o *GetProjectStatsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectStatsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectStatsDefault creates a GetProjectStatsDefault with default headers values
func NewGetProjectStatsDefault(code int) *GetProjectStatsDefault {
	return &GetProjectStatsDefault{
		_statusCode: code,
	}
}

/* GetProjectStatsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetProjectStatsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get project stats default response
func (o *GetProjectStatsDefault) Code() int {
	return o._statusCode
}

func (o *GetProjectStatsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{name}/stats][%d] GetProjectStats default  %+v", o._statusCode, o.Payload)
}
func (o *GetProjectStatsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetProjectStatsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
