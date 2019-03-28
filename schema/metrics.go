package schema

const EndpointMetrics = "/metrics"

type (
	Metrics struct {
		SystemBalance    int         `json:"balance"`
		TransactionCount int         `json:"transactionCount"`
		UserCount        int         `json:"userCount"`
		Days             []DayMetric `json:"days"` // last 30 days
	}

	// single day of metrics
	DayMetric struct {
		Date             string `json:"date"` // YYYY-MM-DD
		TransactionCount int    `json:"count,string"`
		DistinctUsers    int    `json:"distinctUsers,string"`
		SystemBalance    int    `json:"balance"`
		PositiveFlux     int    `json:"positiveBalance"`
		NegativeFlux     int    `json:"negativeBalance"`
	}
)
