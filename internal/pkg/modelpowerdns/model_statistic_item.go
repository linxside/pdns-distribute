/*
 * PowerDNS Authoritative HTTP API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.13
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
//nolint:golint
package modelpowerdns

type StatisticItem struct {
	// Item name
	Name string `json:"name,omitempty"`
	// set to \"StatisticItem\"
	Type_ string `json:"type,omitempty"`
	// Item value
	Value string `json:"value,omitempty"`
}
