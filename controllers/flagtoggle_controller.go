/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	featureflagsv1 "_/Users/danielobrien/Projects/launchdarkly/ld-flagtoggle-controller/api/v1"

	ldapi "github.com/launchdarkly/api-client-go"
)

// FlagToggleReconciler reconciles a FlagToggle object
type FlagToggleReconciler struct {
	client.Client
	Log    logr.Logger
	APIKey *string
}

// +kubebuilder:rbac:groups=featureflags.launchdarkly.com,resources=flagtoggles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=featureflags.launchdarkly.com,resources=flagtoggles/status,verbs=get;update;patch

func (r *FlagToggleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	kctx := context.Background()
	_ = r.Log.WithValues("flagtoggle", req.NamespacedName)

	var flag featureflagsv1.FlagToggle
	if err := r.Get(kctx, req.NamespacedName, &flag); err != nil {
		r.Log.Error(err, "unable to fetch Flag")
	}

	client := ldapi.NewAPIClient(ldapi.NewConfiguration())
	ldctx := context.WithValue(context.Background(), ldapi.ContextAPIKey, ldapi.APIKey{
		Key: *r.APIKey,
	})

	for _, env := range flag.Spec.EnvironmentKey {
		flagPath := "/environments/" + env + "/on"

		var state interface{} = flag.Spec.Enabled
		patchArr := ldapi.PatchOperation{
			Op:    "replace",
			Path:  flagPath,
			Value: &state,
		}
		patch := ldapi.PatchComment{
			Comment: flag.Spec.Comment,
			Patch:   []ldapi.PatchOperation{patchArr},
		}
		ldflag, resp, err := client.FeatureFlagsApi.PatchFeatureFlag(ldctx, flag.Spec.ProjectKey, flag.Spec.FlagKey, patch)
		if resp != nil {
			fmt.Println(resp.Body)
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ldflag)
	}
	return ctrl.Result{}, nil
}

func (r *FlagToggleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&featureflagsv1.FlagToggle{}).
		Complete(r)
}
