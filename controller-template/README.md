# Template Kubernetes Controller
This project provides a template for controllers which watch for new kubernetes *resources* and react to them.

Actually this template supports the following kubernetes *resources*:

- ConfigMap
- Ingress

## Usage
To initialize the controller and start watching resources in kubernetes cluster:
```go
cmd/main.go
//Initialize new k8s ingress-controller from common k8s package
ingressController := &ingress.IngressController{}
ingressController.Controller = controller.New(logger)
ingressController.Initialize(k8sClient)
//Run initiated ingress-controller as go routine
go ingressController.Run(stop, wg)

//Initialize new k8s configmap-controller from common k8s package
configMapController := &configmap.ConfigMapController{}
configMapController.Controller = controller.New(logger)
configMapController.Initialize(k8sClient)
//Run initiated configmap-controller as go routine
go configMapController.Run(stop, wg)
```
To react to creation, update and deletion, implement the following functions:
```go
controller/controller.go

func (c *Controller) Create(obj interface{}) {
	level.Debug(c.logger).Log("msg", "Called create func...")
}

func (c *Controller) Update(oldobj interface{}, newobj interface{}) {
	level.Debug(c.logger).Log("msg", "Called update func...")
}

func (c *Controller) Delete(obj interface{}) {
	level.Debug(c.logger).Log("msg", "Called delete func...")
}
```

