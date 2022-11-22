# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AppliancesGet**](SitesApi.md#V1AppliancesGet) | **Get** /v1/appliances | Gets all Sites.

# **V1AppliancesGet**
> Appliances V1AppliancesGet(ctx, field)
Gets all Sites.

Retrive all appliances on which user has access 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **field** | **string**| field for all query parameters | 

### Return type

[**Appliances**](Appliances.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

