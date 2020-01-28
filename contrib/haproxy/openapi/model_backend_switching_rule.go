/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.2
 * Contact: support@haproxy.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// BackendSwitchingRule HAProxy backend switching rule configuration (corresponds to use_backend directive)
type BackendSwitchingRule struct {
	Cond     string `json:"cond,omitempty"`
	CondTest string `json:"cond_test,omitempty"`
	Id       *int32 `json:"id"`
	Name     string `json:"name"`
}