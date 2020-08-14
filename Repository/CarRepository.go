package Repository

import "materi/Interface/Entity"

type CarRepository interface {
	FindAll() ([]*Entity.Car,error)
	Find(id int) (*Entity.Car,error)
	Create(car *Entity.Car) (*Entity.Car,error)
	Update(car *Entity.Car) (*Entity.Car,error)
	FindByName(name string) ([]*Entity.Car, error)
	FindByModel(model string) ([]*Entity.Car, error)
	FindByColor(color string) ([]*Entity.Car, error)
}
