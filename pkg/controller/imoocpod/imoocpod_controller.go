package imoocpod

import (
	"context"
	"fmt"

	k8sv1alpha1 "github.com/imooc-com/imooc-operator/pkg/apis/k8s/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_imoocpod")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new ImoocPod Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileImoocPod{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("imoocpod-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ImoocPod
	err = c.Watch(&source.Kind{Type: &k8sv1alpha1.ImoocPod{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ImoocPod
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &k8sv1alpha1.ImoocPod{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileImoocPod implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileImoocPod{}

// ReconcileImoocPod reconciles a ImoocPod object
type ReconcileImoocPod struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a ImoocPod object and makes changes based on the state read
// and what is in the ImoocPod.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileImoocPod) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ImoocPod")

	// 定义返回结果，定时执行
	result := reconcile.Result{
		// Requeue:      true,
		RequeueAfter: 1000 * 60 * 2,
	}

	// Fetch the ImoocPod instance
	instance := &k8sv1alpha1.ImoocPod{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return result, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	fmt.Println("watchedNamespace", instance.Spec.WatchedNamespace)

	// 从 instance 中获取
	watchedNamespace := instance.Spec.WatchedNamespace

	// 获取所有存在的 pod
	existingPods := &corev1.PodList{}
	err = r.client.List(
		context.TODO(),
		existingPods,
		&client.ListOptions{
			Namespace: watchedNamespace,
		},
	)
	if err != nil {
		reqLogger.Error(err, "获取已经存在的 pod 失败")
		return result, nil
	}
	var existingPodNames []string
	for _, pod := range existingPods.Items {
		existingPodNames = append(existingPodNames, pod.GetObjectMeta().GetName())
	}
	fmt.Println(existingPodNames)

	// 获取所有存在 deployment
	existingDeployment := &appsv1.DeploymentList{}
	err = r.client.List(
		context.TODO(),
		existingDeployment,
		&client.ListOptions{
			Namespace: watchedNamespace,
		},
	)
	var existingDeploymentNames []string
	for _, deployment := range existingDeployment.Items {
		existingDeploymentNames = append(existingDeploymentNames, deployment.GetObjectMeta().GetName())
	}
	fmt.Println(existingDeploymentNames)

	//获取所有的 service
	existingService := &corev1.ServiceList{}
	r.client.List(
		context.TODO(),
		existingService,
		&client.ListOptions{
			Namespace: watchedNamespace,
		},
	)
	for _, service := range existingService.Items {
		fmt.Println("service:", service.GetObjectMeta().GetName())
	}

	// lbls := labels.Set{
	// 	"app": instance.Name,
	// }
	// existingPods := &corev1.PodList{}
	// err = r.client.List(
	// 	context.TODO(),
	// 	existingPods,
	// 	&client.ListOptions{
	// 		Namespace:     request.Namespace,
	// 		LabelSelector: labels.SelectorFromSet(lbls),
	// 	},
	// )
	// for _, pod := range existingPods.Items {
	// 	fmt.Printf("pod: %+v", pod)
	// }

	/*
		// 通过自定义的资源 ImoocPod 获取已经存在的 pod
		lbls := labels.Set{
			"app": instance.Name,
		}
		existingPods := &corev1.PodList{}
		err = r.client.List(
			context.TODO(),
			existingPods,
			&client.ListOptions{
				Namespace: request.Namespace,
				LabelSelector: labels.SelectorFromSet(lbls),
			},
		)
		if err != nil {
			reqLogger.Error(err, "获取已经存在的 pod 失败")
			return reconcile.Result{}, nil
		}

		// 获取 pod 的名称
		var existingPodNames []string
		for _, pod := range existingPods.Items {
			if pod.GetObjectMeta().GetDeletionTimestamp() != nil {
				continue
			}
			if pod.Status.Phase == corev1.PodPending || pod.Status.Phase == corev1.PodRunning {
				existingPodNames = append(existingPodNames, pod.GetObjectMeta().GetName())
			}
		}
	*/

	// Define a new Pod object
	pod := newPodForCR(instance)

	// Set ImoocPod instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Pod created successfully - don't requeue
		return result, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return result, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *k8sv1alpha1.ImoocPod) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
