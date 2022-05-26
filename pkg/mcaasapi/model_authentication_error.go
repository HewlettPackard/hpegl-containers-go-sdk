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
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
	RecommendedActions []string `json:"recommendedActions,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
}
