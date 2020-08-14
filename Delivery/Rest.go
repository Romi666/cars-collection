package Delivery

import (
	"fmt"
	"github.com/gorilla/mux"
	"materi/Interface/Config"
	"materi/Interface/Repository"
	"materi/Interface/UseCase"
	"materi/Interface/appHttpUtil"
	"net/http"
)

type Server struct {
	useCase UseCase.CarUseCase
	router *mux.Router
	listeningAddress string
	apiResponse appHttpUtil.HttpResponseBuilder
}

func (s *Server) Run() {
	err := s.initRouter()
	if err!=nil{
		panic(err)
	}
	Config.Logger.Debug(fmt.Sprintf("Server run on %s", s.listeningAddress))
	if err := http.ListenAndServe(s.listeningAddress, s.router); err!=nil{
		panic(err)
	}
}

func(s *Server) initRouter() error{
	s.router.HandleFunc("/car-collection", s.carCollectionHandler).Methods(http.MethodGet)
	return nil
}

func (s *Server) carCollectionHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	coll, err:= s.useCase.GetCarCollection()
	if err!=nil{
		s.apiResponse.Error(w, 200, appHttpUtil.NewErrorMessage(100, "Can't get car collection "))
	} else {
		s.apiResponse.Data(w, 200, "Success", coll)
	}
}

func NewRestServer(c *Config.Config) CarDelivery{
	listeningAddress := fmt.Sprintf("%s:%s", c.GetConfigValue("host"), c.GetConfigValue("port"))
	carrepo := Repository.NewFakeAPIRepository(c.GetConfigValue("fake_api_url"))
	carusecase := UseCase.NewCarUseCase(carrepo)
	return &Server{
		useCase: carusecase,
		listeningAddress: listeningAddress,
		router: mux.NewRouter(),
		apiResponse: appHttpUtil.NewDefaultJSONResponder(),
	}
}
