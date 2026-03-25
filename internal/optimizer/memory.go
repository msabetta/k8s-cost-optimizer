package optimizer

import "k8s-cost-optimizer/internal/models"

func memoryRightsizing(m models.ContainerMetrics) *Recommendation {
	if m.MemoryUsageMB < (0.5 * m.MemoryRequestMB) {
		savings := (m.MemoryRequestMB - m.MemoryUsageMB) * 0.0001

		return &Recommendation{
			Container: m.Namespace + "/" + m.Pod + "/" + m.Name,
			Action:    "Reduce Memory request",
			Savings:   savings,
		}
	}
	return nil
}