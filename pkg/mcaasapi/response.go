/*
 * HPE GreenLake for Private Cloud Containers APIs
 *
 * # Table Of Contents - [Introduction](#section/Introduction) - [Authentication](#section/Authentication) - [Get a Space ID](#section/Get-a-Space-ID) - [Roles and Permissions](#section/Roles-and-Permissions) - [Getting Started Guide](#section/Getting-Started-Guide)     - [Machine Blueprint Creation Example](#section/Getting-Started-Guide/Machine-Blueprint-Creation-Example)     - [Cluster Blueprint Creation Example](#section/Getting-Started-Guide/Cluster-Blueprint-Creation-Example)      - [Cluster Creation Example](#section/Getting-Started-Guide/Cluster-Creation-Example)  # Introduction Welcome to the APIs for HPE GreenLake for Private Cloud Enterprise: containers.  This document describes the API protocol and available endpoints for the containers service.   The containers service APIs are built on HTTP. Our API is RESTful. It has predictable resource URLs. It returns HTTP response codes to indicate errors. It also accepts and returns JSON in the HTTP body. You can use your favorite HTTP/REST library for your programming language to use the containers service APIs.  # Authentication  To authenticate with the containers service, you must obtain the access token using either of the following options.   <h2>Option 1: Creating an API client (recommended)</h2>  An API client allows nonhuman entities (an application service account, for instance) programmatic access to a resource on a space.    1. <a     href=\"https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=greenlakecentral-create-api-client.html\"     target=\"_blank\">     Create an API client     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>     **Note:** Make sure to save the **Issuer Url**, **Client ID**, and **Client Secret**.    2. <a     href=\"https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=GUID-1CEA233B-C4B0-41B7-9A25-7A36D9FC0312.html\"     target=\"_blank\">     Create an assignment for the API client     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>  **Note:** Make sure to assign `Private Cloud Cluster Owner`(cluster-admin) or `Private Cloud Cluster Resource Contributor`(developer)  role.    3. Make a REST call to generate the API access token, using the issuer URL, client ID, and client secret from step 1:  ``` curl -i -X POST \\   '{issuerUrl}/v1/token' \\   -H 'Content-Type: application/x-www-form-urlencoded' \\   -d 'client_id={clientID}' \\   -d 'client_secret={clientSecret}' \\   -d grant_type=client_credentials \\   -d scope=hpe-tenant ``` Obtain the `access_token` from the response.  <h2>Option 2: Getting the token directly from the UI</h2>  To authenticate with the containers service API,  you can obtain the access token from <a     href=\"https://client.greenlake.hpe.com\"     target=\"_blank\">     HPE GreenLake Central     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>. Log into **HPE GreenLake Central**. Click the <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/user.svg\"> icon in the top right corner and then click **API Access**. In the API Access page, click the <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/clipboard.svg\"> icon to copy the personal access token.  <br>  > <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/circle-information.svg\"> **INFO**: For both the options, the access token lease expires in **15** minutes.  # Get a Space ID  You must obtain the space ID for the space on which you have access. Use the UI to get the space ID.  ![Alt text](./space.png?raw=true \"Space\")  # Roles and Permissions  The containers service provides two default roles for customers: **Private Cloud Cluster Owner** and **Private Cloud Cluster Resource Contributor**. The **Private Cloud Cluster Owner** role is for a cluster administrator who gets all the required permissions to execute all the APIs documented in this guide. The **Private Cloud Cluster Resource Contributor** role is for developers with only selective permissions, so that they can execute only a selected list of APIs.  Each API has details about \"Required Permissions to access the API\" and \"Default Roles which can access the API\"  # Getting Started Guide  Our getting started guide will demonstrate a machine blueprint creation, a cluster blueprint creation and then a cluster creation using that cluster blueprint. Here are the set of APIs that you will likely execute:  + Sites     + Gets all Sites + Machine Providers     + Gets all Machine Providers + Machine Blueprints     + Create a Machine Blueprint + Cluster Providers     + Gets all Cluster Providers + Cluster Blueprints     + Creates a Cluster Blueprint + Clusters     + Create a Cluster     + Get a specific Cluster  ## Machine Blueprint Creation Example  Use the steps in this section to create a machine blueprint that can be used by the cluster blueprint.  ### Step 1 ###  Obtain the 'applianceID' for the site on which you want to create a cluster by running the '[Gets all Sites](#tag/Sites/paths/~1v1~1appliances/get)' API.   Note the `applianceID`.   ### Step 2 ###  Get a list of all machine providers for the specified 'applianceID' by running the '[Gets all Machine Providers](tag/Machine-Providers/paths/~1v1~1appliances~1%7Bid%7D~1machineproviders/get)' API.  Make a note of the `name`, `workerType`, one of the `osImages`, one of the `computeInstanceTypes` and one of the `storageInstanceTypes`.  ### Step 3 ###  Create a machine blueprint by running the '[Create a Machine Blueprint](#tag/Machine-Blueprints/paths/~1v1~1machineblueprints/post)' API. Designate the machine role as \"worker\", and in the request body, provide the values that you noted in the previous step. Make note of the 'id' and 'name' values of the machine blueprint.   ## Cluster Blueprint Creation Example  ### Step 1 ###  List all the cluster providers for the specified 'applianceId' by running the '[Gets all Cluster Providers](l#tag/Cluster-Providers/paths/~1v1~1appliances~1%7Bid%7D~1clusterproviders/get)' API. Make note of the `name`, one of the `kubernetesVersions` and one of the `storageClasses`.   ### Step 2 ###  Create a cluster blueprint by running the [Creates a Cluster Blueprint](#tag/Cluster-Blueprints/paths/~1v1~1clusterblueprints/post) API, and, in the request body, providing the values noted in the previous step. Take note of the 'id' value in the response body.   ## Cluster Creation Example  ### Step 1 ###  Create the cluster by running the '[Creates a Cluster](#tag/Clusters/paths/~1v1~1clusters/post)' API. In the response body, observe the details of the newly created cluster. Make note of the 'clusterID' value.  ### Step 2 ###  Monitor the cluster provisioning state of the specified 'clusterID' by running the [`Get a specific Cluster`](#tag/Clusters/paths/~1v1~1clusters~1%7Bid%7D/get) API.  After the state turns ready, the cluster is ready for use! `` 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcaasapi

import (
	"net/http"
)

type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the swagger operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}
