# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MachineblueprintsGet**](MachineBlueprintsApi.md#V1MachineblueprintsGet) | **Get** /v1/machineblueprints | Gets all Machine Blueprints
[**V1MachineblueprintsIdDelete**](MachineBlueprintsApi.md#V1MachineblueprintsIdDelete) | **Delete** /v1/machineblueprints/{id} | Deletes a Machine Blueprint
[**V1MachineblueprintsIdGet**](MachineBlueprintsApi.md#V1MachineblueprintsIdGet) | **Get** /v1/machineblueprints/{id} | Gets a specific Machine Blueprint
[**V1MachineblueprintsPost**](MachineBlueprintsApi.md#V1MachineblueprintsPost) | **Post** /v1/machineblueprints | Create a Machine Blueprint

# **V1MachineblueprintsGet**
> MachineBlueprints V1MachineblueprintsGet(ctx, field)
Gets all Machine Blueprints

Retrieves all machine blueprints available for the current tenant  **Required Permissions to access the API**:    - caas.cluster.create  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **field** | **string**| field for all query parameters | 

### Return type

[**MachineBlueprints**](MachineBlueprints.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MachineblueprintsIdDelete**
> V1MachineblueprintsIdDelete(ctx, id)
Deletes a Machine Blueprint

Delete the specified machine blueprint for the current tenant  **Required Permissions to access the API**:    - caas.cluster.create  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| machine blueprint id | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MachineblueprintsIdGet**
> MachineBlueprint V1MachineblueprintsIdGet(ctx, id, field)
Gets a specific Machine Blueprint

Retrieve the specified machine blueprint for the current tenant  **Required Permissions to access the API**:    - caas.cluster.create  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| machine blueprint id | 
  **field** | **string**| field for all query parameters | 

### Return type

[**MachineBlueprint**](MachineBlueprint.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MachineblueprintsPost**
> MachineBlueprint V1MachineblueprintsPost(ctx, body)
Create a Machine Blueprint

Creates a new machine blueprint for the current tenant  **Required Permissions to access the API**:    - caas.cluster.create  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MachineBlueprint**](MachineBlueprint.md)| machine blueprint to create | 

### Return type

[**MachineBlueprint**](MachineBlueprint.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

