package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Crystalix007/semantic-sensei/backend/api/redirect"
	"github.com/Crystalix007/semantic-sensei/backend/openapi"
	"github.com/Crystalix007/semantic-sensei/backend/storage"
	"github.com/huandu/go-sqlbuilder"
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
	var classificationTask storage.ClassificationTask

	if params.JSONBody != nil {
		embedding, err := params.JSONBody.Embedding.Bytes()
		if err != nil {
			return nil, fmt.Errorf(
				"api: error reading embedding: %w",
				err,
			)
		}

		classificationTask = storage.ClassificationTask{
			ProjectID: params.ProjectId,
			LLMInput:  params.JSONBody.LlmInput,
			LLMOutput: params.JSONBody.LlmOutput,
			Embedding: embedding,
		}
	}

	if params.FormdataBody != nil {
		embedding, err := params.FormdataBody.Embedding.Bytes()
		if err != nil {
			return nil, fmt.Errorf(
				"api: error reading embedding: %w",
				err,
			)
		}

		classificationTask = storage.ClassificationTask{
			ProjectID: params.ProjectId,
			LLMInput:  params.FormdataBody.LlmInput,
			LLMOutput: params.FormdataBody.LlmOutput,
			Embedding: embedding,
		}
	}

	taskID, err := a.db.CreateClassificationTask(ctx, classificationTask)
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

// GetProjectProjectIdClassificationTasks gets the project classification tasks
// for the given project, optionally filtered by whether they've been
// completed.
func (a *API) GetProjectProjectIdClassificationTasks(
	ctx context.Context,
	request openapi.GetProjectProjectIdClassificationTasksRequestObject,
) (openapi.GetProjectProjectIdClassificationTasksResponseObject, error) {
	var parameters storage.Parameters

	if request.Params.Labelled != nil {
		var condBuilder sqlbuilder.Cond

		if *request.Params.Labelled {
			parameters.Where = append(parameters.Where, condBuilder.IsNotNull("label_id"))
		} else {
			parameters.Where = append(parameters.Where, condBuilder.IsNull("label_id"))
		}
	}

	if request.Params.Page != nil {
		parameters.Page = *request.Params.Page
		parameters.PageSize = storage.DefaultPageSize
	}

	if request.Params.PageSize != nil {
		parameters.PageSize = *request.Params.PageSize
	}

	tasks, err := a.db.FindClassificationTasksForProject(
		ctx,
		request.ProjectId,
		parameters,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"api: error getting tasks for project %d: %w",
			request.ProjectId,
			err,
		)
	}

	projectTasks := make([]openapi.ClassificationTask, len(tasks))

	for i, task := range tasks {
		projectTasks[i] = openapi.ClassificationTask{
			CreatedAt: task.CreatedAt,
			Embedding: task.Embedding,
			Id:        task.ID,
			LabelId:   task.LabelID,
			LlmInput:  task.LLMInput,
			LlmOutput: task.LLMOutput,
			ProjectId: task.ProjectID,
		}
	}

	return openapi.GetProjectProjectIdClassificationTasks200JSONResponse{
		Data:  projectTasks,
		Total: 0,
	}, nil
}
