// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// ClassificationTask defines model for ClassificationTask.
type ClassificationTask struct {
	CreatedAt time.Time `json:"created_at"`
	Embedding []byte    `json:"embedding"`
	Id        int64     `json:"id"`
	LabelId   *int64    `json:"label_id,omitempty"`
	LlmInput  string    `json:"llm_input"`
	LlmOutput string    `json:"llm_output"`
	ProjectId int64     `json:"project_id"`
}

// ClassificationTaskLabel defines model for ClassificationTaskLabel.
type ClassificationTaskLabel struct {
	CreatedAt time.Time `json:"created_at"`
	Id        int64     `json:"id"`
	Label     string    `json:"label"`
	ProjectId int64     `json:"project_id"`
}

// CreateClassificationTask defines model for CreateClassificationTask.
type CreateClassificationTask struct {
	Embedding []byte `json:"embedding"`
	LlmInput  string `json:"llm_input"`
	LlmOutput string `json:"llm_output"`
}

// CreateClassificationTaskLabel defines model for CreateClassificationTaskLabel.
type CreateClassificationTaskLabel struct {
	Label string `json:"label"`
}

// CreateProject defines model for CreateProject.
type CreateProject struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Location defines model for Location.
type Location = string

// Project defines model for Project.
type Project struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
}

// PostProjectFormdataRequestBody defines body for PostProject for application/x-www-form-urlencoded ContentType.
type PostProjectFormdataRequestBody = CreateProject

// PostProjectProjectIdClassificationTaskFormdataRequestBody defines body for PostProjectProjectIdClassificationTask for application/x-www-form-urlencoded ContentType.
type PostProjectProjectIdClassificationTaskFormdataRequestBody = CreateClassificationTask

// PostProjectProjectIdClassificationTaskLabelFormdataRequestBody defines body for PostProjectProjectIdClassificationTaskLabel for application/x-www-form-urlencoded ContentType.
type PostProjectProjectIdClassificationTaskLabelFormdataRequestBody = CreateClassificationTaskLabel

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /project)
	PostProject(w http.ResponseWriter, r *http.Request)

	// (GET /project/{id})
	GetProjectId(w http.ResponseWriter, r *http.Request, id int64)

	// (POST /project/{project_id}/classification_task)
	PostProjectProjectIdClassificationTask(w http.ResponseWriter, r *http.Request, projectId int64)

	// (GET /project/{project_id}/classification_task/{id})
	GetProjectProjectIdClassificationTaskId(w http.ResponseWriter, r *http.Request, projectId int64, id int64)

	// (POST /project/{project_id}/classification_task_label)
	PostProjectProjectIdClassificationTaskLabel(w http.ResponseWriter, r *http.Request, projectId int64)

	// (GET /project/{project_id}/classification_task_label/{id})
	GetProjectProjectIdClassificationTaskLabelId(w http.ResponseWriter, r *http.Request, projectId int64, id int64)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (POST /project)
func (_ Unimplemented) PostProject(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /project/{id})
func (_ Unimplemented) GetProjectId(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /project/{project_id}/classification_task)
func (_ Unimplemented) PostProjectProjectIdClassificationTask(w http.ResponseWriter, r *http.Request, projectId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /project/{project_id}/classification_task/{id})
func (_ Unimplemented) GetProjectProjectIdClassificationTaskId(w http.ResponseWriter, r *http.Request, projectId int64, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /project/{project_id}/classification_task_label)
func (_ Unimplemented) PostProjectProjectIdClassificationTaskLabel(w http.ResponseWriter, r *http.Request, projectId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /project/{project_id}/classification_task_label/{id})
func (_ Unimplemented) GetProjectProjectIdClassificationTaskLabelId(w http.ResponseWriter, r *http.Request, projectId int64, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostProject operation middleware
func (siw *ServerInterfaceWrapper) PostProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostProject(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetProjectId operation middleware
func (siw *ServerInterfaceWrapper) GetProjectId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetProjectId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostProjectProjectIdClassificationTask operation middleware
func (siw *ServerInterfaceWrapper) PostProjectProjectIdClassificationTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "project_id", runtime.ParamLocationPath, chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "project_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostProjectProjectIdClassificationTask(w, r, projectId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetProjectProjectIdClassificationTaskId operation middleware
func (siw *ServerInterfaceWrapper) GetProjectProjectIdClassificationTaskId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "project_id", runtime.ParamLocationPath, chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "project_id", Err: err})
		return
	}

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetProjectProjectIdClassificationTaskId(w, r, projectId, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostProjectProjectIdClassificationTaskLabel operation middleware
func (siw *ServerInterfaceWrapper) PostProjectProjectIdClassificationTaskLabel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "project_id", runtime.ParamLocationPath, chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "project_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostProjectProjectIdClassificationTaskLabel(w, r, projectId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetProjectProjectIdClassificationTaskLabelId operation middleware
func (siw *ServerInterfaceWrapper) GetProjectProjectIdClassificationTaskLabelId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "project_id", runtime.ParamLocationPath, chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "project_id", Err: err})
		return
	}

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetProjectProjectIdClassificationTaskLabelId(w, r, projectId, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/project", wrapper.PostProject)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/project/{id}", wrapper.GetProjectId)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/project/{project_id}/classification_task", wrapper.PostProjectProjectIdClassificationTask)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/project/{project_id}/classification_task/{id}", wrapper.GetProjectProjectIdClassificationTaskId)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/project/{project_id}/classification_task_label", wrapper.PostProjectProjectIdClassificationTaskLabel)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/project/{project_id}/classification_task_label/{id}", wrapper.GetProjectProjectIdClassificationTaskLabelId)
	})

	return r
}

