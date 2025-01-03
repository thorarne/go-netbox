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

// RackType Adds support for custom fields and tags.
//
// swagger:model RackType
type RackType struct {

	// comments
	Comments string `json:"comments,omitempty"`

	// created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// custom fields
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`

	// Descending units
	//
	// Units are numbered top-to-bottom
	DescUnits bool `json:"desc_units,omitempty"`

	// description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// display url
	// Read Only: true
	// Format: uri
	DisplayURL strfmt.URI `json:"display_url,omitempty"`

	// form factor
	FormFactor *RackTypeFormFactor `json:"form_factor,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// manufacturer
	// Required: true
	Manufacturer *NestedManufacturer `json:"manufacturer"`

	// Maximum load capacity for the rack
	// Maximum: 2.147483647e+09
	// Minimum: 0
	MaxWeight *int64 `json:"max_weight,omitempty"`

	// model
	// Required: true
	// Max Length: 100
	Model *string `json:"model"`

	// Maximum depth of a mounted device, in millimeters. For four-post racks, this is the distance between the front and rear rails.
	// Maximum: 32767
	// Minimum: 0
	MountingDepth *int64 `json:"mounting_depth,omitempty"`

	// Outer dimension of rack (depth)
	// Maximum: 32767
	// Minimum: 0
	OuterDepth *int64 `json:"outer_depth,omitempty"`

	// outer unit
	OuterUnit *RackTypeOuterUnit `json:"outer_unit,omitempty"`

	// Outer dimension of rack (width)
	// Maximum: 32767
	// Minimum: 0
	OuterWidth *int64 `json:"outer_width,omitempty"`

	// slug
	// Required: true
	// Max Length: 100
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

	// url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// weight
	// Maximum: 1e+06
	// Minimum: -1e+06
	Weight *float64 `json:"weight,omitempty"`

	// weight unit
	WeightUnit *RackTypeWeightUnit `json:"weight_unit,omitempty"`

	// width
	Width *RackTypeWidth `json:"width,omitempty"`
}

// Validate validates this rack type
func (m *RackType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormFactor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
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

	if err := m.validateURL(formats); err != nil {
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

func (m *RackType) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RackType) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *RackType) validateDisplayURL(formats strfmt.Registry) error {
	if swag.IsZero(m.DisplayURL) { // not required
		return nil
	}

	if err := validate.FormatOf("display_url", "body", "uri", m.DisplayURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RackType) validateFormFactor(formats strfmt.Registry) error {
	if swag.IsZero(m.FormFactor) { // not required
		return nil
	}

	if m.FormFactor != nil {
		if err := m.FormFactor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("form_factor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("form_factor")
			}
			return err
		}
	}

	return nil
}

func (m *RackType) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RackType) validateManufacturer(formats strfmt.Registry) error {

	if err := validate.Required("manufacturer", "body", m.Manufacturer); err != nil {
		return err
	}

	if m.Manufacturer != nil {
		if err := m.Manufacturer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *RackType) validateMaxWeight(formats strfmt.Registry) error {
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

func (m *RackType) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	if err := validate.MaxLength("model", "body", *m.Model, 100); err != nil {
		return err
	}

	return nil
}

func (m *RackType) validateMountingDepth(formats strfmt.Registry) error {
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

func (m *RackType) validateOuterDepth(formats strfmt.Registry) error {
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

func (m *RackType) validateOuterUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.OuterUnit) { // not required
		return nil
	}

	if m.OuterUnit != nil {
		if err := m.OuterUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outer_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outer_unit")
			}
			return err
		}
	}

	return nil
}

func (m *RackType) validateOuterWidth(formats strfmt.Registry) error {
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

func (m *RackType) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
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

func (m *RackType) validateStartingUnit(formats strfmt.Registry) error {
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

func (m *RackType) validateTags(formats strfmt.Registry) error {
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

func (m *RackType) validateUHeight(formats strfmt.Registry) error {
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

func (m *RackType) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RackType) validateWeight(formats strfmt.Registry) error {
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

func (m *RackType) validateWeightUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.WeightUnit) { // not required
		return nil
	}

	if m.WeightUnit != nil {
		if err := m.WeightUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight_unit")
			}
			return err
		}
	}

	return nil
}

func (m *RackType) validateWidth(formats strfmt.Registry) error {
	if swag.IsZero(m.Width) { // not required
		return nil
	}

	if m.Width != nil {
		if err := m.Width.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("width")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("width")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rack type based on the context it is used
func (m *RackType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplayURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFormFactor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManufacturer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOuterUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeightUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWidth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackType) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *RackType) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *RackType) contextValidateDisplayURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display_url", "body", strfmt.URI(m.DisplayURL)); err != nil {
		return err
	}

	return nil
}

