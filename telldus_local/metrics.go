package telldus_local

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

func AddMetrics() map[string]*prometheus.Desc {

	Metrics := make(map[string]*prometheus.Desc)

	Metrics["Sensors"] = prometheus.NewDesc(
		prometheus.BuildFQName("telldus", "sensors", "count"),
		"Number of Sensors",
		[]string{}, nil,
	)

	Metrics["Devices"] = prometheus.NewDesc(
		prometheus.BuildFQName("telldus", "devices", "count"),
		"Number of Devices",
		[]string{}, nil,
	)

	Metrics["Metric"] = prometheus.NewDesc(
		prometheus.BuildFQName("telldus", "sensors", "metric"),
		"Value of the Sensor",
		[]string{"Sensorname", "type", "unit"}, nil,
	)

	log.Info("Metric Descriptions added!")

	return Metrics

}

// processMetrics - processes the response data and sets the metrics using it as a source
func (e *Exporter) processMetrics(sl *SensorList, dc float64, ch chan<- prometheus.Metric) error {

	for _, s := range sl.Sensor {
		if hasData(s) {
			for _, sd := range s.Data {
				if s.Name != "" {
					v := sd.Value
					ch <- prometheus.MustNewConstMetric(e.Metrics["Metric"], prometheus.GaugeValue, v, s.Name, sd.Name, unitLookup(sd.Name, sd.Scale))
				}
			}
		}
	}

	sc := float64(len(sl.Sensor))

	// Set stats
	ch <- prometheus.MustNewConstMetric(e.Metrics["Sensors"], prometheus.GaugeValue, sc)
	ch <- prometheus.MustNewConstMetric(e.Metrics["Devices"], prometheus.GaugeValue, dc)

	return nil
}

func hasData(si SensorInfo) bool {
	if len(si.Data) == 0 {
		return false
	} else {
		return true
	}
}