type RedirectResponseHeaders struct {
	Location Location
}
type RedirectResponse struct {
	Headers RedirectResponseHeaders
}

type PostProjectRequestObject struct {
	Body *PostProjectFormdataRequestBody
}

type PostProjectResponseObject interface {
	VisitPostProjectResponse(w http.ResponseWriter) error
}

type PostProject201JSONResponse Project

func (response PostProject201JSONResponse) VisitPostProjectResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostProject303Response = RedirectResponse

func (response PostProject303Response) VisitPostProjectResponse(w http.ResponseWriter) error {
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(303)
	return nil
}

type GetProjectIdRequestObject struct {
	Id int64 `json:"id"`
}

type GetProjectIdResponseObject interface {
	VisitGetProjectIdResponse(w http.ResponseWriter) error
}

type GetProjectId200JSONResponse Project

func (response GetProjectId200JSONResponse) VisitGetProjectIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProjectId404Response struct {
}

func (response GetProjectId404Response) VisitGetProjectIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PostProjectProjectIdClassificationTaskRequestObject struct {
	ProjectId int64 `json:"project_id"`
	Body      *PostProjectProjectIdClassificationTaskFormdataRequestBody
}

type PostProjectProjectIdClassificationTaskResponseObject interface {
	VisitPostProjectProjectIdClassificationTaskResponse(w http.ResponseWriter) error
}

type PostProjectProjectIdClassificationTask201JSONResponse ClassificationTask

func (response PostProjectProjectIdClassificationTask201JSONResponse) VisitPostProjectProjectIdClassificationTaskResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostProjectProjectIdClassificationTask303Response = RedirectResponse

func (response PostProjectProjectIdClassificationTask303Response) VisitPostProjectProjectIdClassificationTaskResponse(w http.ResponseWriter) error {
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(303)
	return nil
}

type GetProjectProjectIdClassificationTaskIdRequestObject struct {
	ProjectId int64 `json:"project_id"`
	Id        int64 `json:"id"`
}

type GetProjectProjectIdClassificationTaskIdResponseObject interface {
	VisitGetProjectProjectIdClassificationTaskIdResponse(w http.ResponseWriter) error
}

type GetProjectProjectIdClassificationTaskId200JSONResponse ClassificationTask

func (response GetProjectProjectIdClassificationTaskId200JSONResponse) VisitGetProjectProjectIdClassificationTaskIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProjectProjectIdClassificationTaskId404Response struct {
}

func (response GetProjectProjectIdClassificationTaskId404Response) VisitGetProjectProjectIdClassificationTaskIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PostProjectProjectIdClassificationTaskLabelRequestObject struct {
	ProjectId int64 `json:"project_id"`
	Body      *PostProjectProjectIdClassificationTaskLabelFormdataRequestBody
}

type PostProjectProjectIdClassificationTaskLabelResponseObject interface {
	VisitPostProjectProjectIdClassificationTaskLabelResponse(w http.ResponseWriter) error
}

type PostProjectProjectIdClassificationTaskLabel201JSONResponse ClassificationTaskLabel

func (response PostProjectProjectIdClassificationTaskLabel201JSONResponse) VisitPostProjectProjectIdClassificationTaskLabelResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostProjectProjectIdClassificationTaskLabel303Response = RedirectResponse

func (response PostProjectProjectIdClassificationTaskLabel303Response) VisitPostProjectProjectIdClassificationTaskLabelResponse(w http.ResponseWriter) error {
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(303)
	return nil
}

type GetProjectProjectIdClassificationTaskLabelIdRequestObject struct {
	ProjectId int64 `json:"project_id"`
	Id        int64 `json:"id"`
}

type GetProjectProjectIdClassificationTaskLabelIdResponseObject interface {
	VisitGetProjectProjectIdClassificationTaskLabelIdResponse(w http.ResponseWriter) error
}

type GetProjectProjectIdClassificationTaskLabelId200JSONResponse ClassificationTaskLabel

func (response GetProjectProjectIdClassificationTaskLabelId200JSONResponse) VisitGetProjectProjectIdClassificationTaskLabelIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProjectProjectIdClassificationTaskLabelId404Response struct {
}

