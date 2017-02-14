package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ProjectsSSHTunnelsDeleteReader is a Reader for the ProjectsSSHTunnelsDelete structure.
type ProjectsSSHTunnelsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsSSHTunnelsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewProjectsSSHTunnelsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsSSHTunnelsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsSSHTunnelsDeleteNoContent creates a ProjectsSSHTunnelsDeleteNoContent with default headers values
func NewProjectsSSHTunnelsDeleteNoContent() *ProjectsSSHTunnelsDeleteNoContent {
	return &ProjectsSSHTunnelsDeleteNoContent{}
}

/*ProjectsSSHTunnelsDeleteNoContent handles this case with default header values.

SshTunnel deleted
*/
type ProjectsSSHTunnelsDeleteNoContent struct {
}

func (o *ProjectsSSHTunnelsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/{server_type}/{server_pk}/ssh-tunnels/{id}/][%d] projectsSshTunnelsDeleteNoContent ", 204)
}

func (o *ProjectsSSHTunnelsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsSSHTunnelsDeleteNotFound creates a ProjectsSSHTunnelsDeleteNotFound with default headers values
func NewProjectsSSHTunnelsDeleteNotFound() *ProjectsSSHTunnelsDeleteNotFound {
	return &ProjectsSSHTunnelsDeleteNotFound{}
}

/*ProjectsSSHTunnelsDeleteNotFound handles this case with default header values.

SshTunnel not found
*/
type ProjectsSSHTunnelsDeleteNotFound struct {
}

func (o *ProjectsSSHTunnelsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/{server_type}/{server_pk}/ssh-tunnels/{id}/][%d] projectsSshTunnelsDeleteNotFound ", 404)
}

func (o *ProjectsSSHTunnelsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
