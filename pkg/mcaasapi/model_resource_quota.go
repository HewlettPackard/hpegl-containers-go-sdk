/*
 * MCaaS API
 *
 * APIs and object schemas for MCaaS.
 *
 * API version: v1
 * Contact: Centaur-GLHC@hpe.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcaasapi

type ResourceQuota struct {
	ResourceQuotaName string `json:"resourceQuotaName,omitempty"`
	UsedLimitMemory string `json:"usedLimitMemory,omitempty"`
	UsedRequestMemory string `json:"usedRequestMemory,omitempty"`
	UsedLimitCPU string `json:"usedLimitCPU,omitempty"`
	UsedRequestCPU string `json:"usedRequestCPU,omitempty"`
	HardLimitMemory string `json:"hardLimitMemory,omitempty"`
	HardLimitCPU string `json:"hardLimitCPU,omitempty"`
	HardRequestMemory string `json:"hardRequestMemory,omitempty"`
	HardRequestCPU string `json:"hardRequestCPU,omitempty"`
	HardRequestStorage string `json:"hardRequestStorage,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	ClusterID string `json:"clusterID,omitempty"`
	NamespaceID string `json:"namespaceID,omitempty"`
}