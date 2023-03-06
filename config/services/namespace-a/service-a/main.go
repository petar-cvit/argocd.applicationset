package main

import "petar-cvit/cdk8s-argo/lib/models"

func GetConfig(target string) models.AppProps {
	return models.AppProps{
		Env: map[string]string{
			"TARGET":      target,
			"KAFKA_TOPIC": target + "-offer-kafka-topic",
		},
		Host:     "https://offer.api/" + target,
		Replicas: 3,
	}
}
