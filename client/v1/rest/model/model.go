//go:generate easyjson -all $GOFILE
package model

//easyjson:json
type Any map[string]interface{}

//easyjson:json
type Iserver_Scanner_Params_GET_200_Instrument_list_List_Item_Filters_List []string

//easyjson:json
type Iserver_Scanner_Run_POST_200_List []*Iserver_Scanner_Run_POST_200_List_Item

// Always contains at least one 'tradingTime'  and zero or more 'sessionTime' tags
//easyjson:json
type Trsrv_Secdef_Schedule_GET_200_Schedules_List []*Trsrv_Secdef_Schedule_GET_200_Schedules_List_Item

//easyjson:json
type Iserver_Secdef_Search_POST_200_List_Item_Sections_List []*Iserver_Secdef_Search_POST_200_List_Item_Sections_List_Item

//easyjson:json
type AlertRequest_Conditions_List []*AlertRequest_Conditions_List_Item

// Prompt messages that may affect trading or the account
//easyjson:json
type AuthStatus_Prompts_List []string

// array of contracts from different exchanges
//easyjson:json
type Stocks_List_Item_Contracts_List []*Stocks_List_Item_Contracts_List_Item

//easyjson:json
type Pa_Transactions_POST_Request_AcctIds_List []string

//easyjson:json
type HistoryData_Data_List []*HistoryData_Data_List_Item

//easyjson:json
type Portfolio_Allocation_POST_Request_AcctIds_List []string

//easyjson:json
type Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item_TifTypes_List []int64

//easyjson:json
type Ibcust_Entity_Info_GET_200_List []*Ibcust_Entity_Info_GET_200_List_Item

//easyjson:json
type Trsrv_Secdef_POST_Request_Conids_List []int64

//easyjson:json
type Iserver_Account_Orders_GET_Request_Filters_List []string

//easyjson:json
type Iserver_Scanner_Params_GET_200_Instrument_list_List []*Iserver_Scanner_Params_GET_200_Instrument_list_List_Item

//easyjson:json
type Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item_OrderTypesOutside_List []int64

//easyjson:json
type Futures_List []*Futures_List_Item

//easyjson:json
type Performance_Nav_Data_List []*Performance_Nav_Data_List_Item

//easyjson:json
type Ibcust_Entity_Info_GET_200_List_Item_Entities_List_Item_IdentDocs_List []map[string]interface{}

//easyjson:json
type Pa_Performance_POST_Request_AcctIds_List []string

// Notes for bracket orders: 1. Children orders will not have its own "cOID", so please
// donot pass "cOID"
// parameter in child order.Instead, they will have a "parentId" which must be equal
// to "cOID" of parent.
// 2. When you cancel a parent order, it will cancel all bracket orders, when you cancel
// one child order,
// it will also cancel its sibling order.
//easyjson:json
type Iserver_Account_AccountId_Orders_POST_Request_Orders_List []*OrderRequest

//easyjson:json
type Iserver_Reply_Replyid_POST_200_List []*Iserver_Reply_Replyid_POST_200_List_Item

//easyjson:json
type Iserver_Account_Orders_GET_200_Orders_List []*Iserver_Account_Orders_GET_200_Orders_List_Item

//easyjson:json
type Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List []*Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item

//easyjson:json
type ScannerParams_ScanTypeList_ScanType_List []*ScannerParams_ScanTypeList_ScanType_List_Item

//easyjson:json
type Summary_BalanceByDate_Series_List_Item_Date_List_Item_List []float64

//easyjson:json
type Summary_BalanceByDate_Series_List_Item_Date_List [][]float64

// events
//easyjson:json
type Events_List []*Events_List_Item

// This is an array of object(s), there could be multiple results under same symbol
//easyjson:json
type Stocks_List []*Stocks_List_Item

//easyjson:json
type Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item_IbalgoTypes_List []int64

//easyjson:json
type Pa_Transactions_POST_Request_Conids_List []float64

//easyjson:json
type Summary_ExcludedAccounts_List []*Summary_ExcludedAccounts_List_Item

//easyjson:json
type Contract_Rules_OrderTypes_List []string

//easyjson:json
type Contract_Rules_OrderTypesOutside_List []string

// Unique account id
//easyjson:json
type Iserver_Accounts_GET_200_Accounts_List []string

//easyjson:json
type Iserver_Secdef_Strikes_GET_200_Call_List []string

//easyjson:json
type Iserver_Scanner_Params_GET_200_Location_tree_List_Item_Locations_List []*Iserver_Scanner_Params_GET_200_Location_tree_List_Item_Locations_List_Item

//easyjson:json
type Iserver_Scanner_Params_GET_200_Location_tree_List []*Iserver_Scanner_Params_GET_200_Location_tree_List_Item

//easyjson:json
type Iserver_Scanner_Params_GET_200_Scan_type_list_List []*Iserver_Scanner_Params_GET_200_Scan_type_list_List_Item

//easyjson:json
type Accounts_List []*Account

// each value stands for price change percent of corresponding date in dates array
//easyjson:json
type Performance_Nav_Data_List_Item_Returns_List []float64

// each value stands for price change percent of corresponding date in dates array
//easyjson:json
type Performance_Tpps_Data_List_Item_Returns_List []float64

//easyjson:json
type Fyi_Settings_GET_200_List []*Fyi_Settings_GET_200_List_Item

//easyjson:json
type Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item_OrderTypes_List []int64

//easyjson:json
type Iserver_Account_AccountId_Orders_POST_200_List []*Iserver_Account_AccountId_Orders_POST_200_List_Item

// array of dates, the length should be same as the length of returns inside data.
//easyjson:json
type Performance_Cps_Dates_List []string

// Please note here, if the message is a question, you have to reply to question in
// order to submit
// the order successfully. See more in the "/iserver/reply/{replyid}" endpoint.
//easyjson:json
type Iserver_Account_Orders_FaGroup_POST_200_List_Item_Message_List []string

//easyjson:json
type Iserver_Account_AccountId_Order_OrderId_POST_200_List []*Iserver_Account_AccountId_Order_OrderId_POST_200_List_Item

//easyjson:json
type Iserver_Secdef_Info_GET_200_List []*SecdefInfo

//easyjson:json
type Iserver_Account_AccountId_Order_POST_200_List []*Iserver_Account_AccountId_Order_POST_200_List_Item

//easyjson:json
type Iserver_Secdef_Search_POST_200_List []*Iserver_Secdef_Search_POST_200_List_Item

//easyjson:json
type Secdef_List []*Secdef_List_Item

//easyjson:json
type Events_List_Item_Conids_List []string

// orderType
//easyjson:json
type Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item_OrderDefaults_List_Item_String_List []string

//easyjson:json
type Performance_Tpps_Data_List []*Performance_Tpps_Data_List_Item

//easyjson:json
type Iserver_Marketdata_Snapshot_GET_200_List []*Iserver_Marketdata_Snapshot_GET_200_List_Item

//easyjson:json
type Ibcust_Entity_Info_GET_200_List_Item_Entities_List []*Ibcust_Entity_Info_GET_200_List_Item_Entities_List_Item

//easyjson:json
type Fyi_Deliveryoptions_GET_200_E_List []*Fyi_Deliveryoptions_GET_200_E_List_Item

// Please note here, if the message is a question, you have to reply to question in
// order to submit
// the order successfully. See more in the "/iserver/reply/{replyid}" endpoint.
//easyjson:json
type Iserver_Account_AccountId_Orders_POST_200_List_Item_Message_List []string

// Contains information like name, supported filters, etc. for an instrument
//easyjson:json
type ScannerParams_InstrumentList_Instrument_List []*ScannerParams_InstrumentList_Instrument_List_Item

//easyjson:json
type ScannerResult_Contracts_Contract_List []*ScannerResult_Contracts_Contract_List_Item

//easyjson:json
type AlertResponse_Conditions_List []*AlertResponse_Conditions_List_Item

//easyjson:json
type Iserver_Account_AccountId_Alerts_GET_200_List []*Iserver_Account_AccountId_Alerts_GET_200_List_Item

//easyjson:json
type Summary_BalanceByDate_Series_List []*Summary_BalanceByDate_Series_List_Item

// array of dates, the length should be same as the length of returns inside data.
//easyjson:json
type Performance_Nav_Dates_List []string

//easyjson:json
type Performance_Cps_Data_List []*Performance_Cps_Data_List_Item

// Sorted by date descending
//easyjson:json
type Transactions_Transactions_List []*Transactions_Transactions_List_Item

// array of dates, the length should be same as the length of returns inside data.
//easyjson:json
type Performance_Tpps_Dates_List []string

//easyjson:json
type Position_List_Item_ConExchMap_List []string

//easyjson:json
type Iserver_Account_Orders_FaGroup_POST_200_List []*Iserver_Account_Orders_FaGroup_POST_200_List_Item

// If object returned will provide the defaults based on user settings
//easyjson:json
type Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item_OrderDefaults_List []*Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item_OrderDefaults_List_Item

// Contains information like name, supported filters, etc. for an instrument. A location
// can contain more locations forming a tree-like structure which allows user to control
// the lcoation at more granular level. locationCode has to be used to specify lcoations
// while querying a scanner.
//easyjson:json
type ScannerParams_LocationTree_Location_List []*ScannerParams_LocationTree_Location_List_Item

// Contains list of filters supported for the scanner
//easyjson:json
type ScannerParams_FilterList_List []map[string]interface{}

//easyjson:json
type Wagers_List []*Wagers_List_Item

//easyjson:json
type Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item_CqtTypes_List []int64

// Please note here, if the message is a question, you have to reply to question in
// order to submit
// the order successfully. See more in the "/iserver/reply/{replyid}" endpoint.
//easyjson:json
type Iserver_Account_AccountId_Order_POST_200_List_Item_Message_List []string

//easyjson:json
type Pa_Summary_POST_Request_AcctIds_List []string

//easyjson:json
type Notifications_List []*Notifications_List_Item

// each value stands for price change percent of corresponding date in dates array
//easyjson:json
type Performance_Cps_Data_List_Item_Returns_List []float64

//easyjson:json
type Iserver_Scanner_Params_GET_200_Scan_type_list_List_Item_Instruments_List []string

