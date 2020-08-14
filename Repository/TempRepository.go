package Repository

import (
	"materi/Interface/Entity"
	"math/rand"
)

type TempRepository struct {
	repo []*Entity.Car
	repoOne *Entity.Car
}

func (t *TempRepository) FindByName(name string) ([]*Entity.Car, error) {
	panic("implement me")
}

func (t *TempRepository) FindByModel(model string) ([]*Entity.Car, error) {
	panic("implement me")
}

func (t *TempRepository) FindByColor(color string) ([]*Entity.Car, error) {
	return t.repo, nil
}

func (t *TempRepository) Update(car *Entity.Car) (*Entity.Car, error) {
	return t.repoOne, nil
}

func NewTempRepository() CarRepository  {
	return &TempRepository{}
}

func (t *TempRepository) FindAll() ([]*Entity.Car, error) {
	return t.repo, nil
}

func (t *TempRepository) Find(id int) (*Entity.Car, error) {
	panic("implement me")
}

func (t *TempRepository) Create(car *Entity.Car) (*Entity.Car, error) {
	car.CarDetail.Id = rand.Intn(100)
	t.repo = append(t.repo,car)
	return car,nil
}


