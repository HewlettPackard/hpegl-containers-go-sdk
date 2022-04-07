# {{classname}}

All URIs are relative to *https://mcaas.intg.hpedevops.net/mcaas/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppliancesGet**](SiteApi.md#AppliancesGet) | **Get** /appliances | Retrive all appliances on which user has access

# **AppliancesGet**
> []Appliance AppliancesGet(ctx, spaceID)
Retrive all appliances on which user has access

Retrive all appliances on which user has access 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceID** | [**string**](.md)| space id | 

### Return type

[**[]Appliance**](Appliance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

