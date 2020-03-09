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

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1IO Inputs/Outputs specification
//
// swagger:model v1IO
type V1IO struct {

	// A flag to tell if param validation for this input/output should be delayed
	DelayValidation bool `json:"delay_validation,omitempty"`

	// Description for the input/output
	Description string `json:"description,omitempty"`

	// The type of the input/output
	Iotype string `json:"iotype,omitempty"`

	// A flag to tell if this input/output is flag, only valid for bool type
	IsFlag bool `json:"is_flag,omitempty"`

	// A flag to tell if this input/output is list
	IsList bool `json:"is_list,omitempty"`

	// A flag to tell if this input/output is optional
	IsOptional bool `json:"is_optional,omitempty"`

	// Name for the input/output
	Name string `json:"name,omitempty"`

	// An optional field to provide possible values for validation
	Options []interface{} `json:"options"`

	// The value of the input/output should be compatible with the type
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this v1 i o
func (m *V1IO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1IO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IO) UnmarshalBinary(b []byte) error {
	var res V1IO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
