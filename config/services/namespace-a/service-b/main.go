package main

import "petar-cvit/cdk8s-argo/pipelines/models"

func GetConfig(color string) models.AppProps {
	specificByColor := map[string]string{
		"green":  "red",
		"blue":   "orange",
		"orange": "blue",
	}

	return models.AppProps{
		Name:      "service-b",
		Namespace: "namespace-a",
		Env: map[string]string{
			"COLOR":      color,
			"TENANCY":    "tenancy/" + color,
			"KEY":        "im.service.b",
			"COMPLEMENT": specificByColor[color],
		},
		Replicas: 1,
	}

}
