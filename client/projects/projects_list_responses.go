package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IllumiDesk/go-sdk/models"
)

// ProjectsListReader is a Reader for the ProjectsList structure.
type ProjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsListOK creates a ProjectsListOK with default headers values
func NewProjectsListOK() *ProjectsListOK {
	return &ProjectsListOK{}
}

/*ProjectsListOK handles this case with default header values.

Project list.
*/
type ProjectsListOK struct {
	Payload models.ProjectsListOKBody
}

func (o *ProjectsListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/projects/][%d] projectsListOK  %+v", 200, o.Payload)
}

func (o *ProjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
