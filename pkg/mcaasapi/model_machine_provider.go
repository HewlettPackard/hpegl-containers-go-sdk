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

type MachineProvider struct {
	State string `json:"state,omitempty"`
	Health string `json:"health,omitempty"`
	CreatedDate time.Time `json:"createdDate,omitempty"`
	LastUpdateDate time.Time `json:"lastUpdateDate,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	WorkerType string `json:"workerType,omitempty"`
	OsImages []OsImages `json:"osImages,omitempty"`
	Networks []string `json:"networks,omitempty"`
	ComputeInstanceTypes []ComputeInstanceTypes `json:"computeInstanceTypes,omitempty"`
	StorageInstanceTypes []string `json:"storageInstanceTypes,omitempty"`
}
