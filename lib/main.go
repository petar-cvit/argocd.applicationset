package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"os"
	"petar-cvit/cdk8s-argo/lib/imports/k8s"
)

type MyChartProps struct {
	cdk8s.ChartProps

	Color string
}

func NewMyChart(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	k8s.NewKubeDeployment(chart, jsii.String("aefae"), &k8s.KubeDeploymentProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("service-" + props.Color),
		},
	})

	return chart
}

func main() {
	color := os.Args[1]

	app := cdk8s.NewApp(nil)
	NewMyChart(app, "cdk8s", &MyChartProps{Color: color})
	app.Synth()
}