//easyjson:json
type Performance_Included_List []string

//easyjson:json
type Position_List []*Position_List_Item

//easyjson:json
type Iserver_Account_Trades_GET_200_List []*Trade

//easyjson:json
type Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item_FraqTypes_List []int64

//easyjson:json
type Summary_AccountSummaries_List []*Summary_AccountSummaries_List_Item

//easyjson:json
type Contract_Rules_TifTypes_List []string

//easyjson:json
type Allocation_List []*Allocation_List_Item

//easyjson:json
type Inds_List []*Inds_List_Item

//easyjson:json
type Iserver_Secdef_Strikes_GET_200_Put_List []string

//easyjson:json
type Iserver_Scanner_Params_GET_200_Filter_list_List []*Iserver_Scanner_Params_GET_200_Filter_list_List_Item

type Fyi_Deliveryoptions_Email_PUT_200 struct {
	T int64 `json:"T"`
	V int64 `json:"V"`
}

// confirms session is open
type Tickle_POST_200 struct {
	Collission bool                     `json:"collission"`
	Iserver    *Tickle_POST_200_Iserver `json:"iserver"`
	// SessionID
	Session    string `json:"session"`
	SsoExpires int64  `json:"ssoExpires"`
	UserId     int64  `json:"userId"`
}

type Iserver_Secdef_Search_POST_Request struct {
	// should be true if the search is to be performed by name. false by default.
	Name bool `json:"name"`
	// If search is done by name, only the assets provided in this field will be returned.
	// Currently, only STK is supported.
	SecType string `json:"secType"`
	// symbol or name to be searched
	Symbol string `json:"symbol"`
}

type Trsrv_Stocks_GET_500 struct {
	Error string `json:"error"`
}

type Transactions_Transactions_List_Item struct {
	Acctid string `json:"acctid"`
	// Raw value, no formatting. Net transaction amount (may include commission, tax).
	// In asset currency
	Amt   float64 `json:"amt"`
	Conid float64 `json:"conid"`
	// currency code
	Cur string `json:"cur"`
	// Date of transaction.  Epoch time, GMT
	Date string `json:"date"`
	// Transaction description
	Desc string `json:"desc"`
	// Conversion rate from asset currency to response currency
	FxRate float64 `json:"fxRate"`
	// In asset currency. Not be applicable for all transaction types.
	Pr float64 `json:"pr"`
	// Not applicable for all transaction types
	Qty float64 `json:"qty"`
	// Transaction Type Name: Examples: "Sell", "Buy", "Corporate Action",
	// "Dividend Payment", "Transfer", "Payment in Lieu"
	// Dividends and Transfers do not have price and quantity in response
	Type string `json:"type"`
}

type Performance struct {
	// Cumulative performance data
	Cps          *Performance_Cps `json:"cps"`
	CurrencyType string           `json:"currencyType"`
	Id           string           `json:"id"`
	Included     []string         `json:"included"`
	// Net asset value data for the account or consolidated accounts. NAV data is not applicable
	// to benchmarks.
	Nav *Performance_Nav `json:"nav"`
	Pm  string           `json:"pm"`
	Rc  int64            `json:"rc"`
	// Time period performance data
	Tpps *Performance_Tpps `json:"tpps"`
}

type StatsData struct {
	Conid    float64 `json:"Conid"`
	Exchange string  `json:"Exchange"`
	// Object, payload depends on event type. See confluence page for IGEvntUpd.
	P  string  `json:"P"`
	T  float64 `json:"T"`
	TT float64 `json:"TT"`
	V  float64 `json:"V"`
}

type ScannerParams_ScanTypeList_ScanType_List_Item struct {
	DisplayName string `json:"displayName"`
	// Instrument types separated by a comma which are supported for this scan type
	Instruments string `json:"instruments"`
	// scan code which ahs to be provided while querying scanner with this scan type
	ScanCode string `json:"scanCode"`
}

// contains all the order related info
type Order struct {
	// account id
	Acct string `json:"acct"`
	// back-ground color
	BgColor        string `json:"bgColor"`
	CompanyName    string `json:"companyName"`
	Conid          int64  `json:"conid"`
	Description1   string `json:"description1"`
	FgColor        string `json:"fgColor"`
	FilledQuantity string `json:"filledQuantity"`
	// for example NASDAQ.NMS
	ListingExchange string `json:"listingExchange"`
	OrderDesc       string `json:"orderDesc"`
	OrderId         int64  `json:"orderId"`
	// User defined string used to identify the order. Value is set using "cOID" field while
	// placing an order.
	Order_ref string `json:"order_ref"`
	// for example Limit
	OrigOrderType string `json:"origOrderType"`
	// Only exists in child order of bracket
	ParentId          int64   `json:"parentId"`
	Price             float64 `json:"price"`
	RemainingQuantity string  `json:"remainingQuantity"`
	// for example STK
	SecType string `json:"secType"`
	// BUY or SELL
	Side string `json:"side"`
	// PendingSubmit - Indicates the order was sent, but confirmation has not been received
	// that it has been received by the destination.
	// Occurs most commonly if an exchange is closed.
	// PendingCancel - Indicates that a request has been sent to cancel an order but confirmation
	// has not been received of its cancellation. PreSubmitted - Indicates that a simulated
	// order type has been accepted by the IBKR system and that this order has yet to be
	// elected.
	// The order is held in the IBKR system until the election criteria are
	// met. At that time the order is transmitted to the order destination as specified.
	//
	// Submitted - Indicates that the order has been accepted at the order destination and
	// is working. Cancelled - Indicates that the balance of the order has been confirmed
	// cancelled by the IB system.
	// This could occur unexpectedly when IB or the destination has rejected
	// the order.
	// Filled - Indicates that the order has been completely filled.  Inactive - Indicates
	// the order is not working, for instance if the order was invalid and triggered an
	// error message,
	// or if the order was to short a security and shares have not yet been located.
	//
	Status string `json:"status"`
	// for exmple FB
	Ticker string `json:"ticker"`
}

type Iserver_Account_AccountId_Alert_Activate_POST_Request struct {
	// 1 to activate, 0 to deactivate
	AlertActive int64 `json:"alertActive"`
	// alert id(order id)
	AlertId int64 `json:"alertId"`
}

type Fyi_Disclaimer_Typecode_PUT_200 struct {
	T int64 `json:"T"`
	V int64 `json:"V"`
}

type AlertResponse_Conditions_List_Item struct {
	// "a" means "AND", "o" means "OR", "n" means "END",
	// the last one condition in the condition array should "n"
	Condition_logic_bind string `json:"condition_logic_bind"`
	// optional, operator for the current condition, can be >= or <=
	Condition_operator string `json:"condition_operator"`
	// only needed for some MTA alert condition
	Condition_time_zone string `json:"condition_time_zone"`
	// optional, only some type of conditions have triggerMethod
	Condition_trigger_method string `json:"condition_trigger_method"`
	// Types: 1-Price, 3-Time, 4-Margin, 5-Trade, 6-Volume, 7: MTA market 8: MTA Position,
	// 9: MTA Acc. Daily PN&
	Condition_type int64 `json:"condition_type"`
	// can not be empty, can pass default value "*"
	Condition_value string `json:"condition_value"`
	// format, conid@exchange
	Conidex                string `json:"conidex"`
	Contract_description_1 string `json:"contract_description_1"`
}

// future contract information
type Futures_List_Item struct {
	// conid of the future contract
	Conid          int64  `json:"conid"`
	ExpirationDate string `json:"expirationDate"`
	// last trading day
	Ltd             string `json:"ltd"`
	Symbol          string `json:"symbol"`
	UnderlyingConid int64  `json:"underlyingConid"`
}

type Summary_BalanceByDate_Series_List_Item struct {
	Date    [][]float64 `json:"date"`
	GroupId string      `json:"groupId"`
	Id      string      `json:"id"`
	Name    string      `json:"name"`
}

type Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item struct {
	// Cash currency for the contract
	CashCcy string `json:"cashCcy"`
	// Increment value for cash quantity
	CashQtyIncr float64 `json:"cashQtyIncr"`
	// cash value
	CashSize int64   `json:"cashSize"`
	CqtTypes []int64 `json:"cqtTypes"`
	// Default quantity
	DefaultSize int64 `json:"defaultSize"`
	// Default time in force value
	DefaultTIF  string  `json:"defaultTIF"`
	DisplaySize float64 `json:"displaySize"`
	Error       string  `json:"error"`
	// decimal places for fractional order size
	FraqInt     float64 `json:"fraqInt"`
	FraqTypes   []int64 `json:"fraqTypes"`
	IbalgoTypes []int64 `json:"ibalgoTypes"`
	// Price increment value
	Increment float64 `json:"increment"`
	// Number of digits for price increment
	IncrementDigits int64 `json:"incrementDigits"`
	// Limit price
	LimitPrice float64 `json:"limitPrice"`
	// trading negative price support
	NegativeCapable bool `json:"negativeCapable"`
	// If object returned will provide the defaults based on user settings
	OrderDefaults []*Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item_OrderDefaults_List_Item `json:"orderDefaults"`
	// Order origin designation for US securities options and Options Clearing Corporation
	OrderOrigination  float64 `json:"orderOrigination"`
	OrderTypes        []int64 `json:"orderTypes"`
	OrderTypesOutside []int64 `json:"orderTypesOutside"`
	// order preview required
	Preview bool `json:"preview"`
	// Price Magnifier
	PriceMagnifier float64 `json:"priceMagnifier"`
	// increment quantity value
	SizeIncrement int64 `json:"sizeIncrement"`
	// Stop price
	Stopprice float64 `json:"stopprice"`
	TifTypes  []int64 `json:"tifTypes"`
}

