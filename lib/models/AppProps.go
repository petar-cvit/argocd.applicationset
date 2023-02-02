package models

import "github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"

type AppProps struct {
	cdk8s.ChartProps

	Name      string
	Namespace string
	Color     string
	Version   string

	Env      map[string]string
	Host     string
	Replicas int
}
