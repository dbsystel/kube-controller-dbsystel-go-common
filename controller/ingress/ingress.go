package ingress

import (
	"sync"
	"time"

	"github.com/dbsystel/kube-controller-dbsystel-go-common/controller"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type IngressController struct {
	Controller controller.Controller
	informer   cache.SharedIndexInformer
	kclient    *kubernetes.Clientset
}

func (ic *IngressController) Run(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	go ic.informer.Run(stopCh)
	<-stopCh
}

func (ic *IngressController) Initialize(kclient *kubernetes.Clientset) {

	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return kclient.ExtensionsV1beta1().Ingresses(metav1.NamespaceAll).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return kclient.ExtensionsV1beta1().Ingresses(metav1.NamespaceAll).Watch(options)
			},
		},
		&v1beta1.Ingress{},
		3*time.Minute,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    ic.Controller.Create,
		UpdateFunc: ic.Controller.Update,
		DeleteFunc: ic.Controller.Delete,
	})

	ic.informer = informer
	ic.kclient = kclient

}
