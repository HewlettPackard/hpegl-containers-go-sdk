# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AppliancesIdClusterprovidersGet**](ClusterProvidersApi.md#V1AppliancesIdClusterprovidersGet) | **Get** /v1/appliances/{id}/clusterproviders | Gets all Cluster Providers

# **V1AppliancesIdClusterprovidersGet**
> ClusterProviders V1AppliancesIdClusterprovidersGet(ctx, id, optional)
Gets all Cluster Providers

Retrieve all cluster providers for the current appliance  **Required Permissions to access the API**:    - caas.cluster.create  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| appliance id | 
 **optional** | ***ClusterProvidersApiV1AppliancesIdClusterprovidersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterProvidersApiV1AppliancesIdClusterprovidersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **field** | **optional.String**| field for all query parameters. Optional field to query all cluster providers within a particular space. | 

### Return type

[**ClusterProviders**](ClusterProviders.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