type Ledger struct {
	Acctcode                  string  `json:"acctcode"`
	Cashbalance               float64 `json:"cashbalance"`
	Commoditymarketvalue      float64 `json:"commoditymarketvalue"`
	Corporatebondsmarketvalue float64 `json:"corporatebondsmarketvalue"`
	Currency                  string  `json:"currency"`
	Exchangerate              float64 `json:"exchangerate"`
	Funds                     float64 `json:"funds"`
	Futuremarketvalue         float64 `json:"futuremarketvalue"`
	Interest                  float64 `json:"interest"`
	Issueroptionsmarketvalue  float64 `json:"issueroptionsmarketvalue"`
	Key                       string  `json:"key"`
	Moneyfunds                float64 `json:"moneyfunds"`
	Netliquidationvalue       float64 `json:"netliquidationvalue"`
	Realizedpnl               float64 `json:"realizedpnl"`
	Sessionid                 int64   `json:"sessionid"`
	Settledcash               float64 `json:"settledcash"`
	Severity                  int64   `json:"severity"`
	Stockmarketvalue          float64 `json:"stockmarketvalue"`
	Timestamp                 int64   `json:"timestamp"`
	Unrealizedpnl             float64 `json:"unrealizedpnl"`
	Warrantsmarketvalue       float64 `json:"warrantsmarketvalue"`
}

// List of wagers
type Wagers_List_Item struct {
	Conid float64 `json:"conid"`
	Curr  string  `json:"curr"`
	Desc  string  `json:"desc"`
	Part  string  `json:"part"`
}

type Iserver_Accounts_GET_200 struct {
	// Unique account id
	Accounts []string `json:"accounts"`
	// Account Id and its alias
	Aliases         map[string]interface{} `json:"aliases"`
	SelectedAccount string                 `json:"selectedAccount"`
}

