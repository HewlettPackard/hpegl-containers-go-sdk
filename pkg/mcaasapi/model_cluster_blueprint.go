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
import (
	"time"
)

type ClusterBlueprint struct {
	CreatedDate time.Time `json:"createdDate,omitempty"`
	LastUpdateDate time.Time `json:"lastUpdateDate,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	K8sVersion string `json:"k8sVersion,omitempty"`
	ClusterProvider string `json:"clusterProvider,omitempty"`
	MachineSets []MachineSet `json:"machineSets,omitempty"`
	MachineSetsDetail []MachineSetDetail `json:"machineSetsDetail,omitempty"`
	ApplianceID string `json:"applianceID,omitempty"`
	DefaultStorageClass string `json:"defaultStorageClass,omitempty"`
}