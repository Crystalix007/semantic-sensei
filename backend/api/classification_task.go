package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Crystalix007/semantic-sensei/backend/api/redirect"
	"github.com/Crystalix007/semantic-sensei/backend/openapi"
	"github.com/Crystalix007/semantic-sensei/backend/storage"
)

// GetProjectProjectIdClassificationTaskId retrieves a classification task
// by its ID for a specific project.
func (a API) GetProjectProjectIdClassificationTaskId(
	ctx context.Context,
	params openapi.GetProjectProjectIdClassificationTaskIdRequestObject,
) (openapi.GetProjectProjectIdClassificationTaskIdResponseObject, error) {
	task, err := a.db.GetClassificationTask(ctx, params.Id)
	if errors.Is(err, sql.ErrNoRows) {
		return openapi.GetProjectProjectIdClassificationTaskId404Response{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf(
			"api: error getting classification task %d: %w",
			params.Id,
			err,
		)
	}

	return openapi.GetProjectProjectIdClassificationTaskId200JSONResponse{
		CreatedAt: task.CreatedAt,
		Embedding: task.Embedding,
		Id:        task.ID,
		LabelId:   task.LabelID,
		LlmInput:  task.LLMInput,
		LlmOutput: task.LLMOutput,
		ProjectId: task.ProjectID,
	}, nil
}

// PostProjectProjectIdClassificationTask creates a new classification task
// for a specific project.
func (a API) PostProjectProjectIdClassificationTask(
	ctx context.Context,
	params openapi.PostProjectProjectIdClassificationTaskRequestObject,
) (openapi.PostProjectProjectIdClassificationTaskResponseObject, error) {
	taskID, err := a.db.CreateClassificationTask(ctx, storage.ClassificationTask{
		ProjectID: params.ProjectId,
		LLMInput:  params.Body.LlmInput,
		LLMOutput: params.Body.LlmOutput,
		Embedding: params.Body.Embedding,
	})
	if err != nil {
		return nil, fmt.Errorf(
			"api: error creating classification task: %w",
			err,
		)
	}

	if redirect.Should(ctx) {
		return redirect.To(fmt.Sprintf("/project/%d/task/%d", params.ProjectId, taskID))
	}

	task, err := a.db.GetClassificationTask(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf(
			"api: error getting created classification task %d: %w",
			taskID,
			err,
		)
	}

	return openapi.PostProjectProjectIdClassificationTask201JSONResponse{
		CreatedAt: task.CreatedAt,
		Embedding: task.Embedding,
		Id:        taskID,
		LabelId:   task.LabelID,
		LlmInput:  task.LLMInput,
		LlmOutput: task.LLMOutput,
		ProjectId: task.ProjectID,
	}, nil
}

// PostProjectProjectIdClassificationTaskIdLabel allows a task to be labelled
// with the given task.
func (a *API) PostProjectProjectIdClassificationTaskIdLabel(
	ctx context.Context,
	request openapi.PostProjectProjectIdClassificationTaskIdLabelRequestObject,
) (openapi.PostProjectProjectIdClassificationTaskIdLabelResponseObject, error) {
	task, err := a.db.GetClassificationTask(ctx, request.Id)
	if errors.Is(err, sql.ErrNoRows) {
		return openapi.PostProjectProjectIdClassificationTaskIdLabel404Response{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf(
			"api: error getting classification task: %w",
			err,
		)
	}

	task.LabelID = &request.Body.Label

	fmt.Printf("Updating classification task to %#v, with label %#v\n", *task, *task.LabelID)

	err = a.db.UpdateClassificationTask(ctx, *task)
	if err != nil {
		return nil, fmt.Errorf(
			"api: error updating classification task: %w",
			err,
		)
	}

	if redirect.Should(ctx) {
		return redirect.To(fmt.Sprintf("/project/%d/task/%d", request.ProjectId, request.Id))
	}

	return openapi.PostProjectProjectIdClassificationTaskIdLabel200JSONResponse{
		CreatedAt: task.CreatedAt,
		Embedding: task.Embedding,
		Id:        task.ID,
		LabelId:   task.LabelID,
		LlmInput:  task.LLMInput,
		LlmOutput: task.LLMOutput,
		ProjectId: task.ProjectID,
	}, nil
}
