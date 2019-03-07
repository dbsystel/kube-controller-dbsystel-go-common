package controller

type Controller interface {
	Create(obj interface{})
	Update(oldobj interface{}, newobj interface{})
	Delete(obj interface{})
}
