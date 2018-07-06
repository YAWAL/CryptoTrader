package gdaxClient

type Product struct {
	Id             string `json:"id"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	BaseMinSize    string `json:"base_min_size"`
	BaseMaxSize    string `json:"base_max_size"`
	QuoteIncrement string `json:"quote_increment"`

	CancelOnly     bool   `json:"cancel_only"`
	DisplayName    string `json:"display_name"`
	LimitOnly      bool   `json:"limit_only"`
	MarginEnable   string `json:"margin_enable"`
	MaxMarketFunds string `json:"max_market_funds"`
	MinMarketFunds string `json:"min_market_funds"`
	PostOnly       bool   `json:"post_only"`
	Status         string `json:"status"`
	StatusMessage  string `json:"status_message"`
}

type BookEntry struct {
	Price          string
	Size           string
	NumberOfOrders int
	//OrderId        string
}

type Book struct {
	Sequence int         `json:"sequence"`
	Bids     []BookEntry `json:"bids"`
	Asks     []BookEntry `json:"asks"`
}
