# {{classname}}

All URIs are relative to *https://mcaas.intg.hpedevops.net/mcaas/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppliancesIdClusterprovidersGet**](ClusterAdminApi.md#AppliancesIdClusterprovidersGet) | **Get** /appliances/{id}/clusterproviders | Retrive all cluster providers for the current appliance
[**ClusterblueprintsGet**](ClusterAdminApi.md#ClusterblueprintsGet) | **Get** /clusterblueprints | Retrieve all cluster blueprints
[**ClusterblueprintsIdGet**](ClusterAdminApi.md#ClusterblueprintsIdGet) | **Get** /clusterblueprints/{id} | Retrieves an existing cluster blueprint
[**ClustersGet**](ClusterAdminApi.md#ClustersGet) | **Get** /clusters | Retrieve all clusters currently created
[**ClustersIdDelete**](ClusterAdminApi.md#ClustersIdDelete) | **Delete** /clusters/{id} | Delete a cluster for the current tenant
[**ClustersIdGet**](ClusterAdminApi.md#ClustersIdGet) | **Get** /clusters/{id} | Retrieves an existing cluster for the current user space in a tenant
[**ClustersIdMachineimageversionsGet**](ClusterAdminApi.md#ClustersIdMachineimageversionsGet) | **Get** /clusters/{id}/machineimageversions | Retrieves available machine image version information of an existing cluster for the current user space in a tenant
[**ClustersIdNamespacesGet**](ClusterAdminApi.md#ClustersIdNamespacesGet) | **Get** /clusters/{id}/namespaces | Retrieves namespaces from specified cluster
[**ClustersIdPut**](ClusterAdminApi.md#ClustersIdPut) | **Put** /clusters/{id} | Update a cluster
[**ClustersIdStorageclassesGet**](ClusterAdminApi.md#ClustersIdStorageclassesGet) | **Get** /clusters/{id}/storageclasses | Retrieves storage class information of an existing cluster for the current user space in a tenant
[**ClustersPost**](ClusterAdminApi.md#ClustersPost) | **Post** /clusters | Create a new cluster
[**ConsumptionGet**](ClusterAdminApi.md#ConsumptionGet) | **Get** /consumption | Retrieve all consumption details for the given space
[**MachineblueprintsGet**](ClusterAdminApi.md#MachineblueprintsGet) | **Get** /machineblueprints | Retrieve all machine blueprints
[**MachineblueprintsIdGet**](ClusterAdminApi.md#MachineblueprintsIdGet) | **Get** /machineblueprints/{id} | Retrieves an existing machine blueprint

# **AppliancesIdClusterprovidersGet**
> []ClusterProvider AppliancesIdClusterprovidersGet(ctx, id)
Retrive all cluster providers for the current appliance

Retrive all cluster providers for the current appliance 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| appliance id | 

### Return type

[**[]ClusterProvider**](ClusterProvider.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterblueprintsGet**
> []ClusterBlueprint ClusterblueprintsGet(ctx, spaceID)
Retrieve all cluster blueprints

Retrieves all cluster blueprints available for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceID** | [**string**](.md)| space id | 

### Return type

[**[]ClusterBlueprint**](ClusterBlueprint.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterblueprintsIdGet**
> ClusterBlueprint ClusterblueprintsIdGet(ctx, id, spaceID)
Retrieves an existing cluster blueprint

Retrieve the specified cluster blueprint for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster blueprint id | 
  **spaceID** | [**string**](.md)| Space filter | 

### Return type

[**ClusterBlueprint**](ClusterBlueprint.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **ClustersIdDelete**
> ClustersIdDelete(ctx, id)
Delete a cluster for the current tenant

Delete the specified cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 

### Return type

 (empty response body)

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

# **ClustersIdPut**
> Cluster ClustersIdPut(ctx, body, id)
Update a cluster

Update the specified cluster for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Cluster**](Cluster.md)| cluster to update | 
  **id** | [**string**](.md)| cluster id | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
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

# **ClustersPost**
> Cluster ClustersPost(ctx, body)
Create a new cluster

Creates a new cluster based on the specified cluster blueprint for the current tenant

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

# **MachineblueprintsGet**
> []MachineBlueprint MachineblueprintsGet(ctx, spaceID)
Retrieve all machine blueprints

Retrieves all machine blueprints available for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceID** | [**string**](.md)| space id | 

### Return type

[**[]MachineBlueprint**](MachineBlueprint.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MachineblueprintsIdGet**
> MachineBlueprint MachineblueprintsIdGet(ctx, id, spaceID)
Retrieves an existing machine blueprint

Retrieve the specified machine blueprint for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| machine blueprint id | 
  **spaceID** | [**string**](.md)| Space filter | 

### Return type

[**MachineBlueprint**](MachineBlueprint.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

