// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package waypointaddon

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WaypointAddOnAddOnInputVariablesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WaypointAddOnAddOnInputVariablesList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WaypointAddOnAddOnInputVariablesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WaypointAddOnAddOnInputVariablesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WaypointAddOnAddOnInputVariablesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WaypointAddOnAddOnInputVariablesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WaypointAddOnAddOnInputVariablesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWaypointAddOnAddOnInputVariablesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

