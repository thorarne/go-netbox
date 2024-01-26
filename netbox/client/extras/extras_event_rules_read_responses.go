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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasEventRulesReadReader is a Reader for the ExtrasEventRulesRead structure.
type ExtrasEventRulesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasEventRulesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasEventRulesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasEventRulesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasEventRulesReadOK creates a ExtrasEventRulesReadOK with default headers values
func NewExtrasEventRulesReadOK() *ExtrasEventRulesReadOK {
	return &ExtrasEventRulesReadOK{}
}

/*
ExtrasEventRulesReadOK describes a response with status code 200, with default header values.

ExtrasEventRulesReadOK extras event rules read o k
*/
type ExtrasEventRulesReadOK struct {
	Payload *models.EventRule
}

// IsSuccess returns true when this extras event rules read o k response has a 2xx status code
func (o *ExtrasEventRulesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras event rules read o k response has a 3xx status code
func (o *ExtrasEventRulesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras event rules read o k response has a 4xx status code
func (o *ExtrasEventRulesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras event rules read o k response has a 5xx status code
func (o *ExtrasEventRulesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras event rules read o k response a status code equal to that given
func (o *ExtrasEventRulesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras event rules read o k response
func (o *ExtrasEventRulesReadOK) Code() int {
	return 200
}

func (o *ExtrasEventRulesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/event-rules/{id}/][%d] extrasEventRulesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasEventRulesReadOK) String() string {
	return fmt.Sprintf("[GET /extras/event-rules/{id}/][%d] extrasEventRulesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasEventRulesReadOK) GetPayload() *models.EventRule {
	return o.Payload
}

func (o *ExtrasEventRulesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasEventRulesReadDefault creates a ExtrasEventRulesReadDefault with default headers values
func NewExtrasEventRulesReadDefault(code int) *ExtrasEventRulesReadDefault {
	return &ExtrasEventRulesReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasEventRulesReadDefault describes a response with status code -1, with default header values.

ExtrasEventRulesReadDefault extras event rules read default
*/
type ExtrasEventRulesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras event rules read default response has a 2xx status code
func (o *ExtrasEventRulesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras event rules read default response has a 3xx status code
func (o *ExtrasEventRulesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras event rules read default response has a 4xx status code
func (o *ExtrasEventRulesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras event rules read default response has a 5xx status code
func (o *ExtrasEventRulesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras event rules read default response a status code equal to that given
func (o *ExtrasEventRulesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras event rules read default response
func (o *ExtrasEventRulesReadDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasEventRulesReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/event-rules/{id}/][%d] extras_event_rules_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasEventRulesReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/event-rules/{id}/][%d] extras_event_rules_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasEventRulesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasEventRulesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
