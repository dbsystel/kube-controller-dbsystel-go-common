package service

import (
	"sync"
	"time"

	"github.com/dbsystel/kube-controller-dbsystel-go-common/controller"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type ServiceController struct {
	Controller controller.Controller
	informer   cache.SharedIndexInformer
	kclient    *kubernetes.Clientset
}

func (sc *ServiceController) Run(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	go sc.informer.Run(stopCh)
	<-stopCh
}

func (sc *ServiceController) Initialize(kclient *kubernetes.Clientset) {

	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return kclient.CoreV1().Service(metav1.NamespaceAll).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return kclient.CoreV1().Service(metav1.NamespaceAll).Watch(options)
			},
		},
		&v1.Service{},
		3*time.Minute,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    sc.Controller.Create,
		UpdateFunc: sc.Controller.Update,
		DeleteFunc: sc.Controller.Delete,
	})

	sc.informer = informer
	sc.kclient = kclient

}
