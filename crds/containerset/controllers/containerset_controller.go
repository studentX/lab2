/*
Copyright 2019 K8sland Training.

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

	"github.com/go-logr/logr"
	workloadv1alpha1 "github.com/k8sland.io/crds/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// ContainerSetReconciler reconciles a ContainerSet object
type ContainerSetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

func ignoreNotFound(err error) error {
	if apierrs.IsNotFound(err) {
		return nil
	}
	return err
}

func genDeployment(crd *workloadv1alpha1.ContainerSet, s *runtime.Scheme) (*appsv1.Deployment, error) {
	count := int32(crd.Spec.Replicas)
	dp := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crd.Name,
			Namespace: crd.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"cs": crd.Name},
			},
			Replicas: &count,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"cs": crd.Name},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "c1",
							Image: crd.Spec.Image,
						},
					},
				},
			},
		},
	}
	if err := controllerutil.SetControllerReference(crd, &dp, s); err != nil {
		return nil, err
	}

	return &dp, nil
}

// Reconcile ensures desired == current.
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=workload.k8sland.io,resources=containersets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=workload.k8sland.io,resources=containersets/status,verbs=get;update;patch
func (r *ContainerSetReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("containerset", req.NamespacedName)
	log.V(0).Info("\n\n>>> Reconciling container set")

	var (
		res ctrl.Result
		cs  workloadv1alpha1.ContainerSet
		ctx = context.Background()
	)
	err := r.Get(ctx, req.NamespacedName, &cs)
	if err != nil {
		return res, ignoreNotFound(err)
	}

	ndp, err := genDeployment(&cs, r.Scheme)
	if err != nil {
		return res, err
	}

	var (
		dp  appsv1.Deployment
		fqn = types.NamespacedName{Name: ndp.Name, Namespace: ndp.Namespace}
	)
	// Fetch existing deployment.
	err = r.Get(ctx, fqn, &dp)
	if err != nil && !errors.IsNotFound(err) {
		return res, err
	}

	// No associated deployment found, created it.
	if apierrs.IsNotFound(err) {
		log.V(1).Info("Creating deployment", "fqn", fqn)
		// Create new deployment
		if err = r.Create(ctx, ndp); err != nil {
			return res, err
		}
		return res, r.updateStatus(ctx, &cs, ndp)
	}

	// Check if we have deltas between desired vs current if so update deployment.
	if matchDesiredState(log, ndp.Spec, dp.Spec) {
		return res, r.updateStatus(ctx, &cs, &dp)
	}

	log.V(1).Info("CRD Changed", "fqn", fqn, "cresp", dp.Spec.Replicas, "dresp", ndp.Spec.Replicas)
	*dp.Spec.Replicas = *ndp.Spec.Replicas
	dp.Spec.Template.Spec.Containers[0].Image = ndp.Spec.Template.Spec.Containers[0].Image

	log.V(1).Info("Updating DP", "name", dp.Name)
	if err := r.Update(ctx, &dp); err != nil {
		log.Error(err, "Update DP failed")
		return res, err
	}

	return res, r.updateStatus(ctx, &cs, &dp)
}

func (r *ContainerSetReconciler) updateStatus(ctx context.Context, cs *workloadv1alpha1.ContainerSet, dp *appsv1.Deployment) error {
	log := r.Log.WithValues("containerset", cs.Name)

	log.V(1).Info("Check CS Status Reps", "old", cs.Status.HealthyReplicas, "new", dp.Status.ReadyReplicas)
	if cs.Status.HealthyReplicas == dp.Status.ReadyReplicas {
		return nil
	}
	log.V(1).Info("Updating CS status", "name", cs.Name)
	cs.Status.HealthyReplicas = dp.Status.ReadyReplicas

	return r.Update(ctx, cs)
}

// SetupWithManager adds this controller to the manager.
func (r *ContainerSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&workloadv1alpha1.ContainerSet{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}

func matchDesiredState(log logr.Logger, dSpec, cSpec appsv1.DeploymentSpec) bool {
	return *dSpec.Replicas == *cSpec.Replicas && dSpec.Template.Spec.Containers[0].Image == cSpec.Template.Spec.Containers[0].Image
}
