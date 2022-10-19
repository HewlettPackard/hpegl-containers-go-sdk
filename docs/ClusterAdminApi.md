# {{classname}}

All URIs are relative to *https://mcaas.intg.hpedevops.net/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AppliancesIdClusterprovidersGet**](ClusterAdminApi.md#V1AppliancesIdClusterprovidersGet) | **Get** /v1/appliances/{id}/clusterproviders | Retrive all cluster providers for the current appliance
[**V1AppliancesIdMachineprovidersGet**](ClusterAdminApi.md#V1AppliancesIdMachineprovidersGet) | **Get** /v1/appliances/{id}/machineproviders | Retrive all machine providers for the current appliance
[**V1ClusterblueprintsGet**](ClusterAdminApi.md#V1ClusterblueprintsGet) | **Get** /v1/clusterblueprints | Retrieve all cluster blueprints
[**V1ClusterblueprintsIdDelete**](ClusterAdminApi.md#V1ClusterblueprintsIdDelete) | **Delete** /v1/clusterblueprints/{id} | Delete a cluster blueprint
[**V1ClusterblueprintsIdGet**](ClusterAdminApi.md#V1ClusterblueprintsIdGet) | **Get** /v1/clusterblueprints/{id} | Retrieves an existing cluster blueprint
[**V1ClusterblueprintsPost**](ClusterAdminApi.md#V1ClusterblueprintsPost) | **Post** /v1/clusterblueprints | Create a new cluster blueprint
[**V1ClustersGet**](ClusterAdminApi.md#V1ClustersGet) | **Get** /v1/clusters | Retrieve all clusters currently created
[**V1ClustersIdDelete**](ClusterAdminApi.md#V1ClustersIdDelete) | **Delete** /v1/clusters/{id} | Delete a cluster for the current tenant
[**V1ClustersIdGet**](ClusterAdminApi.md#V1ClustersIdGet) | **Get** /v1/clusters/{id} | Retrieves an existing cluster for the current user space in a tenant
[**V1ClustersIdKubeconfigGet**](ClusterAdminApi.md#V1ClustersIdKubeconfigGet) | **Get** /v1/clusters/{id}/kubeconfig | Retrieves kubeconfig for specified cluster
[**V1ClustersIdNamespacesGet**](ClusterAdminApi.md#V1ClustersIdNamespacesGet) | **Get** /v1/clusters/{id}/namespaces | Retrieves namespaces from specified cluster
[**V1ClustersIdPut**](ClusterAdminApi.md#V1ClustersIdPut) | **Put** /v1/clusters/{id} | Update a cluster
[**V1ClustersIdStorageclassesGet**](ClusterAdminApi.md#V1ClustersIdStorageclassesGet) | **Get** /v1/clusters/{id}/storageclasses | Retrieves storage class information of an existing cluster for the current user space in a tenant
[**V1ClustersPost**](ClusterAdminApi.md#V1ClustersPost) | **Post** /v1/clusters | Create a new cluster
[**V1MachineblueprintsGet**](ClusterAdminApi.md#V1MachineblueprintsGet) | **Get** /v1/machineblueprints | Retrieve all machine blueprints
[**V1MachineblueprintsIdDelete**](ClusterAdminApi.md#V1MachineblueprintsIdDelete) | **Delete** /v1/machineblueprints/{id} | Delete a machine blueprint
[**V1MachineblueprintsIdGet**](ClusterAdminApi.md#V1MachineblueprintsIdGet) | **Get** /v1/machineblueprints/{id} | Retrieves an existing machine blueprint
[**V1MachineblueprintsPost**](ClusterAdminApi.md#V1MachineblueprintsPost) | **Post** /v1/machineblueprints | Create a new machine blueprint

# **V1AppliancesIdClusterprovidersGet**
> ClusterProviders V1AppliancesIdClusterprovidersGet(ctx, id)
Retrive all cluster providers for the current appliance

Retrive all cluster providers for the current appliance 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| appliance id | 

### Return type

[**ClusterProviders**](ClusterProviders.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppliancesIdMachineprovidersGet**
> MachineProviders V1AppliancesIdMachineprovidersGet(ctx, id)
Retrive all machine providers for the current appliance

Retrive all machine providers for the current appliance 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| appliance id | 

### Return type

[**MachineProviders**](MachineProviders.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClusterblueprintsGet**
> ClusterBlueprints V1ClusterblueprintsGet(ctx, field, optional)
Retrieve all cluster blueprints

Retrieves all cluster blueprints available for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **field** | **string**| field for all query parameters | 
 **optional** | ***ClusterAdminApiV1ClusterblueprintsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterAdminApiV1ClusterblueprintsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applianceID** | [**optional.Interface of string**](.md)| Appliance ID | 

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
Delete a cluster blueprint

Delete the specified blueprint for the current tenant 

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
> ClusterBlueprint V1ClusterblueprintsIdGet(ctx, id, field, optional)
Retrieves an existing cluster blueprint

Retrieve the specified cluster blueprint for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster blueprint id | 
  **field** | **string**| field for all query parameters | 
 **optional** | ***ClusterAdminApiV1ClusterblueprintsIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterAdminApiV1ClusterblueprintsIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **spaceID** | [**optional.Interface of string**](.md)| Space filter | 

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
Create a new cluster blueprint

Creates a new cluster blueprint for the current tenant 

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

# **V1ClustersGet**
> Clusters V1ClustersGet(ctx, field, optional)
Retrieve all clusters currently created

Retrieves all clusters currently created for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **field** | **string**| field for all query parameters | 
 **optional** | ***ClusterAdminApiV1ClustersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterAdminApiV1ClustersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceID** | [**optional.Interface of string**](.md)| space id | 

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
Delete a cluster for the current tenant

Delete the specified cluster 

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
> Cluster V1ClustersIdGet(ctx, id, field, optional)
Retrieves an existing cluster for the current user space in a tenant

Retrieve the specified cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 
  **field** | **string**| field for all query parameters | 
 **optional** | ***ClusterAdminApiV1ClustersIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterAdminApiV1ClustersIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **spaceID** | [**optional.Interface of string**](.md)| Space filter | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersIdKubeconfigGet**
> Kubeconfig V1ClustersIdKubeconfigGet(ctx, id)
Retrieves kubeconfig for specified cluster

Retrieve kubeconfig for specified cluster 

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

# **V1ClustersIdPut**
> Cluster V1ClustersIdPut(ctx, body, id)
Update a cluster

Update the specified cluster for the current tenant 

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

# **V1ClustersPost**
> Cluster V1ClustersPost(ctx, body, optional)
Create a new cluster

Creates a new cluster based on the specified cluster blueprint for the current tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCluster**](CreateCluster.md)| cluster to create with cluster blueprint reference | 
 **optional** | ***ClusterAdminApiV1ClustersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterAdminApiV1ClustersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imported** | **optional.**| set to true if existing cluster has to be imported into GL. | 
 **field** | **optional.**| set \&quot;imported\&quot; to true if existing cluster has to be imported into GL. | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MachineblueprintsGet**
> MachineBlueprints V1MachineblueprintsGet(ctx, field, optional)
Retrieve all machine blueprints

Retrieves all machine blueprints available for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **field** | **string**| field for all query parameters | 
 **optional** | ***ClusterAdminApiV1MachineblueprintsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterAdminApiV1MachineblueprintsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applianceID** | [**optional.Interface of string**](.md)| Appliance ID | 

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
Delete a machine blueprint

Delete the specified machine blueprint for the current tenant 

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
> MachineBlueprint V1MachineblueprintsIdGet(ctx, id, field, optional)
Retrieves an existing machine blueprint

Retrieve the specified machine blueprint for the current tenant 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| machine blueprint id | 
  **field** | **string**| field for all query parameters | 
 **optional** | ***ClusterAdminApiV1MachineblueprintsIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterAdminApiV1MachineblueprintsIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **spaceID** | [**optional.Interface of string**](.md)| Space filter | 

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
Create a new machine blueprint

Creates a new machine blueprint for the current tenant 

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

