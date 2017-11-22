package main

import (
	"net/http"

	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	telldusLocal "telldus_exporter/telldus_local"
)

var (
	listenPort           = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	metricsPath          = flag.String("metrics-url", "/metrics", "The url used for the metrics endpoint.")
	localtellstickApiUrl = flag.String("local-tellstick-host", "", "The Tellstick Host to connect to including protocol.")
	telldusapiToken      = flag.String("telldus-token", "", "The API Token for metric collection.")
)

func main() {
	flag.Parse()

	host := telldusLocal.TellStickHost{*localtellstickApiUrl, *telldusapiToken}

	telldusLocal.PublishBuildInfo()

	exporter := telldusLocal.Exporter{
		Metrics:       telldusLocal.AddMetrics(),
		TellStickHost: host,
	}

	prometheus.MustRegister(&exporter)

	// Setup HTTP handler
	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		                <head><title>Telldus Exporter</title></head>
		                <body>
		                   <h1>Telldus Prometheus Metrics Exporter</h1>
						   <p></p>
		                   <p><a href='` + *metricsPath + `'>Metrics</a></p>
		                   </body>
		                </html>
		              `))
	})
	log.Fatal(http.ListenAndServe(*listenPort, nil))
}
