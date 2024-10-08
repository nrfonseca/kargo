package v1alpha1

import (
	"context"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetAnalysisTemplate returns a pointer to the AnalysisTemplate resource
// specified by the namespacedName argument. If no such resource is found, nil
// is returned instead.
func GetAnalysisTemplate(
	ctx context.Context,
	c client.Client,
	namespacedName types.NamespacedName,
) (*AnalysisTemplate, error) {
	at := AnalysisTemplate{}
	if err := c.Get(ctx, namespacedName, &at); err != nil {
		if err = client.IgnoreNotFound(err); err == nil {
			return nil, nil
		}
		return nil, errors.Wrapf(
			err,
			"error getting AnalysisTemplate %q in namespace %q",
			namespacedName.Name,
			namespacedName.Namespace,
		)
	}
	return &at, nil
}

func GetAnalysisRun(
	ctx context.Context,
	c client.Client,
	namespacedName types.NamespacedName,
) (*AnalysisRun, error) {
	ar := AnalysisRun{}
	if err := c.Get(ctx, namespacedName, &ar); err != nil {
		if err = client.IgnoreNotFound(err); err == nil {
			return nil, nil
		}
		return nil, errors.Wrapf(
			err,
			"error getting AnalysisRun %q in namespace %q",
			namespacedName.Name,
			namespacedName.Namespace,
		)
	}
	return &ar, nil
}
