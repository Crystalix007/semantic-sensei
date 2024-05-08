package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	"github.com/Crystalix007/semantic-sensei/backend/api/headers"
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
	pendingTask, err := a.db.GetPendingClassificationTask(ctx, params.Id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf(
			"api: error getting pending classification task %d: %w",
			params.Id,
			err,
		)
	}

	if err == nil {
		var b openapi.ClassificationTaskOrPendingClassificationTask

		if err := b.FromPendingClassificationTask(openapi.PendingClassificationTask{
			CreatedAt: pendingTask.CreatedAt,
			Embedding: pendingTask.Embedding,
			Id:        pendingTask.ID,
			LlmInput:  pendingTask.LLMInput,
			LlmOutput: pendingTask.LLMOutput,
			ProjectId: pendingTask.ProjectID,
		}); err != nil {
			return nil, fmt.Errorf(
				"api: error encoding pending classification task in response: %w",
				err,
			)
		}

		return openapi.GetProjectProjectIdClassificationTaskId200JSONResponse(b), nil
	}

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

	var b openapi.ClassificationTaskOrPendingClassificationTask

	if err := b.FromClassificationTask(openapi.ClassificationTask{
		CreatedAt: task.CreatedAt,
		Embedding: task.Embedding,
		Id:        task.ID,
		LabelId:   task.LabelID,
		LlmInput:  task.LLMInput,
		LlmOutput: task.LLMOutput,
		ProjectId: task.ProjectID,
	}); err != nil {
		return nil, fmt.Errorf(
			"api: error encoding classification task in response: %w",
			err,
		)
	}

	return openapi.GetProjectProjectIdClassificationTaskId200JSONResponse(b), nil
}

// PostProjectProjectIdClassificationTask creates a new classification task
// for a specific project.
func (a API) PostProjectProjectIdClassificationTask(
	ctx context.Context,
	params openapi.PostProjectProjectIdClassificationTaskRequestObject,
) (openapi.PostProjectProjectIdClassificationTaskResponseObject, error) {
	var pendingClassificationTask storage.PendingClassificationTask

	if params.JSONBody != nil {
		embedding, err := params.JSONBody.Embedding.Bytes()
		if err != nil {
			return nil, fmt.Errorf(
				"api: error reading embedding: %w",
				err,
			)
		}

		pendingClassificationTask = storage.PendingClassificationTask{
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

		pendingClassificationTask = storage.PendingClassificationTask{
			ProjectID: params.ProjectId,
			LLMInput:  params.FormdataBody.LlmInput,
			LLMOutput: params.FormdataBody.LlmOutput,
			Embedding: embedding,
		}
	}

	taskID, err := a.db.CreatePendingClassificationTask(ctx, pendingClassificationTask)

	if errors.Is(err, storage.ErrExistingResource) {
		return openapi.PostProjectProjectIdClassificationTask409Response{
			Headers: openapi.ConflictResponseHeaders{
				Location: fmt.Sprintf("/project/%d/task/%d", params.ProjectId, taskID),
			},
		}, nil
	}

	if err != nil {
		return nil, fmt.Errorf(
			"api: error creating classification task: %w",
			err,
		)
	}

	if redirect.Should(ctx) {
		return redirect.To(fmt.Sprintf("/project/%d/task/%d", params.ProjectId, taskID))
	}

	task, err := a.db.GetPendingClassificationTask(ctx, taskID)
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
	pendingTask, err := a.db.GetPendingClassificationTask(ctx, request.Id)
	if errors.Is(err, sql.ErrNoRows) {
		return openapi.PostProjectProjectIdClassificationTaskIdLabel404Response{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf(
			"api: error getting classification task: %w",
			err,
		)
	}

	classificationTask := storage.ClassificationTask{
		ProjectID: pendingTask.ProjectID,
		LLMInput:  pendingTask.LLMInput,
		LLMOutput: pendingTask.LLMOutput,
		Embedding: pendingTask.Embedding,
		LabelID:   request.Body.Label,
	}
	if err != nil {
		return nil, fmt.Errorf(
			"api: error updating classification task: %w",
			err,
		)
	}

	err = a.db.DeletePendingClassificationTask(ctx, request.Id)
	if err != nil {
		return nil, fmt.Errorf(
			"api: error deleting pending classification task: %w",
			err,
		)
	}

	if redirect.Should(ctx) {
		headers := headers.Get(ctx)
		referer, err := url.Parse(headers.Get("Referer"))

		if err != nil {
			return nil, fmt.Errorf(
				"api: error parsing referer header: %w",
				err,
			)
		}

		if referer.Path == fmt.Sprintf("/project/%d/label_batch", request.ProjectId) {
			return redirect.To(fmt.Sprintf("/project/%d/label_batch", request.ProjectId))
		}

		return redirect.To(fmt.Sprintf("/project/%d/task/%d", request.ProjectId, request.Id))
	}

	return openapi.PostProjectProjectIdClassificationTaskIdLabel200JSONResponse{
		CreatedAt: classificationTask.CreatedAt,
		Embedding: classificationTask.Embedding,
		Id:        classificationTask.ID,
		LabelId:   classificationTask.LabelID,
		LlmInput:  classificationTask.LLMInput,
		LlmOutput: classificationTask.LLMOutput,
		ProjectId: classificationTask.ProjectID,
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

// GetProjectProjectIdPendingClassificationTasks gets the pending classification
// tasks for the given project.
func (a *API) GetProjectProjectIdPendingClassificationTasks(
	ctx context.Context,
	request openapi.GetProjectProjectIdPendingClassificationTasksRequestObject,
) (openapi.GetProjectProjectIdPendingClassificationTasksResponseObject, error) {
	var parameters storage.Parameters

	if request.Params.Page != nil {
		parameters.Page = *request.Params.Page
		parameters.PageSize = storage.DefaultPageSize
	}

	if request.Params.PageSize != nil {
		parameters.PageSize = *request.Params.PageSize
	}

	tasks, err := a.db.FindPendingClassificationTasksForProject(ctx, request.ProjectId, parameters)
	if err != nil {
		return nil, fmt.Errorf(
			"api: error getting pending tasks for project %d: %w",
			request.ProjectId,
			err,
		)
	}

	projectTasks := make([]openapi.PendingClassificationTask, len(tasks))

	for i, task := range tasks {
		projectTasks[i] = openapi.PendingClassificationTask{
			CreatedAt: task.CreatedAt,
			Embedding: task.Embedding,
			Id:        task.ID,
			LlmInput:  task.LLMInput,
			LlmOutput: task.LLMOutput,
			ProjectId: task.ProjectID,
		}
	}

	pendingTaskCount, err := a.db.FindPendingClassificationTaskCountForProject(ctx, request.ProjectId)
	if err != nil {
		return nil, fmt.Errorf(
			"api: error getting pending task count for project %d: %w",
			request.ProjectId,
			err,
		)
	}

	return openapi.GetProjectProjectIdPendingClassificationTasks200JSONResponse{
		Data:  projectTasks,
		Total: pendingTaskCount,
	}, nil
}
