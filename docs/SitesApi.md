# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AppliancesGet**](SitesApi.md#V1AppliancesGet) | **Get** /v1/appliances | Gets all Sites

# **V1AppliancesGet**
> Appliances V1AppliancesGet(ctx, field, optional)
Gets all Sites

Retrive all appliances on which user has access  **Required Permissions to access the API**:    Any of:    - caas.cluster.create    - caas.cluster.manage    - caas.cluster-resource.manage\"  **Default Roles which can access the API**:    - Private Cloud Cluster Owner    - Private Cloud Resource Contributor 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **field** | **string**| field for all query parameters | 
 **optional** | ***SitesApiV1AppliancesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesApiV1AppliancesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceID** | [**optional.Interface of string**](.md)| space id | 

### Return type

[**Appliances**](Appliances.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

