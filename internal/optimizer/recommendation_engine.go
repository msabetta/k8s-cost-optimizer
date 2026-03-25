package optimizer

import "k8s-cost-optimizer/internal/models"

func GenerateRecommendations(metrics []models.ContainerMetrics, _ interface{}) []Recommendation {
	var recs []Recommendation

	for _, m := range metrics {

		if r := cpuRightsizing(m); r != nil {
			recs = append(recs, *r)
		}

		if r := memoryRightsizing(m); r != nil {
			recs = append(recs, *r)
		}

		if r := detectIdle(m); r != nil {
			recs = append(recs, *r)
		}

		if r := detectOverprovision(m); r != nil {
			recs = append(recs, *r)
		}
	}

	return recs
}