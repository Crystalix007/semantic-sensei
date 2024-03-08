package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/Crystalix007/semantic-sensei/backend/api/url"
	"github.com/Crystalix007/semantic-sensei/backend/openapi"
	"github.com/getkin/kin-openapi/openapi3"
)

// ErrUnableToGetURL is returned when the URL cannot be retrieved from the
// context.
var ErrUnableToGetURL = errors.New("unable to get URL from context")

// GetOpenapiJSON returns the OpenAPI schema as JSON.
func (a *API) GetOpenapiJson(
	ctx context.Context,
	params openapi.GetOpenapiJsonRequestObject,
) (openapi.GetOpenapiJsonResponseObject, error) {
	schema, err := openapi.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("api: error getting OpenAPI schema: %w", err)
	}

	url := url.Get(ctx)
	if url == nil {
		return nil, ErrUnableToGetURL
	}

	if host := url.Host; host != "" {
		apiURL := *url
		apiURL.Path = "/api"

		schema.Servers = []*openapi3.Server{
			{
				URL: apiURL.String(),
			},
		}
	}

	return openapi.GetOpenapiJson200JSONResponse(*schema), nil
}
