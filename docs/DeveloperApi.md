# {{classname}}

All URIs are relative to *https://mcaas.intg.hpedevops.net/mcaas/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersGet**](DeveloperApi.md#ClustersGet) | **Get** /clusters | Retrieve all clusters currently created
[**ClustersIdGet**](DeveloperApi.md#ClustersIdGet) | **Get** /clusters/{id} | Retrieves an existing cluster for the current user space in a tenant
[**ClustersIdMachineimageversionsGet**](DeveloperApi.md#ClustersIdMachineimageversionsGet) | **Get** /clusters/{id}/machineimageversions | Retrieves available machine image version information of an existing cluster for the current user space in a tenant
[**ClustersIdNamespacesGet**](DeveloperApi.md#ClustersIdNamespacesGet) | **Get** /clusters/{id}/namespaces | Retrieves namespaces from specified cluster
[**ClustersIdStorageclassesGet**](DeveloperApi.md#ClustersIdStorageclassesGet) | **Get** /clusters/{id}/storageclasses | Retrieves storage class information of an existing cluster for the current user space in a tenant
[**ConsumptionGet**](DeveloperApi.md#ConsumptionGet) | **Get** /consumption | Retrieve all consumption details for the given space

# **ClustersGet**
> []Cluster ClustersGet(ctx, spaceID)
Retrieve all clusters currently created

Retrieves all clusters currently created for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceID** | [**string**](.md)| space id | 

### Return type

[**[]Cluster**](Cluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersIdGet**
> Cluster ClustersIdGet(ctx, id, spaceID)
Retrieves an existing cluster for the current user space in a tenant

Retrieve the specified cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 
  **spaceID** | [**string**](.md)| Space filter | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersIdMachineimageversionsGet**
> MachineImageVersion ClustersIdMachineimageversionsGet(ctx, id)
Retrieves available machine image version information of an existing cluster for the current user space in a tenant

Retrieve the available machine image version information of the cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 

### Return type

[**MachineImageVersion**](MachineImageVersion.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersIdNamespacesGet**
> []Namespace ClustersIdNamespacesGet(ctx, id)
Retrieves namespaces from specified cluster

Retrieve namespaces from specified cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 

### Return type

[**[]Namespace**](Namespace.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersIdStorageclassesGet**
> []StorageClass ClustersIdStorageclassesGet(ctx, id)
Retrieves storage class information of an existing cluster for the current user space in a tenant

Retrieve the specified storage class information of the cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 

### Return type

[**[]StorageClass**](StorageClass.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConsumptionGet**
> []SiteConsumption ConsumptionGet(ctx, spaceID)
Retrieve all consumption details for the given space

Retrieves all consumption details for the given space for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceID** | [**string**](.md)| space id | 

### Return type

[**[]SiteConsumption**](SiteConsumption.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

