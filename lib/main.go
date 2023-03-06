package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"io/ioutil"
	"log"
	"os"
	"petar-cvit/cdk8s-argo/lib/imports/argoprojio"
	"strings"
)

type NamespaceApp struct {
	Namespace string
}

type MyChartProps struct {
	cdk8s.ChartProps

	NamespaceApps []NamespaceApp
}

func NewApplication(namespace string, scope constructs.Construct) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	chart := cdk8s.NewChart(scope, jsii.String(namespace), &cprops)

	argoprojio.NewApplication(chart, jsii.String(namespace+"-application"), &argoprojio.ApplicationProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      jsii.String(namespace + "-application"),
			Namespace: jsii.String("argocd"),
		},
		Spec: &argoprojio.ApplicationSpec{
			Destination: &argoprojio.ApplicationSpecDestination{
				Namespace: jsii.String("argocd"),
				Server:    jsii.String("https://kubernetes.default.svc"),
			},
			Project: jsii.String("default"),
			Source: &argoprojio.ApplicationSpecSource{
				RepoUrl:        jsii.String("https://github.com/petar-cvit/argocd.applicationset"),
				TargetRevision: jsii.String("main"),
				Path:           jsii.String("underthehood/02-service-application"),
				Helm: &argoprojio.ApplicationSpecSourceHelm{
					Parameters: &[]*argoprojio.ApplicationSpecSourceHelmParameters{
						{
							Name:  jsii.String("namespace"),
							Value: jsii.String(namespace),
						},
					},
					ValueFiles: &[]*string{
						jsii.String("../../config/prod/" + namespace + ".yaml"),
						jsii.String("colors.yaml"),
					},
				},
			},
		},
	})

	return chart
}

func NewMyChart(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	for _, app := range props.NamespaceApps {
		argoprojio.NewApplication(chart, jsii.String(app.Namespace+"-application"), &argoprojio.ApplicationProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name:      jsii.String(app.Namespace + "-application"),
				Namespace: jsii.String("argocd"),
			},
			Spec: &argoprojio.ApplicationSpec{
				Destination: &argoprojio.ApplicationSpecDestination{
					Namespace: jsii.String("argocd"),
					Server:    jsii.String("https://kubernetes.default.svc"),
				},
				Project: jsii.String("default"),
				Source: &argoprojio.ApplicationSpecSource{
					RepoUrl:        jsii.String("https://github.com/petar-cvit/argocd.applicationset"),
					TargetRevision: jsii.String("main"),
					Path:           jsii.String("underthehood/02-service-application"),
					Helm: &argoprojio.ApplicationSpecSourceHelm{
						Parameters: &[]*argoprojio.ApplicationSpecSourceHelmParameters{
							{
								Name:  jsii.String("namespace"),
								Value: jsii.String(app.Namespace),
							},
						},
						ValueFiles: &[]*string{
							jsii.String("../../config/prod/" + app.Namespace + ".yaml"),
							jsii.String("colors.yaml"),
						},
					},
				},
			},
		})
	}

	return chart
}

func main() {
	if os.Args[1] == "application" {
		app := cdk8s.NewApp(nil)
		NewApplication(os.Args[2], app)

		app.Synth()

		return
	}

	if os.Args[1] == "service" {
		name := os.Args[2]
		namespace := os.Args[3]
		color := os.Args[4]
		version := os.Args[5]

		synthServices(name, namespace, color, version)
		return
	}

	files, err := ioutil.ReadDir("../config/prod")
	if err != nil {
		log.Fatal(err)
	}

	app := cdk8s.NewApp(nil)

	props := make([]NamespaceApp, 0, 0)

	for _, f := range files {
		namespace := strings.Replace(f.Name(), ".yaml", "", -1)
		props = append(props, NamespaceApp{Namespace: namespace})
	}

	NewMyChart(app, "cdk8s", &MyChartProps{
		NamespaceApps: props,
	})

	app.Synth()
}
