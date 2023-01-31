// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsList struct {
	Elements *[]interface{} `field:"required" json:"elements" yaml:"elements"`
	Template *ApplicationSetSpecGeneratorsListTemplate `field:"optional" json:"template" yaml:"template"`
}

