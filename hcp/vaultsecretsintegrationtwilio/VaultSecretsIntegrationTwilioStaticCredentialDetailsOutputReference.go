// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegrationtwilio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/jsii"

	"github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/vaultsecretsintegrationtwilio/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference interface {
	cdktf.ComplexObject
	AccountSid() *string
	SetAccountSid(val *string)
	AccountSidInput() *string
	ApiKeySecret() *string
	SetApiKeySecret(val *string)
	ApiKeySecretInput() *string
	ApiKeySid() *string
	SetApiKeySid(val *string)
	ApiKeySidInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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

// The jsii proxy struct for VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference
type jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) AccountSid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountSid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) AccountSidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountSidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) ApiKeySecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) ApiKeySecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) ApiKeySid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) ApiKeySidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewVaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultSecretsIntegrationTwilio.VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference_Override(v VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultSecretsIntegrationTwilio.VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference)SetAccountSid(val *string) {
	if err := j.validateSetAccountSidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountSid",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference)SetApiKeySecret(val *string) {
	if err := j.validateSetApiKeySecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKeySecret",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference)SetApiKeySid(val *string) {
	if err := j.validateSetApiKeySidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKeySid",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VaultSecretsIntegrationTwilioStaticCredentialDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

