# Tellstick Local Exporter

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
