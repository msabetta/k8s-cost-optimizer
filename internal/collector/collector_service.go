package collector

import (
	"fmt"
	"os"
	"encoding/json"
)

var promURL = getPrometheusURL()


func getPrometheusURL() string{
	url := os.Getenv("PROMETHEUS_URL")
	if url == "" {
		return "http://localhost:9090"
	}
	return url
}

func FetchMetrics() []ContainerMetrics {

	//PromQL queries (filtrate per evitare rumore)
	cpuUsageQuery := `rate(container_cpu_usage_seconds_total{container!="",image!=""}[5m])`
	cpuReqQuery := `kube_pod_container_resource_requests{resource="cpu"}`
	memUsageQuery := `container_memory_usage_bytes{container!=""}`
	memReqQuery := `kube_pod_container_resource_requests{resource="memory"}`

	cpuUsageData, err := queryPrometheus(promURL, cpuUsageQuery)

	if err != nil {
		fmt.Println("Error CPU usage:", err)
		return nil
	}

	cpuReqData, _ := queryPrometheus(promURL, cpuReqQuery)
	memUsageData, _ := queryPrometheus(promURL, memUsageQuery)
	memReqData, _ := queryPrometheus(promURL, memReqQuery)

	metrics := parseMetrics(cpuUsageData, cpuReqData, memUsageData, memReqData)

	if len(metrics) == 0 {
		fmt.Println("Using fake metrics...")
		return loadFakeMetrics()
	}

	//debug
	fmt.Println("Collected metrics:")
	for _, m := range metrics {
		fmt.Printf("%+v\n", m)
	}

	return metrics
}


func loadFakeMetrics() []ContainerMetrics {
	file, err := os.ReadFile("/tmp/fake_metrics.json")
	if err != nil {
		return nil
	}

	var data []ContainerMetrics
	_ = json.Unmarshal(file, &data)
	return data
}
