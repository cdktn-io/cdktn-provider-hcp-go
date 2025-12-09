// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package waypointapplication

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WaypointApplicationOutputValuesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WaypointApplicationOutputValuesList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WaypointApplicationOutputValuesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WaypointApplicationOutputValuesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WaypointApplicationOutputValuesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WaypointApplicationOutputValuesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWaypointApplicationOutputValuesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

