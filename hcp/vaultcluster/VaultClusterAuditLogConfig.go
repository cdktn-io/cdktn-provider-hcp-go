// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultcluster


type VaultClusterAuditLogConfig struct {
	// CloudWatch access key ID for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#cloudwatch_access_key_id VaultCluster#cloudwatch_access_key_id}
	CloudwatchAccessKeyId *string `field:"optional" json:"cloudwatchAccessKeyId" yaml:"cloudwatchAccessKeyId"`
	// CloudWatch region for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#cloudwatch_region VaultCluster#cloudwatch_region}
	CloudwatchRegion *string `field:"optional" json:"cloudwatchRegion" yaml:"cloudwatchRegion"`
	// CloudWatch secret access key for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#cloudwatch_secret_access_key VaultCluster#cloudwatch_secret_access_key}
	CloudwatchSecretAccessKey *string `field:"optional" json:"cloudwatchSecretAccessKey" yaml:"cloudwatchSecretAccessKey"`
	// Datadog api key for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#datadog_api_key VaultCluster#datadog_api_key}
	DatadogApiKey *string `field:"optional" json:"datadogApiKey" yaml:"datadogApiKey"`
	// Datadog region for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#datadog_region VaultCluster#datadog_region}
	DatadogRegion *string `field:"optional" json:"datadogRegion" yaml:"datadogRegion"`
	// ElasticSearch endpoint for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#elasticsearch_endpoint VaultCluster#elasticsearch_endpoint}
	ElasticsearchEndpoint *string `field:"optional" json:"elasticsearchEndpoint" yaml:"elasticsearchEndpoint"`
	// ElasticSearch password for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#elasticsearch_password VaultCluster#elasticsearch_password}
	ElasticsearchPassword *string `field:"optional" json:"elasticsearchPassword" yaml:"elasticsearchPassword"`
	// ElasticSearch user for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#elasticsearch_user VaultCluster#elasticsearch_user}
	ElasticsearchUser *string `field:"optional" json:"elasticsearchUser" yaml:"elasticsearchUser"`
	// Grafana endpoint for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#grafana_endpoint VaultCluster#grafana_endpoint}
	GrafanaEndpoint *string `field:"optional" json:"grafanaEndpoint" yaml:"grafanaEndpoint"`
	// Grafana password for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#grafana_password VaultCluster#grafana_password}
	GrafanaPassword *string `field:"optional" json:"grafanaPassword" yaml:"grafanaPassword"`
	// Grafana user for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#grafana_user VaultCluster#grafana_user}
	GrafanaUser *string `field:"optional" json:"grafanaUser" yaml:"grafanaUser"`
	// HTTP basic authentication password for streaming audit logs, one of the two available authentication methods, can be specified only if http_basic_user is also provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#http_basic_password VaultCluster#http_basic_password}
	HttpBasicPassword *string `field:"optional" json:"httpBasicPassword" yaml:"httpBasicPassword"`
	// HTTP basic authentication username for streaming audit logs, one of the two available authentication methods, can be specified only if http_basic_password is also provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#http_basic_user VaultCluster#http_basic_user}
	HttpBasicUser *string `field:"optional" json:"httpBasicUser" yaml:"httpBasicUser"`
	// HTTP bearer authentication token for streaming audit logs, one of the two available authentication methods, can be specified only if http_basic_user and http_basic_password are not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#http_bearer_token VaultCluster#http_bearer_token}
	HttpBearerToken *string `field:"optional" json:"httpBearerToken" yaml:"httpBearerToken"`
	// HTTP codec for streaming audit logs, allowed values are JSON and NDJSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#http_codec VaultCluster#http_codec}
	HttpCodec *string `field:"optional" json:"httpCodec" yaml:"httpCodec"`
	// HTTP compression flag for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#http_compression VaultCluster#http_compression}
	HttpCompression interface{} `field:"optional" json:"httpCompression" yaml:"httpCompression"`
	// HTTP headers for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#http_headers VaultCluster#http_headers}
	HttpHeaders *map[string]*string `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// HTTP payload method for streaming audit logs, , allowed values are PATCH, POST, or PUT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#http_method VaultCluster#http_method}
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// HTTP payload prefix for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#http_payload_prefix VaultCluster#http_payload_prefix}
	HttpPayloadPrefix *string `field:"optional" json:"httpPayloadPrefix" yaml:"httpPayloadPrefix"`
	// HTTP payload suffix for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#http_payload_suffix VaultCluster#http_payload_suffix}
	HttpPayloadSuffix *string `field:"optional" json:"httpPayloadSuffix" yaml:"httpPayloadSuffix"`
	// HTTP URI for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#http_uri VaultCluster#http_uri}
	HttpUri *string `field:"optional" json:"httpUri" yaml:"httpUri"`
	// NewRelic Account ID for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#newrelic_account_id VaultCluster#newrelic_account_id}
	NewrelicAccountId *string `field:"optional" json:"newrelicAccountId" yaml:"newrelicAccountId"`
	// NewRelic license key for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#newrelic_license_key VaultCluster#newrelic_license_key}
	NewrelicLicenseKey *string `field:"optional" json:"newrelicLicenseKey" yaml:"newrelicLicenseKey"`
	// NewRelic region for streaming audit logs, allowed values are "US" and "EU".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#newrelic_region VaultCluster#newrelic_region}
	NewrelicRegion *string `field:"optional" json:"newrelicRegion" yaml:"newrelicRegion"`
	// Splunk endpoint for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#splunk_hecendpoint VaultCluster#splunk_hecendpoint}
	SplunkHecendpoint *string `field:"optional" json:"splunkHecendpoint" yaml:"splunkHecendpoint"`
	// Splunk token for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_cluster#splunk_token VaultCluster#splunk_token}
	SplunkToken *string `field:"optional" json:"splunkToken" yaml:"splunkToken"`
}

