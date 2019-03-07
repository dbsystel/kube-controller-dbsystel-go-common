package configmap

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

type ConfigMapController struct {
	Controller controller.Controller
	informer   cache.SharedIndexInformer
	kclient    *kubernetes.Clientset
}

func (cc *ConfigMapController) Run(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	go cc.informer.Run(stopCh)
	<-stopCh
}

func (cc *ConfigMapController) Initialize(kclient *kubernetes.Clientset) {

	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return kclient.CoreV1().ConfigMaps(metav1.NamespaceAll).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return kclient.CoreV1().ConfigMaps(metav1.NamespaceAll).Watch(options)
			},
		},
		&v1.ConfigMap{},
		3*time.Minute,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    cc.Controller.Create,
		UpdateFunc: cc.Controller.Update,
		DeleteFunc: cc.Controller.Delete,
	})

	cc.informer = informer
	cc.kclient = kclient

}
