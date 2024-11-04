package afr_kafka

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var chan_threshold_Control = make(chan int, 450)

func (uc *UserControl) Init_Cache() (err error) {
	log.Println("Initializing cache map")
	//MapAPIUsers.Initialize("APIUsers", uc, "APIUsers", true, DB_NAME, "Col_APIUsers")
	return nil
}

func (uc *UserControl) Init_DAO() (err error) {
	log.Println("Initializing DAO")

	return nil
}

func (Uc *UserControl) Post_SSR_totarget(load []byte) {
	post_start := time.Now()
	TransactionsCount.With(prometheus.Labels{"Stream": "SSR", "Type": "Post request", "Description": "SSR Post request"}).Inc()
	//post to target
	chan_threshold_Control <- 1
	//url := "http://" + Configuration.CVM_CDRV_IP + ":" + Configuration.CVM_CDRV_Port + "/API/BonusAllocation"
	url := Configuration.TargetProtocol + "://" + Configuration.TargetIP + ":" + Configuration.TargetPort + Configuration.TargetURI
	//log.Println(url)
	method := "POST"
	//client := &http.Client{}
	client := &http.Client{
		Timeout: 5 * time.Second, //is SMSC not reachable request will time out after "TimeOut" sec
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(load))
	if err != nil {
		log.Println("Error in http post 1: " + err.Error())
		EndToEndLatency.With(prometheus.Labels{"host": Configuration.TargetIP, "path": "", "status": "Failed"}).Observe(float64((time.Since(post_start).Nanoseconds()) / 1000000))
		TransactionsCount.With(prometheus.Labels{"Stream": "SSR", "Type": "Post failed", "Description": "SSR Post failed"}).Inc()
		<-chan_threshold_Control
		<-chan_SSR_controler
		return
	}
	req.Header.Add("Content-Type", "application/xml")
	req.Header.Set("Connection", "close")
	_, err = client.Do(req)
	if err != nil {
		log.Println("Error in http post 2: " + err.Error())
		EndToEndLatency.With(prometheus.Labels{"host": Configuration.TargetIP, "path": "", "status": "Failed"}).Observe(float64((time.Since(post_start).Nanoseconds()) / 1000000))
		TransactionsCount.With(prometheus.Labels{"Stream": "SSR", "Type": "Post failed", "Description": "SSR Post failed"}).Inc()
		<-chan_threshold_Control
		<-chan_SSR_controler
		return
	}
	EndToEndLatency.With(prometheus.Labels{"host": Configuration.TargetIP, "path": "", "status": "Successful"}).Observe(float64((time.Since(post_start).Nanoseconds()) / 1000000))
	TransactionsCount.With(prometheus.Labels{"Stream": "SSR", "Type": "Post successful", "Description": "SSR Post successful"}).Inc()
	<-chan_threshold_Control
	<-chan_SSR_controler
}
