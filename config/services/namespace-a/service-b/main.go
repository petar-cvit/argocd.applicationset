package main

import "petar-cvit/cdk8s-argo/lib/models"

func GetConfig(color string) models.AppProps {
	return models.AppProps{
		Env: map[string]string{
			"COLOR":   color,
			"TENANCY": "tenancy/" + color,
			"KEY":     "im.service.b",
		},
		Replicas: 1,
	}
}
