package api

import (
	"context"
	"fmt"

	"github.com/Crystalix007/semantic-sensei/backend/openapi"
	"github.com/Crystalix007/semantic-sensei/backend/storage"
)

// GetProjectProjectIdClassificationTaskLabelId retrieves a classification task
// label by its ID for a specific project.
func (a API) GetProjectProjectIdClassificationTaskLabelId(
	ctx context.Context,
	params openapi.GetProjectProjectIdClassificationTaskLabelIdRequestObject,
) (openapi.GetProjectProjectIdClassificationTaskLabelIdResponseObject, error) {
	label, err := a.db.GetClassificationTaskLabel(ctx, params.Id)
	if err != nil {
		return nil, fmt.Errorf(
			"api: error getting classification task label %d: %w",
			params.Id,
			err,
		)
	}

	return openapi.GetProjectProjectIdClassificationTaskLabelId200JSONResponse{
		CreatedAt: label.CreatedAt,
		Id:        label.ID,
		Label:     label.Label,
		ProjectId: label.ProjectID,
	}, nil
}

// PostProjectProjectIdClassificationTaskLabel is a function that handles the
// creation of a classification task label for a specific project.
// The function creates a classification task label in the database using the
// provided parameters and returns the created label.
// If an error occurs during the creation or retrieval of the label, an error
// is returned.
func (a API) PostProjectProjectIdClassificationTaskLabel(
	ctx context.Context,
	params openapi.PostProjectProjectIdClassificationTaskLabelRequestObject,
) (openapi.PostProjectProjectIdClassificationTaskLabelResponseObject, error) {
	labelID, err := a.db.CreateClassificationTaskLabel(ctx, storage.ClassificationTaskLabel{
		ProjectID: params.ProjectId,
		Label:     params.Body.Label,
	})
	if err != nil {
		return nil, fmt.Errorf(
			"api: error creating classification task label: %w",
			err,
		)
	}

	label, err := a.db.GetClassificationTaskLabel(ctx, labelID)
	if err != nil {
		return nil, fmt.Errorf(
			"api: error getting created classification task label %d: %w",
			labelID,
			err,
		)
	}

	return openapi.PostProjectProjectIdClassificationTaskLabel201JSONResponse{
		CreatedAt: label.CreatedAt,
		Id:        labelID,
		Label:     label.Label,
		ProjectId: label.ProjectID,
	}, nil
}
