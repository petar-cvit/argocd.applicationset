//go:build no_runtime_type_checking

// argoprojio
package argoprojio

// Building without runtime type checking enabled, so all the below just return nil

func validateAppProject_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAppProject_ManifestParameters(props *AppProjectProps) error {
	return nil
}

func validateAppProject_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewAppProjectParameters(scope constructs.Construct, id *string, props *AppProjectProps) error {
	return nil
}

