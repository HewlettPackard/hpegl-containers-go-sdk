/*
 * HPE GreenLake for Private Cloud Containers APIs
 *
 * # Table Of Contents - [Introduction](#section/Introduction) - [Authentication](#section/Authentication) - [Get a Space ID](#section/Get-a-Space-ID) - [Roles and Permissions](#section/Roles-and-Permissions) - [Getting Started Guide](#section/Getting-Started-Guide)     - [Machine Blueprint Creation Example](#section/Getting-Started-Guide/Machine-Blueprint-Creation-Example)     - [Cluster Blueprint Creation Example](#section/Getting-Started-Guide/Cluster-Blueprint-Creation-Example)      - [Cluster Creation Example](#section/Getting-Started-Guide/Cluster-Creation-Example)  # Introduction Welcome to the HPE GreenLake for Private Cloud Containers APIs.  This document describes the HPE GreenLake for Private Cloud – Containers APIs protocol and the available endpoints.  The HPE GreenLake for Private Cloud Containers APIs is built on HTTP. Our API is RESTful. It has predictable resource URLs. It returns HTTP response codes to indicate errors. It also accepts and returns JSON in the HTTP body. You can use your favorite HTTP/REST library for your programming language to use HPE GreenLake for Private Cloud Containers APIs.  # Authentication  To authenticate with the HPE GreenLake for Private Cloud Containers APIs, you need to obtain the access token using any of the following way.  <h2>Option1: API client(recommended)</h2>  An API client allows nonhuman entities (an application service account, for instance) programmatic access to a resource on a space.    + 1. <a     href=\"https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=greenlakecentral-create-api-client.html\"     target=\"_blank\">     Create an API client     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>  **Note:** Make sure to save the **Issuer Url**, **Client ID**, and **Client Secret**    + 2. <a     href=\"https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=GUID-1CEA233B-C4B0-41B7-9A25-7A36D9FC0312.html\"     target=\"_blank\">     Creating an assignment for an API client     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>  **Note:** Make sure to assign `Private Cloud Cluster Owner`(cluster-admin) or `Private Cloud Cluster Resource Contributor`(developer)  role.    + 3. Make a REST call to generate the API access token:  ``` curl -i -X POST \\   '{issuerUrl}/v1/token' \\   -H 'Content-Type: application/x-www-form-urlencoded' \\   -d 'client_id={clientID}' \\   -d 'client_secret={clientSecret}' \\   -d grant_type=client_credentials \\   -d scope=hpe-tenant ``` Obtain the `access_token` from the response.  <h2>Option2: Getting the token directly from UI</h2>  To authenticate with the HPE GreenLake for private cloud enterprise VMaaS API you need to obtain the access token from <a     href=\"https://client.greenlake.hpe.com\"     target=\"_blank\">     HPE GreenLake Central     <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg\"     height=\"12\"> </a>. Once logged in to **HPE GreenLake Central**, Click the <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/user.svg\"> icon in the top right corner and then click <input     type=\"button\"     value=\"API Access\"     style=\"color: #000;         background-color: #fff;         border: 2px solid #00b388;         border-radius: 4px;\"/> . In the API Access page, click the <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/clipboard.svg\"> icon to copy the personal access token.  <br>  > <img src=\"https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/circle-information.svg\"> **INFO** For both the options, the access token lease expires in **15** minutes.  # Get a Space ID  You will be needing space id for the space on which you have access. You can use UI to get space ID detail. ![Alt text](./space.png?raw=true \"Space\")  # Roles and Permissions  HPE GreenLake for Private Cloud – Containers provides two default roles for customers. **Private Cloud Cluster Owner** and **Private Cloud Cluster Resource Contributor**. The **Private Cloud Cluster Owner** role is for a cluster administrator role with who gets all the required permissions to execute all the APIs documented in this guide. The **Private Cloud Cluster Resource Contributor** role is for developers with only selective permissions, hence it can execute only a selected list of APIs.  Each API has details about \"Required Permissions to access the API\" and \"Default Roles which can access the API\"  # Getting Started Guide  Our getting started guide will demonstrate a machineblueprint creation, a clusterblueprint creation and then a cluster creation using that clusterblueprint. Here are the set of APIs that you will likely execute:  + Sites     + Gets all Sites + Machine Providers     + Gets all Machine Providers + Machine Blueprints     + Create a Machine Blueprint + Cluster Providers     + Gets all Cluster Providers + Cluster Blueprints     + Creates a Cluster Blueprint + Clusters     + Create a Cluster     + Get a specific Cluster  ## Machine Blueprint Creation Example  ### [Gets all Sites](#tag/Sites/paths/~1v1~1appliances/get)  Note the `applianceID` for the site on which you want to create cluster.  ### [Gets all Machine Providers](tag/Machine-Providers/paths/~1v1~1appliances~1%7Bid%7D~1machineproviders/get)  Provide the applianceID in the URL path. Note the `name`, `workerType`, one of the `osImages`, one of the `computeInstanceTypes` and one of the `storageInstanceTypes`  ### [Create a Machine Blueprint](#tag/Machine-Blueprints/paths/~1v1~1machineblueprints/post)  Create Machine Blueprint add worker as machineRoles and pass the values noted above in request body. Note the `id` and `name` of the machine blueprint.  ## Cluster Blueprint Creation Example  ### [Gets all Cluster Providers](l#tag/Cluster-Providers/paths/~1v1~1appliances~1%7Bid%7D~1clusterproviders/get)  Provide the applianceID in the URL path. Note the  `name`, one of the `kubernetesVersions` and one of the `storageClasses`  ### [Creates a Cluster Blueprint](#tag/Cluster-Blueprints/paths/~1v1~1clusterblueprints/post)  Create Cluster Blueprint pass the values noted in request body and note the `id` from response body  ## Cluster Creation Example  ### [Creates a Cluster](#tag/Clusters/paths/~1v1~1clusters/post) Create cluster will return the newly created cluster details. Note down the cluster ID and make the [`Get a specific Cluster`](#tag/Clusters/paths/~1v1~1clusters~1%7Bid%7D/get) API call to check the cluster provisioning state. Once the state turns ready, the cluster is ready for use!
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcaasapi
import (
	"time"
)

type Cluster struct {
	State string `json:"state,omitempty"`
	Health string `json:"health,omitempty"`
	UserID string `json:"userID,omitempty"`
	UserDisplayName string `json:"userDisplayName,omitempty"`
	UserName string `json:"userName,omitempty"`
	CreatedDate time.Time `json:"createdDate,omitempty"`
	LastUpdateDate time.Time `json:"lastUpdateDate,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	ClusterBlueprintId string `json:"clusterBlueprintId,omitempty"`
	ClusterProvider string `json:"clusterProvider,omitempty"`
	K8sVersion string `json:"k8sVersion,omitempty"`
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
	MachineSets []MachineSet `json:"machineSets,omitempty"`
	MachineSetsDetail []MachineSetDetail `json:"machineSetsDetail,omitempty"`
	Utilization *ClusterUtilization `json:"utilization,omitempty"`
	ApiEndpoint string `json:"apiEndpoint,omitempty"`
	ServiceEndpoints []ServiceEndpoints `json:"serviceEndpoints,omitempty"`
	ApplianceID string `json:"applianceID,omitempty"`
	ApplianceName string `json:"applianceName,omitempty"`
	SpaceID string `json:"spaceID,omitempty"`
	DefaultStorageClass string `json:"defaultStorageClass,omitempty"`
	DefaultStorageClassDescription string `json:"defaultStorageClassDescription,omitempty"`
	ProviderData *ClusterProviderData `json:"providerData,omitempty"`
}
