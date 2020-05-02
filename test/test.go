package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		//for {
		opsProcessed.Inc()
		//	time.Sleep(2 * time.Second)
		//}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)
var (
	messageChannel = make(chan string)
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	port1 := 8000
	port2 := 8001
	finish := make(chan bool)

	http.Handle("/metrics", promhttp.Handler())

	server8001 := http.NewServeMux()
	server8001.HandleFunc("/foo", foo8001)
	server8001.HandleFunc("/bar", bar8001)
	recordMetrics()

	go func() {
		log.Printf("Starting web server metrics on port %d", port2)
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port2), nil))
	}()
	go func() {
		log.Printf("Starting web server port %d", port1)
		http.Handle('/metrics', promhttp.Handler())
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port1), server8001))
	}()

	<-finish
}

func logger(text string) {
	log.Print("hit" + text)
}

func channelLogger(c chan string) {
	log.Printf("channel says %s", <-c)
}

func foo8001(w http.ResponseWriter, r *http.Request) {
	recordMetrics()
	w.Write([]byte("Listening on 8001: foo "))
}

func bar8001(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on 8001: bar "))
}