type HistoryResult_Bars struct {
	Close       float64 `json:"close"`
	Count       float64 `json:"count"`
	EndTime     string  `json:"endTime"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	Open        float64 `json:"open"`
	Time        string  `json:"time"`
	Volume      float64 `json:"volume"`
	WeightedAvg float64 `json:"weightedAvg"`
}

type Calendar_request_Date struct {
	// end date of a period. for example 20180808-0400
	End string `json:"end"`
	// start date of a period. for example 20180808-0400
	Start string `json:"start"`
}

type SystemError struct {
	Error string `json:"error"`
}

type Summary_AccountSummaries_List_Item struct {
	AccountTypeCode string `json:"accountTypeCode"`
	AccountTypeName string `json:"accountTypeName"`
	Chg             string `json:"chg"`
	EndVal          string `json:"endVal"`
	HasAccounts     string `json:"hasAccounts"`
	Rtn             string `json:"rtn"`
	StartVal        string `json:"startVal"`
}

// Cumulative performance data
type Performance_Cps struct {
	Data []*Performance_Cps_Data_List_Item `json:"data"`
	// array of dates, the length should be same as the length of returns inside data.
	Dates []string `json:"dates"`
	// D means Day
	Freq string `json:"freq"`
}

type Ibcust_Entity_Info_GET_200_List_Item_Entities_List_Item_Name struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Salutation string `json:"salutation"`
}

type Iserver_Account_Orders_GET_200 struct {
	Orders []*Iserver_Account_Orders_GET_200_Orders_List_Item `json:"orders"`
	// If live order update is a snapshot
	Snapshot bool `json:"snapshot"`
}

type Trsrv_Secdef_Schedule_GET_200 struct {
	Id string `json:"id"`
	// Always contains at least one 'tradingTime'  and zero or more 'sessionTime' tags
	Schedules []*Trsrv_Secdef_Schedule_GET_200_Schedules_List_Item `json:"schedules"`
	// null is returned.
	TradeVenueId string `json:"tradeVenueId"`
}

type Tickle_POST_200_Iserver_AuthStatus struct {
	MAC           string `json:"MAC"`
	Authenticated bool   `json:"authenticated"`
	Competing     bool   `json:"competing"`
	Connected     bool   `json:"connected"`
	Message       string `json:"message"`
}

// Contains list of instruments for which scanner can be ran
type ScannerParams_LocationTree struct {
	// Contains information like name, supported filters, etc. for an instrument. A location
	// can contain more locations forming a tree-like structure which allows user to control
	// the lcoation at more granular level. locationCode has to be used to specify lcoations
	// while querying a scanner.
	Location []*ScannerParams_LocationTree_Location_List_Item `json:"Location"`
}

// If the LIQUID hours differs from the total trading day then a separate 'session'
// tag is returned.
type Trsrv_Secdef_Schedule_GET_200_Schedules_List_Item_Sessions struct {
	ClosingTime int64 `json:"closingTime"`
	OpeningTime int64 `json:"openingTime"`
	// If the whole trading day is considered LIQUID then the value 'LIQUID' is returned.
	Prop string `json:"prop"`
}

// short positions allocation
type Allocation_List_Item_Sector_Short struct {
	Consumer    float64 `json:"Consumer"`
	Diversified float64 `json:"Diversified"`
	Industrial  float64 `json:"Industrial"`
}

type Iserver_Account_AccountId_Alert_POST_200 struct {
	Order_id        int64  `json:"order_id"`
	Order_status    string `json:"order_status"`
	Request_id      int64  `json:"request_id"`
	Success         bool   `json:"success"`
	Text            string `json:"text"`
	Warning_message string `json:"warning_message"`
}

type HistoryResult struct {
	Bars *HistoryResult_Bars `json:"bars"`
}

type Trade struct {
	Account                string  `json:"account"`
	Clearing_id            string  `json:"clearing_id"`
	Clearing_name          string  `json:"clearing_name"`
	Comission              float64 `json:"comission"`
	Company_name           string  `json:"company_name"`
	Conidex                string  `json:"conidex"`
	Contract_description_1 string  `json:"contract_description_1"`
	Exchange               string  `json:"exchange"`
	Execution_id           string  `json:"execution_id"`
	Net_amount             float64 `json:"net_amount"`
	Order_description      string  `json:"order_description"`
	// User defined string used to identify the order. Value is set using "cOID" field while
	// placing an order.
	Order_ref    string  `json:"order_ref"`
	Position     string  `json:"position"`
	Price        string  `json:"price"`
	Sec_type     string  `json:"sec_type"`
	Side         string  `json:"side"`
	Size         string  `json:"size"`
	Submitter    string  `json:"submitter"`
	Symbol       string  `json:"symbol"`
	Trade_time   string  `json:"trade_time"`
	Trade_time_r float64 `json:"trade_time_r"`
}

type Fyi_Deliveryoptions_Device_POST_200 struct {
	T int64 `json:"T"`
	V int64 `json:"V"`
}

// Contains list of contracts matching the scanner query
type ScannerResult_Contracts struct {
	Contract []*ScannerResult_Contracts_Contract_List_Item `json:"Contract"`
}

// account information
type Summary struct {
	AccountSummaries []*Summary_AccountSummaries_List_Item `json:"accountSummaries"`
	BalanceByDate    *Summary_BalanceByDate                `json:"balanceByDate"`
	Currency         string                                `json:"currency"`
	EndDate          string                                `json:"endDate"`
	ExcludedAccounts []*Summary_ExcludedAccounts_List_Item `json:"excludedAccounts"`
	// indicator of user having configured any external accounts
	HasExternalAccounts  bool   `json:"hasExternalAccounts"`
	LastSuccessfulUpdate string `json:"lastSuccessfulUpdate"`
	Pm                   string `json:"pm"`
	Rc                   int64  `json:"rc"`
	// date format-- yyyy-MM-dd
	StartDate string         `json:"startDate"`
	Total     *Summary_Total `json:"total"`
	UserId    string         `json:"userId"`
	View      string         `json:"view"`
}

// Contains all details of the contract, including rules you can use when placing orders
type Contract struct {
	Category     string `json:"category"`
	CompanyName  string `json:"companyName"`
	Company_name string `json:"company_name"`
	// same as that in request
	Con_id   string `json:"con_id"`
	Currency string `json:"currency"`
	Exchange string `json:"exchange"`
	Industry string `json:"industry"`
	// for example STK
	Instrument_type string `json:"instrument_type"`
	// for exmple FB
	Local_symbol string `json:"local_symbol"`
	// true means you can trade outside RTH(regular trading hours)
	R_t_h bool            `json:"r_t_h"`
	Rules *Contract_Rules `json:"rules"`
}

type Fyi_Disclaimer_Typecode_GET_200 struct {
	// disclaimer message
	DT string `json:"DT"`
	// fyi code
	FC string `json:"FC"`
}

type Iserver_Scanner_Params_GET_200_Location_tree_List_Item struct {
	Display_name string                                                                        `json:"display_name"`
	Locations    []*Iserver_Scanner_Params_GET_200_Location_tree_List_Item_Locations_List_Item `json:"locations"`
	Type         string                                                                        `json:"type"`
}

type Trsrv_Futures_GET_500 struct {
	Error string `json:"error"`
}

type Iserver_Account_AccountId_Order_POST_200_List_Item struct {
	Id string `json:"id"`
	// Please note here, if the message is a question, you have to reply to question in
	// order to submit
	// the order successfully. See more in the "/iserver/reply/{replyid}" endpoint.
	Message []string `json:"message"`
}

type Iserver_Account_Pnl_Partitioned_GET_200 struct {
	AcctId map[string]interface{} `json:"acctId"`
}

// Contains list of scan types for which scanner can be ran
type ScannerParams_ScanTypeList struct {
	ScanType []*ScannerParams_ScanTypeList_ScanType_List_Item `json:"ScanType"`
}

type ScannerParams struct {
	// Contains list of filters supported for the scanner
	FilterList []map[string]interface{} `json:"FilterList"`
	// Contains list of instruments for which scanner can be ran
	InstrumentList *ScannerParams_InstrumentList `json:"InstrumentList"`
	// Contains list of instruments for which scanner can be ran
	LocationTree *ScannerParams_LocationTree `json:"LocationTree"`
	// Contains list of scan types for which scanner can be ran
	ScanTypeList *ScannerParams_ScanTypeList `json:"ScanTypeList"`
}

type AlertRequest_Conditions_List_Item struct {
	// format, conid@exchange
	Conidex string `json:"conidex"`
	// "a" means "AND", "o" means "OR", "n" means "END",
	// the last one condition in the condition array should "n"
	LogicBind string `json:"logicBind"`
	// optional, operator for the current condition, can be >= or <=
	Operator string `json:"operator"`
	// only needed for some MTA alert condition
	TimeZone string `json:"timeZone"`
	// optional, only some type of conditions have triggerMethod
	TriggerMethod string `json:"triggerMethod"`
	// Types: 1-Price, 3-Time, 4-Margin, 5-Trade, 6-Volume, 7: MTA market 8: MTA Position,
	// 9: MTA Acc. Daily PN&
	Type int64 `json:"type"`
	// can not be empty, can pass default value "*"
	Value string `json:"value"`
}

// long positions allocation
type Allocation_List_Item_Group_Long struct {
	Apparel        float64 `json:"Apparel"`
	Chemicals      float64 `json:"Chemicals"`
	Communications float64 `json:"Communications"`
	Computers      float64 `json:"Computers"`
	Others         float64 `json:"Others"`
	Semiconductors float64 `json:"Semiconductors"`
}

// security definition information
type Secdef_List_Item struct {
	AssetClass     string  `json:"assetClass"`
	Conid          int64   `json:"conid"`
	Expiry         string  `json:"expiry"`
	FullName       string  `json:"fullName"`
	Group          string  `json:"group"`
	LastTradingDay string  `json:"lastTradingDay"`
	Name           string  `json:"name"`
	PageSize       int64   `json:"pageSize"`
	PutOrCall      string  `json:"putOrCall"`
	Sector         string  `json:"sector"`
	SectorGroup    string  `json:"sectorGroup"`
	Strike         float64 `json:"strike"`
	Ticker         string  `json:"ticker"`
	UndConid       int64   `json:"undConid"`
}

type SetAccount struct {
	// Account ID
	AcctId string `json:"acctId"`
}

type Iserver_Account_AccountId_Order_Whatif_POST_200 struct {
	Amount      *Iserver_Account_AccountId_Order_Whatif_POST_200_Amount      `json:"amount"`
	Equity      *Iserver_Account_AccountId_Order_Whatif_POST_200_Equity      `json:"equity"`
	Error       string                                                       `json:"error"`
	Initial     *Iserver_Account_AccountId_Order_Whatif_POST_200_Initial     `json:"initial"`
	Maintenance *Iserver_Account_AccountId_Order_Whatif_POST_200_Maintenance `json:"maintenance"`
	Warn        string                                                       `json:"warn"`
}

type Pa_Performance_POST_Request struct {
	AcctIds []string `json:"acctIds"`
	// Frequency of cumulative performance data points: 'D'aily, 'M'onthly,'Q'uarterly.
	Freq string `json:"freq"`
}

type Iserver_Account_POST_200 struct {
	AcctId string `json:"acctId"`
	Set    bool   `json:"set"`
}

type Iserver_Marketdata_Unsubscribeall_GET_200 struct {
	// true means market data is canceled, false means it is not.
	Confirmed bool `json:"confirmed"`
}

type Iserver_Reply_Replyid_POST_400 struct {
	// for example-order not confirmed
	Error      string `json:"error"`
	StatusCode int64  `json:"statusCode"`
}

type ScannerResult_Contracts_Contract_List_Item struct {
	ContractID int64  `json:"contractID"`
	Distance   int64  `json:"distance"`
	InScanTime string `json:"inScanTime"`
}

type Iserver_Scanner_Params_GET_200_Location_tree_List_Item_Locations_List_Item struct {
	Display_name string `json:"display_name"`
	Type         string `json:"type"`
}

type Iserver_Scanner_Params_GET_200 struct {
	Filter_list     []*Iserver_Scanner_Params_GET_200_Filter_list_List_Item     `json:"filter_list"`
	Instrument_list []*Iserver_Scanner_Params_GET_200_Instrument_list_List_Item `json:"instrument_list"`
	Location_tree   []*Iserver_Scanner_Params_GET_200_Location_tree_List_Item   `json:"location_tree"`
	Scan_type_list  []*Iserver_Scanner_Params_GET_200_Scan_type_list_List_Item  `json:"scan_type_list"`
}

type Iserver_Contract_Conid_InfoAndRules_GET_200 struct {
	// Allowed to sell shares that you own
	Allow_sell_long bool `json:"allow_sell_long"`
	// Classification of Financial Instrument codes
	Cfi_code   string `json:"cfi_code"`
	Classifier string `json:"classifier"`
	// Contracts company name
	Company_name string `json:"company_name"`
	// IBKRs contract identifier
	Con_id float64 `json:"con_id"`
	// Month the contract must be satisfied by making or accepting delivery
	Contract_month string `json:"contract_month"`
	// Currency contract trades in
	Currency string `json:"currency"`
	Cusip    string `json:"cusip"`
	// Primary Exchange, Routing or Trading Venue
	Exchange string `json:"exchange"`
	// Expiration Date in the format YYYYMMDD
	Expiry_full float64 `json:"expiry_full"`
	// Specific group of companies or businesses.
	Industry string `json:"industry"`
	// Asset Class of the contract
	Instrument_type string `json:"instrument_type"`
	// Supports zero commission trades
	Is_zero_commission_security bool `json:"is_zero_commission_security"`
	// Contracts symbol from primary exchange. For options it is the OCC symbol.
	Local_symbol string `json:"local_symbol"`
	// Date on which the underlying transaction settles if the option is exercised
	Maturity_date float64 `json:"maturity_date"`
	// numerical value of each point of price movement
	Multiplier string `json:"multiplier"`
	// Provides trading outside of Regular Trading Hours
	R_t_h bool `json:"r_t_h"`
	// Put or Call of the option
	Right string                                                         `json:"right"`
	Rules []*Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item `json:"rules"`
	// Support IBKRs SMART routing
	Smart_available bool `json:"smart_available"`
	// fixed price at which the owner of the option buys or sells the underlying
	Strike string `json:"strike"`
	// Underlying Symbol for contract
	Symbol string `json:"symbol"`
	// Formatted contract parameters
	Text string `json:"text"`
	// Designation of the contract
	Trading_class string `json:"trading_class"`
	// IBKRs contract identifier for the underlying instrument
	Underlying_con_id float64 `json:"underlying_con_id"`
	// Legal entity for underlying contract
	Underlying_issuer string `json:"underlying_issuer"`
	// Comma separated list of exchanges or trading venues
	Valid_exchanges string `json:"valid_exchanges"`
}

type Iserver_Account_AccountId_Alert_AlertId_DELETE_200 struct {
	Failure_list string `json:"failure_list"`
	Order_id     int64  `json:"order_id"`
	Order_status string `json:"order_status"`
	Request_id   int64  `json:"request_id"`
	Success      bool   `json:"success"`
	Text         string `json:"text"`
}

type Iserver_Account_AccountId_Orders_POST_200_List_Item struct {
	Id string `json:"id"`
	// Please note here, if the message is a question, you have to reply to question in
	// order to submit
	// the order successfully. See more in the "/iserver/reply/{replyid}" endpoint.
	Message []string `json:"message"`
}

type Iserver_Scanner_Params_GET_200_Filter_list_List_Item struct {
	Code         string `json:"code"`
	Display_name string `json:"display_name"`
	Group        string `json:"group"`
	Type         string `json:"type"`
}

type Iserver_Account_Orders_FaGroup_POST_200_List_Item struct {
	Id string `json:"id"`
	// Please note here, if the message is a question, you have to reply to question in
	// order to submit
	// the order successfully. See more in the "/iserver/reply/{replyid}" endpoint.
	Message []string `json:"message"`
}

type Fyi_Settings_Typecode_POST_Request struct {
	Enabled bool `json:"enabled"`
}

type Iserver_Scanner_Params_GET_200_Instrument_list_List_Item struct {
	Display_name string   `json:"display_name"`
	Filters      []string `json:"filters"`
	Type         string   `json:"type"`
}

type Iserver_Account_AccountId_Orders_POST_Request struct {
	// Notes for bracket orders: 1. Children orders will not have its own "cOID", so please
	// donot pass "cOID"
	// parameter in child order.Instead, they will have a "parentId" which must be equal
	// to "cOID" of parent.
	// 2. When you cancel a parent order, it will cancel all bracket orders, when you cancel
	// one child order,
	// it will also cancel its sibling order.
	Orders []*OrderRequest `json:"orders"`
}

type Tickle_POST_200_Iserver struct {
	AuthStatus *Tickle_POST_200_Iserver_AuthStatus `json:"authStatus"`
	Tickle     bool                                `json:"tickle"`
}

// portfolio allocation by asset class
type Allocation_List_Item_AssetClass struct {
	// long positions allocation
	Long *Allocation_List_Item_AssetClass_Long `json:"long"`
	// short positions allocation
	Short *Allocation_List_Item_AssetClass_Short `json:"short"`
}

type Events_List_Item struct {
	Conids []string `json:"conids"`
	// will be different for different event types
	Data map[string]interface{} `json:"data"`
	// for example 11662135
	Event_key  string `json:"event_key"`
	Event_type string `json:"event_type"`
	// for exmple 20180817T040000+0000
	Index_date      string `json:"index_date"`
	Index_date_type string `json:"index_date_type"`
	// for example RSE
	Source   string                 `json:"source"`
	Status   string                 `json:"status"`
	Tooltips map[string]interface{} `json:"tooltips"`
}

// long positions allocation
type Allocation_List_Item_Sector_Long struct {
	Communications float64 `json:"Communications"`
	Energy         float64 `json:"Energy"`
	Financial      float64 `json:"Financial"`
	Others         float64 `json:"Others"`
	Technology     float64 `json:"Technology"`
	Utilities      float64 `json:"Utilities"`
}

type Ibcust_Entity_Info_GET_200_List_Item struct {
	ApplicantId float64                                                    `json:"applicantId"`
	Entities    []*Ibcust_Entity_Info_GET_200_List_Item_Entities_List_Item `json:"entities"`
}

type Portfolio_Allocation_POST_Request struct {
	AcctIds []string `json:"acctIds"`
}

type Iserver_Secdef_Strikes_GET_500 struct {
	Error string `json:"error"`
}

// Returns tradingTime in exchange time zone.
type Trsrv_Secdef_Schedule_GET_200_Schedules_List_Item_TradingTimes struct {
	CancelDayOrders string `json:"cancelDayOrders"`
	ClosingTime     int64  `json:"closingTime"`
	OpeningTime     int64  `json:"openingTime"`
}

type AlertRequest struct {
	// The message you want to receive via email or text message
	AlertMessage string `json:"alertMessage"`
	// name of alert
	AlertName string `json:"alertName"`
	// whether alert is repeatable or not, so value can only be 0 or 1, this has to be 1
	// for MTA alert
	AlertRepeatable int64                                `json:"alertRepeatable"`
	Conditions      []*AlertRequest_Conditions_List_Item `json:"conditions"`
	// email address to receive alert
	Email string `json:"email"`
	// format, YYYYMMDD-HH:mm:ss, please NOTE this will only work when tif is GTD
	ExpireTime string `json:"expireTime"`
	// value can only be 0 or 1, set to 1 to enable the alert only in IBKR mobile
	ITWSOrdersOnly int64 `json:"iTWSOrdersOnly"`
	// orderId is required when modifying alert. You can get it from /iserver/account/:accountId/alerts
	OrderId int64 `json:"orderId"`
	// value can only be 0 or 1, set to 1 if the alert can be triggered outside regular
	// trading hours.
	OutsideRth int64 `json:"outsideRth"`
	// audio message to play when alert is triggered
	PlayAudio string `json:"playAudio"`
	// whether allowing to send email or not, so value can only be 0 or 1,
	SendMessage int64 `json:"sendMessage"`
	// value can only be 0 or 1, set to 1 to allow to show alert in pop-ups
	ShowPopup int64 `json:"showPopup"`
	// time in force, can only be GTC or GTD
	Tif string `json:"tif"`
	// for MTA alert only, each user has a unique toolId and it will stay the same, do not
	// send for normal alert
	ToolId int64 `json:"toolId"`
}

type AlertResponse struct {
	// account id
	Account string `json:"account"`
	// whether alert is active or not, so value can only be 0 or 1
	Alert_active int64 `json:"alert_active"`
	// MTA alert only
	Alert_default_type string `json:"alert_default_type"`
	// email address to receive alert
	Alert_email string `json:"alert_email"`
	// The message you want to receive via email or text message
	Alert_message string `json:"alert_message"`
	// MTA alert only
	Alert_mta_currency string `json:"alert_mta_currency"`
	// MTA alert only
	Alert_mta_defaults string `json:"alert_mta_defaults"`
	// name of alert
	Alert_name string `json:"alert_name"`
	// audio message to play when alert is triggered
	Alert_play_audio string `json:"alert_play_audio"`
	// whether alert is repeatable or not, so value can only be 0 or 1
	Alert_repeatable int64 `json:"alert_repeatable"`
	// whether allowing to send email or not, so value can only be 0 or 1,
	Alert_send_message int64 `json:"alert_send_message"`
	// value can only be 0 or 1, set to 1 to allow to show alert in pop-ups
	Alert_show_popup int64 `json:"alert_show_popup"`
	// whether the alert has been triggered
	Alert_triggered bool `json:"alert_triggered"`
	// whether allowing the condition can be triggered outside of regular trading hours,
	// 1 means allow
	Condition_outside_rth int64 `json:"condition_outside_rth"`
	// size of conditions array
	Condition_size int64                                 `json:"condition_size"`
	Conditions     []*AlertResponse_Conditions_List_Item `json:"conditions"`
	// format, YYYYMMDD-HH:mm:ss
	Expire_time string `json:"expire_time"`
	// value can only be 0 or 1, set to 1 to enable the alert only in IBKR mobile
	Itws_orders_only int64 `json:"itws_orders_only"`
	Order_id         int64 `json:"order_id"`
	// whether the alert can be edited
	Order_not_editable bool `json:"order_not_editable"`
	// status of alert
	Order_status string `json:"order_status"`
	// value can only be 0 or 1, set to 1 if the alert can be triggered outside regular
	// trading hours.
	OutsideRth int64 `json:"outsideRth"`
	// time in force, can only be GTC or GTD
	Tif string `json:"tif"`
	// MTA alert only
	Time_zone string `json:"time_zone"`
	// for MTA alert only, each user has a unique toolId and it will stay the same, do not
	// send for normal alert
	Tool_id int64 `json:"tool_id"`
}

type AuthStatus struct {
	// Brokerage session is authenticated
	Authenticated bool `json:"authenticated"`
	// Brokerage session is competing, e.g. user is logged in to IBKR Mobile, WebTrader,
	// TWS or other trading platforms.
	Competing bool `json:"competing"`
	// Connected to backend
	Connected bool `json:"connected"`
	// Authentication failed, why.
	Fail string `json:"fail"`
	// System messages that may affect trading
	Message string `json:"message"`
	// Prompt messages that may affect trading or the account
	Prompts []string `json:"prompts"`
}

type Iserver_Marketdata_Snapshot_GET_200_List_Item struct {
	// Last Price
	LastPrice string `json:"31"`
	// Symbol
	Symbol string `json:"55"`
	// Text
	Text string `json:"58"`
	// High
	High string `json:"70"`
	// Low
	Low string `json:"71"`
	// Position
	Position string `json:"72"`
	// Market Value
	MarketValue string `json:"73"`
	// Average Price
	AveragePrice string `json:"74"`
	// Unrealized PnL
	UnrealizedPnL string `json:"75"`
	// Formatted position
	FormattedPosition string `json:"76"`
	// Formatted Unrealized PnL
	FormattedUnrealizedPnL string `json:"77"`
	// Daily PnL
	DailyPnL string `json:"78"`
	// Change Price
	ChangePrice string `json:"82"`
	// Change Percent
	ChangePercent string `json:"83"`
	// Bid Price
	BidPrice string `json:"84"`
	// Ask Size
	AskSize string `json:"85"`
	// Ask Price
	AskPrice string `json:"86"`
	// Volume
	Volume string `json:"87"`
	// Bid Size
	BidSize string `json:"88"`
	// Exchange
	Exchange string `json:"6004"`
	// Conid
	Conid6008 string `json:"6008"`
	// Security Type
	SecurityType string `json:"6070"`
	// Months
	Months string `json:"6072"`
	// Regular Expiry
	RegularExpiry string `json:"6073"`
	// Marker for market data delivery method (similar to request id)
	Marker string `json:"6119"`
	// Underlying Conid. Use /trsrv/secdef to get more information about the security
	UnderlyingConid string `json:"6457"`
	// Market Data Availability. The field may contain two chars. The first char is the
	// primary code: R = Realtime, D = Delayed,
	// Z = Frozen, Y = Frozen Delayed. The second char is the secondary code: P = Snapshot
	// Available, p = Consolidated.
	MarketDataAvailability string `json:"6509"`
	// Company name
	CompanyName string `json:"7051"`
	// Last Size
	LastSize string `json:"7059"`
	// Conid + Exchange
	ConidAndExchange string `json:"7094"`
	// Contract Description
	ContractDescription string `json:"7219"`
	// Contract Description
	ContractDescription2 string `json:"7220"`
	// Listing Exchange
	ListingExchange string `json:"7221"`
	// Industry
	Industry string `json:"7280"`
	// Category
	Category string `json:"7281"`
	// Average Daily Volume
	AverageDailyVolume string `json:"7282"`
	// Historic Volume (30d)
	HistoricVolume30d string `json:"7284"`
	// Put/Call Ratio
	PutCallRatio string `json:"7285"`
	// Dividend Amount
	DividendAmount string `json:"7286"`
	// Dividend Yield %
	DividendYieldPercent string `json:"7287"`
	// Ex-date of the dividend
	ExDateOfTheDividend string `json:"7288"`
	// Market Cap
	MarketCap string `json:"7289"`
	// P/E
	PE string `json:"7290"`
	// EPS
	EPS string `json:"7291"`
	// Cost Basis
	CostBasis string `json:"7292"`
	// 52 Week High
	High52Week string `json:"7293"`
	// 52 Week Low
	Low52Week string `json:"7294"`
	// Open Price
	OpenPrice string `json:"7295"`
	// Close Price
	ClosePrice string `json:"7296"`
	// Delta
	Delta string `json:"7308"`
	// Gamma
	Gamma string `json:"7309"`
	// Theta
	Theta string `json:"7310"`
	// Vega
	Vega string `json:"7311"`
	// Implied volatility of the option
	ImpliedVolatility string `json:"7633"`
	Updated           int64  `json:"_updated"`
	Conid             int64  `json:"conid"`
	Server_id         string `json:"server_id"`
}

type Iserver_Marketdata_Snapshot_GET_400 struct {
	Error      string `json:"error"`
	StatusCode int64  `json:"statusCode"`
}

type Iserver_Scanner_Run_POST_200_List_Item struct {
	Available_chart_periods string  `json:"available_chart_periods"`
	Column_name             string  `json:"column_name"`
	Company_name            string  `json:"company_name"`
	Con_id                  float64 `json:"con_id"`
	Conidex                 string  `json:"conidex"`
	Contract_description_1  string  `json:"contract_description_1"`
	Listing_exchange        string  `json:"listing_exchange"`
	Sec_type                string  `json:"sec_type"`
	Server_id               string  `json:"server_id"`
	Symbol                  string  `json:"symbol"`
}

type Calendar_request_Filters struct {
	// value can be 'true' or 'false'.
	DivExDates string `json:"DivExDates"`
	// value can be 'true' or 'false'.
	Corporate_earnings string `json:"corporate_earnings"`
	// value can be 'true' or 'false'.
	Corporate_events string `json:"corporate_events"`
	// default is 'All'.
	Country string `json:"country"`
	// value can be 'true' or 'false'.
	Economic_events string `json:"economic_events"`
	// value can be 'true' or 'false'.
	Ipo string `json:"ipo"`
	// default is '250'.
	Limit string `json:"limit"`
	// default is '50'.
	Limit_region string `json:"limit_region"`
	// value can be 'true' or 'false'.
	Option_show_monthly string `json:"option_show_monthly"`
	// value can be 'true' or 'false'.
	Option_show_weekly string `json:"option_show_weekly"`
	// value can be 'true' or 'false'.
	Recently_held string `json:"recently_held"`
	// value can be 'true' or 'false'.
	Splits string `json:"splits"`
}

type Summary_Total struct {
	// total change amount
	Chg    string `json:"chg"`
	EndVal string `json:"endVal"`
	// set to true if any external account data is not available for starting or ending
	// date, resulting in potentially unusual total values.
	IncompleteData bool `json:"incompleteData"`
	// change percent
	Rtn      string `json:"rtn"`
	StartVal string `json:"startVal"`
}

type Performance_Tpps_Data_List_Item struct {
	BaseCurrency string `json:"baseCurrency"`
	// end date-- yyyyMMdd
	End string `json:"end"`
	Id  string `json:"id"`
	// for example-- acctid
	IdType string `json:"idType"`
	// each value stands for price change percent of corresponding date in dates array
	Returns []float64 `json:"returns"`
	// start date-- yyyyMMdd
	Start string `json:"start"`
}

// Time period performance data
type Performance_Tpps struct {
	Data []*Performance_Tpps_Data_List_Item `json:"data"`
	// array of dates, the length should be same as the length of returns inside data.
	Dates []string `json:"dates"`
	// M means Month
	Freq string `json:"freq"`
}

// future contract information
type Stocks_List_Item struct {
	AssetClass string `json:"assetClass"`
	// company name in Chinese
	ChineseName string `json:"chineseName"`
	// array of contracts from different exchanges
	Contracts []*Stocks_List_Item_Contracts_List_Item `json:"contracts"`
	// company name
	Name string `json:"name"`
}

// notification
type Notifications_List_Item struct {
	// notification date
	D string `json:"D"`
	// FYI code, we can use it to find whether the disclaimer is accepted or not in settings
	FC string `json:"FC"`
	// unique way to reference this notification
	ID string `json:"ID"`
	// content of notification
	MD string `json:"MD"`
	// title of notification
	MS string `json:"MS"`
	// 0-unread, 1-read
	R string `json:"R"`
}

type ScannerResult struct {
	// Contains list of contracts matching the scanner query
	Contracts *ScannerResult_Contracts `json:"Contracts"`
	Id        float64                  `json:"id"`
	Offset    int64                    `json:"offset"`
	Position  string                   `json:"position"`
	ScanTime  string                   `json:"scanTime"`
	Size      int64                    `json:"size"`
	Total     int64                    `json:"total"`
}

type OrderRequest struct {
	// acctId is optional. It should be one of the accounts returned by
	// /iserver/accounts. If not passed, the first one in the list is selected.
	AcctId string `json:"acctId"`
	// Set the allocation method when placing an order using an FA account for a group
	// Possible allocation methods are "NetLiquidity", "AvailableEquity", "EqualQuantity"
	// and "PctChange".
	AllocationMethod string `json:"allocationMethod"`
	// Customer Order ID. An arbitraty string that can be used to identify the order, e.g
	// "my-fb-order". The
	// value must be unique for a 24h span. Please do not set this value for child orders
	// when placing a bracket order.
	COID string `json:"cOID"`
	// conid is the identifier of the security you want to trade, you can find the
	// conid with /iserver/secdef/search.
	Conid int64 `json:"conid"`
	// double number, this is the cash quantity field which can only be used for FX conversion
	// order.
	FxQty float64 `json:"fxQty"`
	// set to true if the order is a FX conversion order
	IsCurrencyConversion bool `json:"isCurrencyConversion"`
	// listingExchange is optional. By default we use "SMART" routing. Possible values are
	// available via this end
	// point: /v1/portal/iserver/contract/{{conid}}/info, see valid_exchange: e.g: SMART,AMEX,NYSE,
	// CBOE,ISE,CHX,ARCA,ISLAND,DRCTEDGE,BEX,BATS,EDGEA,CSFBALGO,JE FFALGO,BYX,IEX,FOXRIVER,TPLUS1,NYSENAT,PSX
	ListingExchange string `json:"listingExchange"`
	// orderType can be one of MKT (Market), LMT (Limit), STP (Stop) or STP_LIMIT (stop
	// limit)
	OrderType string `json:"orderType"`
	// set to true if the order can be executed outside regular trading hours.
	OutsideRTH bool `json:"outsideRTH"`
	// When placing bracket orders, the child parentId must be equal to the cOId (customer
	// order id) of the parent.
	ParentId string `json:"parentId"`
	// optional if order is MKT, for LMT, this is the limit price. For STP this is the stop
	// price.
	Price float64 `json:"price"`
	// usually integer, for some special cases can be float numbers
	Quantity float64 `json:"quantity"`
	// for example QuickTrade
	Referrer string `json:"referrer"`
	// conid:type for example 265598:STK
	SecType string `json:"secType"`
	// SELL or BUY
	Side   string `json:"side"`
	Ticker string `json:"ticker"`
	// GTC (Good Till Cancel) or DAY. DAY orders are automatically cancelled at the end
	// of the Day or Trading hours.
	Tif string `json:"tif"`
	// If true, the system will use the Adaptive Algo to submit the order
	// https://www.interactivebrokers.com/en/index.php?f=19091
	UseAdaptive bool `json:"useAdaptive"`
}

type HistoryData struct {
	// The number of seconds in a bar
	BarLength int64                         `json:"barLength"`
	Data      []*HistoryData_Data_List_Item `json:"data"`
	// High value during this time series with format %h/%v/%t. %h is the high price (scaled
	// by priceFactor), %v is volume
	// (volume factor will always be 100 (reported volume = actual volume/100)) and %t is
	// minutes from start time of the chart
	High string `json:"high"`
	// Low value during this time series with format %l/%v/%t. %l is the low price (scaled
	// by priceFactor), %v is volume
	// (volume factor will always be 100 (reported volume = actual volume/100)) and %t is
	// minutes from start time of the chart
	Low string `json:"low"`
	// Market Data Availability. The field may contain two chars. The first char is the
	// primary code: S = Streaming, R = Realtime, D = Delayed,
	// Z = Frozen, Y = Frozen Delayed. The second char is the secondary code: P = Snapshot
	// Available, p = Consolidated.
	MdAvailability string `json:"mdAvailability"`
	MessageVersion int64  `json:"messageVersion"`
	// The time it takes, in milliseconds, to process the historical data request
	MktDataDelay    int64 `json:"mktDataDelay"`
	NegativeCapable bool  `json:"negativeCapable"`
	// The historical data returned includes outside of regular trading hours
	OutsideRth bool `json:"outsideRth"`
	// total number of points
	Points            int64  `json:"points"`
	PriceDisplayRule  int64  `json:"priceDisplayRule"`
	PriceDisplayValue string `json:"priceDisplayValue"`
	// priceFactor is price increment obtained from display rule
	PriceFactor int64 `json:"priceFactor"`
	// start date time in the format YYYYMMDD-HH:mm:ss
	StartTime string `json:"startTime"`
	// Underlying Symbol of the corresponding contract
	Symbol string `json:"symbol"`
	// companyName of the corresponding contract
	Text string `json:"text"`
	// The duration for the historical data request
	TimePeriod string `json:"timePeriod"`
	// The number of seconds in the trading day
	TradingDayDuration int64 `json:"tradingDayDuration"`
	TravelTime         int64 `json:"travelTime"`
	VolumeFactor       int64 `json:"volumeFactor"`
}

// Net asset value data for the account or consolidated accounts. NAV data is not applicable
// to benchmarks.
type Performance_Nav struct {
	Data []*Performance_Nav_Data_List_Item `json:"data"`
	// array of dates, the length should be same as the length of returns inside data.
	Dates []string `json:"dates"`
	// D means Day
	Freq string `json:"freq"`
}

type Iserver_Account_AccountId_Alert_Activate_POST_200 struct {
	Failure_list string `json:"failure_list"`
	Order_id     int64  `json:"order_id"`
	Order_status string `json:"order_status"`
	Request_id   int64  `json:"request_id"`
	Success      bool   `json:"success"`
	Text         string `json:"text"`
}

type Iserver_Account_Orders_GET_200_Orders_List_Item struct {
	// Account number
	Acct string `json:"acct"`
	// background color in hex format
	BgColor string `json:"bgColor"`
	// Cash currency
	CashCcy string `json:"cashCcy"`
	// Company Name
	CompanyName string `json:"companyName"`
	// Contract identifier
	Conid float64 `json:"conid"`
	// Formatted ticker description
	Description1 string `json:"description1"`
	// Exchange or trading venue
	Exchange string `json:"exchange"`
	// foreground color in hex format
	FgColor string `json:"fgColor"`
	// Quantity filled
	FilledQuantity float64 `json:"filledQuantity"`
	// Last status update in format YYMMDDhhmms based in GMT
	LastExecutionTime float64 `json:"lastExecutionTime"`
	// Last status update unix time in ms
	LastExecutionTime_r float64 `json:"lastExecutionTime_r"`
	// Listing Exchange
	ListingExchange string `json:"listingExchange"`
	// Order description
	OrderDesc string `json:"orderDesc"`
	// Order identifier
	OrderId string `json:"orderId"`
	// Order type
	OrderType string `json:"orderType"`
	// Order reference
	Order_ref string `json:"order_ref"`
	// Original order type
	OrigOrderType string `json:"origOrderType"`
	// Price of order
	Price float64 `json:"price"`
	// Quantity remaining
	RemainingQuantity float64 `json:"remainingQuantity"`
	// Asset class
	SecType string `json:"secType"`
	// Side of order
	Side string `json:"side"`
	// Quantity outstanding and total quantity concatenated with forward slash separator
	SizeAndFills float64 `json:"sizeAndFills"`
	// Status of the order
	Status string `json:"status"`
	// Supports Tax Optimization with 0 for no and 1 for yes
	SupportsTaxOpt float64 `json:"supportsTaxOpt"`
	// Underlying symbol
	Ticker string `json:"ticker"`
	// Time in force
	TimeInForce string `json:"timeInForce"`
}

type Fyi_Deliveryoptions_GET_200 struct {
	E []*Fyi_Deliveryoptions_GET_200_E_List_Item `json:"E"`
	// Email option is enabled or not 0-off, 1-on.
	M int64 `json:"M"`
}

type Inds_List_Item struct {
	Conid float64 `json:"conid"`
}

type Iserver_Scanner_Params_GET_200_Scan_type_list_List_Item struct {
	Code         string   `json:"code"`
	Display_name string   `json:"display_name"`
	Instruments  []string `json:"instruments"`
}

type Trsrv_Secdef_POST_Request struct {
	Conids []int64 `json:"conids"`
}

// Contains some basic info of contract
type SecdefInfo struct {
	Conid float64 `json:"conid"`
	// Annual interest rate paid on a bond
	Coupon string `json:"coupon"`
	// Currency the contract trades in
	Currency string `json:"currency"`
	// Committee on Uniform Securities Identification Procedures number
	Cusip string `json:"cusip"`
	// Formatted symbol
	Desc1 string `json:"desc1"`
	// Formatted expiration, strike and right
	Desc2           string `json:"desc2"`
	Exchange        string `json:"exchange"`
	ListingExchange string `json:"listingExchange"`
	// Format YYYYMMDD, the date on which the underlying transaction settles if the option
	// is exercised
	MaturityDate string `json:"maturityDate"`
	// total premium paid or received for an option contract
	Multiplier string `json:"multiplier"`
	// C = Call Option, P = Put Option
	Right   string `json:"right"`
	SecType string `json:"secType"`
	// The strike price also known as exercise price
	Strike string `json:"strike"`
	// For example IBKR
	Symbol         string `json:"symbol"`
	TradingClass   string `json:"tradingClass"`
	ValidExchanges string `json:"validExchanges"`
}

type Iserver_Marketdata_Conid_Unsubscribe_GET_200 struct {
	// success means market data was canceled.
	Confirmed string `json:"confirmed"`
}

type Logout_POST_200 struct {
	// true means username is still logged in, false means it is not
	Confirmed bool `json:"confirmed"`
}

type ScannerParams_LocationTree_Location_List_Item struct {
	DisplayName   string `json:"displayName"`
	Instruments   string `json:"instruments"`
	LocationCode  string `json:"locationCode"`
	RouteExchange string `json:"routeExchange"`
}

// portfolio allocation by group
type Allocation_List_Item_Group struct {
	// long positions allocation
	Long *Allocation_List_Item_Group_Long `json:"long"`
	// short positions allocation
	Short *Allocation_List_Item_Group_Short `json:"short"`
}

// short positions allocation
type Allocation_List_Item_Group_Short struct {
	Airlines float64 `json:"Airlines"`
	Banks    float64 `json:"Banks"`
	Internet float64 `json:"Internet"`
}

type Contract_Rules struct {
	// default quantity you can use to place an order
	DefaultSize float64 `json:"defaultSize"`
	DisplaySize string  `json:"displaySize"`
	Increment   string  `json:"increment"`
	// default limit price you can use to prefill your order
	LimitPrice        float64  `json:"limitPrice"`
	OrderTypes        []string `json:"orderTypes"`
	OrderTypesOutside []string `json:"orderTypesOutside"`
	// if you can preview the order or not with the whatif endpoint
	Preview       bool    `json:"preview"`
	SizeIncrement float64 `json:"sizeIncrement"`
	// default stop price you can use to prefill your order
	Stopprice float64  `json:"stopprice"`
	TifTypes  []string `json:"tifTypes"`
}

type Performance_Cps_Data_List_Item struct {
	BaseCurrency string `json:"baseCurrency"`
	// end date-- yyyyMMdd
	End string `json:"end"`
	Id  string `json:"id"`
	// for example-- acctid
	IdType string `json:"idType"`
	// each value stands for price change percent of corresponding date in dates array
	Returns []float64 `json:"returns"`
	// start date-- yyyyMMdd
	Start string `json:"start"`
}

type Iserver_Account_AccountId_Order_Whatif_POST_200_Maintenance struct {
	After   string `json:"after"`
	Change  string `json:"change"`
	Current string `json:"current"`
}

type Iserver_Account_AccountId_Order_Whatif_POST_200_Amount struct {
	// for example 23,000 USD
	Amount string `json:"amount"`
	// for example 1.1 ... 1.2 USD
	Commission string `json:"commission"`
	Total      string `json:"total"`
}

type Iserver_Secdef_Info_GET_500 struct {
	Error string `json:"error"`
}

type Iserver_Account_AccountId_Order_OrderId_DELETE_200 struct {
	Account  string `json:"account"`
	Conid    int64  `json:"conid"`
	Msg      string `json:"msg"`
	Order_id string `json:"order_id"`
}

type Iserver_Secdef_Search_POST_200_List_Item struct {
	// Company Name - Exchange
	CompanyHeader string `json:"companyHeader"`
	CompanyName   string `json:"companyName"`
	// Contract Identifier
	Conid int64 `json:"conid"`
	// Exchange
	Description string                                                         `json:"description"`
	Fop         string                                                         `json:"fop"`
	Opt         string                                                         `json:"opt"`
	Restricted  string                                                         `json:"restricted"`
	Sections    []*Iserver_Secdef_Search_POST_200_List_Item_Sections_List_Item `json:"sections"`
	Symbol      string                                                         `json:"symbol"`
	War         string                                                         `json:"war"`
}

// Contains list of instruments for which scanner can be ran
type ScannerParams_InstrumentList struct {
	// Contains information like name, supported filters, etc. for an instrument
	Instrument []*ScannerParams_InstrumentList_Instrument_List_Item `json:"Instrument"`
}

// account transactions
type Transactions struct {
	// same as request
	Currency string `json:"currency"`
	// Period start date. Epoch time, GMT
	From float64 `json:"from"`
	// will always be getTransactions
	Id string `json:"id"`
	// Indicates whether current day and realtime data is included in the result
	IncludesRealTime bool `json:"includesRealTime"`
	// Period end date. Epoch time, GMT
	To float64 `json:"to"`
	// Sorted by date descending
	Transactions []*Transactions_Transactions_List_Item `json:"transactions"`
}

// long positions allocation
type Allocation_List_Item_AssetClass_Long struct {
	BOND float64 `json:"BOND"`
	CASH float64 `json:"CASH"`
	FUT  float64 `json:"FUT"`
	OPT  float64 `json:"OPT"`
	STK  float64 `json:"STK"`
	WAR  float64 `json:"WAR"`
}

// allocation
type Allocation_List_Item struct {
	// portfolio allocation by asset class
	AssetClass *Allocation_List_Item_AssetClass `json:"assetClass"`
	// portfolio allocation by group
	Group *Allocation_List_Item_Group `json:"group"`
	// portfolio allocation by sector
	Sector *Allocation_List_Item_Sector `json:"sector"`
}

type Iserver_Secdef_Strikes_GET_200 struct {
	Call []string `json:"call"`
	Put  []string `json:"put"`
}

type Fyi_Unreadnumber_GET_200 struct {
	// unread number
	BN int64 `json:"BN"`
}

type Iserver_Account_AccountId_Alerts_GET_200_List_Item struct {
	// account id
	Account string `json:"account"`
	// Value can only be 0 or 1, 1 means active
	Alert_active int64  `json:"alert_active"`
	Alert_name   string `json:"alert_name"`
	// whether the alert can be repeatable or not, value can be 1 or 0. 1 means true
	Alert_repeatable int64 `json:"alert_repeatable"`
	// whether the alert has been triggered or not
	Alert_triggered bool  `json:"alert_triggered"`
	Order_id        int64 `json:"order_id"`
	// format, YYYYMMDD-HH:mm:ss, the time when you created the alert
	Order_time string `json:"order_time"`
}

type Iserver_Secdef_Search_POST_500 struct {
	Error string `json:"error"`
}

type Summary_BalanceByDate struct {
	Series []*Summary_BalanceByDate_Series_List_Item `json:"series"`
}

// Account Information
type Position_List_Item struct {
	AcctId            string   `json:"acctId"`
	AssetClass        string   `json:"assetClass"`
	AvgCost           float64  `json:"avgCost"`
	AvgPrice          float64  `json:"avgPrice"`
	BaseAvgCost       float64  `json:"baseAvgCost"`
	BaseAvgPrice      float64  `json:"baseAvgPrice"`
	BaseMktPrice      float64  `json:"baseMktPrice"`
	BaseMktValue      float64  `json:"baseMktValue"`
	BaseRealizedPnl   float64  `json:"baseRealizedPnl"`
	BaseUnrealizedPnl float64  `json:"baseUnrealizedPnl"`
	ConExchMap        []string `json:"conExchMap"`
	Conid             int64    `json:"conid"`
	ContractDesc      string   `json:"contractDesc"`
	Currency          string   `json:"currency"`
	Exchs             string   `json:"exchs"`
	ExerciseStyle     string   `json:"exerciseStyle"`
	Expiry            string   `json:"expiry"`
	FullName          string   `json:"fullName"`
	Group             string   `json:"group"`
	LastTradingDay    string   `json:"lastTradingDay"`
	MktPrice          float64  `json:"mktPrice"`
	MktValue          float64  `json:"mktValue"`
	Model             string   `json:"model"`
	Multiplier        float64  `json:"multiplier"`
	Name              string   `json:"name"`
	PageSize          int64    `json:"pageSize"`
	Position          float64  `json:"position"`
	PutOrCall         string   `json:"putOrCall"`
	RealizedPnl       float64  `json:"realizedPnl"`
	Sector            string   `json:"sector"`
	SectorGroup       string   `json:"sectorGroup"`
	Strike            float64  `json:"strike"`
	Ticker            string   `json:"ticker"`
	UndComp           string   `json:"undComp"`
	UndConid          int64    `json:"undConid"`
	UndSym            string   `json:"undSym"`
	UnrealizedPnl     float64  `json:"unrealizedPnl"`
}

type Ibcust_Entity_Info_GET_200_List_Item_Entities_List_Item_Address struct {
	City        string `json:"city"`
	Compact     string `json:"compact"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	PostalCode  string `json:"postalCode"`
	State       string `json:"state"`
	Street      string `json:"street"`
	Street2     string `json:"street2"`
}

