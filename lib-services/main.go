package main

import (
	"example.com/lib-services/imports/k8s"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"os"
)

type MyChartProps struct {
	cdk8s.ChartProps

	Color     string
	Version   string
	Name      string
	Namespace string
}

func NewMyChart(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	k8s.NewKubeDeployment(chart, jsii.String("application"), &k8s.KubeDeploymentProps{
		Metadata: &k8s.ObjectMeta{
			Name:      jsii.String(props.Name),
			Namespace: jsii.String("guestbook"),
		},
		Spec: &k8s.DeploymentSpec{
			Selector: &k8s.LabelSelector{
				MatchLabels: &map[string]*string{
					"app": jsii.String(props.Name),
				},
			},
			Template: &k8s.PodTemplateSpec{
				Metadata: &k8s.ObjectMeta{
					Labels: &map[string]*string{
						"app": jsii.String(props.Name),
					},
				},
				Spec: &k8s.PodSpec{
					Containers: &[]*k8s.Container{
						{
							Name:  jsii.String(props.Name),
							Image: jsii.String(props.Name + ":" + props.Version),
						},
					},
				},
			},
		},
	})

	return chart
}

func main() {
	name := os.Args[1]
	namespace := os.Args[2]
	color := os.Args[3]
	version := os.Args[4]

	app := cdk8s.NewApp(nil)

	//props := make([]NamespaceApp, 0, 0)
	//
	//for _, f := range files {
	//	namespace := strings.Replace(f.Name(), ".yaml", "", -1)
	//	props = append(props, NamespaceApp{Namespace: namespace})
	//}

	NewMyChart(app, "cdk8s", &MyChartProps{
		Color:     color,
		Version:   version,
		Name:      name,
		Namespace: namespace,
	})

	app.Synth()
}
