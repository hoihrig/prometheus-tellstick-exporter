package telldus_local

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
)

func PublishBuildInfo() {
	buildInfo := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "telldus_exporter_build_info",
		Help: "telldus_exporter build_info",
	}, []string{"version", "golang_version"})
	buildInfo.WithLabelValues("0.1", runtime.Version()).Set(1)
	prometheus.MustRegister(buildInfo)
}

func gatherData(host TellStickHost) SensorList {
	var list SensorList

	getSensorList(host, &list)

	for index, item := range list.Sensor {
		list.Sensor[index].Data = getSensorData(host, item.ID)
		log.Info(list.Sensor[index].Data)
	}

	return list
}

func getSensorData(host TellStickHost, id int) []SensorData {

	body, err := makeRequest(host, "sensor/info?id="+strconv.Itoa(id))

	if err != nil {
		log.Error(err)
	}

	si := SensorInfo{}
	json.Unmarshal(body, &si)

	return si.Data
}

func getSensorList(host TellStickHost, list *SensorList) {
	body, err := makeRequest(host, "sensors/list")

	if err != nil {
		log.Error(err)
	}

	json.Unmarshal(body, list)
}

func getDeviceCount(host TellStickHost) float64 {
	body, err := makeRequest(host, "devices/list")

	if err != nil {
		log.Error(err)
	}

	dl := DeviceList{}
	json.Unmarshal(body, &dl)

	return float64(len(dl.Device))
}

func makeRequest(host TellStickHost, endpoint string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", host.Address+"/api/"+endpoint, nil)
	if err != nil {
		log.Error("Could not create http request", err)
	}

	req.Header.Add("Authorization", `Bearer `+host.Token)
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
