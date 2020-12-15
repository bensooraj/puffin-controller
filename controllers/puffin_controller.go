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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	birdsv1beta1 "github.com/bensooraj/puffin-controller/api/v1beta1"
)

// PuffinReconciler reconciles a Puffin object
type PuffinReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	// My changes
	Logged bool
}

// +kubebuilder:rbac:groups=birds.bensooraj.com,resources=puffins,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=birds.bensooraj.com,resources=puffins/status,verbs=get;update;patch

func (r *PuffinReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	var err error
	ctx := context.Background()
	log := r.Log.WithValues("puffin", req.NamespacedName)

	if !r.Logged {
		log.Info("Setting logged to true")
		r.Logged = true
	}

	// your logic here
	puffin := &birdsv1beta1.Puffin{}
	err = r.Get(ctx, req.NamespacedName, puffin)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request - return and don't requeue:
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request:
		return ctrl.Result{}, err
	}

	if puffin.Status.Message == "" {
		puffin.Status.Message = birdsv1beta1.ColorPhasePending
	}

	switch puffin.Status.Message {
	case birdsv1beta1.ColorPhasePending:
		log.Info(fmt.Sprintf("[PENDING] %s will be assigned a color shortly!\n", puffin.Name))
		// Assign the color
		puffin.Status.Message = birdsv1beta1.ColorPhaseColored
	case birdsv1beta1.ColorPhaseColored:
		log.Info(fmt.Sprintf("[COLORED] %s is assigned the color %s!\n", puffin.Name, puffin.Spec.Color))
		return ctrl.Result{}, nil
	default:
		log.Info(fmt.Sprintf("[NOP] %s has nothing to do!\n", puffin.Name))
		return ctrl.Result{}, nil
	}

	// Update the Puffinf instance, setting the status to the respective phase:
	err = r.Update(ctx, puffin)
	// err = r.Status().Update(ctx, puffin)
	if err != nil {
		log.Error(err, "Error updating the status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *PuffinReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&birdsv1beta1.Puffin{}).
		Complete(r)
}
