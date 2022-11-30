# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClustersIdStorageclassesGet**](StorageClassesApi.md#V1ClustersIdStorageclassesGet) | **Get** /v1/clusters/{id}/storageclasses | Gets all StorageClasses from Cluster

# **V1ClustersIdStorageclassesGet**
> StorageClasses V1ClustersIdStorageclassesGet(ctx, id)
Gets all StorageClasses from Cluster

Retrieve the specified storage class information of the cluster  **Required Permissions to access the API**:    - caas.cluster.read  **Default Roles which can access the API**:    - Private Cloud Cluster Owner    - Private Cloud Resource Contributor 

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

