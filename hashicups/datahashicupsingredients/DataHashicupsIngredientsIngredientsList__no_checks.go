// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datahashicupsingredients

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataHashicupsIngredientsIngredientsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataHashicupsIngredientsIngredientsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataHashicupsIngredientsIngredientsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataHashicupsIngredientsIngredientsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataHashicupsIngredientsIngredientsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataHashicupsIngredientsIngredientsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

