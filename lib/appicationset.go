package main

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"petar-cvit/cdk8s-argo/lib/imports/argoprojio"
	"petar-cvit/cdk8s-argo/lib/models"
)

func NewApplicationsSet(scope constructs.Construct, id string, props *models.ApplicationSet) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	argoprojio.NewApplicationSet(chart, jsii.String(id), &argoprojio.ApplicationSetProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      jsii.String(props.Name),
			Namespace: jsii.String(props.Namespace),
			Labels: &map[string]*string{
				"app.kubernetes.io/instance": jsii.String("colors-service"),
			},
		},
		Spec: &argoprojio.ApplicationSetSpec{
			Generators: &[]*argoprojio.ApplicationSetSpecGenerators{
				{
					List: &argoprojio.ApplicationSetSpecGeneratorsList{
						Elements: &[]interface{}{
							jsii.String("color"),
						},
					},
				},
			},
			Template: &argoprojio.ApplicationSetSpecTemplate{
				Metadata: &argoprojio.ApplicationSetSpecTemplateMetadata{
					Name: jsii.String(props.Name + "-" + "green"),
				},
				Spec: &argoprojio.ApplicationSetSpecTemplateSpec{
					Destination: &argoprojio.ApplicationSetSpecTemplateSpecDestination{
						Namespace: jsii.String(props.Namespace),
						Server:    jsii.String("https://kubernetes.default.svc"),
					},
					Project: jsii.String("default"),
					Source: &argoprojio.ApplicationSetSpecTemplateSpecSource{
						RepoUrl:        jsii.String("https://github.com/petar-cvit/argocd.applicationset"),
						Path:           jsii.String(fmt.Sprintf("synth/services/%s/%s", props.Namespace, props.Name)),
						TargetRevision: jsii.String("HEAD"),
					},
				},
			},
		},
	})

	return chart
}
