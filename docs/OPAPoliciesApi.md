# {{classname}}

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClustersIdPoliciesGet**](OPAPoliciesApi.md#V1ClustersIdPoliciesGet) | **Get** /v1/clusters/{id}/policies | Gets all policies for the cluster
[**V1ClustersIdPoliciesPolicyNameGet**](OPAPoliciesApi.md#V1ClustersIdPoliciesPolicyNameGet) | **Get** /v1/clusters/{id}/policies/{policy_name} | Gets a specific policy details

# **V1ClustersIdPoliciesGet**
> Policies V1ClustersIdPoliciesGet(ctx, id)
Gets all policies for the cluster

Retrieve all OPA policies for the cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 

### Return type

[**Policies**](Policies.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersIdPoliciesPolicyNameGet**
> Policy V1ClustersIdPoliciesPolicyNameGet(ctx, id, policyName)
Gets a specific policy details

Retrieves an existing OPA policy violating details 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| cluster id | 
  **policyName** | **string**| name of the policy | 

### Return type

[**Policy**](Policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

