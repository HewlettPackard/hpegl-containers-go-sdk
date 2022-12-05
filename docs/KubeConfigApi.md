# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClustersIdKubeconfigGet**](KubeConfigApi.md#V1ClustersIdKubeconfigGet) | **Get** /v1/clusters/{id}/kubeconfig | Gets a kubeconfig for Cluster

# **V1ClustersIdKubeconfigGet**
> Kubeconfig V1ClustersIdKubeconfigGet(ctx, id)
Gets a kubeconfig for Cluster

Retrieve kubeconfig for specified cluster  **Required Permissions to access the API**:    - caas.cluster.manage  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 

### Return type

[**Kubeconfig**](Kubeconfig.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

