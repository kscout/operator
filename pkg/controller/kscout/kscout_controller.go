package kscout

import (
	"context"

	kscoutv1 "github.com/kscout/operator/pkg/apis/kscout/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	knv1alpha1 "knative.dev/serving/pkg/apis/serving/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_kscout")

// Add creates a KScout controller and adds it to the manager.
func Add(mgr manager.Manager) error {
	r := &ReconcileKScout{client: mgr.GetClient(), scheme: mgr.GetScheme()}

	// Create a new controller
	c, err := controller.New("kscout-controller", mgr,
		controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource KScout
	err = c.Watch(&source.Kind{Type: &kscoutv1.KScout{}},
		&handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resources and requeue the
	// owner KScout
	watchTypes := []runtime.Object{knv1alpha1.Service{}}
	for _, watchT := range watchTypes {
		err = c.Watch(&source.Kind{Type: &watchT},
			&handler.EnqueueRequestForOwner{
				IsController: true,
				OwnerType:    &kscoutv1.KScout{},
			})
		if err != nil {
			return err
		}
	}

	return nil
}

// blank assignment to verify that ReconcileKScout
// implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileKScout{}

// ReconcileKScout reconciles a KScout object
type ReconcileKScout struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a KScout object and makes
// changes based on the state read
// and what is in the KScout.Spec
// Note:
// The Controller will requeue the Request to be processed again if the
// returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work
// from the queue.
func (r *ReconcileKScout) Reconcile(request reconcile.Request) (
	reconcile.Result, error) {

	reqLogger := log.WithValues("Request.Namespace", request.Namespace,
		"Request.Name", request.Name)
	reqLogger.Info("Reconciling KScout")

	// Fetch the KScout instance
	instance := &kscoutv1.KScout{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after
			// reconcile request.
			// Owned objects are automatically garbage collected. For
			// additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define desired resources
	wantRes := []runtime.Object{}
	wantRes = append(wantRes, getCatalogAPI(instance)...)

	// Set KScout instance as the owner and controller of each resource
	for _, want := range wantRes {
		err = controllerutil.SetControllerReference(instance, want, r.scheme)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	// Check if desired resources already exists
	for _, want := range wantRes {
		objKey := types.NamespacedName{}
		if wantObjMeta, ok := want.(v1meta.ObjectMeta); !ok {
			return reconcile.Result{}, fmt.Errorf("desired object %#v "+
				"did not have v1meta.ObjectMeta", want)
		} else {
			objKey.Name = wantObjMeta.Name
			objKey.Namspace = wantObjMeta.Namespace
		}

		var found runtime.Object
		err = r.client.Get(context.TODO(), objKey, found)
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Creating a new %T", want, "Namespace",
				objKey.Namespace, "Name", objKey.Name)
			err = r.client.Create(context.TODO(), pod)
			if err != nil {
				return reconcile.Result{}, err
			}

			// Created successfully - don't requeue
			return reconcile.Result{}, nil
		} else if err != nil {
			return reconcile.Result{}, err
		}

		// Already exists - don't requeue
		reqLogger.Info("Skip reconcile: %T already exists", want,
			"Namespace", objKey.Namespace, "Name", objKey.Name)
	}

	return reconcile.Result{}, nil
}

// getCatalogAPI returns resources which run the Catalog API
func getCatalogAPI(instance kscoutv1.KScout) []runtime.Object {
	svcImg := "quay.io/kscout/catalog-api:" + instance.CatalogAPI.ImageVersion
	return []runtime.Object{
		knv1alpha1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      instance.Name + "-catalog-api",
				Namespace: instance.Namespace,
			},
			Spec: knv1alpha1.ServiceSpec{
				ConfigurationSpec: knv1alpha1.ConfigurationSpec{
					Template: knv1alpha1.RevisionTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name:      instance.Name + "-catalog-api",
							Namespace: instance.Namespace,
						},
						Spec: knv1alpha1.RevisionSpec{
							PodSpec: corev1.PodSpec{
								Containers: []corev1.Container{
									corev1.Container{
										Name:  "app",
										Image: svcImg,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
