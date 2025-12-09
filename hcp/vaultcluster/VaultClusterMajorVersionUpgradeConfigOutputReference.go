// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/jsii"

	"github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/vaultcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultClusterMajorVersionUpgradeConfigOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *VaultClusterMajorVersionUpgradeConfig
	SetInternalValue(val *VaultClusterMajorVersionUpgradeConfig)
	MaintenanceWindowDay() *string
	SetMaintenanceWindowDay(val *string)
	MaintenanceWindowDayInput() *string
	MaintenanceWindowTime() *string
	SetMaintenanceWindowTime(val *string)
	MaintenanceWindowTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpgradeType() *string
	SetUpgradeType(val *string)
	UpgradeTypeInput() *string
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
	ResetMaintenanceWindowDay()
	ResetMaintenanceWindowTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VaultClusterMajorVersionUpgradeConfigOutputReference
type jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) InternalValue() *VaultClusterMajorVersionUpgradeConfig {
	var returns *VaultClusterMajorVersionUpgradeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) MaintenanceWindowDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) MaintenanceWindowDayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) MaintenanceWindowTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) MaintenanceWindowTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) UpgradeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) UpgradeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeTypeInput",
		&returns,
	)
	return returns
}


func NewVaultClusterMajorVersionUpgradeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VaultClusterMajorVersionUpgradeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVaultClusterMajorVersionUpgradeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterMajorVersionUpgradeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVaultClusterMajorVersionUpgradeConfigOutputReference_Override(v VaultClusterMajorVersionUpgradeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterMajorVersionUpgradeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference)SetInternalValue(val *VaultClusterMajorVersionUpgradeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference)SetMaintenanceWindowDay(val *string) {
	if err := j.validateSetMaintenanceWindowDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindowDay",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference)SetMaintenanceWindowTime(val *string) {
	if err := j.validateSetMaintenanceWindowTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindowTime",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference)SetUpgradeType(val *string) {
	if err := j.validateSetUpgradeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeType",
		val,
	)
}

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) ResetMaintenanceWindowDay() {
	_jsii_.InvokeVoid(
		v,
		"resetMaintenanceWindowDay",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) ResetMaintenanceWindowTime() {
	_jsii_.InvokeVoid(
		v,
		"resetMaintenanceWindowTime",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

