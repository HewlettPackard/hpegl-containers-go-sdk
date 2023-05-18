# Go API client for mcaasapi

# Table Of Contents - [Introduction](#section/Introduction) - [Authentication](#section/Authentication) - [Get a Space ID](#section/Get-a-Space-ID) - [Roles and Permissions](#section/Roles-and-Permissions) - [Getting Started Guide](#section/Getting-Started-Guide)     - [Machine Blueprint Creation Example](#section/Getting-Started-Guide/Machine-Blueprint-Creation-Example)     - [Cluster Blueprint Creation Example](#section/Getting-Started-Guide/Cluster-Blueprint-Creation-Example)      - [Cluster Creation Example](#section/Getting-Started-Guide/Cluster-Creation-Example)  # Introduction Welcome to the APIs for HPE GreenLake for Private Cloud Enterprise: containers.  This document describes the API protocol and available endpoints for the containers service.   The containers service APIs are built on HTTP. Our API is RESTful. It has predictable resource URLs. It returns HTTP response codes to indicate errors. It also accepts and returns JSON in the HTTP body. You can use your favorite HTTP/REST library for your programming language to use the containers service APIs.  # Authentication  To authenticate with the containers service, you must obtain the access token using either of the following options.   <h2>Option 1: Creating an API client (recommended)</h2>  An API client allows nonhuman entities (an application service account, for instance) programmatic access to a resource on a space.    1. <a     href=\"https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=greenlakecentral-create-api-client.html\"     target=\"_blank\">     Create an API client     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>     **Note:** Make sure to save the **Issuer Url**, **Client ID**, and **Client Secret**.    2. <a     href=\"https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=GUID-1CEA233B-C4B0-41B7-9A25-7A36D9FC0312.html\"     target=\"_blank\">     Create an assignment for the API client     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>  **Note:** Make sure to assign `Private Cloud Cluster Owner`(cluster-admin) or `Private Cloud Cluster Resource Contributor`(developer)  role.    3. Make a REST call to generate the API access token, using the issuer URL, client ID, and client secret from step 1:  ``` curl -i -X POST \\   '{issuerUrl}/v1/token' \\   -H 'Content-Type: application/x-www-form-urlencoded' \\   -d 'client_id={clientID}' \\   -d 'client_secret={clientSecret}' \\   -d grant_type=client_credentials \\   -d scope=hpe-tenant ``` Obtain the `access_token` from the response.  <h2>Option 2: Getting the token directly from the UI</h2>  To authenticate with the containers service API,  you can obtain the access token from <a     href=\"https://client.greenlake.hpe.com\"     target=\"_blank\">     HPE GreenLake Central     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>. Log into **HPE GreenLake Central**. Click the <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/user.svg\"> icon in the top right corner and then click **API Access**. In the API Access page, click the <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/clipboard.svg\"> icon to copy the personal access token.  <br>  > <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/circle-information.svg\"> **INFO**: For both the options, the access token lease expires in **15** minutes.  # Get a Space ID  You must obtain the space ID for the space on which you have access. Use the UI to get the space ID.  ![Alt text](./space.png?raw=true \"Space\")  # Roles and Permissions  The containers service provides two default roles for customers: **Private Cloud Cluster Owner** and **Private Cloud Cluster Resource Contributor**. The **Private Cloud Cluster Owner** role is for a cluster administrator who gets all the required permissions to execute all the APIs documented in this guide. The **Private Cloud Cluster Resource Contributor** role is for developers with only selective permissions, so that they can execute only a selected list of APIs.  Each API has details about \"Required Permissions to access the API\" and \"Default Roles which can access the API\"  # Getting Started Guide  Our getting started guide will demonstrate a machine blueprint creation, a cluster blueprint creation and then a cluster creation using that cluster blueprint. Here are the set of APIs that you will likely execute:  + Sites     + Gets all Sites + Machine Providers     + Gets all Machine Providers + Machine Blueprints     + Create a Machine Blueprint + Cluster Providers     + Gets all Cluster Providers + Cluster Blueprints     + Creates a Cluster Blueprint + Clusters     + Create a Cluster     + Get a specific Cluster  ## Machine Blueprint Creation Example  Use the steps in this section to create a machine blueprint that can be used by the cluster blueprint.  ### Step 1 ###  Obtain the 'applianceID' for the site on which you want to create a cluster by running the '[Gets all Sites](#tag/Sites/paths/~1v1~1appliances/get)' API.   Note the `applianceID`.   ### Step 2 ###  Get a list of all machine providers for the specified 'applianceID' by running the '[Gets all Machine Providers](tag/Machine-Providers/paths/~1v1~1appliances~1%7Bid%7D~1machineproviders/get)' API.  Make a note of the `name`, `workerType`, one of the `osImages`, one of the `computeInstanceTypes` and one of the `storageInstanceTypes`.  ### Step 3 ###  Create a machine blueprint by running the '[Create a Machine Blueprint](#tag/Machine-Blueprints/paths/~1v1~1machineblueprints/post)' API. Designate the machine role as \"worker\", and in the request body, provide the values that you noted in the previous step. Make note of the 'id' and 'name' values of the machine blueprint.   ## Cluster Blueprint Creation Example  ### Step 1 ###  List all the cluster providers for the specified 'applianceId' by running the '[Gets all Cluster Providers](l#tag/Cluster-Providers/paths/~1v1~1appliances~1%7Bid%7D~1clusterproviders/get)' API. Make note of the `name`, one of the `kubernetesVersions` and one of the `storageClasses`.   ### Step 2 ###  Create a cluster blueprint by running the [Creates a Cluster Blueprint](#tag/Cluster-Blueprints/paths/~1v1~1clusterblueprints/post) API, and, in the request body, providing the values noted in the previous step. Take note of the 'id' value in the response body.   ## Cluster Creation Example  ### Step 1 ###  Create the cluster by running the '[Creates a Cluster](#tag/Clusters/paths/~1v1~1clusters/post)' API. In the response body, observe the details of the newly created cluster. Make note of the 'clusterID' value.  ### Step 2 ###  Monitor the cluster provisioning state of the specified 'clusterID' by running the [`Get a specific Cluster`](#tag/Clusters/paths/~1v1~1clusters~1%7Bid%7D/get) API.  After the state turns ready, the cluster is ready for use! `` 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: v0.5.7
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
*SitesApi* | [**V1AppliancesGet**](docs/SitesApi.md#v1appliancesget) | **Get** /v1/appliances | Gets all Sites
*StorageClassesApi* | [**V1ClustersIdStorageclassesGet**](docs/StorageClassesApi.md#v1clustersidstorageclassesget) | **Get** /v1/clusters/{id}/storageclasses | Gets all StorageClasses from Cluster

## Documentation For Models

 - [AllOfClusterProviderMinMasterSize](docs/AllOfClusterProviderMinMasterSize.md)
 - [AllOfClusterProviderMinWorkerSize](docs/AllOfClusterProviderMinWorkerSize.md)
 - [AllOfMachineBlueprintSizeDetail](docs/AllOfMachineBlueprintSizeDetail.md)
 - [AllOfMachineSetDetailSizeDetail](docs/AllOfMachineSetDetailSizeDetail.md)
 - [Appliance](docs/Appliance.md)
 - [Appliances](docs/Appliances.md)
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
 - [CreateCluster](docs/CreateCluster.md)
 - [EmptyBody](docs/EmptyBody.md)
 - [Health](docs/Health.md)
 - [Kubeconfig](docs/Kubeconfig.md)
 - [Licenses](docs/Licenses.md)
 - [Machine](docs/Machine.md)
 - [MachineBlueprint](docs/MachineBlueprint.md)
 - [MachineBlueprints](docs/MachineBlueprints.md)
 - [MachineProvider](docs/MachineProvider.md)
 - [MachineProviderName](docs/MachineProviderName.md)
 - [MachineProviders](docs/MachineProviders.md)
 - [MachineRolesType](docs/MachineRolesType.md)
 - [MachineSet](docs/MachineSet.md)
 - [MachineSetDetail](docs/MachineSetDetail.md)
 - [MachineWorkerType](docs/MachineWorkerType.md)
 - [ModelError](docs/ModelError.md)
 - [Namespace](docs/Namespace.md)
 - [Namespaces](docs/Namespaces.md)
 - [OsImages](docs/OsImages.md)
 - [Pagination](docs/Pagination.md)
 - [ResourceQuota](docs/ResourceQuota.md)
 - [ServiceEndpoints](docs/ServiceEndpoints.md)
 - [SizeDetail](docs/SizeDetail.md)
 - [State](docs/State.md)
 - [StorageClass](docs/StorageClass.md)
 - [StorageClasses](docs/StorageClasses.md)
 - [UpdateCluster](docs/UpdateCluster.md)
 - [UpdateClusterMachineSet](docs/UpdateClusterMachineSet.md)
 - [User](docs/User.md)