func (m *RackType) contextValidateFormFactor(ctx context.Context, formats strfmt.Registry) error {

	if m.FormFactor != nil {

		if swag.IsZero(m.FormFactor) { // not required
			return nil
		}

		if err := m.FormFactor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("form_factor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("form_factor")
			}
			return err
		}
	}

	return nil
}

func (m *RackType) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *RackType) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *RackType) contextValidateManufacturer(ctx context.Context, formats strfmt.Registry) error {

	if m.Manufacturer != nil {

		if err := m.Manufacturer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *RackType) contextValidateOuterUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.OuterUnit != nil {

		if swag.IsZero(m.OuterUnit) { // not required
			return nil
		}

		if err := m.OuterUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outer_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outer_unit")
			}
			return err
		}
	}

	return nil
}

func (m *RackType) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RackType) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *RackType) contextValidateWeightUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.WeightUnit != nil {

		if swag.IsZero(m.WeightUnit) { // not required
			return nil
		}

		if err := m.WeightUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight_unit")
			}
			return err
		}
	}

	return nil
}

func (m *RackType) contextValidateWidth(ctx context.Context, formats strfmt.Registry) error {

	if m.Width != nil {

		if swag.IsZero(m.Width) { // not required
			return nil
		}

		if err := m.Width.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("width")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("width")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackType) UnmarshalBinary(b []byte) error {
	var res RackType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackTypeFormFactor rack type form factor
//
// swagger:model RackTypeFormFactor
type RackTypeFormFactor struct {

	// label
	// Enum: ["2-post frame","4-post frame","4-post cabinet","Wall-mounted frame","Wall-mounted frame (vertical)","Wall-mounted cabinet","Wall-mounted cabinet (vertical)"]
	Label string `json:"label,omitempty"`

	// * `2-post-frame` - 2-post frame
	// * `4-post-frame` - 4-post frame
	// * `4-post-cabinet` - 4-post cabinet
	// * `wall-frame` - Wall-mounted frame
	// * `wall-frame-vertical` - Wall-mounted frame (vertical)
	// * `wall-cabinet` - Wall-mounted cabinet
	// * `wall-cabinet-vertical` - Wall-mounted cabinet (vertical)
	// Enum: ["2-post-frame","4-post-frame","4-post-cabinet","wall-frame","wall-frame-vertical","wall-cabinet","wall-cabinet-vertical",""]
	Value string `json:"value,omitempty"`
}

// Validate validates this rack type form factor
func (m *RackTypeFormFactor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rackTypeFormFactorTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["2-post frame","4-post frame","4-post cabinet","Wall-mounted frame","Wall-mounted frame (vertical)","Wall-mounted cabinet","Wall-mounted cabinet (vertical)"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTypeFormFactorTypeLabelPropEnum = append(rackTypeFormFactorTypeLabelPropEnum, v)
	}
}

const (

	// RackTypeFormFactorLabelNr2DashPostFrame captures enum value "2-post frame"
	RackTypeFormFactorLabelNr2DashPostFrame string = "2-post frame"

	// RackTypeFormFactorLabelNr4DashPostFrame captures enum value "4-post frame"
	RackTypeFormFactorLabelNr4DashPostFrame string = "4-post frame"

	// RackTypeFormFactorLabelNr4DashPostCabinet captures enum value "4-post cabinet"
	RackTypeFormFactorLabelNr4DashPostCabinet string = "4-post cabinet"

	// RackTypeFormFactorLabelWallDashMountedFrame captures enum value "Wall-mounted frame"
	RackTypeFormFactorLabelWallDashMountedFrame string = "Wall-mounted frame"

	// RackTypeFormFactorLabelWallDashMountedFrameVertical captures enum value "Wall-mounted frame (vertical)"
	RackTypeFormFactorLabelWallDashMountedFrameVertical string = "Wall-mounted frame (vertical)"

	// RackTypeFormFactorLabelWallDashMountedCabinet captures enum value "Wall-mounted cabinet"
	RackTypeFormFactorLabelWallDashMountedCabinet string = "Wall-mounted cabinet"

	// RackTypeFormFactorLabelWallDashMountedCabinetVertical captures enum value "Wall-mounted cabinet (vertical)"
	RackTypeFormFactorLabelWallDashMountedCabinetVertical string = "Wall-mounted cabinet (vertical)"
)

// prop value enum
func (m *RackTypeFormFactor) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackTypeFormFactorTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackTypeFormFactor) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	// value enum
	if err := m.validateLabelEnum("form_factor"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

var rackTypeFormFactorTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["2-post-frame","4-post-frame","4-post-cabinet","wall-frame","wall-frame-vertical","wall-cabinet","wall-cabinet-vertical",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTypeFormFactorTypeValuePropEnum = append(rackTypeFormFactorTypeValuePropEnum, v)
	}
}

const (

	// RackTypeFormFactorValueNr2DashPostDashFrame captures enum value "2-post-frame"
	RackTypeFormFactorValueNr2DashPostDashFrame string = "2-post-frame"

	// RackTypeFormFactorValueNr4DashPostDashFrame captures enum value "4-post-frame"
	RackTypeFormFactorValueNr4DashPostDashFrame string = "4-post-frame"

	// RackTypeFormFactorValueNr4DashPostDashCabinet captures enum value "4-post-cabinet"
	RackTypeFormFactorValueNr4DashPostDashCabinet string = "4-post-cabinet"

	// RackTypeFormFactorValueWallDashFrame captures enum value "wall-frame"
	RackTypeFormFactorValueWallDashFrame string = "wall-frame"

	// RackTypeFormFactorValueWallDashFrameDashVertical captures enum value "wall-frame-vertical"
	RackTypeFormFactorValueWallDashFrameDashVertical string = "wall-frame-vertical"

	// RackTypeFormFactorValueWallDashCabinet captures enum value "wall-cabinet"
	RackTypeFormFactorValueWallDashCabinet string = "wall-cabinet"

	// RackTypeFormFactorValueWallDashCabinetDashVertical captures enum value "wall-cabinet-vertical"
	RackTypeFormFactorValueWallDashCabinetDashVertical string = "wall-cabinet-vertical"

	// RackTypeFormFactorValueEmpty captures enum value ""
	RackTypeFormFactorValueEmpty string = ""
)

// prop value enum
func (m *RackTypeFormFactor) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackTypeFormFactorTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackTypeFormFactor) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	// value enum
	if err := m.validateValueEnum("form_factor"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rack type form factor based on context it is used
func (m *RackTypeFormFactor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RackTypeFormFactor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackTypeFormFactor) UnmarshalBinary(b []byte) error {
	var res RackTypeFormFactor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackTypeOuterUnit rack type outer unit
//
// swagger:model RackTypeOuterUnit
type RackTypeOuterUnit struct {

	// label
	// Enum: ["Millimeters","Inches"]
	Label string `json:"label,omitempty"`

	// * `mm` - Millimeters
	// * `in` - Inches
	// Enum: ["mm","in",""]
	Value string `json:"value,omitempty"`
}

// Validate validates this rack type outer unit
func (m *RackTypeOuterUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rackTypeOuterUnitTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Millimeters","Inches"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTypeOuterUnitTypeLabelPropEnum = append(rackTypeOuterUnitTypeLabelPropEnum, v)
	}
}

const (

	// RackTypeOuterUnitLabelMillimeters captures enum value "Millimeters"
	RackTypeOuterUnitLabelMillimeters string = "Millimeters"

	// RackTypeOuterUnitLabelInches captures enum value "Inches"
	RackTypeOuterUnitLabelInches string = "Inches"
)

// prop value enum
func (m *RackTypeOuterUnit) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackTypeOuterUnitTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackTypeOuterUnit) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	// value enum
	if err := m.validateLabelEnum("outer_unit"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

var rackTypeOuterUnitTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["mm","in",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTypeOuterUnitTypeValuePropEnum = append(rackTypeOuterUnitTypeValuePropEnum, v)
	}
}

