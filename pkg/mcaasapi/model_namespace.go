/*
 * HPE GreenLake for Private Cloud Containers APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcaasapi
import (
	"time"
)

type Namespace struct {
	CreatedDate time.Time `json:"createdDate,omitempty"`
	LastUpdateDate time.Time `json:"lastUpdateDate,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ApplianceID string `json:"applianceID,omitempty"`
	ClusterID string `json:"clusterID,omitempty"`
	TenantID string `json:"tenantID,omitempty"`
	ResourceQuota []ResourceQuota `json:"resourceQuota,omitempty"`
}
