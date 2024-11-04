package main

import (
	gw "afr_DRC_INLiveFeed_consumer_SSR_ActivationPromoEast/afr_kafka"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kardianos/service"
)

const ApplicationName = "AFR INLiveFeed Consumer Activation Promo East"
const ApplicationReleaseNumber = "0.1.0"
const ApplicationReleaseDate = "22/08/2024" //"dd/MM/YYYY"

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	log.Println("-------------------------------------------------------------------")
	log.Println("Application name: ", ApplicationName)
	log.Println("Application release number: ", ApplicationReleaseNumber)
	log.Println("Application release date: ", ApplicationReleaseDate)
	log.Println("-------------------------------------------------------------------")

	//read and parse configuration
	err := gw.GetDefaultConfiguration()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Establishing connections...")
	UserControl := gw.NewUserControl()
	time.Sleep(5 * time.Second)
	UserControl.Init_Cache()
	UserControl.Init_DAO()

	gw.Init_Prometheus_Metrics()
	go gw.PortlinkInquiry()
	go gw.Reset_Prometheus_Metrics()
	//latency prometheus
	gw.Init_Prometheus_Metrics_Latency()
	go gw.Reset_Prometheus_Metrics_Latency()

	//start the topic readers

	time.Sleep(15 * time.Second)

	go UserControl.KafkaClient.RunKafkaReader_SSR_4()
	go UserControl.KafkaClient.RunKafkaReader_SSR_26()

	///
	go UserControl.SSR_Process()
	go UserControl.SSR_ReadersStats()

	//Add user routers to the web service
	log.Println("Add routers to the web service")
	router := mux.NewRouter().StrictSlash(true)
	UserControl.AddToRouter(router, UserControl)
	log.Println("HTTP listen and serve on port " + gw.Configuration.HttpServicePort)
	log.Fatal(http.ListenAndServe(":"+gw.Configuration.HttpServicePort, router))
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        ApplicationName,
		DisplayName: ApplicationName,
		Description: ApplicationName + " service",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
