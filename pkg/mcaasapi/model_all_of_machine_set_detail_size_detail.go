/*
 * HPE GreenLake for Private Cloud Containers APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcaasapi

type AllOfMachineSetDetailSizeDetail struct {
	Name string `json:"name,omitempty"`
	Cpu int32 `json:"cpu,omitempty"`
	Memory int32 `json:"memory,omitempty"`
	RootDisk int32 `json:"rootDisk,omitempty"`
	EphemeralDisk int32 `json:"ephemeralDisk,omitempty"`
	PersistentDisk int32 `json:"persistentDisk,omitempty"`
}
