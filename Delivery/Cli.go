package Delivery

import (
	"fmt"
	"materi/Interface/Config"
	"materi/Interface/Entity"
	"materi/Interface/Repository"
	"materi/Interface/UseCase"
	"strings"
)

type Cli struct {
	useCase UseCase.CarUseCase
}

func (c *Cli) Run() {
	//c.PrintAllCarCollection()
	c.PrintById()
}

func (c *Cli) init(uc UseCase.CarUseCase) error{
	fmt.Println("Application Started")
	fmt.Printf("%s\n", strings.Repeat("-",40))
	c.useCase = uc
	c.PrintAllCarCollection()
	//c.PrintById()
	//c.PrintByColor()


	return nil
}

func (c *Cli) RegisterCar(car *Entity.Car){
	newCar, err := c.useCase.RegisterCar(car)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Printf("Success register : %v\n",*newCar)
}


func (c *Cli) PrintAllCarCollection(){
	coll,err := c.useCase.GetCarCollection()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", strings.Repeat("=", 55))
	fmt.Printf("%-5s %-10s %-10s %-10s %-10s\n", "No.", "ID", "Car", "Model", "Colour")
	fmt.Printf("%s\n", strings.Repeat("-", 55))
	for idx,c := range coll{
		carResult := (*c).ToString()
		car := strings.Split(carResult, "-")
		fmt.Printf("%-5d %-10s %-10s %-10s %-10s\n", idx+1, car[0], car[1], car[2], car[3])
		fmt.Printf("%s\n", strings.Repeat("-", 55))
	}
}

func (c *Cli) PrintById(){
	obj, err := c.useCase.GetCarById(81)
	if err!=nil{
		panic("Error")
	}
	fmt.Printf("%s\n", strings.Repeat("=", 55))
	fmt.Printf("%-10s %-10s %-10s %-10s\n", "ID", "Car", "Model", "Colour")
	fmt.Printf("%s\n", strings.Repeat("-", 55))
	fmt.Printf("%-10d %-10s %-10s %-10s\n", obj.CarDetail.Id, obj.CarDetail.Car, obj.CarDetail.CarModel, obj.CarDetail.CarColor)
}

func (c *Cli) UpdateColor(){
	obj, err := c.useCase.UpdateCar(81, "Red")
	if err!=nil{
		panic("Error")
	}

	fmt.Printf("%s\n", (*obj).ToString())
}

func (c *Cli) PrintByColor() {
	obj, err := c.useCase.GetByColor("Red")
	if err!=nil{
		panic("Error")
	}
	fmt.Printf("%s\n", strings.Repeat("=", 55))
	fmt.Printf("%-5s %-10s %-10s %-10s %-10s\n", "No.", "ID", "Car", "Model", "Colour")
	fmt.Printf("%s\n", strings.Repeat("-", 55))
	for idx,c := range obj{
		carResult := (*c).ToString()
		car := strings.Split(carResult, "-")
		fmt.Printf("%-5d %-10s %-10s %-10s %-10s\n", idx+1, car[0], car[1], car[2], car[3])
		fmt.Printf("%s\n", strings.Repeat("-", 55))
	}


}

func NewCliDelivery(c *Config.Config) CarDelivery  {
	Config.Logger.Debug("Run Fake API")
	carrepo := Repository.NewFakeAPIRepository(c.GetConfigValue("fake_api_url"))
	carusecase := UseCase.NewCarUseCase(carrepo)
	return &Cli{
		useCase: carusecase,
	}
}

func NewCliDeliveryTemp(c *Config.Config) CarDelivery{
	Config.Logger.Debug("Run CLI Temp")
	carrepo := Repository.NewTempRepository()
	carusecase := UseCase.NewCarUseCase(carrepo)
	cli := &Cli{
		useCase: carusecase,
	}
	car01 := Entity.Car{CarDetail: Entity.CarDetail{
		Car:          "",
		CarModel:     "Brio",
		CarColor:     "",
		CarModelYear: 1900,
		CarVin:       "",
		Price:        "",
		Availability: false,
	}}
	cli.RegisterCar(&car01)
	return cli
}