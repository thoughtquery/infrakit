package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// RemoveUserReader is a Reader for the RemoveUser structure.
type RemoveUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRemoveUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRemoveUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewRemoveUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveUserNoContent creates a RemoveUserNoContent with default headers values
func NewRemoveUserNoContent() *RemoveUserNoContent {
	return &RemoveUserNoContent{}
}

/*RemoveUserNoContent handles this case with default header values.

Successfully deleted the specified user
*/
type RemoveUserNoContent struct {
}

func (o *RemoveUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{name}][%d] removeUserNoContent ", 204)
}

func (o *RemoveUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveUserUnauthorized creates a RemoveUserUnauthorized with default headers values
func NewRemoveUserUnauthorized() *RemoveUserUnauthorized {
	return &RemoveUserUnauthorized{}
}

/*RemoveUserUnauthorized handles this case with default header values.

Unauthorized
*/
type RemoveUserUnauthorized struct {
	Payload RemoveUserUnauthorizedBody
}

func (o *RemoveUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /users/{name}][%d] removeUserUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserForbidden creates a RemoveUserForbidden with default headers values
func NewRemoveUserForbidden() *RemoveUserForbidden {
	return &RemoveUserForbidden{}
}

/*RemoveUserForbidden handles this case with default header values.

Forbidden
*/
type RemoveUserForbidden struct {
	Payload RemoveUserForbiddenBody
}

func (o *RemoveUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/{name}][%d] removeUserForbidden  %+v", 403, o.Payload)
}

func (o *RemoveUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserDefault creates a RemoveUserDefault with default headers values
func NewRemoveUserDefault(code int) *RemoveUserDefault {
	return &RemoveUserDefault{
		_statusCode: code,
	}
}

/*RemoveUserDefault handles this case with default header values.

Unexpected error
*/
type RemoveUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the remove user default response
func (o *RemoveUserDefault) Code() int {
	return o._statusCode
}

func (o *RemoveUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/{name}][%d] removeUser default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveUserForbiddenBody remove user forbidden body
swagger:model RemoveUserForbiddenBody
*/
type RemoveUserForbiddenBody interface{}

/*RemoveUserUnauthorizedBody remove user unauthorized body
swagger:model RemoveUserUnauthorizedBody
*/
type RemoveUserUnauthorizedBody interface{}