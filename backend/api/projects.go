package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Crystalix007/semantic-sensei/backend/openapi"
)

// GetProjects retrieves a list of all projects.
func (a *API) GetProjects(
	ctx context.Context,
	request openapi.GetProjectsRequestObject,
) (openapi.GetProjectsResponseObject, error) {
	projects, err := a.db.FindProjects(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return openapi.GetProjects200JSONResponse{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get projects: %w", err)
	}

	var apiProjects []openapi.Project

	for _, project := range projects {
		apiProjects = append(apiProjects, openapi.Project{
			CreatedAt:   project.CreatedAt,
			Description: project.Description,
			Id:          project.ID,
			Labels:      nil,
			Name:        project.Name,
		})
	}

	return openapi.GetProjects200JSONResponse{
		Projects: &apiProjects,
	}, err
}
