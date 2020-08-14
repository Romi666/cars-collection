package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"materi/Interface/Config"
	"materi/Interface/Delivery"
	"os"
)


const (
	AppName    = "Enigma Car Collection"
	AppTagLine = "Car Collection Sample Project"
	Version    = "0.0.1"
)

type app struct {
 	appConfig *Config.Config
}

func (a app) runServer(){
	a.run(Delivery.NewRestServer(a.appConfig))
}

func (a app) runFakeApi(){
	a.run(Delivery.NewCliDelivery(a.appConfig))
}


func (a app) runTemp(){
	a.run(Delivery.NewCliDeliveryTemp(a.appConfig))
}

func (a app) run(delivery Delivery.CarDelivery){
	delivery.Run()
}

func newApp(configPath string) *app {
	return &app{appConfig: Config.NewConfig(configPath)}
}

func main()  {
	var language string
	appConfig := &cli.App{
		Name: AppName,
		Usage: AppTagLine,
		Version: Version,
		Action: func(c *cli.Context)error{
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config c",
				Usage: "Load configuration from `FILE`",
			},&cli.StringFlag{
				Name: "lang",
				Value: "english",
				Usage: "language for the greeting",
				Destination: &language,
			},
		},

		Commands: []*cli.Command{
			{
				Name:    "runfakeapi",
				Aliases: []string{"f"},
				Usage:   "Run with Fake API",
				Action: func(c *cli.Context) error {
					newApp(c.String("config")).runFakeApi()
					return nil
				},
			},
			{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "Run with Server API",
				Action: func(c *cli.Context) error {
					newApp(c.String("config")).runServer()
					return nil
				},
			},
			{
				Name:    "runtemp",
				Aliases: []string{"t"},
				Usage:   "Run with Temporary Slice",
				Action: func(c *cli.Context) error {
					newApp(c.String("config")).runTemp()
					return nil
				},
			},
		},

	}
	err := appConfig.Run(os.Args)
	if err!=nil{
		log.Fatal(err)
	}
}