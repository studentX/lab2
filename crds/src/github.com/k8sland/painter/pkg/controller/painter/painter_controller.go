// Copyright 2018 Imhotep Software LLC. Apache 2.0 Licence

package painter

import (
	"context"
	"log"
	"reflect"

	workloadv1alpha1 "github.com/k8sland/painter/pkg/apis/workload/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// Add creates a new Painter Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
// USER ACTION REQUIRED: update cmd/manager/main.go to call this workload.Add(mgr) to install this Controller
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("painter-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Painter
	if err = c.Watch(&source.Kind{Type: &workloadv1alpha1.Painter{}}, &handler.EnqueueRequestForObject{}); err != nil {
		return err
	}

	// Watch for pods changes
	if err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForObject{}); err != nil {
		return err
	}
	return nil
}

// ReconcilePainter reconciles a Painter object
type ReconcilePainter struct {
	client.Client
	scheme *runtime.Scheme
	colors map[string]string
}

var _ reconcile.Reconciler = &ReconcilePainter{}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcilePainter{
		Client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
		colors: map[string]string{},
	}
}

func skipSystemNS(ns string) bool {
	switch ns {
	case "kube-system", "kube-public":
		return true
	}
	return false
}

// Reconcile reads that state of the cluster for a Painter object and makes changes based on the state read
// and what is in the Painter.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Pods
// +kubebuilder:rbac:groups=apps,resources=pods,verbs=get;list;watch;create;update;patch;
// +kubebuilder:rbac:groups=workload.k8sland.io,resources=painters,verbs=get;list;watch;create;update;patch;delete
func (r *ReconcilePainter) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	var (
		res reconcile.Result
		ns  = request.NamespacedName.Namespace
	)

	// Skip system namespaces...
	if skipSystemNS(ns) {
		return res, nil
	}

	log.Println("Reconciling", request.NamespacedName)

	p, err := r.findCRD(request)
	if err == nil {
		r.colors[ns] = p.Spec.Color
		return res, r.colorPods(ns, r.colors[ns])
	}

	po, e := r.findPod(request)
	if e != nil {
		if errors.IsNotFound(e) {
			log.Println("Resetting pods color")
			delete(r.colors, ns)
			return res, r.colorPods(ns, r.colors[ns])
		}
		return res, e
	}
	return res, r.colorPod(po, r.colors[ns])
}

// Check if this is a pod event. Returns a pod or error out otherwise.
func (r *ReconcilePainter) findPod(req reconcile.Request) (corev1.Pod, error) {
	var p corev1.Pod
	return p, r.Get(context.TODO(), req.NamespacedName, &p)
}

// Check if event is a paint crd. Returns the crd or error out otherwise.
func (r *ReconcilePainter) findCRD(req reconcile.Request) (workloadv1alpha1.Painter, error) {
	var p workloadv1alpha1.Painter
	return p, r.Get(context.TODO(), req.NamespacedName, &p)
}

// Color all pods in given namespace
func (r *ReconcilePainter) colorPods(ns, color string) error {
	var pp corev1.PodList
	if err := r.List(context.TODO(), &client.ListOptions{Namespace: ns}, &pp); err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}

	for _, p := range pp.Items {
		if err := r.colorPod(p, color); err != nil {
			return err
		}
	}
	return nil
}

// Color a given pod
func (r *ReconcilePainter) colorPod(p corev1.Pod, color string) error {
	desired := p.DeepCopy()
	if color != "" {
		desired.ObjectMeta.Labels["color"] = color
	} else {
		delete(desired.ObjectMeta.Labels, "color")
	}

	if !reflect.DeepEqual(desired, p) {
		log.Printf(">> Coloring Pod %s/%s: '%s'", p.ObjectMeta.Namespace, p.ObjectMeta.Name, color)
		return r.Update(context.TODO(), desired)
	}
	return nil
}
