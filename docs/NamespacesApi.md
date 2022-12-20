# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClustersIdNamespacesGet**](NamespacesApi.md#V1ClustersIdNamespacesGet) | **Get** /v1/clusters/{id}/namespaces | Gets all Namespaces for Cluster

# **V1ClustersIdNamespacesGet**
> Namespaces V1ClustersIdNamespacesGet(ctx, id)
Gets all Namespaces for Cluster

Retrieve namespaces from specified cluster  **Required Permissions to access the API**:    - caas.cluster.read  **Default Roles which can access the API**:    - Private Cloud Cluster Owner    - Private Cloud Resource Contributor 

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

