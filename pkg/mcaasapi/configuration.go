/*
 * HPE GreenLake for Private Cloud Containers APIs
 *
 * # Table Of Contents - [Introduction](#introduction) - [Authentication](#authentication) - [Get a Space ID](#get-a-space-id)  # Introduction Welcome to the HPE GreenLake for Private Cloud Containers APIs.  This document describes the HPE GreenLake for Private Cloud – Containers APIs protocol and the available endpoints.  The HPE GreenLake for Private Cloud Containers APIs is built on HTTP. Our API is RESTful. It has predictable resource URLs. It returns HTTP response codes to indicate errors. It also accepts and returns JSON in the HTTP body. You can use your favorite HTTP/REST library for your programming language to use HPE GreenLake for Private Cloud Containers APIs.  # Authentication  To authenticate with the HPE GreenLake for Private Cloud Containers APIs, you need to obtain the access token using any of the following way.  <h2>Option1: API client(recommended)</h2>  An API client allows nonhuman entities (an application service account, for instance) programmatic access to a resource on a space.    + 1. <a     href=\"https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=greenlakecentral-create-api-client.html\"     target=\"_blank\">     Create an API client     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>  **Note:** Make sure to save the **Issuer Url**, **Client ID**, and **Client Secret**    + 2. <a     href=\"https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=GUID-1CEA233B-C4B0-41B7-9A25-7A36D9FC0312.html\"     target=\"_blank\">     Creating an assignment for an API client     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>  **Note:** Make sure to assign `Private Cloud Cluster Owner`(cluster-admin) or `Private Cloud Cluster Resource Contributor`(developer)  role.    + 3. Make a REST call to generate the API access token:  ``` curl -i -X POST \\   '{issuerUrl}/v1/token' \\   -H 'Content-Type: application/x-www-form-urlencoded' \\   -d 'client_id={clientID}' \\   -d 'client_secret={clientSecret}' \\   -d grant_type=client_credentials \\   -d scope=hpe-tenant ``` Obtain the `access_token` from the response.  <h2>Option2: Getting the token directly from UI</h2>  To authenticate with the HPE GreenLake for private cloud enterprise VMaaS API you need to obtain the access token from <a     href=\"https://client.greenlake.hpe.com\"     target=\"_blank\">     HPE GreenLake Central     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>. Once logged in to **HPE GreenLake Central**, Click the <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/user.svg\"> icon in the top right corner and then click <input     type=\"button\"     value=\"API Access\"     style=\"color: #000;         background-color: #fff;         border: 2px solid #00b388;         border-radius: 4px;\"/> . In the API Access page, click the <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/clipboard.svg\"> icon to copy the personal access token.  <br>  > <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/circle-information.svg\"> **INFO** For both the options, the access token lease expires in **15** minutes.  # Get a Space ID  You will be needing space id for the space on which you have access. You can use UI to get space ID detail. ![Alt text](./space.png?raw=true \"Space\")  # Roles and Permissions  HPE GreenLake for Private Cloud – Containers provides two default roles for customers. **Private Cloud Cluster Owner** and **Private Cloud Cluster Resource Contributor**. The **Private Cloud Cluster Owner** role is for cluster admin role with all the required permissions to execute all the documented APIs. The **Private Cloud Cluster Resource Contributor** role is for developers with only selective permissions, hence it can execute only a selected list of APIs.  Each API has details about \"Required Permissions to access the API\" and \"Default Roles which can access the API\"  # Getting Started Guide  Our getting started guide will demonstrate a machineblueprint creation, a clusterblueprint creation and then a cluster creation using that clusterblueprint. Here are the set of APIs that you will likely execute:  + Sites     + Gets all Sites + Machine Providers     + Gets all Machine Providers + Machine Blueprints     + Create a Machine Blueprint + Cluster Providers     + Gets all Cluster Providers + Cluster Blueprints     + Creates a Cluster Blueprint + Clusters     + Create a Cluster  ## Machine Blueprint Creation Example  ### Gets all Sites  Note the `applianceID` for the site on which you want to create cluster.  ### Gets all Machine Providers  Provide the applianceID in the URL path. Note the `name`, `workerType`, one of the `osImages`, one of the `computeInstanceTypes` and one of the `storageInstanceTypes`  ### Create a Machine Blueprint  Create Machine Blueprint add worker as machineRoles and pass the values noted above in request body. Note the `id` and `name` of the machine blueprint.  ## Cluster Blueprint Creation Example  ### Gets all Cluster Providers  Provide the applianceID in the URL path. Note the  `name`, one of the `kubernetesVersions` and one of the `storageClasses`  ### Creates a Cluster Blueprint  Create Cluster Blueprint pass the values noted in request body and note the `id` from response body  ## Cluster Creation Example  Create cluster and wait until it becomes `ready`
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcaasapi

import (
	"net/http"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://mcaas.us1.greenlake-hpe.com/mcaas",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/v0.4.65/go",
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
