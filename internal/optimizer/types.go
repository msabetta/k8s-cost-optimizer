package optimizer

type Recommendation struct {
	Container string  `json:"container"`
	Action    string  `json:"action"`
	Savings   float64 `json:"savings"`
}