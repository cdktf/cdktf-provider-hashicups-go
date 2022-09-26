//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt hashicups Provider for Terraform CDK (cdktf)
package hashicups

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataHashicupsOrderItemsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataHashicupsOrderItemsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataHashicupsOrderItemsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataHashicupsOrderItemsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataHashicupsOrderItemsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataHashicupsOrderItemsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

