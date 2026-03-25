package optimizer

import "k8s-cost-optimizer/internal/models"

func cpuRightsizing(m models.ContainerMetrics) *Recommendation {
	if m.CPUUsage < (0.3 * m.CPURequest) {
		savings := (m.CPURequest - m.CPUUsage) * 0.2

		return &Recommendation{
			Container: m.Namespace + "/" + m.Pod + "/" + m.Name,
			Action:    "Reduce CPU request",
			Savings:   savings,
		}
	}
	return nil
}