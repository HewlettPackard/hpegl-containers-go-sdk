# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClusterblueprintsGet**](ClusterBlueprintsApi.md#V1ClusterblueprintsGet) | **Get** /v1/clusterblueprints | Gets all Cluster Blueprints
[**V1ClusterblueprintsIdDelete**](ClusterBlueprintsApi.md#V1ClusterblueprintsIdDelete) | **Delete** /v1/clusterblueprints/{id} | Delete a Cluster Blueprint
[**V1ClusterblueprintsIdGet**](ClusterBlueprintsApi.md#V1ClusterblueprintsIdGet) | **Get** /v1/clusterblueprints/{id} | Gets a specific Cluster Blueprint
[**V1ClusterblueprintsPost**](ClusterBlueprintsApi.md#V1ClusterblueprintsPost) | **Post** /v1/clusterblueprints | Creates a Cluster Blueprint

# **V1ClusterblueprintsGet**
> ClusterBlueprints V1ClusterblueprintsGet(ctx, field)
Gets all Cluster Blueprints

Retrieves all cluster blueprints available for the site. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **field** | **string**| field for all query parameters | 

### Return type

[**ClusterBlueprints**](ClusterBlueprints.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClusterblueprintsIdDelete**
> V1ClusterblueprintsIdDelete(ctx, id)
Delete a Cluster Blueprint

Delete the specified blueprint from the site. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster blueprint id | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClusterblueprintsIdGet**
> ClusterBlueprint V1ClusterblueprintsIdGet(ctx, id, field)
Gets a specific Cluster Blueprint

Retrieve the specified cluster blueprint for the site. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster blueprint id | 
  **field** | **string**| field for all query parameters | 

### Return type

[**ClusterBlueprint**](ClusterBlueprint.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClusterblueprintsPost**
> ClusterBlueprint V1ClusterblueprintsPost(ctx, body)
Creates a Cluster Blueprint

Creates a new cluster blueprint for the site. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClusterBlueprint**](ClusterBlueprint.md)| cluster blueprint to create | 

### Return type

[**ClusterBlueprint**](ClusterBlueprint.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

