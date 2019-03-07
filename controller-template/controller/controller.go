package controller

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Controller struct {
	logger log.Logger
}

func (c *Controller) Create(obj interface{}) {
	level.Debug(c.logger).Log("msg", "Called create func...")
	//configmapObj := obj.(*v1.ConfigMap)
	//level.Debug(c.logger).Log("name: ", configmapObj.Name, "namespace: ", configmapObj.Namespace)
}

func (c *Controller) Update(oldobj interface{}, newobj interface{}) {
	level.Debug(c.logger).Log("msg", "Called update func...")
}

func (c *Controller) Delete(obj interface{}) {
	level.Debug(c.logger).Log("msg", "Called delete func...")
}

func New(logger log.Logger) *Controller {
	controller := &Controller{}
	controller.logger = logger
	return controller
}
