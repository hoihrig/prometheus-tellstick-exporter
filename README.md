# Tellstick Local Exporter
[![Circle CI](https://circleci.com/gh/hoihrig/prometheus-tellstick-exporter.svg?style=shield)](https://circleci.com/gh/hoihrig/prometheus-tellstick-exporter)

Prometheus exporter collecting Sensor information from TellDus TellStick devices having a local API, e.g. TellStick Net v2 and TellStick ZNet v1/2

It currently only works with the local API and TellDus Live is not supported.

## Obtaining the local API token

	1. curl -X PUT http://<TellStick_IP>/api/token -d 'app="exporter"'
	2. Open the URL in the authUrl field from Step 1
	3. curl -i -X GET http://<TellStick_IP>/api/token?token=<token_Step_1>
	4. Use the token received in Step 3

## Building and running

Prerequisites:

* [Go compiler](https://golang.org/dl/)

Building:

    go get github.com/hoihrig/prometheus-tellstick-exporter
    cd ${GOPATH-$HOME/go}/src/github.com/hoihrig/prometheus-tellstick-exporter
    go build
    ./telldus_exporter <flags>

To see all available configuration flags:

    ./telldus_exporter -h

## Metrics

Metrics are made available at port 9317 by default.
The metrics follow this format:

    # HELP telldus_devices_count Number of Devices
    # TYPE telldus_devices_count gauge
    telldus_devices_count 4
    # HELP telldus_exporter_build_info telldus_exporter build_info
    # TYPE telldus_exporter_build_info gauge
    telldus_exporter_build_info{golang_version="go1.9.2",version="0.1"} 1
    # HELP telldus_sensors_count Number of Sensors
    # TYPE telldus_sensors_count gauge
    telldus_sensors_count 7
    # HELP telldus_sensors_metric Value of the Sensor
    # TYPE telldus_sensors_metric gauge
    telldus_sensors_metric{Sensorname="Office A",type="humidity"} 29
    telldus_sensors_metric{Sensorname="Office A",type="temp"} 24.1
    telldus_sensors_metric{Sensorname="Outside",type="temp"} 1
    telldus_sensors_metric{Sensorname="Kitchen B",type="humidity"} 22
    telldus_sensors_metric{Sensorname="Kitchen B",type="temp"} 22.6
