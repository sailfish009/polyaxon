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

// V1Mapping Parallelism based on map/array either specified manually by the user or passed from another step
//
// swagger:model v1Mapping
type V1Mapping struct {

	// Number of concurrent runs
	Concurrency int32 `json:"concurrency,omitempty"`

	// A list of Early stopping objects, accpets both metric and failure early stopping mechanisms
	EarlyStopping []interface{} `json:"early_stopping"`

	// Kind of parallel, should be equal to "iterative"
	Kind *string `json:"kind,omitempty"`

	// The mapping values
	Values string `json:"values,omitempty"`
}

// Validate validates this v1 mapping
func (m *V1Mapping) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Mapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Mapping) UnmarshalBinary(b []byte) error {
	var res V1Mapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
