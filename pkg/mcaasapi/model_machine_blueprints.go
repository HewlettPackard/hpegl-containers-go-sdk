/*
 * HPE GreenLake for Private Cloud Containers APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcaasapi

type MachineBlueprints struct {
	Count int32 `json:"count,omitempty"`
	Total int32 `json:"total,omitempty"`
	Items []MachineBlueprint `json:"items,omitempty"`
}