const (

	// RackTypeOuterUnitValueMm captures enum value "mm"
	RackTypeOuterUnitValueMm string = "mm"

	// RackTypeOuterUnitValueIn captures enum value "in"
	RackTypeOuterUnitValueIn string = "in"

	// RackTypeOuterUnitValueEmpty captures enum value ""
	RackTypeOuterUnitValueEmpty string = ""
)

// prop value enum
func (m *RackTypeOuterUnit) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackTypeOuterUnitTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackTypeOuterUnit) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	// value enum
	if err := m.validateValueEnum("outer_unit"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rack type outer unit based on context it is used
func (m *RackTypeOuterUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RackTypeOuterUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackTypeOuterUnit) UnmarshalBinary(b []byte) error {
	var res RackTypeOuterUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackTypeWeightUnit rack type weight unit
//
// swagger:model RackTypeWeightUnit
type RackTypeWeightUnit struct {

	// label
	// Enum: ["Kilograms","Grams","Pounds","Ounces"]
	Label string `json:"label,omitempty"`

	// * `kg` - Kilograms
	// * `g` - Grams
	// * `lb` - Pounds
	// * `oz` - Ounces
	// Enum: ["kg","g","lb","oz",""]
	Value string `json:"value,omitempty"`
}

// Validate validates this rack type weight unit
func (m *RackTypeWeightUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rackTypeWeightUnitTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Kilograms","Grams","Pounds","Ounces"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTypeWeightUnitTypeLabelPropEnum = append(rackTypeWeightUnitTypeLabelPropEnum, v)
	}
}

const (

	// RackTypeWeightUnitLabelKilograms captures enum value "Kilograms"
	RackTypeWeightUnitLabelKilograms string = "Kilograms"

	// RackTypeWeightUnitLabelGrams captures enum value "Grams"
	RackTypeWeightUnitLabelGrams string = "Grams"

	// RackTypeWeightUnitLabelPounds captures enum value "Pounds"
	RackTypeWeightUnitLabelPounds string = "Pounds"

	// RackTypeWeightUnitLabelOunces captures enum value "Ounces"
	RackTypeWeightUnitLabelOunces string = "Ounces"
)

// prop value enum
func (m *RackTypeWeightUnit) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackTypeWeightUnitTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackTypeWeightUnit) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	// value enum
	if err := m.validateLabelEnum("weight_unit"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

var rackTypeWeightUnitTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kg","g","lb","oz",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTypeWeightUnitTypeValuePropEnum = append(rackTypeWeightUnitTypeValuePropEnum, v)
	}
}

