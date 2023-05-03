package main

import (
	"os"

	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/akuity/kargo/internal/cmd/controlplane"
	"github.com/akuity/kargo/internal/logging"
)

func main() {
	ctx := signals.SetupSignalHandler()
	if err := controlplane.Execute(ctx); err != nil {
		logging.LoggerFromContext(ctx).Error(err)
		os.Exit(1)
	}
}
