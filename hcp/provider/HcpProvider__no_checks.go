// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HcpProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_HcpProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateHcpProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateHcpProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHcpProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateHcpProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_HcpProvider) validateSetSkipStatusCheckParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HcpProvider) validateSetWorkloadIdentityParameters(val interface{}) error {
	return nil
}

func validateNewHcpProviderParameters(scope constructs.Construct, id *string, config *HcpProviderConfig) error {
	return nil
}

