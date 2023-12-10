/*
CloudTruth Management API

CloudTruth centralizes your configuration parameters and secrets making them easier to manage and use as a team.

API version: v1
Contact: support@cloudtruth.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudtruthapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// UtilsAPIService UtilsAPI service
type UtilsAPIService service

type ApiUtilsGeneratePasswordCreateRequest struct {
	ctx context.Context
	ApiService *UtilsAPIService
	length *int32
	requireHardwareGeneration *bool
	requireLowercase *bool
	requireNumbers *bool
	requireSpaces *bool
	requireSymbols *bool
	requireUppercase *bool
}

// The length of the password to generate.  Minimum of 8, maximum of 4095.
func (r ApiUtilsGeneratePasswordCreateRequest) Length(length int32) ApiUtilsGeneratePasswordCreateRequest {
	r.length = &length
	return r
}

// Default behavior is to fallback to /dev/urandom if we fail to get a random password from AWS Secrets Manager.  If set to &#39;True&#39;, we will not fallback to local password generation using /dev/urandom.  Default: False
func (r ApiUtilsGeneratePasswordCreateRequest) RequireHardwareGeneration(requireHardwareGeneration bool) ApiUtilsGeneratePasswordCreateRequest {
	r.requireHardwareGeneration = &requireHardwareGeneration
	return r
}

// The password must include lowercase letters [a-z]. Default: True.
func (r ApiUtilsGeneratePasswordCreateRequest) RequireLowercase(requireLowercase bool) ApiUtilsGeneratePasswordCreateRequest {
	r.requireLowercase = &requireLowercase
	return r
}

// The password must include numbers [0-9].  Default: True.
func (r ApiUtilsGeneratePasswordCreateRequest) RequireNumbers(requireNumbers bool) ApiUtilsGeneratePasswordCreateRequest {
	r.requireNumbers = &requireNumbers
	return r
}

// The password must include spaces [ ].  Default: False.
func (r ApiUtilsGeneratePasswordCreateRequest) RequireSpaces(requireSpaces bool) ApiUtilsGeneratePasswordCreateRequest {
	r.requireSpaces = &requireSpaces
	return r
}

// The password must include symbols [!\&quot;#$%&amp;&#39;()*+,-./:;&lt;&#x3D;&gt;?@[\\]^_&#x60;{|}~].  Default: False.
func (r ApiUtilsGeneratePasswordCreateRequest) RequireSymbols(requireSymbols bool) ApiUtilsGeneratePasswordCreateRequest {
	r.requireSymbols = &requireSymbols
	return r
}

// The password must include uppercase letters [A-Z].  Default: True.
func (r ApiUtilsGeneratePasswordCreateRequest) RequireUppercase(requireUppercase bool) ApiUtilsGeneratePasswordCreateRequest {
	r.requireUppercase = &requireUppercase
	return r
}

func (r ApiUtilsGeneratePasswordCreateRequest) Execute() (*GeneratedPasswordResponse, *http.Response, error) {
	return r.ApiService.UtilsGeneratePasswordCreateExecute(r)
}

/*
UtilsGeneratePasswordCreate Get a randomly generated password using AWS Secrets Manager, with fallback to /dev/urandom.

Endpoint for accessing utility functions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUtilsGeneratePasswordCreateRequest
*/
func (a *UtilsAPIService) UtilsGeneratePasswordCreate(ctx context.Context) ApiUtilsGeneratePasswordCreateRequest {
	return ApiUtilsGeneratePasswordCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GeneratedPasswordResponse
func (a *UtilsAPIService) UtilsGeneratePasswordCreateExecute(r ApiUtilsGeneratePasswordCreateRequest) (*GeneratedPasswordResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GeneratedPasswordResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UtilsAPIService.UtilsGeneratePasswordCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/utils/generate_password/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.length == nil {
		return localVarReturnValue, nil, reportError("length is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "length", r.length, "")
	if r.requireHardwareGeneration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "require_hardware_generation", r.requireHardwareGeneration, "")
	} else {
		var defaultValue bool = false
		r.requireHardwareGeneration = &defaultValue
	}
	if r.requireLowercase != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "require_lowercase", r.requireLowercase, "")
	} else {
		var defaultValue bool = true
		r.requireLowercase = &defaultValue
	}
	if r.requireNumbers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "require_numbers", r.requireNumbers, "")
	} else {
		var defaultValue bool = true
		r.requireNumbers = &defaultValue
	}
	if r.requireSpaces != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "require_spaces", r.requireSpaces, "")
	} else {
		var defaultValue bool = false
		r.requireSpaces = &defaultValue
	}
	if r.requireSymbols != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "require_symbols", r.requireSymbols, "")
	} else {
		var defaultValue bool = false
		r.requireSymbols = &defaultValue
	}
	if r.requireUppercase != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "require_uppercase", r.requireUppercase, "")
	} else {
		var defaultValue bool = true
		r.requireUppercase = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWTAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
