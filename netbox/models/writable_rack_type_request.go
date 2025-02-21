// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableRackTypeRequest Adds support for custom fields and tags.
//
// swagger:model WritableRackTypeRequest
type WritableRackTypeRequest struct {

	// comments
	Comments string `json:"comments,omitempty"`

	// custom fields
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`

	// Descending units
	//
	// Units are numbered top-to-bottom
	DescUnits bool `json:"desc_units,omitempty"`

	// description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// * `2-post-frame` - 2-post frame
	// * `4-post-frame` - 4-post frame
	// * `4-post-cabinet` - 4-post cabinet
	// * `wall-frame` - Wall-mounted frame
	// * `wall-frame-vertical` - Wall-mounted frame (vertical)
	// * `wall-cabinet` - Wall-mounted cabinet
	// * `wall-cabinet-vertical` - Wall-mounted cabinet (vertical)
	// Required: true
	// Enum: ["2-post-frame","4-post-frame","4-post-cabinet","wall-frame","wall-frame-vertical","wall-cabinet","wall-cabinet-vertical"]
	FormFactor *string `json:"form_factor"`

	// Manufacturer
	// Required: true
	Manufacturer *int64 `json:"manufacturer"`

	// Maximum load capacity for the rack
	// Maximum: 2.147483647e+09
	// Minimum: 0
	MaxWeight *int64 `json:"max_weight,omitempty"`

	// model
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Model *string `json:"model"`

	// Maximum depth of a mounted device, in millimeters.
	// Maximum: 32767
	// Minimum: 0
	MountingDepth *int64 `json:"mounting_depth,omitempty"`

	// Outer dimension of rack (depth)
	// Maximum: 32767
	// Minimum: 0
	OuterDepth *int64 `json:"outer_depth,omitempty"`

	// * `mm` - Millimeters
	// * `in` - Inches
	// Enum: ["mm","in",""]
	OuterUnit string `json:"outer_unit,omitempty"`

	// Outer dimension of rack (width)
	// Maximum: 32767
	// Minimum: 0
	OuterWidth *int64 `json:"outer_width,omitempty"`

	// slug
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// Starting unit for rack
	// Maximum: 32767
	// Minimum: 1
	StartingUnit int64 `json:"starting_unit,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Height (U)
	//
	// Height in rack units
	// Maximum: 100
	// Minimum: 1
	UHeight int64 `json:"u_height,omitempty"`

	// Weight of the rack
	// Maximum: 1e+06
	// Minimum: -1e+06
	Weight *float64 `json:"weight,omitempty"`

	// * `kg` - Kilograms
	// * `g` - Grams
	// * `lb` - Pounds
	// * `oz` - Ounces
	// Enum: ["kg","g","lb","oz",""]
	WeightUnit string `json:"weight_unit,omitempty"`

	// Rail-to-rail width
	//
	// * `10` - 10 inches
	// * `19` - 19 inches
	// * `21` - 21 inches
	// * `23` - 23 inches
	// Maximum: 32767
	// Minimum: 0
	// Enum: [10,19,21,23]
	Width *int64 `json:"width,omitempty"`
}

// Validate validates this writable rack type request
func (m *WritableRackTypeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormFactor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountingDepth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterDepth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartingUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeightUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableRackTypeRequest) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

var writableRackTypeRequestTypeFormFactorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["2-post-frame","4-post-frame","4-post-cabinet","wall-frame","wall-frame-vertical","wall-cabinet","wall-cabinet-vertical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeRequestTypeFormFactorPropEnum = append(writableRackTypeRequestTypeFormFactorPropEnum, v)
	}
}

const (

	// WritableRackTypeRequestFormFactorNr2DashPostDashFrame captures enum value "2-post-frame"
	WritableRackTypeRequestFormFactorNr2DashPostDashFrame string = "2-post-frame"

	// WritableRackTypeRequestFormFactorNr4DashPostDashFrame captures enum value "4-post-frame"
	WritableRackTypeRequestFormFactorNr4DashPostDashFrame string = "4-post-frame"

	// WritableRackTypeRequestFormFactorNr4DashPostDashCabinet captures enum value "4-post-cabinet"
	WritableRackTypeRequestFormFactorNr4DashPostDashCabinet string = "4-post-cabinet"

	// WritableRackTypeRequestFormFactorWallDashFrame captures enum value "wall-frame"
	WritableRackTypeRequestFormFactorWallDashFrame string = "wall-frame"

	// WritableRackTypeRequestFormFactorWallDashFrameDashVertical captures enum value "wall-frame-vertical"
	WritableRackTypeRequestFormFactorWallDashFrameDashVertical string = "wall-frame-vertical"

	// WritableRackTypeRequestFormFactorWallDashCabinet captures enum value "wall-cabinet"
	WritableRackTypeRequestFormFactorWallDashCabinet string = "wall-cabinet"

	// WritableRackTypeRequestFormFactorWallDashCabinetDashVertical captures enum value "wall-cabinet-vertical"
	WritableRackTypeRequestFormFactorWallDashCabinetDashVertical string = "wall-cabinet-vertical"
)

