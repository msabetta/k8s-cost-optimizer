package optimizer

import "k8s-cost-optimizer/internal/models"

// detectOverprovision identifica container con spreco significativo combinato
func detectOverprovision(m models.ContainerMetrics) *Recommendation {

	// percentuali utilizzo
	cpuUtil := m.CPUUsage / m.CPURequest
	memUtil := m.MemoryUsageMB / m.MemoryRequestMB

	// soglie MVP
	if cpuUtil < 0.4 && memUtil < 0.5 {

		// stima semplice risparmio combinato
		cpuSavings := (m.CPURequest - m.CPUUsage) * 0.2
		memSavings := (m.MemoryRequestMB - m.MemoryUsageMB) * 0.0001

		totalSavings := cpuSavings + memSavings

		return &Recommendation{
			Container: m.Namespace + "/" + m.Pod + "/" + m.Name,
			Action:    "Overprovisioned container - reduce CPU and Memory",
			Savings:   totalSavings,
		}
	}

	return nil
}