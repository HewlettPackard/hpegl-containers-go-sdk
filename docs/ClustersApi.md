# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClustersGet**](ClustersApi.md#V1ClustersGet) | **Get** /v1/clusters | Gets all Clusters
[**V1ClustersIdDelete**](ClustersApi.md#V1ClustersIdDelete) | **Delete** /v1/clusters/{id} | Deletes a Cluster
[**V1ClustersIdGet**](ClustersApi.md#V1ClustersIdGet) | **Get** /v1/clusters/{id} | Get a specific Cluster
[**V1ClustersIdPut**](ClustersApi.md#V1ClustersIdPut) | **Put** /v1/clusters/{id} | Updates a Cluster
[**V1ClustersPost**](ClustersApi.md#V1ClustersPost) | **Post** /v1/clusters | Create a Cluster

# **V1ClustersGet**
> Clusters V1ClustersGet(ctx, field)
Gets all Clusters

Retrieves all clusters currently created for the current tenant.  **Required Permissions to access the API**:    - caas.cluster.read  **Default Roles which can access the API**:    - Private Cloud Cluster Owner    - Private Cloud Resource Contributor 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **field** | **string**| field for all query parameters | 

### Return type

[**Clusters**](Clusters.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersIdDelete**
> EmptyBody V1ClustersIdDelete(ctx, id)
Deletes a Cluster

Delete the specified cluster  **Required Permissions to access the API**:    - caas.cluster.delete  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 

### Return type

[**EmptyBody**](EmptyBody.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersIdGet**
> Cluster V1ClustersIdGet(ctx, id, field)
Get a specific Cluster

Retrieve the specified cluster  **Required Permissions to access the API**:    - caas.cluster.read  **Default Roles which can access the API**:    - Private Cloud Cluster Owner    - Private Cloud Resource Contributor 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 
  **field** | **string**| field for all query parameters | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersIdPut**
> Cluster V1ClustersIdPut(ctx, body, id)
Updates a Cluster

Update the specified cluster for the current tenant  **Required Permissions to access the API**:    - caas.cluster.update  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCluster**](UpdateCluster.md)| cluster to update | 
  **id** | [**string**](.md)| cluster id | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersPost**
> Cluster V1ClustersPost(ctx, body)
Create a Cluster

Creates a new cluster based on the specified cluster blueprint for the current tenant. **Required Permissions to access the API**:    - caas.cluster.create  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCluster**](CreateCluster.md)| cluster to create with cluster blueprint reference | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

