// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datahcpiampolicy

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataHcpIamPolicyBindingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataHcpIamPolicyBindingsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataHcpIamPolicyBindingsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataHcpIamPolicyBindingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataHcpIamPolicyBindingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataHcpIamPolicyBindingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataHcpIamPolicyBindingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataHcpIamPolicyBindingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

