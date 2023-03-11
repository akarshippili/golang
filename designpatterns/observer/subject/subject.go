package subject

import "github.com/akarshippili/golang/designpatterns/observer/observers"

type Subject interface {
	Register(observer observers.Investor)
	Remove(observers.Investor)
	Notify()
}
