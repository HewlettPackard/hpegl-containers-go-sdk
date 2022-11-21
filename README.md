# Go API client for mcaasapi

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: v0.4.62
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation

### Example Usage

```
    import "github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"

    ...

    apiUrl := "<actual url>"
    accessToken := "<actual token"

    cfg := &mcaasapi.Configuration{
        BasePath:      apiUrl,
        DefaultHeader: make(map[string]string),
        UserAgent:     "hpecli",
    }

    client = mcaasapi.NewAPIClient(cfg)

    ctx = context.WithValue(gocontext.Background(), mcaasapi.ContextAccessToken, accessToken)
    myThing, _, err := client.<XYZ>API.<Function>(ctx)
```

```golang
import "./mcaasapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://mcaas.us1.greenlake-hpe.com/mcaas*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ClusterBlueprintsApi* | [**V1ClusterblueprintsGet**](docs/ClusterBlueprintsApi.md#v1clusterblueprintsget) | **Get** /v1/clusterblueprints | Gets all Cluster Blueprints
*ClusterBlueprintsApi* | [**V1ClusterblueprintsIdDelete**](docs/ClusterBlueprintsApi.md#v1clusterblueprintsiddelete) | **Delete** /v1/clusterblueprints/{id} | Delete a Cluster Blueprint
*ClusterBlueprintsApi* | [**V1ClusterblueprintsIdGet**](docs/ClusterBlueprintsApi.md#v1clusterblueprintsidget) | **Get** /v1/clusterblueprints/{id} | Gets a specific Cluster Blueprint
*ClusterBlueprintsApi* | [**V1ClusterblueprintsPost**](docs/ClusterBlueprintsApi.md#v1clusterblueprintspost) | **Post** /v1/clusterblueprints | Creates a Cluster Blueprint
*ClusterProvidersApi* | [**V1AppliancesIdClusterprovidersGet**](docs/ClusterProvidersApi.md#v1appliancesidclusterprovidersget) | **Get** /v1/appliances/{id}/clusterproviders | Gets all Cluster Providers
*ClustersApi* | [**V1ClustersGet**](docs/ClustersApi.md#v1clustersget) | **Get** /v1/clusters | Gets all Clusters
*ClustersApi* | [**V1ClustersIdDelete**](docs/ClustersApi.md#v1clustersiddelete) | **Delete** /v1/clusters/{id} | Deletes a Cluster
*ClustersApi* | [**V1ClustersIdGet**](docs/ClustersApi.md#v1clustersidget) | **Get** /v1/clusters/{id} | Get a specific Cluster
*ClustersApi* | [**V1ClustersIdPut**](docs/ClustersApi.md#v1clustersidput) | **Put** /v1/clusters/{id} | Updates a Cluster
*ClustersApi* | [**V1ClustersPost**](docs/ClustersApi.md#v1clusterspost) | **Post** /v1/clusters | Create a Cluster
*KubeConfigApi* | [**V1ClustersIdKubeconfigGet**](docs/KubeConfigApi.md#v1clustersidkubeconfigget) | **Get** /v1/clusters/{id}/kubeconfig | Gets a kubeconfig for Cluster
*MachineBlueprintsApi* | [**V1MachineblueprintsGet**](docs/MachineBlueprintsApi.md#v1machineblueprintsget) | **Get** /v1/machineblueprints | Gets all Machine Blueprints
*MachineBlueprintsApi* | [**V1MachineblueprintsIdDelete**](docs/MachineBlueprintsApi.md#v1machineblueprintsiddelete) | **Delete** /v1/machineblueprints/{id} | Deletes a Machine Blueprint
*MachineBlueprintsApi* | [**V1MachineblueprintsIdGet**](docs/MachineBlueprintsApi.md#v1machineblueprintsidget) | **Get** /v1/machineblueprints/{id} | Gets a specific Machine Blueprint
*MachineBlueprintsApi* | [**V1MachineblueprintsPost**](docs/MachineBlueprintsApi.md#v1machineblueprintspost) | **Post** /v1/machineblueprints | Create a Machine Blueprint
*MachineProvidersApi* | [**V1AppliancesIdMachineprovidersGet**](docs/MachineProvidersApi.md#v1appliancesidmachineprovidersget) | **Get** /v1/appliances/{id}/machineproviders | Gets all Machine Providers
*NamespacesApi* | [**V1ClustersIdNamespacesGet**](docs/NamespacesApi.md#v1clustersidnamespacesget) | **Get** /v1/clusters/{id}/namespaces | Gets all Namespaces for Cluster
*OPAPoliciesApi* | [**V1ClustersIdPoliciesGet**](docs/OPAPoliciesApi.md#v1clustersidpoliciesget) | **Get** /v1/clusters/{id}/policies | Gets all policies for the cluster
*OPAPoliciesApi* | [**V1ClustersIdPoliciesPolicyNameGet**](docs/OPAPoliciesApi.md#v1clustersidpoliciespolicynameget) | **Get** /v1/clusters/{id}/policies/{policy_name} | Gets a specific policy details
*SitesApi* | [**V1AppliancesGet**](docs/SitesApi.md#v1appliancesget) | **Get** /v1/appliances | Gets all Sites.
*StorageClassesApi* | [**V1ClustersIdStorageclassesGet**](docs/StorageClassesApi.md#v1clustersidstorageclassesget) | **Get** /v1/clusters/{id}/storageclasses | Gets all StorageClasses from Cluster

## Documentation For Models

 - [AllOfClusterProviderMinMasterSize](docs/AllOfClusterProviderMinMasterSize.md)
 - [AllOfClusterProviderMinWorkerSize](docs/AllOfClusterProviderMinWorkerSize.md)
 - [AllOfMachineBlueprintSizeDetail](docs/AllOfMachineBlueprintSizeDetail.md)
 - [AllOfMachineSetDetailSizeDetail](docs/AllOfMachineSetDetailSizeDetail.md)
 - [Appliance](docs/Appliance.md)
 - [Appliances](docs/Appliances.md)
 - [AuthenticationError](docs/AuthenticationError.md)
 - [BadRequestError](docs/BadRequestError.md)
 - [Cluster](docs/Cluster.md)
 - [ClusterBlueprint](docs/ClusterBlueprint.md)
 - [ClusterBlueprints](docs/ClusterBlueprints.md)
 - [ClusterProvider](docs/ClusterProvider.md)
 - [ClusterProviderAvailableCapacity](docs/ClusterProviderAvailableCapacity.md)
 - [ClusterProviderData](docs/ClusterProviderData.md)
 - [ClusterProviderLicenseInfo](docs/ClusterProviderLicenseInfo.md)
 - [ClusterProviders](docs/ClusterProviders.md)
 - [ClusterUtilization](docs/ClusterUtilization.md)
 - [Clusters](docs/Clusters.md)
 - [Common](docs/Common.md)
 - [ComputeInstanceTypes](docs/ComputeInstanceTypes.md)
 - [ConflictError](docs/ConflictError.md)
 - [CreateCluster](docs/CreateCluster.md)
 - [EmptyBody](docs/EmptyBody.md)
 - [ForbiddenError](docs/ForbiddenError.md)
 - [Health](docs/Health.md)
 - [InternalError](docs/InternalError.md)
 - [Kubeconfig](docs/Kubeconfig.md)
 - [Licenses](docs/Licenses.md)
 - [Machine](docs/Machine.md)
 - [MachineBlueprint](docs/MachineBlueprint.md)
 - [MachineBlueprints](docs/MachineBlueprints.md)
 - [MachineProvider](docs/MachineProvider.md)
 - [MachineProviders](docs/MachineProviders.md)
 - [MachineRolesType](docs/MachineRolesType.md)
 - [MachineSet](docs/MachineSet.md)
 - [MachineSetDetail](docs/MachineSetDetail.md)
 - [MachineWorkerType](docs/MachineWorkerType.md)
 - [Namespace](docs/Namespace.md)
 - [Namespaces](docs/Namespaces.md)
 - [NotFoundError](docs/NotFoundError.md)
 - [OsImages](docs/OsImages.md)
 - [Pagination](docs/Pagination.md)
 - [Policies](docs/Policies.md)
 - [Policy](docs/Policy.md)
 - [ResourceQuota](docs/ResourceQuota.md)
 - [ServiceEndpoints](docs/ServiceEndpoints.md)
 - [SizeDetail](docs/SizeDetail.md)
 - [State](docs/State.md)
 - [StorageClass](docs/StorageClass.md)
 - [StorageClasses](docs/StorageClasses.md)
 - [UnprocessingEntityError](docs/UnprocessingEntityError.md)
 - [UpdateCluster](docs/UpdateCluster.md)
 - [Violation](docs/Violation.md)

