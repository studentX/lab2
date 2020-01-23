/*
Copyright 2020 K8sland Training.

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
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"

	paintv1alpha1 "github.com/k8sland/crds/api/v1alpha1"
)

const cdFinalizer = "clusterdepot.finalizers.paint.k8sland.io"

// ClusterDepotReconciler reconciles a ClusterDepot object
type ClusterDepotReconciler struct {
	client.Client
	Log     logr.Logger
	Scheme  *runtime.Scheme
	watcher podWatcher
	color   string
}

// +kubebuilder:rbac:groups=paint.k8sland.io,resources=clusterdepots,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=paint.k8sland.io,resources=clusterdepots/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=list;watch;update

func (r *ClusterDepotReconciler) registerOrRunFinalizer(ctx context.Context, cd paintv1alpha1.ClusterDepot) (bool, error) {
	if cd.DeletionTimestamp.IsZero() {
		if !strContains(cd.ObjectMeta.Finalizers, cdFinalizer) {
			cd.ObjectMeta.Finalizers = append(cd.ObjectMeta.Finalizers, cdFinalizer)
			return true, r.Update(ctx, &cd)
		}
		return false, nil
	}

	if strContains(cd.ObjectMeta.Finalizers, cdFinalizer) {
		if _, err := r.paintPods(ctx, cd.Namespace, nil); err != nil {
			return false, err
		}
		cd.ObjectMeta.Finalizers = strRemove(cd.ObjectMeta.Finalizers, cdFinalizer)
		return true, r.Update(ctx, &cd)
	}

	return false, nil
}

// Reconcile
func (r *ClusterDepotReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	l := r.Log.WithValues("clusterdepot", req.NamespacedName)

	ctx := context.Background()
	var res ctrl.Result

	var cd paintv1alpha1.ClusterDepot
	if err := r.Get(context.Background(), req.NamespacedName, &cd); err != nil {
		return res, client.IgnoreNotFound(err)
	}

	updated, err := r.registerOrRunFinalizer(ctx, cd)
	if err != nil {
		return res, err
	}
	if updated {
		return res, nil
	}

	l.Info("Current Color", "color", cd.Spec.Color)
	painted, err := r.paintPods(ctx, req.Namespace, &cd.Spec.Color)
	if err != nil {
		return res, err
	}

	return ctrl.Result{}, r.updateStatus(ctx, cd, painted)
}

func (r *ClusterDepotReconciler) paintPods(ctx context.Context, ns string, color *string) (int32, error) {
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
			delete(p.Labels, "color")
		} else {
			c, ok := p.Labels["color"]
			if ok && c == *color {
				continue
			}
			r.Log.Info("Painting pod", "color", *color, "namespace", p.Namespace, "name", p.Name)
			p.Labels["color"] = *color
		}

		if err := r.Update(ctx, &p); err != nil {
			r.Log.Error(err, "Pod Painting Failed", "namespace", p.Namespace, "name", p.Name)
			continue
		}
	}

	return painted, nil
}

func (r *ClusterDepotReconciler) updateStatus(ctx context.Context, c paintv1alpha1.ClusterDepot, count int32) error {
	if c.Status.PaintedPods == count {
		return nil
	}

	r.Log.Info("CD-UPDATE", "painted", count)
	c.Status.PaintedPods = count
	return r.Status().Update(ctx, &c)
}

func (r *ClusterDepotReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&paintv1alpha1.ClusterDepot{}).
		// Watches(&source.Kind{Type: &v1.Pod{}}, &handler.EnqueueRequestForObject{}).
		// Watches(&source.Kind{Type: &v1.Pod{}}, r.watcher).
		Complete(r)
}

type podWatcher struct{}

func (w podWatcher) Create(evt event.CreateEvent, wk workqueue.RateLimitingInterface) {
	// fmt.Printf("POD ADDED %#v\n", evt)
}
func (w podWatcher) Update(evt event.UpdateEvent, wk workqueue.RateLimitingInterface) {
	// fmt.Printf("POD UPDATED %#v\n", evt)
}
func (w podWatcher) Delete(evt event.DeleteEvent, wk workqueue.RateLimitingInterface) {
	// fmt.Printf("POD DELETED %#v\n", evt)
}
func (w podWatcher) Generic(event.GenericEvent, workqueue.RateLimitingInterface) {}

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
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
