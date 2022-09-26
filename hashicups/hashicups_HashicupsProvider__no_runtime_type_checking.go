//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt hashicups Provider for Terraform CDK (cdktf)
package hashicups

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HashicupsProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_HashicupsProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateHashicupsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewHashicupsProviderParameters(scope constructs.Construct, id *string, config *HashicupsProviderConfig) error {
	return nil
}

