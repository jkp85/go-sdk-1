package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ProjectsServersIsAllowedListReader is a Reader for the ProjectsServersIsAllowedList structure.
type ProjectsServersIsAllowedListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersIsAllowedListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersIsAllowedListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersIsAllowedListOK creates a ProjectsServersIsAllowedListOK with default headers values
func NewProjectsServersIsAllowedListOK() *ProjectsServersIsAllowedListOK {
	return &ProjectsServersIsAllowedListOK{}
}

/*ProjectsServersIsAllowedListOK handles this case with default header values.

Is Allowed
*/
type ProjectsServersIsAllowedListOK struct {
}

func (o *ProjectsServersIsAllowedListOK) Error() string {
	return fmt.Sprintf("[GET /{namespace}/projects/{project_pk}/servers/{server_pk}/is-allowed/][%d] projectsServersIsAllowedListOK ", 200)
}

func (o *ProjectsServersIsAllowedListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
