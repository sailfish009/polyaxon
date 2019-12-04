// Copyright 2019 Polyaxon, Inc.
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
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V1AverageStoppingPolicy Early stopping with average stopping, this policy computes running averages across
// all runs and stops those whose best performance is worse than the median
// of the running runs.
// swagger:model v1AverageStoppingPolicy
type V1AverageStoppingPolicy struct {

	// Interval/Frequency for applying the policy.
	EvaluationInterval int32 `json:"evaluation_interval,omitempty"`

	// Kind of this stopping policy, should be equal to "average"
	Kind string `json:"kind,omitempty"`
}

// Validate validates this v1 average stopping policy
func (m *V1AverageStoppingPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AverageStoppingPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AverageStoppingPolicy) UnmarshalBinary(b []byte) error {
	var res V1AverageStoppingPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
