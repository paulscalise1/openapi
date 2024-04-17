/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository

import (
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"

	"context"
	"io/ioutil"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type SDMSubscriptionDocumentApiService service

/*
SDMSubscriptionDocumentApiService Modify an individual sdm subscription
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId - UE id
 * @param SubsId -
 * @param PatchItem -
 * @param SupportedFeatures - Features required to be supported by the target NF

@return ModifysdmSubscriptionResponse
*/

// ModifysdmSubscriptionRequest
type ModifysdmSubscriptionRequest struct {
	UeId              *string
	SubsId            *string
	PatchItem         []models.PatchItem
	SupportedFeatures *string
}

func (r *ModifysdmSubscriptionRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *ModifysdmSubscriptionRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}
func (r *ModifysdmSubscriptionRequest) SetPatchItem(PatchItem []models.PatchItem) {
	r.PatchItem = PatchItem
}
func (r *ModifysdmSubscriptionRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type ModifysdmSubscriptionResponse struct {
	PatchResult models.PatchResult
}

type ModifysdmSubscriptionError struct {
	ProblemDetails models.ProblemDetails
}

func (a *SDMSubscriptionDocumentApiService) ModifysdmSubscription(ctx context.Context, request *ModifysdmSubscriptionRequest) (*ModifysdmSubscriptionResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModifysdmSubscriptionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}
	localVarHTTPContentTypes := []string{"application/json-patch+json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.PatchItem

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return &localVarReturnValue, nil
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.PatchResult, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 403:
		var v ModifysdmSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v ModifysdmSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return nil, apiError
	}
}

/*
SDMSubscriptionDocumentApiService Retrieves a individual sdmSubscription identified by subsId
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId - Unique ID of the subscription to retrieve

@return QuerysdmSubscriptionResponse
*/

// QuerysdmSubscriptionRequest
type QuerysdmSubscriptionRequest struct {
	UeId   *string
	SubsId *string
}

func (r *QuerysdmSubscriptionRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *QuerysdmSubscriptionRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}

type QuerysdmSubscriptionResponse struct {
	QuerysdmSubscriptionResponse200 map[string]interface{}
}

type QuerysdmSubscriptionError struct {
}

func (a *SDMSubscriptionDocumentApiService) QuerysdmSubscription(ctx context.Context, request *QuerysdmSubscriptionRequest) (*QuerysdmSubscriptionResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  QuerysdmSubscriptionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.QuerysdmSubscriptionResponse200, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
SDMSubscriptionDocumentApiService Deletes a sdmsubscriptions
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId - Unique ID of the subscription to remove

@return RemovesdmSubscriptionsResponse
*/

// RemovesdmSubscriptionsRequest
type RemovesdmSubscriptionsRequest struct {
	UeId   *string
	SubsId *string
}

func (r *RemovesdmSubscriptionsRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *RemovesdmSubscriptionsRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}

type RemovesdmSubscriptionsResponse struct {
}

type RemovesdmSubscriptionsError struct {
	ProblemDetails models.ProblemDetails
}

func (a *SDMSubscriptionDocumentApiService) RemovesdmSubscriptions(ctx context.Context, request *RemovesdmSubscriptionsRequest) (*RemovesdmSubscriptionsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RemovesdmSubscriptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return &localVarReturnValue, nil
	case 404:
		var v RemovesdmSubscriptionsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return nil, apiError
	}
}

/*
SDMSubscriptionDocumentApiService Update an individual sdm subscriptions of a UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId -
 * @param SdmSubscription -

@return UpdatesdmsubscriptionsResponse
*/

// UpdatesdmsubscriptionsRequest
type UpdatesdmsubscriptionsRequest struct {
	UeId            *string
	SubsId          *string
	SdmSubscription *models.SdmSubscription
}

func (r *UpdatesdmsubscriptionsRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *UpdatesdmsubscriptionsRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}
func (r *UpdatesdmsubscriptionsRequest) SetSdmSubscription(SdmSubscription models.SdmSubscription) {
	r.SdmSubscription = &SdmSubscription
}

type UpdatesdmsubscriptionsResponse struct {
}

type UpdatesdmsubscriptionsError struct {
	ProblemDetails models.ProblemDetails
}

func (a *SDMSubscriptionDocumentApiService) Updatesdmsubscriptions(ctx context.Context, request *UpdatesdmsubscriptionsRequest) (*UpdatesdmsubscriptionsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UpdatesdmsubscriptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.SdmSubscription

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return &localVarReturnValue, nil
	case 404:
		var v UpdatesdmsubscriptionsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return nil, apiError
	}
}
