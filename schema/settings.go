package schema

const EndpointSettings = "/settings"

type Limit struct {
	Lower int `json:"lower"`
	Upper int `json:"upper"`
}

type AmountPreset struct {
	IsEnabled         bool  `json:"enabled"`
	AllowCustomAmount bool  `json:"custom"`
	PresetAmounts     []int `json:"steps"`
}

type Settings struct {
	Common struct {
		IdleTimeout int `json:"idleTimeout"`
	} `json:"common"`

	Paypal struct {
		IsEnabled  bool   `json:"enabled"`
		Recipient  string `json:"recipient"`
		PercentFee int    `json:"fee"`
	} `json:"paypal"`

	User struct {
		StalePeriod string `json:"stalePeriod"`
	} `json:"user"`

	I18n struct {
		DateFormat string `json:"dateFormat"`
		Timezone   string `json:"timezone"`
		Language   string `json:"language"`
		Currency   struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
			Alpha3 string `json:"alpha3"`
		} `json:"currency"`
	} `json:"i18n"`

	Account struct {
		Limit Limit `json:"boundary"`
	} `json:"account"`

	Payment struct {
		Reverse struct {
			IsEnabled bool   `json:"enabled"`
			Deletes   bool   `json:"delete"`
			Timeout   string `json:"timeout"`
		} `json:"undo"`
		Limit         Limit `json:"boundary"`
		TransferFunds struct {
			IsEnabled bool `json:"enabled"`
		} `json:"transactions"`
		Deposit  AmountPreset `json:"deposit"`
		Withdraw AmountPreset `json:"dispense"`
	} `json:"payment"`
}

type SettingsResponse struct {
	Settings Settings `json:"settings"`
}
