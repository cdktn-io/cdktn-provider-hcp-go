// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/jsii"

	"github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/vaultcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultClusterAuditLogConfigOutputReference interface {
	cdktf.ComplexObject
	CloudwatchAccessKeyId() *string
	SetCloudwatchAccessKeyId(val *string)
	CloudwatchAccessKeyIdInput() *string
	CloudwatchGroupName() *string
	CloudwatchRegion() *string
	SetCloudwatchRegion(val *string)
	CloudwatchRegionInput() *string
	CloudwatchSecretAccessKey() *string
	SetCloudwatchSecretAccessKey(val *string)
	CloudwatchSecretAccessKeyInput() *string
	CloudwatchStreamName() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatadogApiKey() *string
	SetDatadogApiKey(val *string)
	DatadogApiKeyInput() *string
	DatadogRegion() *string
	SetDatadogRegion(val *string)
	DatadogRegionInput() *string
	ElasticsearchDataset() *string
	ElasticsearchEndpoint() *string
	SetElasticsearchEndpoint(val *string)
	ElasticsearchEndpointInput() *string
	ElasticsearchPassword() *string
	SetElasticsearchPassword(val *string)
	ElasticsearchPasswordInput() *string
	ElasticsearchUser() *string
	SetElasticsearchUser(val *string)
	ElasticsearchUserInput() *string
	// Experimental.
	Fqn() *string
	GrafanaEndpoint() *string
	SetGrafanaEndpoint(val *string)
	GrafanaEndpointInput() *string
	GrafanaPassword() *string
	SetGrafanaPassword(val *string)
	GrafanaPasswordInput() *string
	GrafanaUser() *string
	SetGrafanaUser(val *string)
	GrafanaUserInput() *string
	HttpBasicPassword() *string
	SetHttpBasicPassword(val *string)
	HttpBasicPasswordInput() *string
	HttpBasicUser() *string
	SetHttpBasicUser(val *string)
	HttpBasicUserInput() *string
	HttpBearerToken() *string
	SetHttpBearerToken(val *string)
	HttpBearerTokenInput() *string
	HttpCodec() *string
	SetHttpCodec(val *string)
	HttpCodecInput() *string
	HttpCompression() interface{}
	SetHttpCompression(val interface{})
	HttpCompressionInput() interface{}
	HttpHeaders() *map[string]*string
	SetHttpHeaders(val *map[string]*string)
	HttpHeadersInput() *map[string]*string
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	HttpPayloadPrefix() *string
	SetHttpPayloadPrefix(val *string)
	HttpPayloadPrefixInput() *string
	HttpPayloadSuffix() *string
	SetHttpPayloadSuffix(val *string)
	HttpPayloadSuffixInput() *string
	HttpUri() *string
	SetHttpUri(val *string)
	HttpUriInput() *string
	InternalValue() *VaultClusterAuditLogConfig
	SetInternalValue(val *VaultClusterAuditLogConfig)
	NewrelicAccountId() *string
	SetNewrelicAccountId(val *string)
	NewrelicAccountIdInput() *string
	NewrelicLicenseKey() *string
	SetNewrelicLicenseKey(val *string)
	NewrelicLicenseKeyInput() *string
	NewrelicRegion() *string
	SetNewrelicRegion(val *string)
	NewrelicRegionInput() *string
	SplunkHecendpoint() *string
	SetSplunkHecendpoint(val *string)
	SplunkHecendpointInput() *string
	SplunkToken() *string
	SetSplunkToken(val *string)
	SplunkTokenInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetCloudwatchAccessKeyId()
	ResetCloudwatchRegion()
	ResetCloudwatchSecretAccessKey()
	ResetDatadogApiKey()
	ResetDatadogRegion()
	ResetElasticsearchEndpoint()
	ResetElasticsearchPassword()
	ResetElasticsearchUser()
	ResetGrafanaEndpoint()
	ResetGrafanaPassword()
	ResetGrafanaUser()
	ResetHttpBasicPassword()
	ResetHttpBasicUser()
	ResetHttpBearerToken()
	ResetHttpCodec()
	ResetHttpCompression()
	ResetHttpHeaders()
	ResetHttpMethod()
	ResetHttpPayloadPrefix()
	ResetHttpPayloadSuffix()
	ResetHttpUri()
	ResetNewrelicAccountId()
	ResetNewrelicLicenseKey()
	ResetNewrelicRegion()
	ResetSplunkHecendpoint()
	ResetSplunkToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VaultClusterAuditLogConfigOutputReference