type Iserver_Marketdata_History_GET_429 struct {
	Error string `json:"error"`
}

type Pa_Transactions_POST_Request struct {
	AcctIds []string  `json:"acctIds"`
	Conids  []float64 `json:"conids"`
	// optional defaults to USD.
	Currency string `json:"currency"`
	// optional, default value is 90
	Days float64 `json:"days"`
}

type Fyi_Settings_GET_200_List_Item struct {
	// optional, if A doesn't exist, it means user can't toggle this option. 0-off, 1-on.
	A int64 `json:"A"`
	// fyi code
	FC string `json:"FC"`
	// detailed description
	FD string `json:"FD"`
	// title
	FN string `json:"FN"`
	// disclaimer read, 1 = yes, = 0 no.
	H int64 `json:"H"`
}

// account information
type Account struct {
	AccountAlias  string          `json:"accountAlias"`
	AccountId     string          `json:"accountId"`
	AccountStatus float64         `json:"accountStatus"`
	AccountTitle  string          `json:"accountTitle"`
	AccountVan    string          `json:"accountVan"`
	Covestor      bool            `json:"covestor"`
	Currency      string          `json:"currency"`
	Desc          string          `json:"desc"`
	DisplayName   string          `json:"displayName"`
	Faclient      bool            `json:"faclient"`
	Id            string          `json:"id"`
	Master        *Account_Master `json:"master"`
	Parent        string          `json:"parent"`
	TradingType   string          `json:"tradingType"`
	Type          string          `json:"type"`
}

