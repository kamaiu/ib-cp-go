package ws

/*

 */
type MarketData struct {
	LastPrice              string `json:"31"`
	Symbol                 string `json:"55"`
	Text                   string `json:"58"`
	High                   string `json:"70"`
	Low                    string `json:"71"`
	Position               string `json:"72"`
	MarketValue            string `json:"73"`
	AveragePrice           string `json:"74"`
	UnrealizedPnL          string `json:"75"`
	FormattedPosition      string `json:"76"`
	FormattedUnrealizedPnL string `json:"77"`
	DailyPnL               string `json:"78"`
	ChangePrice            string `json:"82"`
	ChangePercent          string `json:"83"`
	BidPrice               string `json:"84"`
	AskSize                string `json:"85"`
	AskPrice               string `json:"86"`
	Volume                 string `json:"87"`
	BidSize                string `json:"88"`

	Exchange        string `json:"6004"`
	Conid6008       string `json:"6008"`
	SecurityType    string `json:"6070"`
	Months          string `json:"6072"`
	RegularExpiry   string `json:"6073"`
	Marker          string `json:"6119"` // Marker for market data delivery method (similar to request id)
	UnderlyingConid string `json:"6457"` // Underlying Conid. Use /trsrv/secdef to get more information about the security

	// Market Data Availability. The field may contain two chars.
	//
	// The first char is the primary code:
	// R = Realtime, D = Delayed, Z = Frozen, Y = Frozen Delayed.
	//
	// The second char is the secondary code:
	// P = Snapshot Available, p = Consolidated.
	MarketDataAvailability string `json:"6509"`
	CompanyName            string `json:"7051"`
	LastSize               string `json:"7059"`
	ConidAndExchange       string `json:"7094"`
	ContractDescription    string `json:"7219"`
	ContractDescription2   string `json:"7220"`
	ListingExchange        string `json:"7221"`
	Industry               string `json:"7280"`
	Category               string `json:"7281"`
	AverageDailyVolume     string `json:"7282"`
	HistoricVolume30d      string `json:"7284"`
	PutCallRatio           string `json:"7285"`
	DividendAmount         string `json:"7286"`
	DividendYieldPercent   string `json:"7287"`
	ExDateOfTheDividend    string `json:"7288"`
	MarketCap              string `json:"7289"`
	PE                     string `json:"7290"`
	EPS                    string `json:"7291"`
	CostBasis              string `json:"7292"`
	High52Week             string `json:"7293"`
	Low52Week              string `json:"7294"`
	OpenPrice              string `json:"7295"`
	ClosePrice             string `json:"7296"`

	Delta             string `json:"7308"`
	Gamma             string `json:"7309"`
	Theta             string `json:"7310"`
	Vega              string `json:"7311"`
	ImpliedVolatility string `json:"7633"`
	ServerId          string `json:"serverId"`
	Conid             int64  `json:"conid"`
	Updated           int64  `json:"_updated"`
}
