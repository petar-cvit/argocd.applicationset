package main

import "petar-cvit/cdk8s-argo/pipelines/models"

func GetConfig(color string) models.AppProps {
	return models.AppProps{
		Name:      "demo-service",
		Namespace: "demo",
		Env: map[string]string{
			"COLOR":       color,
			"TENANCY":     "tenancy/" + color,
			"KAFKA_TOPI":  color + ".rifgjealf",
			"KAFKA_TOPIC": color + ".rifgjealf",
		},
		Replicas: 5,
		Host:     "https://my-demo.com",
	}
}
