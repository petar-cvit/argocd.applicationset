package main

import "github.com/superbet-group/offer.infrastructure.cdk8s/pkg"

func GetConfig(ctx pkg.Context) pkg.AppProps {
	return pkg.AppProps{
		Name:      "demo-service",
		Namespace: "demo",
		Env: map[string]string{
			"CLUSTER":     ctx.GetCluster(),
			"ENV":         ctx.GetEnvironment(),
			"TARGET":      ctx.GetTarget(),
			"TENANCY":     "tenancy/" + ctx.GetTarget(),
			"KAFKA_TOPIC": ctx.GetTarget() + ".rifgjealf",
		},
		Replicas: 3,
	}
}
