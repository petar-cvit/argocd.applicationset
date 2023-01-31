//go:build no_runtime_type_checking

// argoprojio
package argoprojio

// Building without runtime type checking enabled, so all the below just return nil

func validateApplicationSet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApplicationSet_ManifestParameters(props *ApplicationSetProps) error {
	return nil
}

func validateApplicationSet_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewApplicationSetParameters(scope constructs.Construct, id *string, props *ApplicationSetProps) error {
	return nil
}

