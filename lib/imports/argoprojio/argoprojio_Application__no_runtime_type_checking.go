//go:build no_runtime_type_checking

// argoprojio
package argoprojio

// Building without runtime type checking enabled, so all the below just return nil

func validateApplication_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApplication_ManifestParameters(props *ApplicationProps) error {
	return nil
}

func validateApplication_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewApplicationParameters(scope constructs.Construct, id *string, props *ApplicationProps) error {
	return nil
}

