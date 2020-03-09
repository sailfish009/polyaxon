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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EntityNotificationBody Request data to create/update entity status
//
// swagger:model v1EntityNotificationBody
type V1EntityNotificationBody struct {

	// Status condition
	Condition *V1StatusCondition `json:"condition,omitempty"`

	// Connections to notify
	Connections []string `json:"connections"`

	// Optional name of the entity
	Name string `json:"name,omitempty"`

	// Na,espace
	Namespace string `json:"namespace,omitempty"`

	// Owner of the namespace
	Owner string `json:"owner,omitempty"`

	// Project where the run will be assigned
	Project string `json:"project,omitempty"`

	// Uuid identifier of the entity
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this v1 entity notification body
func (m *V1EntityNotificationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EntityNotificationBody) validateCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	if m.Condition != nil {
		if err := m.Condition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("condition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EntityNotificationBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EntityNotificationBody) UnmarshalBinary(b []byte) error {
	var res V1EntityNotificationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
