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

type AuthenticationError struct {
	HttpStatusCode int32 `json:"httpStatusCode,omitempty"`
	Message string `json:"message,omitempty"`
	DebugId string `json:"debugId,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
}
