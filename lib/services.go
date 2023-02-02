package main

import (
	"fmt"
	"petar-cvit/cdk8s-argo/lib/imports/k8s"
	"plugin"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"

	"petar-cvit/cdk8s-argo/lib/models"
)

func NewService(scope constructs.Construct, id string, props *models.AppProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// region map env var
	envVars := make([]*k8s.EnvVar, 0, len(props.Env))
	for key, value := range props.Env {
		envVars = append(envVars, &k8s.EnvVar{
			Name:  jsii.String(key),
			Value: jsii.String(value),
		})
	}

	envVars = append(envVars, &k8s.EnvVar{
		Name:  jsii.String("VERSION"),
		Value: jsii.String(props.Version),
	})
	// endregion

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
							Image: jsii.String("gcr.io/google-containers/busybox"),
							Env:   &envVars,
						},
					},
				},
			},
		},
	})

	k8s.NewKubeService(chart, jsii.String("service"), &k8s.KubeServiceProps{
		Metadata: &k8s.ObjectMeta{
			Name:      jsii.String("service-" + props.Name),
			Namespace: jsii.String("guestbook"),
		},
		Spec: &k8s.ServiceSpec{
			Type: jsii.String("LoadBalancer"),
			Ports: &[]*k8s.ServicePort{
				{
					Port:       jsii.Number(80),
					Protocol:   jsii.String("TCP"),
					TargetPort: k8s.IntOrString_FromNumber(jsii.Number(8888)),
				},
			},
		},
	})

	return chart
}

func synthServices(name, namespace, color, version string) {
	app := cdk8s.NewApp(nil)

	mod := fmt.Sprintf("../config/services/%s/%s/main.so", namespace, name)

	plug, err := plugin.Open(mod)
	if err != nil {
		panic(err)
	}

	symExporter, err := plug.Lookup("GetConfig")
	if err != nil {
		panic(err)
	}

	exp, ok := symExporter.(func(string) models.AppProps)
	if !ok {
		panic(symExporter)
	}

	props := exp(color)
	props.Name = name
	props.Namespace = namespace
	props.Version = version

	//props := make([]NamespaceApp, 0, 0)
	//
	//for _, f := range files {
	//	namespace := strings.Replace(f.Name(), ".yaml", "", -1)
	//	props = append(props, NamespaceApp{Namespace: namespace})
	//}

	NewService(app, "cdk8s", &props)

	//	&MyChartProps{
	//	Color:     color,
	//	Version:   version,
	//	Name:      name,
	//	Namespace: namespace,
	//})

	app.Synth()
}
