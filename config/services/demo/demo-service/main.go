package main

import "github.com/superbet-group/offer.infrastructure.cdk8s/pkg"

func GetConfig(color string) pkg.AppProps {
	return pkg.AppProps{
		Name:      "demo-service",
		Namespace: "demo",
		Env: map[string]string{
			"COLORa":        color,
			"TENANCY":       "tenancy/" + color,
			"KAFKA_TOPI":    color + ".rifgjealf",
			"KAFKA_TOPIC":   color + ".rifgjealf",
			"KAFKA_TOPIC_2": color + ".rifgjealf",
		},
		Replicas: 3,
	}
}
