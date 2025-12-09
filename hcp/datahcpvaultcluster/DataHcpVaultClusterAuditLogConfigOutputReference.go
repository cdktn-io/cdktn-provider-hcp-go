// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahcpvaultcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/jsii"

	"github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/datahcpvaultcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpVaultClusterAuditLogConfigOutputReference interface {
	cdktf.ComplexObject
	CloudwatchAccessKeyId() *string
	CloudwatchGroupName() *string
	CloudwatchRegion() *string
	CloudwatchSecretAccessKey() *string
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
	DatadogRegion() *string
	ElasticsearchDataset() *string
	ElasticsearchEndpoint() *string
	ElasticsearchPassword() *string
	ElasticsearchUser() *string
	// Experimental.
	Fqn() *string
	GrafanaEndpoint() *string
	GrafanaUser() *string
	HttpBasicPassword() *string
	HttpBasicUser() *string
	HttpBearerToken() *string
	HttpCodec() *string
	HttpCompression() cdktf.IResolvable
	HttpHeaders() cdktf.StringMap
	HttpMethod() *string
	HttpPayloadPrefix() *string
	HttpPayloadSuffix() *string
	HttpUri() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NewrelicAccountId() *string
	NewrelicLicenseKey() *string
	NewrelicRegion() *string
	SplunkHecendpoint() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataHcpVaultClusterAuditLogConfigOutputReference
type jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) CloudwatchAccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAccessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) CloudwatchGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) CloudwatchRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) CloudwatchSecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchSecretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) CloudwatchStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) DatadogRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) ElasticsearchDataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) ElasticsearchEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) ElasticsearchPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) ElasticsearchUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) GrafanaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) GrafanaUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) HttpBasicPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBasicPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) HttpBasicUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBasicUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) HttpBearerToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBearerToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) HttpCodec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCodec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) HttpCompression() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"httpCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) HttpHeaders() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"httpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) HttpPayloadPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPayloadPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) HttpPayloadSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPayloadSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) HttpUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) NewrelicAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) NewrelicLicenseKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicLicenseKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) NewrelicRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newrelicRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) SplunkHecendpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkHecendpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataHcpVaultClusterAuditLogConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataHcpVaultClusterAuditLogConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataHcpVaultClusterAuditLogConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-hcp.dataHcpVaultCluster.DataHcpVaultClusterAuditLogConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataHcpVaultClusterAuditLogConfigOutputReference_Override(d DataHcpVaultClusterAuditLogConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcp.dataHcpVaultCluster.DataHcpVaultClusterAuditLogConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHcpVaultClusterAuditLogConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

