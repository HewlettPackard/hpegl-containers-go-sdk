/*
 * HPE GreenLake for Private Cloud Containers APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcaasapi

type StorageClass struct {
	Name string `json:"name,omitempty"`
	Id string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	GlStorageType string `json:"glStorageType,omitempty"`
	AccessProtocol string `json:"accessProtocol,omitempty"`
	Iops string `json:"iops,omitempty"`
	Encryption string `json:"encryption,omitempty"`
	Dedupe string `json:"dedupe,omitempty"`
	CostPerGB string `json:"costPerGB,omitempty"`
}
