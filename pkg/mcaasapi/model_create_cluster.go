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

type CreateCluster struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ClusterBlueprintId string `json:"clusterBlueprintId,omitempty"`
	ApplianceID string `json:"applianceID,omitempty"`
	SpaceID string `json:"spaceID,omitempty"`
}