type Performance_Nav_Data_List_Item struct {
	BaseCurrency string `json:"baseCurrency"`
	// end date-- yyyyMMdd
	End string `json:"end"`
	Id  string `json:"id"`
	// for example-- acctid
	IdType string `json:"idType"`
	// each value stands for price change percent of corresponding date in dates array
	Returns []float64 `json:"returns"`
	// start date-- yyyyMMdd
	Start string `json:"start"`
}

type Iserver_Account_AccountId_Order_Whatif_POST_200_Initial struct {
	After   string `json:"after"`
	Change  string `json:"change"`
	Current string `json:"current"`
}

type Iserver_Account_AccountId_Order_OrderId_POST_200_List_Item struct {
	Local_order_id string `json:"local_order_id"`
	Order_id       string `json:"order_id"`
	Order_status   string `json:"order_status"`
}

type Fyi_Deliveryoptions_Device_POST_Request struct {
	DeviceId   string `json:"deviceId"`
	Devicename string `json:"devicename"`
	Enabled    bool   `json:"enabled"`
	UiName     string `json:"uiName"`
}

type Trsrv_Secdef_Schedule_GET_200_Schedules_List_Item struct {
	ClearingCycleEndTime int64 `json:"clearingCycleEndTime"`
	// If the LIQUID hours differs from the total trading day then a separate 'session'
	// tag is returned.
	Sessions *Trsrv_Secdef_Schedule_GET_200_Schedules_List_Item_Sessions `json:"sessions"`
	// 20000101 stands for any Sat, 20000102 stands for any Sun, ... 20000107 stands for
	// any Fri. Any other date stands for itself.
	TradingScheduleDate int64 `json:"tradingScheduleDate"`
	// Returns tradingTime in exchange time zone.
	TradingTimes *Trsrv_Secdef_Schedule_GET_200_Schedules_List_Item_TradingTimes `json:"tradingTimes"`
}