const (

	// RackTypeWeightUnitValueKg captures enum value "kg"
	RackTypeWeightUnitValueKg string = "kg"

	// RackTypeWeightUnitValueG captures enum value "g"
	RackTypeWeightUnitValueG string = "g"

	// RackTypeWeightUnitValueLb captures enum value "lb"
	RackTypeWeightUnitValueLb string = "lb"

	// RackTypeWeightUnitValueOz captures enum value "oz"
	RackTypeWeightUnitValueOz string = "oz"

	// RackTypeWeightUnitValueEmpty captures enum value ""
	RackTypeWeightUnitValueEmpty string = ""
)

// prop value enum
func (m *RackTypeWeightUnit) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackTypeWeightUnitTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackTypeWeightUnit) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	// value enum
	if err := m.validateValueEnum("weight_unit"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rack type weight unit based on context it is used
func (m *RackTypeWeightUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RackTypeWeightUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackTypeWeightUnit) UnmarshalBinary(b []byte) error {
	var res RackTypeWeightUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackTypeWidth rack type width
//
// swagger:model RackTypeWidth
type RackTypeWidth struct {

	// label
	// Enum: ["10 inches","19 inches","21 inches","23 inches"]
	Label string `json:"label,omitempty"`

	// * `10` - 10 inches
	// * `19` - 19 inches
	// * `21` - 21 inches
	// * `23` - 23 inches
	// Enum: [10,19,21,23]
	Value int64 `json:"value,omitempty"`
}

// Validate validates this rack type width
func (m *RackTypeWidth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rackTypeWidthTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["10 inches","19 inches","21 inches","23 inches"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTypeWidthTypeLabelPropEnum = append(rackTypeWidthTypeLabelPropEnum, v)
	}
}

const (

	// RackTypeWidthLabelNr10Inches captures enum value "10 inches"
	RackTypeWidthLabelNr10Inches string = "10 inches"

	// RackTypeWidthLabelNr19Inches captures enum value "19 inches"
	RackTypeWidthLabelNr19Inches string = "19 inches"

	// RackTypeWidthLabelNr21Inches captures enum value "21 inches"
	RackTypeWidthLabelNr21Inches string = "21 inches"

	// RackTypeWidthLabelNr23Inches captures enum value "23 inches"
	RackTypeWidthLabelNr23Inches string = "23 inches"
)

// prop value enum
func (m *RackTypeWidth) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackTypeWidthTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackTypeWidth) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	// value enum
	if err := m.validateLabelEnum("width"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

var rackTypeWidthTypeValuePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[10,19,21,23]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTypeWidthTypeValuePropEnum = append(rackTypeWidthTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *RackTypeWidth) validateValueEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, rackTypeWidthTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackTypeWidth) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	// value enum
	if err := m.validateValueEnum("width"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rack type width based on context it is used
func (m *RackTypeWidth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RackTypeWidth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackTypeWidth) UnmarshalBinary(b []byte) error {
	var res RackTypeWidth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
