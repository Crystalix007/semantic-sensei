package api

import (
	"context"
	"fmt"

	"github.com/Crystalix007/semantic-sensei/backend/openapi"
)

// GetOpenapiJSON returns the OpenAPI schema as JSON.
func (a *API) GetOpenapiJson(
	ctx context.Context,
	params openapi.GetOpenapiJsonRequestObject,
) (openapi.GetOpenapiJsonResponseObject, error) {
	schema, err := openapi.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("api: error getting OpenAPI schema: %w", err)
	}

	return openapi.GetOpenapiJson200JSONResponse(*schema), nil
}
