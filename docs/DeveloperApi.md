# {{classname}}

All URIs are relative to *https://mcaas.intg.hpedevops.net/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClustersGet**](DeveloperApi.md#V1ClustersGet) | **Get** /v1/clusters | Retrieve all clusters currently created
[**V1ClustersIdGet**](DeveloperApi.md#V1ClustersIdGet) | **Get** /v1/clusters/{id} | Retrieves an existing cluster for the current user space in a tenant
[**V1ClustersIdNamespacesGet**](DeveloperApi.md#V1ClustersIdNamespacesGet) | **Get** /v1/clusters/{id}/namespaces | Retrieves namespaces from specified cluster
[**V1ClustersIdStorageclassesGet**](DeveloperApi.md#V1ClustersIdStorageclassesGet) | **Get** /v1/clusters/{id}/storageclasses | Retrieves storage class information of an existing cluster for the current user space in a tenant

# **V1ClustersGet**
> Clusters V1ClustersGet(ctx, spaceID, field)
Retrieve all clusters currently created

Retrieves all clusters currently created for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceID** | [**string**](.md)| space id | 
  **field** | **string**| field for all query parameters | 

### Return type

[**Clusters**](Clusters.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersIdGet**
> Cluster V1ClustersIdGet(ctx, id, spaceID, field)
Retrieves an existing cluster for the current user space in a tenant

Retrieve the specified cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 
  **spaceID** | [**string**](.md)| Space filter | 
  **field** | **string**| field for all query parameters | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersIdNamespacesGet**
> Namespaces V1ClustersIdNamespacesGet(ctx, id)
Retrieves namespaces from specified cluster

Retrieve namespaces from specified cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 

### Return type

[**Namespaces**](Namespaces.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersIdStorageclassesGet**
> StorageClasses V1ClustersIdStorageclassesGet(ctx, id)
Retrieves storage class information of an existing cluster for the current user space in a tenant

Retrieve the specified storage class information of the cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 

### Return type

[**StorageClasses**](StorageClasses.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

