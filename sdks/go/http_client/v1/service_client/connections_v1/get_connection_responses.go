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

package connections_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetConnectionReader is a Reader for the GetConnection structure.
type GetConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetConnectionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetConnectionOK creates a GetConnectionOK with default headers values
func NewGetConnectionOK() *GetConnectionOK {
	return &GetConnectionOK{}
}

/*GetConnectionOK handles this case with default header values.

A successful response.
*/
type GetConnectionOK struct {
	Payload *service_model.V1ConnectionResponse
}

func (o *GetConnectionOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections/{uuid}][%d] getConnectionOK  %+v", 200, o.Payload)
}

func (o *GetConnectionOK) GetPayload() *service_model.V1ConnectionResponse {
	return o.Payload
}

func (o *GetConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ConnectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionNoContent creates a GetConnectionNoContent with default headers values
func NewGetConnectionNoContent() *GetConnectionNoContent {
	return &GetConnectionNoContent{}
}

/*GetConnectionNoContent handles this case with default header values.

No content.
*/
type GetConnectionNoContent struct {
	Payload interface{}
}

func (o *GetConnectionNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections/{uuid}][%d] getConnectionNoContent  %+v", 204, o.Payload)
}

func (o *GetConnectionNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetConnectionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionForbidden creates a GetConnectionForbidden with default headers values
func NewGetConnectionForbidden() *GetConnectionForbidden {
	return &GetConnectionForbidden{}
}

/*GetConnectionForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetConnectionForbidden struct {
	Payload interface{}
}

func (o *GetConnectionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections/{uuid}][%d] getConnectionForbidden  %+v", 403, o.Payload)
}

func (o *GetConnectionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionNotFound creates a GetConnectionNotFound with default headers values
func NewGetConnectionNotFound() *GetConnectionNotFound {
	return &GetConnectionNotFound{}
}

/*GetConnectionNotFound handles this case with default header values.

Resource does not exist.
*/
type GetConnectionNotFound struct {
	Payload interface{}
}

func (o *GetConnectionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections/{uuid}][%d] getConnectionNotFound  %+v", 404, o.Payload)
}

func (o *GetConnectionNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionDefault creates a GetConnectionDefault with default headers values
func NewGetConnectionDefault(code int) *GetConnectionDefault {
	return &GetConnectionDefault{
		_statusCode: code,
	}
}

/*GetConnectionDefault handles this case with default header values.

An unexpected error response
*/
type GetConnectionDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get connection default response
func (o *GetConnectionDefault) Code() int {
	return o._statusCode
}

func (o *GetConnectionDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/connections/{uuid}][%d] GetConnection default  %+v", o._statusCode, o.Payload)
}

func (o *GetConnectionDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
