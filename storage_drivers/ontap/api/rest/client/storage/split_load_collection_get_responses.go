// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SplitLoadCollectionGetReader is a Reader for the SplitLoadCollectionGet structure.
type SplitLoadCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SplitLoadCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSplitLoadCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSplitLoadCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSplitLoadCollectionGetOK creates a SplitLoadCollectionGetOK with default headers values
func NewSplitLoadCollectionGetOK() *SplitLoadCollectionGetOK {
	return &SplitLoadCollectionGetOK{}
}

/* SplitLoadCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SplitLoadCollectionGetOK struct {
	Payload *models.SplitLoadResponse
}

func (o *SplitLoadCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/file/clone/split-loads][%d] splitLoadCollectionGetOK  %+v", 200, o.Payload)
}
func (o *SplitLoadCollectionGetOK) GetPayload() *models.SplitLoadResponse {
	return o.Payload
}

func (o *SplitLoadCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SplitLoadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSplitLoadCollectionGetDefault creates a SplitLoadCollectionGetDefault with default headers values
func NewSplitLoadCollectionGetDefault(code int) *SplitLoadCollectionGetDefault {
	return &SplitLoadCollectionGetDefault{
		_statusCode: code,
	}
}

/* SplitLoadCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SplitLoadCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the split load collection get default response
func (o *SplitLoadCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SplitLoadCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/file/clone/split-loads][%d] split_load_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *SplitLoadCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SplitLoadCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
