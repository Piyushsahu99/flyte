/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type CoreEnvironmentAssignment struct {
	Id string `json:"id,omitempty"`
	NodeIds []string `json:"node_ids,omitempty"`
	Environment *CoreEnvironment `json:"environment,omitempty"`
}