// prop value enum
func (m *WritableRackTypeRequest) validateFormFactorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeRequestTypeFormFactorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRackTypeRequest) validateFormFactor(formats strfmt.Registry) error {

	if err := validate.Required("form_factor", "body", m.FormFactor); err != nil {
		return err
	}

	// value enum
	if err := m.validateFormFactorEnum("form_factor", "body", *m.FormFactor); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackTypeRequest) validateManufacturer(formats strfmt.Registry) error {

	if err := validate.Required("manufacturer", "body", m.Manufacturer); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackTypeRequest) validateMaxWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxWeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_weight", "body", *m.MaxWeight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_weight", "body", *m.MaxWeight, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackTypeRequest) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	if err := validate.MinLength("model", "body", *m.Model, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("model", "body", *m.Model, 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackTypeRequest) validateMountingDepth(formats strfmt.Registry) error {
	if swag.IsZero(m.MountingDepth) { // not required
		return nil
	}

	if err := validate.MinimumInt("mounting_depth", "body", *m.MountingDepth, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mounting_depth", "body", *m.MountingDepth, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackTypeRequest) validateOuterDepth(formats strfmt.Registry) error {
	if swag.IsZero(m.OuterDepth) { // not required
		return nil
	}

	if err := validate.MinimumInt("outer_depth", "body", *m.OuterDepth, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("outer_depth", "body", *m.OuterDepth, 32767, false); err != nil {
		return err
	}

	return nil
}

var writableRackTypeRequestTypeOuterUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["mm","in",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeRequestTypeOuterUnitPropEnum = append(writableRackTypeRequestTypeOuterUnitPropEnum, v)
	}
}

const (

	// WritableRackTypeRequestOuterUnitMm captures enum value "mm"
	WritableRackTypeRequestOuterUnitMm string = "mm"

	// WritableRackTypeRequestOuterUnitIn captures enum value "in"
	WritableRackTypeRequestOuterUnitIn string = "in"

	// WritableRackTypeRequestOuterUnitEmpty captures enum value ""
	WritableRackTypeRequestOuterUnitEmpty string = ""
)

// prop value enum
func (m *WritableRackTypeRequest) validateOuterUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeRequestTypeOuterUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRackTypeRequest) validateOuterUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.OuterUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateOuterUnitEnum("outer_unit", "body", m.OuterUnit); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackTypeRequest) validateOuterWidth(formats strfmt.Registry) error {
	if swag.IsZero(m.OuterWidth) { // not required
		return nil
	}

	if err := validate.MinimumInt("outer_width", "body", *m.OuterWidth, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("outer_width", "body", *m.OuterWidth, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackTypeRequest) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MinLength("slug", "body", *m.Slug, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", *m.Slug, 100); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", *m.Slug, `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackTypeRequest) validateStartingUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.StartingUnit) { // not required
		return nil
	}

	if err := validate.MinimumInt("starting_unit", "body", m.StartingUnit, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("starting_unit", "body", m.StartingUnit, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackTypeRequest) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableRackTypeRequest) validateUHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.UHeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("u_height", "body", m.UHeight, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("u_height", "body", m.UHeight, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableRackTypeRequest) validateWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.Minimum("weight", "body", *m.Weight, -1e+06, false); err != nil {
		return err
	}

	if err := validate.Maximum("weight", "body", *m.Weight, 1e+06, false); err != nil {
		return err
	}

	return nil
}

var writableRackTypeRequestTypeWeightUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kg","g","lb","oz",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeRequestTypeWeightUnitPropEnum = append(writableRackTypeRequestTypeWeightUnitPropEnum, v)
	}
}

const (

	// WritableRackTypeRequestWeightUnitKg captures enum value "kg"
	WritableRackTypeRequestWeightUnitKg string = "kg"

	// WritableRackTypeRequestWeightUnitG captures enum value "g"
	WritableRackTypeRequestWeightUnitG string = "g"

	// WritableRackTypeRequestWeightUnitLb captures enum value "lb"
	WritableRackTypeRequestWeightUnitLb string = "lb"

	// WritableRackTypeRequestWeightUnitOz captures enum value "oz"
	WritableRackTypeRequestWeightUnitOz string = "oz"

	// WritableRackTypeRequestWeightUnitEmpty captures enum value ""
	WritableRackTypeRequestWeightUnitEmpty string = ""
)

// prop value enum
func (m *WritableRackTypeRequest) validateWeightUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeRequestTypeWeightUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRackTypeRequest) validateWeightUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.WeightUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateWeightUnitEnum("weight_unit", "body", m.WeightUnit); err != nil {
		return err
	}

	return nil
}

var writableRackTypeRequestTypeWidthPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[10,19,21,23]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeRequestTypeWidthPropEnum = append(writableRackTypeRequestTypeWidthPropEnum, v)
	}
}

// prop value enum
func (m *WritableRackTypeRequest) validateWidthEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeRequestTypeWidthPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRackTypeRequest) validateWidth(formats strfmt.Registry) error {
	if swag.IsZero(m.Width) { // not required
		return nil
	}

	if err := validate.MinimumInt("width", "body", *m.Width, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("width", "body", *m.Width, 32767, false); err != nil {
		return err
	}

	// value enum
	if err := m.validateWidthEnum("width", "body", *m.Width); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable rack type request based on the context it is used
func (m *WritableRackTypeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableRackTypeRequest) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableRackTypeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableRackTypeRequest) UnmarshalBinary(b []byte) error {
	var res WritableRackTypeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
