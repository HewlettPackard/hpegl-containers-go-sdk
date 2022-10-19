# {{classname}}

All URIs are relative to *https://mcaas.intg.hpedevops.net/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AppliancesGet**](SiteApi.md#V1AppliancesGet) | **Get** /v1/appliances | Retrive all appliances on which user has access

# **V1AppliancesGet**
> Appliances V1AppliancesGet(ctx, field, optional)
Retrive all appliances on which user has access

Retrive all appliances on which user has access 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **field** | **string**| field for all query parameters | 
 **optional** | ***SiteApiV1AppliancesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SiteApiV1AppliancesGetOpts struct
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

