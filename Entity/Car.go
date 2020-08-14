package Entity

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Car struct {
	CarDetail CarDetail
}

func (c *Car) ToString() string {
	return fmt.Sprintf("%d-%s-%s-%s", c.CarDetail.Id, c.CarDetail.Car, c.CarDetail.CarModel, c.CarDetail.CarColor)
}

type CarDetail struct {
	Id           int    `json:"id"`
	Car          string `json:"car"`
	CarModel     string `json:"car_model"`
	CarColor     string `json:"car_color"`
	CarModelYear int    `json:"car_model_year"`
	CarVin       string `json:"car_vin"`
	Price        string `json:"price"`
	Availability bool   `json:"availability"`
}

type Cars struct {
	Cars []CarDetail
}

func(cd *Car) Validate() error{
	return validation.ValidateStruct(cd,
		validation.Field(&cd.CarDetail.Car, validation.Required.Error("Nama Mobil Kosong")))
}