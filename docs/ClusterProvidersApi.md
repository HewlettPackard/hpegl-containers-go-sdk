# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AppliancesIdClusterprovidersGet**](ClusterProvidersApi.md#V1AppliancesIdClusterprovidersGet) | **Get** /v1/appliances/{id}/clusterproviders | Gets all Cluster Providers

# **V1AppliancesIdClusterprovidersGet**
> ClusterProviders V1AppliancesIdClusterprovidersGet(ctx, id)
Gets all Cluster Providers

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

