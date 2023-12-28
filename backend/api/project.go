package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Crystalix007/semantic-sensei/backend/openapi"
	"github.com/Crystalix007/semantic-sensei/backend/storage"
)

// GetProjectId retrieves a project by its ID.
// It takes a context.Context and a GetProjectIdRequestObject as input parameters.
// It returns a GetProjectIdResponseObject and an error.
// The error will be non-nil if there was an error retrieving the project.
func (a API) GetProjectId(ctx context.Context, params openapi.GetProjectIdRequestObject) (openapi.GetProjectIdResponseObject, error) {
	// Retrieve the project from the database using the provided ID
	project, err := a.db.GetProject(ctx, params.Id)
	if errors.Is(err, sql.ErrNoRows) {
		return openapi.GetProjectId404Response{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf(
			"api: error getting project %d: %w",
			params.Id,
			err,
		)
	}

	// Create and return the response object with the project details
	return openapi.GetProjectId200JSONResponse{
		CreatedAt:   project.CreatedAt,
		Description: project.Description,
		Id:          project.ID,
		Name:        project.Name,
	}, nil
}

// PostProject creates a new project in the database based on the provided
// parameters.
// It returns the created project details in the form of a
// PostProjectResponseObject.
// If an error occurs during project creation or retrieval, it returns an
// error.
func (a API) PostProject(ctx context.Context, params openapi.PostProjectRequestObject) (openapi.PostProjectResponseObject, error) {
	// Create the project in the database
	projectID, err := a.db.CreateProject(ctx, storage.Project{
		Name:        params.Body.Name,
		Description: params.Body.Description,
	})
	if err != nil {
		return nil, fmt.Errorf(
			"api: error creating project: %w",
			err,
		)
	}

	project, err := a.db.GetProject(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf(
			"api: error getting created project %d: %w",
			projectID,
			err,
		)
	}

	// Create and return the response object with the project details
	return openapi.PostProject201JSONResponse{
		CreatedAt:   project.CreatedAt,
		Description: project.Description,
		Id:          project.ID,
		Name:        project.Name,
	}, nil
}