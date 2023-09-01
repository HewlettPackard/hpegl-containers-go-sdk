# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AppliancesIdMachineprovidersGet**](MachineProvidersApi.md#V1AppliancesIdMachineprovidersGet) | **Get** /v1/appliances/{id}/machineproviders | Gets all Machine Providers

# **V1AppliancesIdMachineprovidersGet**
> MachineProviders V1AppliancesIdMachineprovidersGet(ctx, id, optional)
Gets all Machine Providers

Retrieve all machine providers for the current appliance  **Required Permissions to access the API**:    - caas.cluster.create  **Default Roles which can access the API**:    - Private Cloud Cluster Owner 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| appliance id | 
 **optional** | ***MachineProvidersApiV1AppliancesIdMachineprovidersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MachineProvidersApiV1AppliancesIdMachineprovidersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **field** | **optional.String**| field for all query parameters. Optional field to query all machine providers within a particular space. | 

### Return type

[**MachineProviders**](MachineProviders.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

