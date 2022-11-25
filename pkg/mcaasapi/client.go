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
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the HPE GreenLake for Private Cloud Containers APIs API vv1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	ClusterBlueprintsApi *ClusterBlueprintsApiService

	ClusterProvidersApi *ClusterProvidersApiService

	ClustersApi *ClustersApiService

	KubeConfigApi *KubeConfigApiService

	MachineBlueprintsApi *MachineBlueprintsApiService

	MachineProvidersApi *MachineProvidersApiService

	NamespacesApi *NamespacesApiService

	SitesApi *SitesApiService

	StorageClassesApi *StorageClassesApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.ClusterBlueprintsApi = (*ClusterBlueprintsApiService)(&c.common)
	c.ClusterProvidersApi = (*ClusterProvidersApiService)(&c.common)
	c.ClustersApi = (*ClustersApiService)(&c.common)
	c.KubeConfigApi = (*KubeConfigApiService)(&c.common)
	c.MachineBlueprintsApi = (*MachineBlueprintsApiService)(&c.common)
	c.MachineProvidersApi = (*MachineProvidersApiService)(&c.common)
	c.NamespacesApi = (*NamespacesApiService)(&c.common)
	c.SitesApi = (*SitesApiService)(&c.common)
	c.StorageClassesApi = (*StorageClassesApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
		if strings.Contains(contentType, "application/xml") {
			if err = xml.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		} else if strings.Contains(contentType, "application/json") {
			if err = json.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericSwaggerError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericSwaggerError) Model() interface{} {
	return e.model
}
