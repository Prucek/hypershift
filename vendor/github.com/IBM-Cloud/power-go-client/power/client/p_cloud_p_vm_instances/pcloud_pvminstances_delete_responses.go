// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesDeleteReader is a Reader for the PcloudPvminstancesDelete structure.
type PcloudPvminstancesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudPvminstancesDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesDeleteOK creates a PcloudPvminstancesDeleteOK with default headers values
func NewPcloudPvminstancesDeleteOK() *PcloudPvminstancesDeleteOK {
	return &PcloudPvminstancesDeleteOK{}
}

/* PcloudPvminstancesDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesDeleteOK struct {
	Payload models.Object
}

func (o *PcloudPvminstancesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteOK  %+v", 200, o.Payload)
}
func (o *PcloudPvminstancesDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesDeleteBadRequest creates a PcloudPvminstancesDeleteBadRequest with default headers values
func NewPcloudPvminstancesDeleteBadRequest() *PcloudPvminstancesDeleteBadRequest {
	return &PcloudPvminstancesDeleteBadRequest{}
}

/* PcloudPvminstancesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesDeleteBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudPvminstancesDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesDeleteUnauthorized creates a PcloudPvminstancesDeleteUnauthorized with default headers values
func NewPcloudPvminstancesDeleteUnauthorized() *PcloudPvminstancesDeleteUnauthorized {
	return &PcloudPvminstancesDeleteUnauthorized{}
}

/* PcloudPvminstancesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesDeleteUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudPvminstancesDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesDeleteNotFound creates a PcloudPvminstancesDeleteNotFound with default headers values
func NewPcloudPvminstancesDeleteNotFound() *PcloudPvminstancesDeleteNotFound {
	return &PcloudPvminstancesDeleteNotFound{}
}

/* PcloudPvminstancesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesDeleteNotFound struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteNotFound  %+v", 404, o.Payload)
}
func (o *PcloudPvminstancesDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesDeleteGone creates a PcloudPvminstancesDeleteGone with default headers values
func NewPcloudPvminstancesDeleteGone() *PcloudPvminstancesDeleteGone {
	return &PcloudPvminstancesDeleteGone{}
}

/* PcloudPvminstancesDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudPvminstancesDeleteGone struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteGone  %+v", 410, o.Payload)
}
func (o *PcloudPvminstancesDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesDeleteInternalServerError creates a PcloudPvminstancesDeleteInternalServerError with default headers values
func NewPcloudPvminstancesDeleteInternalServerError() *PcloudPvminstancesDeleteInternalServerError {
	return &PcloudPvminstancesDeleteInternalServerError{}
}

/* PcloudPvminstancesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudPvminstancesDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}