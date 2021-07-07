package influxdb

import (
	"context"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func WriteData(serviceName string, containerID string, hostName string, servicePort string) {

	bucketName := os.Getenv("bucket")
	influxToken := os.Getenv("influx_token")
	orgName := os.Getenv("org_name")
	influxURL := os.Getenv("influx_url")

	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient(influxURL, influxToken)
	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(orgName, bucketName)

	// write point immediately
	//	writeAPI.WritePoint(context.Background(), p)
	// Create point using fluent style
	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "serviceName").
		AddField("containerID", containerID).
		AddField("hostName", hostName).
		AddField("servicePort", servicePort).
		SetTime(time.Now())
	writeAPI.WritePoint(context.Background(), p)

	// Ensures background processes finishes
	client.Close()
}
