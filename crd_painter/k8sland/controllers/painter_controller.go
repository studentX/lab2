/*
Copyright 2020 K8sland Labs.

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
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterdepotv1alpha1 "github.com/k8sland/crds/painter/api/v1alpha1"
)

const painterFinalizer = "painter.finalizers.clusterdepot.k8sland.io"

// PainterReconciler reconciles a Painter object
type PainterReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=clusterdepot.k8sland.io,resources=painters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=clusterdepot.k8sland.io,resources=painters/status,verbs=get;update;patch
!!YOUR_CODE!! Add your rbac policies for your controller

// Reconcile
func (r *PainterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	l := r.Log.WithValues("painter", req.NamespacedName)

	ctx := context.Background()
	var res ctrl.Result

	var p clusterdepotv1alpha1.Painter
	if err := r.Get(context.Background(), req.NamespacedName, &p); err != nil {
		return res, client.IgnoreNotFound(err)
	}

	updated, err := r.registerOrRunFinalizer(ctx, p)
	if err != nil {
		return res, err
	}
	if updated {
		return res, nil
	}

	l.Info("Current Color", "color", p.Spec.Color)
	painted, err := r.paintPods(ctx, req.Namespace, &p.Spec.Color)
	if err != nil {
		return res, err
	}

	return ctrl.Result{}, r.updateStatus(ctx, p, painted)
}

func (r *PainterReconciler) registerOrRunFinalizer(ctx context.Context, p clusterdepotv1alpha1.Painter) (bool, error) {
	if p.DeletionTimestamp.IsZero() {
		if !strContains(p.ObjectMeta.Finalizers, painterFinalizer) {
			p.ObjectMeta.Finalizers = append(p.ObjectMeta.Finalizers, painterFinalizer)
			return true, r.Update(ctx, &p)
		}
		return false, nil
	}

	if strContains(p.ObjectMeta.Finalizers, painterFinalizer) {
		if _, err := r.paintPods(ctx, p.Namespace, nil); err != nil {
			return false, err
		}
		p.ObjectMeta.Finalizers = strRemove(p.ObjectMeta.Finalizers, painterFinalizer)
		return true, r.Update(ctx, &p)
	}

	return false, nil
}

func (r *PainterReconciler) paintPods(ctx context.Context, ns string, color *string) (int32, error) {
	var pp v1.PodList
	err := r.List(ctx, &pp, client.InNamespace(ns))
	if err != nil {
		return 0, err
	}

	var painted int32
	for _, p := range pp.Items {
		painted++
		if color == nil {
			r.Log.Info("Unpainting pod", "namespace", p.Namespace, "name", p.Name)
			!!YOUR_CODE!! Unpaint the pod
		} else {
			!!YOUR_CODE!! Paint the pod iff it does is not already painted with the correct color!
			r.Log.Info("Painting pod", "color", *color, "namespace", p.Namespace, "name", p.Name)
		}

		if err := r.Update(ctx, &p); err != nil {
			r.Log.Error(err, "Pod Painting Failed", "namespace", p.Namespace, "name", p.Name)
			continue
		}
	}

	return painted, nil
}

func (r *PainterReconciler) updateStatus(ctx context.Context, p clusterdepotv1alpha1.Painter, count int32) error {
	if p.Status.PaintedPods == count {
		return nil
	}

	r.Log.Info("p-UPDATE", "painted", count)
	!!YOUR_CODE!! Update the CRD status to reflect the number of painted pods
	return r.Status().Update(ctx, &p)
}

func (r *PainterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&clusterdepotv1alpha1.Painter{}).
		Complete(r)
}

// ----------------------------------------------------------------------------
// Helpers...

func strContains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func strRemove(slice []string, s string) (result []string) {
	victim := -1
	for i, item := range slice {
		if item == s {
			victim = i
			break
		}
	}

	if victim == -1 {
		return slice
	}
	return append(slice[:victim], slice[victim+1:]...)
}
