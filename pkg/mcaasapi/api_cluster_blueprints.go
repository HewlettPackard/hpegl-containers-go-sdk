
/*
 * HPE GreenLake for Private Cloud Containers APIs
 *
 * # Table Of Contents - [Introduction](#introduction) - [Authentication](#authentication) - [Get a Space ID](#get-a-space-id)  # Introduction Welcome to the HPE GreenLake for Private Cloud Containers APIs.  This document describes the HPE GreenLake for Private Cloud – Containers APIs protocol and the available endpoints.  The HPE GreenLake for Private Cloud Containers APIs is built on HTTP. Our API is RESTful. It has predictable resource URLs. It returns HTTP response codes to indicate errors. It also accepts and returns JSON in the HTTP body. You can use your favorite HTTP/REST library for your programming language to use HPE GreenLake for Private Cloud Containers APIs.  # Authentication  To authenticate with the HPE GreenLake for Private Cloud Containers APIs, you need to obtain the access token using any of the following way.  <h2>Option1: API client(recommended)</h2>  An API client allows nonhuman entities (an application service account, for instance) programmatic access to a resource on a space.    + 1. <a     href=\"https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=greenlakecentral-create-api-client.html\"     target=\"_blank\">     Create an API client     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>  **Note:** Make sure to save the **Issuer Url**, **Client ID**, and **Client Secret**    + 2. <a     href=\"https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=GUID-1CEA233B-C4B0-41B7-9A25-7A36D9FC0312.html\"     target=\"_blank\">     Creating an assignment for an API client     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>  **Note:** Make sure to assign `Private Cloud Cluster Owner`(cluster-admin) or `Private Cloud Cluster Resource Contributor`(developer)  role.    + 3. Make a REST call to generate the API access token:  ``` curl -i -X POST \\   '{issuerUrl}/v1/token' \\   -H 'Content-Type: application/x-www-form-urlencoded' \\   -d 'client_id={clientID}' \\   -d 'client_secret={clientSecret}' \\   -d grant_type=client_credentials \\   -d scope=hpe-tenant ``` Obtain the `access_token` from the response.  <h2>Option2: Getting the token directly from UI</h2>  To authenticate with the HPE GreenLake for private cloud enterprise VMaaS API you need to obtain the access token from <a     href=\"https://client.greenlake.hpe.com\"     target=\"_blank\">     HPE GreenLake Central     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>. Once logged in to **HPE GreenLake Central**, Click the <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/user.svg\"> icon in the top right corner and then click <input     type=\"button\"     value=\"API Access\"     style=\"color: #000;         background-color: #fff;         border: 2px solid #00b388;         border-radius: 4px;\"/> . In the API Access page, click the <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/clipboard.svg\"> icon to copy the personal access token.  <br>  > <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/circle-information.svg\"> **INFO** For both the options, the access token lease expires in **15** minutes.  # Get a Space ID  You will be needing space id for the space on which you have access. You can use UI to get space ID detail. ![Alt text](./space.png?raw=true \"Space\")  # Roles and Permissions  HPE GreenLake for Private Cloud – Containers provides two default roles for customers. **Private Cloud Cluster Owner** and **Private Cloud Cluster Resource Contributor**. The **Private Cloud Cluster Owner** role is for a cluster administrator role with who gets all the required permissions to execute all the APIs documented in this guide. The **Private Cloud Cluster Resource Contributor** role is for developers with only selective permissions, hence it can execute only a selected list of APIs.  Each API has details about \"Required Permissions to access the API\" and \"Default Roles which can access the API\"  # Getting Started Guide  Our getting started guide will demonstrate a machineblueprint creation, a clusterblueprint creation and then a cluster creation using that clusterblueprint. Here are the set of APIs that you will likely execute:  + Sites     + Gets all Sites + Machine Providers     + Gets all Machine Providers + Machine Blueprints     + Create a Machine Blueprint + Cluster Providers     + Gets all Cluster Providers + Cluster Blueprints     + Creates a Cluster Blueprint + Clusters     + Create a Cluster     + Get a specific Cluster  ## Machine Blueprint Creation Example  ### Gets all Sites  Note the `applianceID` for the site on which you want to create cluster.  ### Gets all Machine Providers  Provide the applianceID in the URL path. Note the `name`, `workerType`, one of the `osImages`, one of the `computeInstanceTypes` and one of the `storageInstanceTypes`  ### Create a Machine Blueprint  Create Machine Blueprint add worker as machineRoles and pass the values noted above in request body. Note the `id` and `name` of the machine blueprint.  ## Cluster Blueprint Creation Example  ### Gets all Cluster Providers  Provide the applianceID in the URL path. Note the  `name`, one of the `kubernetesVersions` and one of the `storageClasses`  ### Creates a Cluster Blueprint  Create Cluster Blueprint pass the values noted in request body and note the `id` from response body  ## Cluster Creation Example  Create cluster will return the newly created cluster details. Note down the cluster ID and make the `Get a specific Cluster` API call to check the cluster provisioning state. Once the state turns ready, the cluster is ready for use!
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcaasapi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ClusterBlueprintsApiService service
/*
ClusterBlueprintsApiService Gets all Cluster Blueprints
Retrieves all cluster blueprints available for the current tenant  **Required Permissions to access the API**:    - caas.cluster.create  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param field field for all query parameters
 * @param optional nil or *ClusterBlueprintsApiV1ClusterblueprintsGetOpts - Optional Parameters:
     * @param "ApplianceID" (optional.Interface of string) -  applianceID
@return ClusterBlueprints
*/

