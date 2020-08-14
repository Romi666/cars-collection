package Repository

import (
	"encoding/json"
	"fmt"
	"materi/Interface/Entity"
	"math/rand"
	"net/http"
)

type FakeApiRepository struct {
	url string
	defaultClient *http.Client
}

func (f FakeApiRepository) FindByName(name string) ([]*Entity.Car, error) {
	panic("implement me")
}

func (f FakeApiRepository) FindByModel(model string) ([]*Entity.Car, error) {
	panic("implement me")
}

func (f FakeApiRepository) FindByColor(color string) ([]*Entity.Car, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cars", f.url), nil)
	if err!=nil{
		return nil, err
	}
	resp, err := (*f.defaultClient).Do(req)
	if err!=nil{
		return nil, err
	}
	defer resp.Body.Close()
	var carResponse Entity.Cars
	json.NewDecoder(resp.Body).Decode(&carResponse)
	var result []*Entity.Car
	for _, val := range carResponse.Cars{
		if val.CarColor == color{
			result = append(result, &Entity.Car{CarDetail: Entity.CarDetail{
				Id : val.Id,
				Car: val.Car,
				CarModel: val.CarModel,
				CarColor: val.CarColor,

			}})
		}

	}
	return result, nil
}

func (f FakeApiRepository) FindAll() ([]*Entity.Car, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cars", f.url), nil)
	if err!=nil{
		return nil, err
	}
	resp, err := (*f.defaultClient).Do(req)
	if err!=nil{
		return nil, err
	}
	defer resp.Body.Close()
	var carResponse Entity.Cars
	json.NewDecoder(resp.Body).Decode(&carResponse)
	var result []*Entity.Car
	for _, val := range carResponse.Cars{
		result = append(result, &Entity.Car{CarDetail: Entity.CarDetail{
			Id : val.Id,
			Car: val.Car,
			CarModel: val.CarModel,
			CarColor: val.CarColor,
			CarModelYear: val.CarModelYear,
			CarVin: val.CarVin,
			Price: val.Price,
			Availability: val.Availability,

		}})
	}
	return result, nil
}

func (f FakeApiRepository) Find(id int) (*Entity.Car, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cars/%d", f.url, id), nil)
	if err!=nil{
		return nil, err
	}

	resp, err := (*f.defaultClient).Do(req)
	if err!=nil{
		return nil, err
	}

	defer resp.Body.Close()

	var carResponse Entity.Car


	json.NewDecoder(resp.Body).Decode(&carResponse)
	//fmt.Println("TEST"+carResponse.CarDetail.CarColor)
	return &carResponse, nil
}

func (f FakeApiRepository) Create(car *Entity.Car) (*Entity.Car, error) {
	car.CarDetail.Id = rand.Intn(100)
	return car, nil
}

func (f FakeApiRepository) Update(car *Entity.Car) (*Entity.Car, error) {
	panic("implement me")
}

func NewFakeAPIRepository(url string) CarRepository{
	return &FakeApiRepository{
		url,&http.Client{

		},
	}
}