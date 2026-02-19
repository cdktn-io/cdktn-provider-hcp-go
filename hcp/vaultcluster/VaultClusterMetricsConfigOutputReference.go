// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-hcp-go/hcp/v11/jsii"

	"github.com/cdktn-io/cdktn-provider-hcp-go/hcp/v11/vaultcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VaultClusterMetricsConfigOutputReference interface {
	cdktn.ComplexObject
	CloudwatchAccessKeyId() *string
	SetCloudwatchAccessKeyId(val *string)
	CloudwatchAccessKeyIdInput() *string
	CloudwatchNamespace() *string
	CloudwatchRegion() *string
	SetCloudwatchRegion(val *string)
	CloudwatchRegionInput() *string
	CloudwatchSecretAccessKey() *string
	SetCloudwatchSecretAccessKey(val *string)
	CloudwatchSecretAccessKeyInput() *string
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
	InternalValue() *VaultClusterMetricsConfig
	SetInternalValue(val *VaultClusterMetricsConfig)
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
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
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
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VaultClusterMetricsConfigOutputReference
type jsiiProxy_VaultClusterMetricsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) CloudwatchAccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAccessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) CloudwatchAccessKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAccessKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) CloudwatchNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) CloudwatchRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) CloudwatchRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) CloudwatchSecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchSecretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) CloudwatchSecretAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchSecretAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) DatadogApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) DatadogApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) DatadogRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) DatadogRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) ElasticsearchDataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) ElasticsearchEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) ElasticsearchEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) ElasticsearchPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) ElasticsearchPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) ElasticsearchUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) ElasticsearchUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) GrafanaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) GrafanaEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) GrafanaPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) GrafanaPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) GrafanaUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) GrafanaUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpBasicPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBasicPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpBasicPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBasicPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpBasicUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBasicUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpBasicUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBasicUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpBearerToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBearerToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpBearerTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBearerTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpCodec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCodec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpCodecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCodecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpCompression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpCompressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"httpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"httpHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpPayloadPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPayloadPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpPayloadPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPayloadPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpPayloadSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPayloadSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpPayloadSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPayloadSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) HttpUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) InternalValue() *VaultClusterMetricsConfig {
	var returns *VaultClusterMetricsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) NewrelicAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) NewrelicAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) NewrelicLicenseKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicLicenseKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) NewrelicLicenseKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicLicenseKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) NewrelicRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) NewrelicRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) SplunkHecendpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkHecendpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) SplunkHecendpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkHecendpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) SplunkToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) SplunkTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVaultClusterMetricsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VaultClusterMetricsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVaultClusterMetricsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VaultClusterMetricsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-hcp.vaultCluster.VaultClusterMetricsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVaultClusterMetricsConfigOutputReference_Override(v VaultClusterMetricsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-hcp.vaultCluster.VaultClusterMetricsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetCloudwatchAccessKeyId(val *string) {
	if err := j.validateSetCloudwatchAccessKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchAccessKeyId",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetCloudwatchRegion(val *string) {
	if err := j.validateSetCloudwatchRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchRegion",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetCloudwatchSecretAccessKey(val *string) {
	if err := j.validateSetCloudwatchSecretAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchSecretAccessKey",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetDatadogApiKey(val *string) {
	if err := j.validateSetDatadogApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datadogApiKey",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetDatadogRegion(val *string) {
	if err := j.validateSetDatadogRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datadogRegion",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetElasticsearchEndpoint(val *string) {
	if err := j.validateSetElasticsearchEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchEndpoint",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetElasticsearchPassword(val *string) {
	if err := j.validateSetElasticsearchPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchPassword",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetElasticsearchUser(val *string) {
	if err := j.validateSetElasticsearchUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchUser",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetGrafanaEndpoint(val *string) {
	if err := j.validateSetGrafanaEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaEndpoint",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetGrafanaPassword(val *string) {
	if err := j.validateSetGrafanaPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaPassword",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetGrafanaUser(val *string) {
	if err := j.validateSetGrafanaUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaUser",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetHttpBasicPassword(val *string) {
	if err := j.validateSetHttpBasicPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpBasicPassword",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetHttpBasicUser(val *string) {
	if err := j.validateSetHttpBasicUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpBasicUser",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetHttpBearerToken(val *string) {
	if err := j.validateSetHttpBearerTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpBearerToken",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetHttpCodec(val *string) {
	if err := j.validateSetHttpCodecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCodec",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetHttpCompression(val interface{}) {
	if err := j.validateSetHttpCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCompression",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetHttpHeaders(val *map[string]*string) {
	if err := j.validateSetHttpHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHeaders",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetHttpPayloadPrefix(val *string) {
	if err := j.validateSetHttpPayloadPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPayloadPrefix",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetHttpPayloadSuffix(val *string) {
	if err := j.validateSetHttpPayloadSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPayloadSuffix",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetHttpUri(val *string) {
	if err := j.validateSetHttpUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpUri",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetInternalValue(val *VaultClusterMetricsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetNewrelicAccountId(val *string) {
	if err := j.validateSetNewrelicAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newrelicAccountId",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetNewrelicLicenseKey(val *string) {
	if err := j.validateSetNewrelicLicenseKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newrelicLicenseKey",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetNewrelicRegion(val *string) {
	if err := j.validateSetNewrelicRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newrelicRegion",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetSplunkHecendpoint(val *string) {
	if err := j.validateSetSplunkHecendpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splunkHecendpoint",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetSplunkToken(val *string) {
	if err := j.validateSetSplunkTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splunkToken",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMetricsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetCloudwatchAccessKeyId() {
	_jsii_.InvokeVoid(
		v,
		"resetCloudwatchAccessKeyId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetCloudwatchRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetCloudwatchRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetCloudwatchSecretAccessKey() {
	_jsii_.InvokeVoid(
		v,
		"resetCloudwatchSecretAccessKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetDatadogApiKey() {
	_jsii_.InvokeVoid(
		v,
		"resetDatadogApiKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetDatadogRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetDatadogRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetElasticsearchEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetElasticsearchEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetElasticsearchPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetElasticsearchPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetElasticsearchUser() {
	_jsii_.InvokeVoid(
		v,
		"resetElasticsearchUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetGrafanaEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetGrafanaEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetGrafanaPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetGrafanaPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetGrafanaUser() {
	_jsii_.InvokeVoid(
		v,
		"resetGrafanaUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetHttpBasicPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpBasicPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetHttpBasicUser() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpBasicUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetHttpBearerToken() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpBearerToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetHttpCodec() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpCodec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetHttpCompression() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpCompression",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetHttpHeaders() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpHeaders",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetHttpPayloadPrefix() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpPayloadPrefix",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetHttpPayloadSuffix() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpPayloadSuffix",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetHttpUri() {
	_jsii_.InvokeVoid(
		v,
		"resetHttpUri",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetNewrelicAccountId() {
	_jsii_.InvokeVoid(
		v,
		"resetNewrelicAccountId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetNewrelicLicenseKey() {
	_jsii_.InvokeVoid(
		v,
		"resetNewrelicLicenseKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetNewrelicRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetNewrelicRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetSplunkHecendpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetSplunkHecendpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ResetSplunkToken() {
	_jsii_.InvokeVoid(
		v,
		"resetSplunkToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VaultClusterMetricsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

