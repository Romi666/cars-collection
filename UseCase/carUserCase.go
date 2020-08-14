package UseCase

import (
	"materi/Interface/Entity"
	"materi/Interface/Repository"
)

type CarUseCase interface {
	RegisterCar(car *Entity.Car) (*Entity.Car,error)
	GetCarCollection() ([]*Entity.Car,error)
	GetCarById(id int) (*Entity.Car, error)
	UpdateCar(id int, color string) (*Entity.Car, error)
	GetByColor(color string)([]*Entity.Car, error)
}

type CarUseCaseImplementation struct{
	repo Repository.CarRepository
}

func (c *CarUseCaseImplementation) UpdateCar(id int, color string) (*Entity.Car, error) {
	 obj, err := c.GetCarById(id)
	 if err!=nil{
	 	panic("Error")
	 }

	 obj.CarDetail.CarColor = color

	 return obj, nil
}

func (c *CarUseCaseImplementation) GetByColor(color string) ([]*Entity.Car, error){
	obj, err := c.repo.FindByColor(color)
	if err!=nil{
		return nil, err
	}

	return obj, nil
}

func (c *CarUseCaseImplementation) GetCarById(id int) (*Entity.Car, error) {
	coll, err := c.repo.Find(id)
	if err !=nil{
		panic("Error")
	}
	return coll, nil
}

func NewCarUseCase(repo Repository.CarRepository) CarUseCase  {
	return &CarUseCaseImplementation{repo: repo}
}

func (c *CarUseCaseImplementation) RegisterCar(car *Entity.Car) (*Entity.Car, error) {
	car,err := c.repo.Create(car)
	if err != nil {
		return nil, err
	}
	return car,nil
}

func (c *CarUseCaseImplementation) GetCarCollection() ([]*Entity.Car, error) {
	coll,err := c.repo.FindAll()
	return coll, err
}