package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// ProjectsFilesListReader is a Reader for the ProjectsFilesList structure.
type ProjectsFilesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsFilesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsFilesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsFilesListOK creates a ProjectsFilesListOK with default headers values
func NewProjectsFilesListOK() *ProjectsFilesListOK {
	return &ProjectsFilesListOK{}
}

/*ProjectsFilesListOK handles this case with default header values.

File list
*/
type ProjectsFilesListOK struct {
	Payload []*models.File
}

func (o *ProjectsFilesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/files/][%d] projectsFilesListOK  %+v", 200, o.Payload)
}

func (o *ProjectsFilesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
