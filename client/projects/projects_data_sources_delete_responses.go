package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ProjectsDataSourcesDeleteReader is a Reader for the ProjectsDataSourcesDelete structure.
type ProjectsDataSourcesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDataSourcesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewProjectsDataSourcesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsDataSourcesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsDataSourcesDeleteNoContent creates a ProjectsDataSourcesDeleteNoContent with default headers values
func NewProjectsDataSourcesDeleteNoContent() *ProjectsDataSourcesDeleteNoContent {
	return &ProjectsDataSourcesDeleteNoContent{}
}

/*ProjectsDataSourcesDeleteNoContent handles this case with default header values.

DataSource deleted
*/
type ProjectsDataSourcesDeleteNoContent struct {
}

func (o *ProjectsDataSourcesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/data-sources/{server}/][%d] projectsDataSourcesDeleteNoContent ", 204)
}

func (o *ProjectsDataSourcesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsDataSourcesDeleteNotFound creates a ProjectsDataSourcesDeleteNotFound with default headers values
func NewProjectsDataSourcesDeleteNotFound() *ProjectsDataSourcesDeleteNotFound {
	return &ProjectsDataSourcesDeleteNotFound{}
}

/*ProjectsDataSourcesDeleteNotFound handles this case with default header values.

DataSource not found
*/
type ProjectsDataSourcesDeleteNotFound struct {
}

func (o *ProjectsDataSourcesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/data-sources/{server}/][%d] projectsDataSourcesDeleteNotFound ", 404)
}

func (o *ProjectsDataSourcesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