type Iserver_Secdef_Search_POST_200_List_Item_Sections_List_Item struct {
	// Listing Exchange
	Exchange string `json:"exchange"`
	// For combo's defines the asset class for each leg
	LegSecType string `json:"legSecType"`
	// List of expiration month(s) and year(s) in MMMYY format separated by semicolon
	Months string `json:"months"`
	// Asset Class
	SecType string `json:"secType"`
	Symbol  string `json:"symbol"`
}

type ModifyOrder struct {
	AcctId   string  `json:"acctId"`
	AuxPrice float64 `json:"auxPrice"`
	Conid    int64   `json:"conid"`
	// optional, not required
	ListingExchange string `json:"listingExchange"`
	// for example LMT
	OrderType  string  `json:"orderType"`
	OutsideRTH bool    `json:"outsideRTH"`
	Price      float64 `json:"price"`
	// usually integer, for some special cases can be float numbers
	Quantity float64 `json:"quantity"`
	// SELL or BUY
	Side   string `json:"side"`
	Ticker string `json:"ticker"`
	// for example DAY
	Tif string `json:"tif"`
}

type Iserver_Reply_Replyid_POST_200_List_Item struct {
	Local_order_id string `json:"local_order_id"`
	Order_id       string `json:"order_id"`
	Order_status   string `json:"order_status"`
}

