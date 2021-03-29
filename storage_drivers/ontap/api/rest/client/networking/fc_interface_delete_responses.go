// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FcInterfaceDeleteReader is a Reader for the FcInterfaceDelete structure.
type FcInterfaceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcInterfaceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcInterfaceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcInterfaceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcInterfaceDeleteOK creates a FcInterfaceDeleteOK with default headers values
func NewFcInterfaceDeleteOK() *FcInterfaceDeleteOK {
	return &FcInterfaceDeleteOK{}
}

/* FcInterfaceDeleteOK describes a response with status code 200, with default header values.

OK
*/
type FcInterfaceDeleteOK struct {
}

func (o *FcInterfaceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /network/fc/interfaces/{uuid}][%d] fcInterfaceDeleteOK ", 200)
}

func (o *FcInterfaceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFcInterfaceDeleteDefault creates a FcInterfaceDeleteDefault with default headers values
func NewFcInterfaceDeleteDefault(code int) *FcInterfaceDeleteDefault {
	return &FcInterfaceDeleteDefault{
		_statusCode: code,
	}
}

/* FcInterfaceDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 53280992 | The FC interface could not be deleted because it is enabled. |

*/
type FcInterfaceDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fc interface delete default response
func (o *FcInterfaceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *FcInterfaceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /network/fc/interfaces/{uuid}][%d] fc_interface_delete default  %+v", o._statusCode, o.Payload)
}
func (o *FcInterfaceDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcInterfaceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}