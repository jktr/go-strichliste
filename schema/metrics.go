package schema

const EndpointMetrics = "/metrics"

type (
	UserMetrics struct {
		Balance      int             `json:"balance"`
		Articles     []ArticleMetric `json:"articles"`
		Transactions struct {
			Count    int `json:"count"`
			Outgoing struct {
				Count    int `json:"count"`
				Cashflow int `json:"amount"`
			} `json:"outgoing"`
			Incoming struct {
				Count    int `json:"count"`
				Cashflow int `json:"amount"`
			} `json:"incoming"`
		} `json:"transactions"`
	}

	ArticleMetric struct {
		Article Article `json:"article"`
		Count   int     `json:"count"`
		Spent   int     `json:"amount"`
	}

	SystemMetrics struct {
		Balance      int         `json:"balance"`
		Transactions int         `json:"transactionCount"`
		Users        int         `json:"userCount"`
		Days         []DayMetric `json:"days"` // last 30 days
	}

	// single day of metrics
	DayMetric struct {
		Date             string `json:"date"` // YYYY-MM-DD
		Transactions     int    `json:"count,string"`
		DistinctUsers    int    `json:"distinctUsers,string"`
		Balance          int    `json:"balance"`
		IncomingCashflow int    `json:"positiveBalance"`
		OutgoingCashflow int    `json:"negativeBalance"`
	}
)
