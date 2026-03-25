package collector

import (
	"strconv"
)

// helper: crea chiave unica
func buildKey(metric map[string]string) string {
	return metric["namespace"] + "|" + metric["pod"] + "|" + metric["container"]
}

func parseMetrics(
	cpuUsageData, cpuReqData,
	memUsageData, memReqData []map[string]interface{},
) []ContainerMetrics {

	metricsMap := make(map[string]*ContainerMetrics)

	// 🔹 CPU USAGE
	for _, item := range cpuUsageData {
		metric := item["metric"].(map[string]string)
		valueStr := item["value"].(string)
		value, _ := strconv.ParseFloat(valueStr, 64)

		key := buildKey(metric)

		metricsMap[key] = &ContainerMetrics{
			Name:      metric["container"],
			Namespace: metric["namespace"],
			Pod:       metric["pod"],
			CPUUsage:  value,
		}
	}

	// 🔹 CPU REQUEST
	for _, item := range cpuReqData {
		metric := item["metric"].(map[string]string)
		valueStr := item["value"].(string)
		value, _ := strconv.ParseFloat(valueStr, 64)

		key := buildKey(metric)

		if m, ok := metricsMap[key]; ok {
			m.CPURequest = value
		}
	}

	// 🔹 MEMORY USAGE
	for _, item := range memUsageData {
		metric := item["metric"].(map[string]string)
		valueStr := item["value"].(string)
		value, _ := strconv.ParseFloat(valueStr, 64)

		key := buildKey(metric)

		if m, ok := metricsMap[key]; ok {
			m.MemoryUsageMB = value / (1024 * 1024)
		}
	}

	// 🔹 MEMORY REQUEST
	for _, item := range memReqData {
		metric := item["metric"].(map[string]string)
		valueStr := item["value"].(string)
		value, _ := strconv.ParseFloat(valueStr, 64)

		key := buildKey(metric)

		if m, ok := metricsMap[key]; ok {
			m.MemoryRequestMB = value / (1024 * 1024)
		}
	}

	// convert map → slice
	var result []ContainerMetrics
	for _, v := range metricsMap {
		result = append(result, *v)
	}

	return result
}