type Account_Master struct {
	OfficialTitle string `json:"officialTitle"`
	Title         string `json:"title"`
}

type Iserver_Account_AccountId_Order_Whatif_POST_200_Equity struct {
	After   string `json:"after"`
	Change  string `json:"change"`
	Current string `json:"current"`
}

type Stocks_List_Item_Contracts_List_Item struct {
	// conid of the stock contract
	Conid    int64  `json:"conid"`
	Exchange string `json:"exchange"`
}

type HistoryData_Data_List_Item struct {
	// close price
	C float64 `json:"c"`
	// high price
	H float64 `json:"h"`
	// low price
	L float64 `json:"l"`
	// open price
	O float64 `json:"o"`
	// unix time stamp
	T float64 `json:"t"`
	// volume
	V float64 `json:"v"`
}

// short positions allocation
type Allocation_List_Item_AssetClass_Short struct {
	BOND float64 `json:"BOND"`
	CASH float64 `json:"CASH"`
	FUT  float64 `json:"FUT"`
	OPT  float64 `json:"OPT"`
	STK  float64 `json:"STK"`
	WAR  float64 `json:"WAR"`
}

type Iserver_Account_Orders_GET_Request struct {
	Filters []string `json:"filters"`
}

type MarketData struct {
	Ask     float64 `json:"Ask"`
	AskSize float64 `json:"AskSize"`
	Bid     float64 `json:"Bid"`
	BidSize float64 `json:"BidSize"`
	// IBKR Contract ID
	Conid float64 `json:"Conid"`
	// Exchange
	Exchange string  `json:"Exchange"`
	Last     float64 `json:"Last"`
	LastSize float64 `json:"LastSize"`
	MinTick  float64 `json:"minTick"`
}

// portfolio allocation by sector
type Allocation_List_Item_Sector struct {
	// long positions allocation
	Long *Allocation_List_Item_Sector_Long `json:"long"`
	// short positions allocation
	Short *Allocation_List_Item_Sector_Short `json:"short"`
}

type Iserver_Contract_Conid_InfoAndRules_GET_200_Rules_List_Item_OrderDefaults_List_Item struct {
	// orderType
	String []string `json:"string"`
}

type Iserver_Reply_Replyid_POST_Request struct {
	// answer to question, true means yes, false means no
	Confirmed bool `json:"confirmed"`
}

type Sso_Validate_GET_200 struct {
	// Time of session validation
	AUTH_TIME int64 `json:"AUTH_TIME"`
	// Time in milliseconds until session expires. Caller needs to call the again to re-validate
	// session
	EXPIRES int64 `json:"EXPIRES"`
	// IP Address
	IP string `json:"IP"`
	// 1 for Live, 2 for Paper
	LOGIN_TYPE int64 `json:"LOGIN_TYPE"`
	// Paper Username
	PAPER_USER_NAME string `json:"PAPER_USER_NAME"`
	// true if session was validated; false if not.
	RESULT     bool `json:"RESULT"`
	SF_ENABLED bool `json:"SF_ENABLED"`
	// User ID
	USER_ID int64 `json:"USER_ID"`
	// Username
	USER_NAME string `json:"USER_NAME"`
	// Time in milliseconds until session expires. Caller needs to call the again to re-validate
	// session
	Expire float64 `json:"expire"`
	// Time of session validation
	Features map[string]interface{} `json:"features"`
	// Time session was last accessed
	LastAccessed int64 `json:"lastAccessed"`
	// 1 for Live, 2 for Paper
	LoginType int64 `json:"loginType"`
}

type ScannerParams_InstrumentList_Instrument_List_Item struct {
	// Supported filters for the instrument separated by a comma
	Filters string `json:"filters"`
	// Display name for the instrument
	Name string `json:"name"`
	// Sec Type of the instrument. This field is not provided if its value is same as 'type'
	SecType   string `json:"secType"`
	ShortName string `json:"shortName"`
	// Type of the asset
	Type string `json:"type"`
}

type Summary_ExcludedAccounts_List_Item struct {
	AcctId               string `json:"acctId"`
	AcctNumAtFI          string `json:"acctNumAtFI"`
	AcctTitle            string `json:"acctTitle"`
	FiName               string `json:"fiName"`
	HarvestCode          int64  `json:"harvestCode"`
	LastUpdate           string `json:"lastUpdate"`
	LastUpdateAttempt    string `json:"lastUpdateAttempt"`
	LastUpdateStatusCode string `json:"lastUpdateStatusCode"`
	Rc                   int64  `json:"rc"`
}

type Ibcust_Entity_Info_GET_200_List_Item_Entities_List_Item struct {
	Address   *Ibcust_Entity_Info_GET_200_List_Item_Entities_List_Item_Address `json:"address"`
	CanSign   bool                                                             `json:"canSign"`
	CanTrade  bool                                                             `json:"canTrade"`
	IdentDocs []map[string]interface{}                                         `json:"identDocs"`
	Name      *Ibcust_Entity_Info_GET_200_List_Item_Entities_List_Item_Name    `json:"name"`
	Type      string                                                           `json:"type"`
}

// device
type Fyi_Deliveryoptions_GET_200_E_List_Item struct {
	// device is enabled or not 0-true, 1-false.
	A string `json:"A"`
	// device id
	I string `json:"I"`
	// device name
	NM string `json:"NM"`
	// unique device id
	UI string `json:"UI"`
}

type Pa_Summary_POST_Request struct {
	AcctIds []string `json:"acctIds"`
}

type Calendar_request struct {
	Date    *Calendar_request_Date    `json:"date"`
	Filters *Calendar_request_Filters `json:"filters"`
}