type ClusterBlueprintsApiV1ClusterblueprintsGetOpts struct {
    ApplianceID optional.Interface
}

func (a *ClusterBlueprintsApiService) V1ClusterblueprintsGet(ctx context.Context, field string, localVarOptionals *ClusterBlueprintsApiV1ClusterblueprintsGetOpts) (ClusterBlueprints, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ClusterBlueprints
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/clusterblueprints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ApplianceID.IsSet() {
		localVarQueryParams.Add("applianceID", parameterToString(localVarOptionals.ApplianceID.Value(), ""))
	}
	localVarQueryParams.Add("field", parameterToString(field, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ClusterBlueprints
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v InternalError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ClusterBlueprintsApiService Delete a Cluster Blueprint
Delete the specified blueprint for the current tenant  **Required Permissions to access the API**:    - caas.cluster.create  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id cluster blueprint id

*/
func (a *ClusterBlueprintsApiService) V1ClusterblueprintsIdDelete(ctx context.Context, id string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/clusterblueprints/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ForbiddenError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v NotFoundError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v InternalError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarHttpResponse, newErr
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
ClusterBlueprintsApiService Gets a specific Cluster Blueprint
Retrieve the specified cluster blueprint for the current tenant  **Required Permissions to access the API**:    - caas.cluster.create  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id cluster blueprint id
 * @param field field for all query parameters
 * @param optional nil or *ClusterBlueprintsApiV1ClusterblueprintsIdGetOpts - Optional Parameters:
     * @param "SpaceID" (optional.Interface of string) -  Space filter
@return ClusterBlueprint
*/

type ClusterBlueprintsApiV1ClusterblueprintsIdGetOpts struct {
    SpaceID optional.Interface
}

func (a *ClusterBlueprintsApiService) V1ClusterblueprintsIdGet(ctx context.Context, id string, field string, localVarOptionals *ClusterBlueprintsApiV1ClusterblueprintsIdGetOpts) (ClusterBlueprint, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ClusterBlueprint
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/clusterblueprints/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.SpaceID.IsSet() {
		localVarQueryParams.Add("spaceID", parameterToString(localVarOptionals.SpaceID.Value(), ""))
	}
	localVarQueryParams.Add("field", parameterToString(field, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ClusterBlueprint
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v NotFoundError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v InternalError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ClusterBlueprintsApiService Creates a Cluster Blueprint
Creates a new cluster blueprint for the current tenant  **Required Permissions to access the API**:    - caas.cluster.create  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body cluster blueprint to create
@return ClusterBlueprint
*/
func (a *ClusterBlueprintsApiService) V1ClusterblueprintsPost(ctx context.Context, body ClusterBlueprint) (ClusterBlueprint, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ClusterBlueprint
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/clusterblueprints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v ClusterBlueprint
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v BadRequestError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 422 {
			var v UnprocessingEntityError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v InternalError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
