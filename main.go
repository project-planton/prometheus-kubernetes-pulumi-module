package main

import (
	"github.com/pkg/errors"
	prometheuskubernetesv1 "github.com/project-planton/project-planton/apis/go/project/planton/provider/kubernetes/prometheuskubernetes/v1"
	"github.com/project-planton/project-planton/pkg/pulmod/stackinput"
	"github.com/project-planton/prometheus-kubernetes-pulumi-module/pkg"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &prometheuskubernetesv1.PrometheusKubernetesStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