type jsiiProxy_VaultClusterAuditLogConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) CloudwatchAccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAccessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) CloudwatchAccessKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAccessKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) CloudwatchGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) CloudwatchRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) CloudwatchRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) CloudwatchSecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchSecretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) CloudwatchSecretAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchSecretAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) CloudwatchStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) DatadogApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) DatadogApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) DatadogRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) DatadogRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ElasticsearchDataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ElasticsearchEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ElasticsearchEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ElasticsearchPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ElasticsearchPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ElasticsearchUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ElasticsearchUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpBasicPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBasicPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpBasicPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBasicPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpBasicUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBasicUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpBasicUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBasicUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpBearerToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBearerToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpBearerTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBearerTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpCodec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCodec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpCodecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCodecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpCompression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpCompressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"httpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"httpHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpPayloadPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPayloadPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpPayloadPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPayloadPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpPayloadSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPayloadSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpPayloadSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPayloadSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) HttpUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) InternalValue() *VaultClusterAuditLogConfig {
	var returns *VaultClusterAuditLogConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) NewrelicAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) NewrelicAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) NewrelicLicenseKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicLicenseKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) NewrelicLicenseKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicLicenseKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) NewrelicRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) NewrelicRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) SplunkHecendpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkHecendpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) SplunkHecendpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkHecendpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) SplunkToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) SplunkTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVaultClusterAuditLogConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VaultClusterAuditLogConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVaultClusterAuditLogConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VaultClusterAuditLogConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterAuditLogConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVaultClusterAuditLogConfigOutputReference_Override(v VaultClusterAuditLogConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterAuditLogConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetCloudwatchAccessKeyId(val *string) {
	if err := j.validateSetCloudwatchAccessKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchAccessKeyId",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetCloudwatchRegion(val *string) {
	if err := j.validateSetCloudwatchRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchRegion",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetCloudwatchSecretAccessKey(val *string) {
	if err := j.validateSetCloudwatchSecretAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchSecretAccessKey",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetDatadogApiKey(val *string) {
	if err := j.validateSetDatadogApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datadogApiKey",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetDatadogRegion(val *string) {
	if err := j.validateSetDatadogRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datadogRegion",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetElasticsearchEndpoint(val *string) {
	if err := j.validateSetElasticsearchEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchEndpoint",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetElasticsearchPassword(val *string) {
	if err := j.validateSetElasticsearchPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchPassword",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetElasticsearchUser(val *string) {
	if err := j.validateSetElasticsearchUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchUser",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetGrafanaEndpoint(val *string) {
	if err := j.validateSetGrafanaEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaEndpoint",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetGrafanaPassword(val *string) {
	if err := j.validateSetGrafanaPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaPassword",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetGrafanaUser(val *string) {
	if err := j.validateSetGrafanaUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaUser",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetHttpBasicPassword(val *string) {
	if err := j.validateSetHttpBasicPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpBasicPassword",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetHttpBasicUser(val *string) {
	if err := j.validateSetHttpBasicUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpBasicUser",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetHttpBearerToken(val *string) {
	if err := j.validateSetHttpBearerTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpBearerToken",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetHttpCodec(val *string) {
	if err := j.validateSetHttpCodecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCodec",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetHttpCompression(val interface{}) {
	if err := j.validateSetHttpCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCompression",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetHttpHeaders(val *map[string]*string) {
	if err := j.validateSetHttpHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHeaders",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetHttpPayloadPrefix(val *string) {
	if err := j.validateSetHttpPayloadPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPayloadPrefix",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetHttpPayloadSuffix(val *string) {
	if err := j.validateSetHttpPayloadSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPayloadSuffix",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetHttpUri(val *string) {
	if err := j.validateSetHttpUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpUri",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetInternalValue(val *VaultClusterAuditLogConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetNewrelicAccountId(val *string) {
	if err := j.validateSetNewrelicAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newrelicAccountId",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetNewrelicLicenseKey(val *string) {
	if err := j.validateSetNewrelicLicenseKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newrelicLicenseKey",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetNewrelicRegion(val *string) {
	if err := j.validateSetNewrelicRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newrelicRegion",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetSplunkHecendpoint(val *string) {
	if err := j.validateSetSplunkHecendpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splunkHecendpoint",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetSplunkToken(val *string) {
	if err := j.validateSetSplunkTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splunkToken",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetCloudwatchAccessKeyId() {
	_jsii_.InvokeVoid(
		v,
		"resetCloudwatchAccessKeyId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetCloudwatchRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetCloudwatchRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetCloudwatchSecretAccessKey() {
	_jsii_.InvokeVoid(
		v,
		"resetCloudwatchSecretAccessKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetDatadogApiKey() {
	_jsii_.InvokeVoid(
		v,
		"resetDatadogApiKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetDatadogRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetDatadogRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetElasticsearchEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetElasticsearchEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetElasticsearchPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetElasticsearchPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetElasticsearchUser() {
	_jsii_.InvokeVoid(
		v,
		"resetElasticsearchUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetGrafanaEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetGrafanaEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetGrafanaPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetGrafanaPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetGrafanaUser() {
	_jsii_.InvokeVoid(
		v,
		"resetGrafanaUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetHttpBasicPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpBasicPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetHttpBasicUser() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpBasicUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetHttpBearerToken() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpBearerToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetHttpCodec() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpCodec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetHttpCompression() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpCompression",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetHttpHeaders() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpHeaders",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetHttpPayloadPrefix() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpPayloadPrefix",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetHttpPayloadSuffix() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpPayloadSuffix",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetHttpUri() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpUri",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetNewrelicAccountId() {
	_jsii_.InvokeVoid(
		v,
		"resetNewrelicAccountId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetNewrelicLicenseKey() {
	_jsii_.InvokeVoid(
		v,
		"resetNewrelicLicenseKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetNewrelicRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetNewrelicRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetSplunkHecendpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetSplunkHecendpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetSplunkToken() {
	_jsii_.InvokeVoid(
		v,
		"resetSplunkToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

