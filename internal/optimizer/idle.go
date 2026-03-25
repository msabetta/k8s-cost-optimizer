package optimizer

import "k8s-cost-optimizer/internal/models"

func detectIdle(m models.ContainerMetrics) *Recommendation {
	if m.CPUUsage < 0.05 && m.MemoryUsageMB < 50 {
		return &Recommendation{
			Container: m.Namespace + "/" + m.Pod + "/" + m.Name,
			Action:    "Container idle - consider scaling to zero",
			Savings:   m.CPURequest * 0.2,
		}
	}
	return nil
}