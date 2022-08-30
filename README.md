# Go API client for mcaasapi

APIs and object schemas for MCaaS.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: v0.4.41
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

All URIs are relative to *https://mcaas.intg.hpedevops.net/mcaas*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ClusterAdminApi* | [**V1AppliancesIdClusterprovidersGet**](docs/ClusterAdminApi.md#v1appliancesidclusterprovidersget) | **Get** /v1/appliances/{id}/clusterproviders | Retrive all cluster providers for the current appliance
*ClusterAdminApi* | [**V1AppliancesIdMachineprovidersGet**](docs/ClusterAdminApi.md#v1appliancesidmachineprovidersget) | **Get** /v1/appliances/{id}/machineproviders | Retrive all machine providers for the current appliance
*ClusterAdminApi* | [**V1ClusterblueprintsGet**](docs/ClusterAdminApi.md#v1clusterblueprintsget) | **Get** /v1/clusterblueprints | Retrieve all cluster blueprints
*ClusterAdminApi* | [**V1ClusterblueprintsIdDelete**](docs/ClusterAdminApi.md#v1clusterblueprintsiddelete) | **Delete** /v1/clusterblueprints/{id} | Delete a cluster blueprint
*ClusterAdminApi* | [**V1ClusterblueprintsIdGet**](docs/ClusterAdminApi.md#v1clusterblueprintsidget) | **Get** /v1/clusterblueprints/{id} | Retrieves an existing cluster blueprint
*ClusterAdminApi* | [**V1ClusterblueprintsPost**](docs/ClusterAdminApi.md#v1clusterblueprintspost) | **Post** /v1/clusterblueprints | Create a new cluster blueprint
*ClusterAdminApi* | [**V1ClustersGet**](docs/ClusterAdminApi.md#v1clustersget) | **Get** /v1/clusters | Retrieve all clusters currently created
*ClusterAdminApi* | [**V1ClustersIdDelete**](docs/ClusterAdminApi.md#v1clustersiddelete) | **Delete** /v1/clusters/{id} | Delete a cluster for the current tenant
*ClusterAdminApi* | [**V1ClustersIdGet**](docs/ClusterAdminApi.md#v1clustersidget) | **Get** /v1/clusters/{id} | Retrieves an existing cluster for the current user space in a tenant
*ClusterAdminApi* | [**V1ClustersIdKubeconfigGet**](docs/ClusterAdminApi.md#v1clustersidkubeconfigget) | **Get** /v1/clusters/{id}/kubeconfig | Retrieves kubeconfig for specified cluster
*ClusterAdminApi* | [**V1ClustersIdNamespacesGet**](docs/ClusterAdminApi.md#v1clustersidnamespacesget) | **Get** /v1/clusters/{id}/namespaces | Retrieves namespaces from specified cluster
*ClusterAdminApi* | [**V1ClustersIdPut**](docs/ClusterAdminApi.md#v1clustersidput) | **Put** /v1/clusters/{id} | Update a cluster
*ClusterAdminApi* | [**V1ClustersIdStorageclassesGet**](docs/ClusterAdminApi.md#v1clustersidstorageclassesget) | **Get** /v1/clusters/{id}/storageclasses | Retrieves storage class information of an existing cluster for the current user space in a tenant
*ClusterAdminApi* | [**V1ClustersPost**](docs/ClusterAdminApi.md#v1clusterspost) | **Post** /v1/clusters | Create a new cluster
*ClusterAdminApi* | [**V1MachineblueprintsGet**](docs/ClusterAdminApi.md#v1machineblueprintsget) | **Get** /v1/machineblueprints | Retrieve all machine blueprints
*ClusterAdminApi* | [**V1MachineblueprintsIdDelete**](docs/ClusterAdminApi.md#v1machineblueprintsiddelete) | **Delete** /v1/machineblueprints/{id} | Delete a machine blueprint
*ClusterAdminApi* | [**V1MachineblueprintsIdGet**](docs/ClusterAdminApi.md#v1machineblueprintsidget) | **Get** /v1/machineblueprints/{id} | Retrieves an existing machine blueprint
*ClusterAdminApi* | [**V1MachineblueprintsPost**](docs/ClusterAdminApi.md#v1machineblueprintspost) | **Post** /v1/machineblueprints | Create a new machine blueprint
*DeveloperApi* | [**V1ClustersGet**](docs/DeveloperApi.md#v1clustersget) | **Get** /v1/clusters | Retrieve all clusters currently created
*DeveloperApi* | [**V1ClustersIdGet**](docs/DeveloperApi.md#v1clustersidget) | **Get** /v1/clusters/{id} | Retrieves an existing cluster for the current user space in a tenant
*DeveloperApi* | [**V1ClustersIdNamespacesGet**](docs/DeveloperApi.md#v1clustersidnamespacesget) | **Get** /v1/clusters/{id}/namespaces | Retrieves namespaces from specified cluster
*DeveloperApi* | [**V1ClustersIdStorageclassesGet**](docs/DeveloperApi.md#v1clustersidstorageclassesget) | **Get** /v1/clusters/{id}/storageclasses | Retrieves storage class information of an existing cluster for the current user space in a tenant
*SiteApi* | [**V1AppliancesGet**](docs/SiteApi.md#v1appliancesget) | **Get** /v1/appliances | Retrive all appliances on which user has access

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
 - [ClusterProviderLicenseInfo](docs/ClusterProviderLicenseInfo.md)
 - [ClusterProviders](docs/ClusterProviders.md)
 - [Clusters](docs/Clusters.md)
 - [Common](docs/Common.md)
 - [ComputeInstanceTypes](docs/ComputeInstanceTypes.md)
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
 - [MachineSet](docs/MachineSet.md)
 - [MachineSetDetail](docs/MachineSetDetail.md)
 - [Namespace](docs/Namespace.md)
 - [Namespaces](docs/Namespaces.md)
 - [NotFoundError](docs/NotFoundError.md)
 - [OsImages](docs/OsImages.md)
 - [Pagination](docs/Pagination.md)
 - [ServiceEndpoints](docs/ServiceEndpoints.md)
 - [SizeDetail](docs/SizeDetail.md)
 - [State](docs/State.md)
 - [StorageClass](docs/StorageClass.md)
 - [StorageClasses](docs/StorageClasses.md)
 - [UnprocessingEntityError](docs/UnprocessingEntityError.md)
 - [UpdateCluster](docs/UpdateCluster.md)