func (response GetProjectProjectIdClassificationTaskLabelId404Response) VisitGetProjectProjectIdClassificationTaskLabelIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /project)
	PostProject(ctx context.Context, request PostProjectRequestObject) (PostProjectResponseObject, error)

	// (GET /project/{id})
	GetProjectId(ctx context.Context, request GetProjectIdRequestObject) (GetProjectIdResponseObject, error)

	// (POST /project/{project_id}/classification_task)
	PostProjectProjectIdClassificationTask(ctx context.Context, request PostProjectProjectIdClassificationTaskRequestObject) (PostProjectProjectIdClassificationTaskResponseObject, error)

	// (GET /project/{project_id}/classification_task/{id})
	GetProjectProjectIdClassificationTaskId(ctx context.Context, request GetProjectProjectIdClassificationTaskIdRequestObject) (GetProjectProjectIdClassificationTaskIdResponseObject, error)

	// (POST /project/{project_id}/classification_task_label)
	PostProjectProjectIdClassificationTaskLabel(ctx context.Context, request PostProjectProjectIdClassificationTaskLabelRequestObject) (PostProjectProjectIdClassificationTaskLabelResponseObject, error)

	// (GET /project/{project_id}/classification_task_label/{id})
	GetProjectProjectIdClassificationTaskLabelId(ctx context.Context, request GetProjectProjectIdClassificationTaskLabelIdRequestObject) (GetProjectProjectIdClassificationTaskLabelIdResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHttpHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHttpMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// PostProject operation middleware
func (sh *strictHandler) PostProject(w http.ResponseWriter, r *http.Request) {
	var request PostProjectRequestObject

	if err := r.ParseForm(); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode formdata: %w", err))
		return
	}
	var body PostProjectFormdataRequestBody
	if err := runtime.BindForm(&body, r.Form, nil, nil); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't bind formdata: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostProject(ctx, request.(PostProjectRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProject")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostProjectResponseObject); ok {
		if err := validResponse.VisitPostProjectResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetProjectId operation middleware
func (sh *strictHandler) GetProjectId(w http.ResponseWriter, r *http.Request, id int64) {
	var request GetProjectIdRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetProjectId(ctx, request.(GetProjectIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProjectId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetProjectIdResponseObject); ok {
		if err := validResponse.VisitGetProjectIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostProjectProjectIdClassificationTask operation middleware
func (sh *strictHandler) PostProjectProjectIdClassificationTask(w http.ResponseWriter, r *http.Request, projectId int64) {
	var request PostProjectProjectIdClassificationTaskRequestObject

	request.ProjectId = projectId

	if err := r.ParseForm(); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode formdata: %w", err))
		return
	}
	var body PostProjectProjectIdClassificationTaskFormdataRequestBody
	if err := runtime.BindForm(&body, r.Form, nil, nil); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't bind formdata: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostProjectProjectIdClassificationTask(ctx, request.(PostProjectProjectIdClassificationTaskRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProjectProjectIdClassificationTask")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostProjectProjectIdClassificationTaskResponseObject); ok {
		if err := validResponse.VisitPostProjectProjectIdClassificationTaskResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetProjectProjectIdClassificationTaskId operation middleware
func (sh *strictHandler) GetProjectProjectIdClassificationTaskId(w http.ResponseWriter, r *http.Request, projectId int64, id int64) {
	var request GetProjectProjectIdClassificationTaskIdRequestObject

	request.ProjectId = projectId
	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetProjectProjectIdClassificationTaskId(ctx, request.(GetProjectProjectIdClassificationTaskIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProjectProjectIdClassificationTaskId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetProjectProjectIdClassificationTaskIdResponseObject); ok {
		if err := validResponse.VisitGetProjectProjectIdClassificationTaskIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostProjectProjectIdClassificationTaskLabel operation middleware
func (sh *strictHandler) PostProjectProjectIdClassificationTaskLabel(w http.ResponseWriter, r *http.Request, projectId int64) {
	var request PostProjectProjectIdClassificationTaskLabelRequestObject

	request.ProjectId = projectId

	if err := r.ParseForm(); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode formdata: %w", err))
		return
	}
	var body PostProjectProjectIdClassificationTaskLabelFormdataRequestBody
	if err := runtime.BindForm(&body, r.Form, nil, nil); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't bind formdata: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostProjectProjectIdClassificationTaskLabel(ctx, request.(PostProjectProjectIdClassificationTaskLabelRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProjectProjectIdClassificationTaskLabel")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostProjectProjectIdClassificationTaskLabelResponseObject); ok {
		if err := validResponse.VisitPostProjectProjectIdClassificationTaskLabelResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetProjectProjectIdClassificationTaskLabelId operation middleware
func (sh *strictHandler) GetProjectProjectIdClassificationTaskLabelId(w http.ResponseWriter, r *http.Request, projectId int64, id int64) {
	var request GetProjectProjectIdClassificationTaskLabelIdRequestObject

	request.ProjectId = projectId
	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetProjectProjectIdClassificationTaskLabelId(ctx, request.(GetProjectProjectIdClassificationTaskLabelIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProjectProjectIdClassificationTaskLabelId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetProjectProjectIdClassificationTaskLabelIdResponseObject); ok {
		if err := validResponse.VisitGetProjectProjectIdClassificationTaskLabelIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
