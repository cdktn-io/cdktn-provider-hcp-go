// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package vaultplugin

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VaultPlugin) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateImportFromParameters(id *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateMoveToIdParameters(id *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (v *jsiiProxy_VaultPlugin) validatePutTimeoutsParameters(value *VaultPluginTimeouts) error {
	return nil
}

func validateVaultPlugin_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateVaultPlugin_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVaultPlugin_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateVaultPlugin_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultPlugin) validateSetClusterIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VaultPlugin) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultPlugin) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultPlugin) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VaultPlugin) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_VaultPlugin) validateSetPluginNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VaultPlugin) validateSetPluginTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VaultPlugin) validateSetProjectIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VaultPlugin) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewVaultPluginParameters(scope constructs.Construct, id *string, config *VaultPluginConfig) error {
	return nil
}

