package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ActionsCancelCreateReader is a Reader for the ActionsCancelCreate structure.
type ActionsCancelCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionsCancelCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewActionsCancelCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewActionsCancelCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActionsCancelCreateCreated creates a ActionsCancelCreateCreated with default headers values
func NewActionsCancelCreateCreated() *ActionsCancelCreateCreated {
	return &ActionsCancelCreateCreated{}
}

/*ActionsCancelCreateCreated handles this case with default header values.

Cancel
*/
type ActionsCancelCreateCreated struct {
}

func (o *ActionsCancelCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/actions/{id}/cancel/][%d] actionsCancelCreateCreated ", 201)
}

func (o *ActionsCancelCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsCancelCreateBadRequest creates a ActionsCancelCreateBadRequest with default headers values
func NewActionsCancelCreateBadRequest() *ActionsCancelCreateBadRequest {
	return &ActionsCancelCreateBadRequest{}
}

/*ActionsCancelCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ActionsCancelCreateBadRequest struct {
}

func (o *ActionsCancelCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/actions/{id}/cancel/][%d] actionsCancelCreateBadRequest ", 400)
}

func (o *ActionsCancelCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
