package rest

import (
	. "github.com/kamaiu/ib-cp-go/client/v1/rest/model"
	"github.com/valyala/bytebufferpool"
	"net/url"
	"strconv"
)

// Get Market Data for the given conid(s). The endpoint will return by default bid,
// ask, last, change, change pct, close, listing exchange.
// See response fields for a list of available fields that can be request via fields
// argument.
// The endpoint /iserver/accounts must be called prior to /iserver/marketdata/snapshot.
// For derivative contracts the endpoint /iserver/secdef/search must be called first.
// First /snapshot endpoint call for given conid will initiate the market data request.
//
// To receive all available fields the /snapshot endpoint will need to be called several
// times.
// To receive streaming market data the endpoint /ws can be used. Refer to [Streaming
// WebSocket Data](https://interactivebrokers.github.io/cpwebapi/RealtimeSubscription.html)
// for details.
func (c *Conn) Iserver_Marketdata_Snapshot_GET(
	// list of conids separated by comma
	conids string,
	// time period since which updates are required. uses epoch time with milliseconds.
	since int64,
	// list of fields separated by comma
	fields string,
) (
	_200 Iserver_Marketdata_Snapshot_GET_200_List,
	_400 *Iserver_Marketdata_Snapshot_GET_400,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/marketdata/snapshot")
	_, _ = _u.WriteString("?conids=")
	_, _ = _u.WriteString(url.PathEscape(conids))
	_, _ = _u.WriteString("&since=")
	_, _ = _u.WriteString(url.PathEscape(strconv.FormatInt(int64(since), 10)))
	_, _ = _u.WriteString("&fields=")
	_, _ = _u.WriteString(url.PathEscape(fields))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Iserver_Marketdata_Snapshot_GET_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		case 400:
			_400 = &Iserver_Marketdata_Snapshot_GET_400{}
			if body != nil {
				return _400.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Account information related to account Id /portfolio/accounts or /portfolio/subaccounts
// must be called prior to this endpoint.
func (c *Conn) Portfolio_AccountId_Meta_GET(
	// account id
	accountId string,
) (
	_200 Accounts_List,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/portfolio/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/meta")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Accounts_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns an object contains four lists contain all parameters for scanners
func (c *Conn) Iserver_Scanner_Params_GET() (
	_200 *Iserver_Scanner_Params_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/scanner/params")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Scanner_Params_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// You can pass a list of orders here
func (c *Conn) Iserver_Account_AccountId_Orders_POST(
	// account id
	accountId string,
	request *Iserver_Account_AccountId_Orders_POST_Request,
) (
	_200 Iserver_Account_AccountId_Orders_POST_200_List,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/orders")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Iserver_Account_AccountId_Orders_POST_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// options for sending fyis to email and other devices
func (c *Conn) Fyi_Deliveryoptions_GET() (
	_200 *Fyi_Deliveryoptions_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/deliveryoptions")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Fyi_Deliveryoptions_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

func (c *Conn) Fyi_Deliveryoptions_DeviceId_DELETE(
	// device ID
	deviceId string,
) (
	_200 Any,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/deliveryoptions/")
	_, _ = _u.WriteString(url.PathEscape(deviceId))
	err = c.Do("DELETE", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Any{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

func (c *Conn) Portfolio_AccountId_Positions_Invalidate_POST(
	// account id
	accountId string,
) (
	_200 Any,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/portfolio/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/positions/invalidate")
	err = c.Do("POST", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Any{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// The endpoint is meant to be used in polling mode, e.g. requesting every x seconds.
// The response will contain two objects, one is notification, the other is orders.
//
// Orders is the list of live orders (cancelled, filled, submitted).
// Notifications contains information about execute orders as they happen, see status
// field.
// To receive streaming live orders the endpoint /ws can be used. Refer to [Streaming
// WebSocket Data](https://interactivebrokers.github.io/cpwebapi/RealtimeSubscription.html)
// for details.
func (c *Conn) Iserver_Account_Orders_GET(
	request *Iserver_Account_Orders_GET_Request,
) (
	_200 *Iserver_Account_Orders_GET_200,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/orders")
	err = c.DoRequest("GET", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Account_Orders_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// The streaming API is documented under [Streaming WebSocket Data](https://interactivebrokers.github.io/cpwebapi/RealtimeSubscription.html)
// for details.
func (c *Conn) Ws_POST() (
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/ws")
	err = c.Do("POST", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns the performance (MTM) for the given accounts, if more than one account is
// passed, the result is consolidated.
func (c *Conn) Pa_Performance_POST(
	request *Pa_Performance_POST_Request,
) (
	_200 *Performance,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/pa/performance")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Performance{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

func (c *Conn) Fyi_Notifications_More_GET(
	// id of last notification in the list
	id string,
) (
	_200 Notifications_List,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/notifications/more")
	_, _ = _u.WriteString("?id=")
	_, _ = _u.WriteString(url.PathEscape(id))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Notifications_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// transaction history for a given number of conids and accounts.
// Types of transactions include dividend payments, buy and sell transactions, transfers.
func (c *Conn) Pa_Transactions_POST(
	request *Pa_Transactions_POST_Request,
) (
	_200 *Transactions,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/pa/transactions")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Transactions{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Financial Advisors can use this endpoint to place an order for a specified group.
// To place an order for a specified account use the endpoint /iserver/account/{accountId}/order.
// More information about groups can be found in the [TWS Users' Guide:](https://guides.interactivebrokers.com/twsguide/twsguide.htm#usersguidebook/financialadvisors/create_an_account_group_for_share_allocation.htm).
func (c *Conn) Iserver_Account_Orders_FaGroup_POST(
	// financial advisor group
	faGroup string,
	request *OrderRequest,
) (
	_200 Iserver_Account_Orders_FaGroup_POST_200_List,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/orders/")
	_, _ = _u.WriteString(url.PathEscape(faGroup))
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Iserver_Account_Orders_FaGroup_POST_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Modifies an open order. Must call /iserver/accounts endpoint prior to modifying an
// order. Use /iservers/account/orders endpoint to review open-order(s).
func (c *Conn) Iserver_Account_AccountId_Order_OrderId_POST(
	// account id, or fa group if modifying a group order
	accountId string,
	// Customer order id, use /iservers/account/orders endpoint to query orderId.
	orderId string,
	request *ModifyOrder,
) (
	_200 Iserver_Account_AccountId_Order_OrderId_POST_200_List,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/order/")
	_, _ = _u.WriteString(url.PathEscape(orderId))
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Iserver_Account_AccountId_Order_OrderId_POST_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Reply to questions when placing orders and submit orders
func (c *Conn) Iserver_Reply_Replyid_POST(
	// Please use the "id" from the response of "Place Order" endpoint
	replyid string,
	request *Iserver_Reply_Replyid_POST_Request,
) (
	_200 Iserver_Reply_Replyid_POST_200_List,
	_400 *Iserver_Reply_Replyid_POST_400,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/reply/")
	_, _ = _u.WriteString(url.PathEscape(replyid))
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Iserver_Reply_Replyid_POST_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		case 400:
			_400 = &Iserver_Reply_Replyid_POST_400{}
			if body != nil {
				return _400.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Information regarding settled cash, cash balances, etc. in the account's base currency
// and any other cash balances hold in other currencies.  /portfolio/accounts or /portfolio/subaccounts
// must be called prior to this endpoint. The list of supported currencies is available
// at https://www.interactivebrokers.com/en/index.php?f=3185.
func (c *Conn) Portfolio_AccountId_Ledger_GET(
	// account id
	accountId string,
) (
	_200 *Portfolio_AccountId_Ledger_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/portfolio/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/ledger")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Portfolio_AccountId_Ledger_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Cancel market data for given conid. To cancel all market data request(s), see /iserver/marketdata/unsubscribeall.
func (c *Conn) Iserver_Marketdata_Conid_Unsubscribe_GET(
	// contract id
	conid string,
) (
	_200 *Iserver_Marketdata_Conid_Unsubscribe_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/marketdata/")
	_, _ = _u.WriteString(url.PathEscape(conid))
	_, _ = _u.WriteString("/unsubscribe")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Marketdata_Conid_Unsubscribe_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Configure which typecode you would like to enable/disable.
func (c *Conn) Fyi_Settings_Typecode_POST(
	// fyi code
	typecode string,
	request *Fyi_Settings_Typecode_POST_Request,
) (
	_200 Any,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/settings/")
	_, _ = _u.WriteString(url.PathEscape(typecode))
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Any{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Current Authentication status to the Brokerage system. Market Data and Trading is
// not possible if not authenticated, e.g. authenticated shows false
func (c *Conn) Iserver_Auth_Status_POST() (
	_200 *AuthStatus,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/auth/status")
	err = c.Do("POST", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &AuthStatus{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// If an user has multiple accounts, and user wants to get orders, trades, etc. of an
// account other than currently selected account, then user can update the currently
// selected account using this API and then can fetch required information for the newly
// updated account.
func (c *Conn) Iserver_Account_POST(
	request *SetAccount,
) (
	_200 *Iserver_Account_POST_200,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Account_POST_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

func (c *Conn) Fyi_Disclaimer_Typecode_GET(
	// fyi code, for example --M8, EA
	typecode string,
) (
	_200 *Fyi_Disclaimer_Typecode_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/disclaimer/")
	_, _ = _u.WriteString(url.PathEscape(typecode))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Fyi_Disclaimer_Typecode_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns the trading schedule up to a month for the requested contract
func (c *Conn) Trsrv_Secdef_Schedule_GET(
	// specify the asset class of the contract.
	// Available values-- Stock: STK, Option: OPT, Future: FUT, Contract For Difference:
	// CFD, Warrant: WAR, Forex: SWP, Mutual Fund: FND, Bond: BND, Inter-Commodity Spreads:
	// ICS
	assetClass string,
	// Underlying Symbol for specified contract, for example 'AAPL' for US Stock - Apple
	// Inc.
	symbol string,
	// Native exchange for contract, for example 'NASDAQ' for US Stock - Apple Inc.
	exchange string,
) (
	_200 *Trsrv_Secdef_Schedule_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/trsrv/secdef/schedule")
	_, _ = _u.WriteString("?assetClass=")
	_, _ = _u.WriteString(url.PathEscape(assetClass))
	_, _ = _u.WriteString("&symbol=")
	_, _ = _u.WriteString(url.PathEscape(symbol))
	_, _ = _u.WriteString("&exchange=")
	_, _ = _u.WriteString(url.PathEscape(exchange))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Trsrv_Secdef_Schedule_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

func (c *Conn) Fyi_Deliveryoptions_Device_POST(
	request *Fyi_Deliveryoptions_Device_POST_Request,
) (
	_200 *Fyi_Deliveryoptions_Device_POST_200,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/deliveryoptions/device")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Fyi_Deliveryoptions_Device_POST_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// When using the CP Gateway, this endpoint provides a way to reauthenticate to the
// Brokerage system as long as there is a valid SSO session, see /sso/validate.
func (c *Conn) Iserver_Reauthenticate_POST() (
	_200 *AuthStatus,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/reauthenticate")
	err = c.Do("POST", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &AuthStatus{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns a list of security definitions for the given conids
func (c *Conn) Trsrv_Secdef_POST(
	request *Trsrv_Secdef_POST_Request,
) (
	_200 Secdef_List,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/trsrv/secdef")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Secdef_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

func (c *Conn) Fyi_Notifications_GET(
	// if set, don't set include
	exclude string,
	// if set, don't set exclude
	include string,
	// max number of fyis in response
	max string,
) (
	_200 Notifications_List,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/notifications")
	_, _ = _u.WriteString("?exclude=")
	_, _ = _u.WriteString(url.PathEscape(exclude))
	_, _ = _u.WriteString("&include=")
	_, _ = _u.WriteString(url.PathEscape(include))
	_, _ = _u.WriteString("&max=")
	_, _ = _u.WriteString(url.PathEscape(max))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Notifications_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns a list of accounts the user has trading access to, their respective aliases
// and the currently selected account. Note this endpoint must be called before modifying
// an order or querying open orders.
func (c *Conn) Iserver_Accounts_GET() (
	_200 *Iserver_Accounts_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/accounts")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Accounts_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns Applicant Id with all owner related entities
func (c *Conn) Ibcust_Entity_Info_GET() (
	_200 Ibcust_Entity_Info_GET_200_List,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/ibcust/entity/info")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Ibcust_Entity_Info_GET_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Search by underlying or company name. Relays back what derivative contract(s) it
// has. This endpoint must be called before using /secdef/info
func (c *Conn) Iserver_Secdef_Search_POST(
	request *Iserver_Secdef_Search_POST_Request,
) (
	_200 Iserver_Secdef_Search_POST_200_List,
	_500 *Iserver_Secdef_Search_POST_500,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/secdef/search")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Iserver_Secdef_Search_POST_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		case 500:
			_500 = &Iserver_Secdef_Search_POST_500{}
			if body != nil {
				return _500.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns a list of non-expired future contracts for given symbol(s)
func (c *Conn) Trsrv_Futures_GET(
	// list of case-sensitive symbols separated by comma
	symbols string,
) (
	_200 *Trsrv_Futures_GET_200,
	_500 *Trsrv_Futures_GET_500,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/trsrv/futures")
	_, _ = _u.WriteString("?symbols=")
	_, _ = _u.WriteString(url.PathEscape(symbols))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Trsrv_Futures_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		case 500:
			_500 = &Trsrv_Futures_GET_500{}
			if body != nil {
				return _500.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

func (c *Conn) Fyi_Disclaimer_Typecode_PUT(
	// fyi code, for example --M8, EA
	typecode string,
) (
	_200 *Fyi_Disclaimer_Typecode_PUT_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/disclaimer/")
	_, _ = _u.WriteString(url.PathEscape(typecode))
	err = c.Do("PUT", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Fyi_Disclaimer_Typecode_PUT_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Please note here, DO NOT pass orderId when creating a new alert, toolId is only required
// for MTA alert
func (c *Conn) Iserver_Account_AccountId_Alert_POST(
	// account id
	accountId string,
	request *AlertRequest,
) (
	_200 *Iserver_Account_AccountId_Alert_POST_200,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/alert")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Account_AccountId_Alert_POST_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Please be careful, if alertId is 0, it will delete all alerts
func (c *Conn) Iserver_Account_AccountId_Alert_AlertId_DELETE(
	// account id
	accountId string,
	// alert id
	alertId string,
) (
	_200 *Iserver_Account_AccountId_Alert_AlertId_DELETE_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/alert/")
	_, _ = _u.WriteString(url.PathEscape(alertId))
	err = c.Do("DELETE", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Account_AccountId_Alert_AlertId_DELETE_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns an object containing PnL for the selected account and its models (if any).
// To receive streaming PnL the endpoint /ws can be used. Refer to [Streaming WebSocket
// Data](https://interactivebrokers.github.io/cpwebapi/RealtimeSubscription.html) for
// details.
func (c *Conn) Iserver_Account_Pnl_Partitioned_GET() (
	_200 *Iserver_Account_Pnl_Partitioned_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/pnl/partitioned")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Account_Pnl_Partitioned_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns a list of positions for the given account. The endpoint supports paging,
// page's default size is 30 positions. /portfolio/accounts or /portfolio/subaccounts
// must be called prior to this endpoint.
func (c *Conn) Portfolio_AccountId_Positions_PageId_GET(
	// account id
	accountId string,
	// page id
	pageId string,
	// optional
	model string,
	// declare the table to be sorted by which column
	sort string,
	// in which order, a means ascending - d means descending
	direction string,
	// period for pnl column, can be 1D, 7D, 1M...
	period string,
) (
	_200 Position_List,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/portfolio/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/positions/")
	_, _ = _u.WriteString(url.PathEscape(pageId))
	_, _ = _u.WriteString("?model=")
	_, _ = _u.WriteString(url.PathEscape(model))
	_, _ = _u.WriteString("&sort=")
	_, _ = _u.WriteString(url.PathEscape(sort))
	_, _ = _u.WriteString("&direction=")
	_, _ = _u.WriteString(url.PathEscape(direction))
	_, _ = _u.WriteString("&period=")
	_, _ = _u.WriteString(url.PathEscape(period))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Position_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// This endpoint allows you to preview order without actually submitting the order and
// you can get
// commission information in the response.
func (c *Conn) Iserver_Account_AccountId_Order_Whatif_POST(
	// account id
	accountId string,
	request *OrderRequest,
) (
	_200 *Iserver_Account_AccountId_Order_Whatif_POST_200,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/order/whatif")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Account_AccountId_Order_Whatif_POST_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Use the endpoint /iserver/account/:accountId/alerts to receive the alert id.
// The order_id in the response is the alert id.
func (c *Conn) Iserver_Account_Alert_Id_GET(
	// alert id
	id string,
) (
	_200 *AlertResponse,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/alert/")
	_, _ = _u.WriteString(url.PathEscape(id))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &AlertResponse{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Please note here, sometimes this endpoint alone can't make sure you submit the order
// successfully,
// you could receive some questions in the response, you have to to answer them in order
// to submit the order
// successfully. You can use "/iserver/reply/{replyid}" endpoint to answer questions
func (c *Conn) Iserver_Account_AccountId_Order_POST(
	// account id
	accountId string,
	request *OrderRequest,
) (
	_200 Iserver_Account_AccountId_Order_POST_200_List,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/order")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Iserver_Account_AccountId_Order_POST_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// In non-tiered account structures, returns a list of accounts for which the user can
// view position and account information. This endpoint must be called prior  to calling
// other /portfolio endpoints for those accounts. For querying a list of accounts  which
// the user can trade, see /iserver/accounts. For a list of subaccounts in tiered  account
// structures (e.g. financial advisor or ibroker accounts) see /portfolio/subaccounts.
func (c *Conn) Portfolio_Accounts_GET() (
	_200 Accounts_List,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/portfolio/accounts")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Accounts_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

func (c *Conn) Fyi_Notifications_NotificationId_PUT(
	// mark a notification read
	notificationId string,
) (
	_200 Any,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/notifications/")
	_, _ = _u.WriteString(url.PathEscape(notificationId))
	err = c.Do("PUT", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Any{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns a list of trades for the currently selected account for current day and six
// previous days. It is advised to call this endpoint once per session.
func (c *Conn) Iserver_Account_Trades_GET() (
	_200 Iserver_Account_Trades_GET_200_List,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/trades")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Iserver_Account_Trades_GET_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

func (c *Conn) Iserver_Scanner_Run_POST(
	request *ScannerParams,
) (
	_200 Iserver_Scanner_Run_POST_200_List,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/scanner/run")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Iserver_Scanner_Run_POST_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Validates the current session for the SSO user
func (c *Conn) Sso_Validate_GET() (
	_200 *Sso_Validate_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/sso/validate")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Sso_Validate_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Cancels an open order. Must call /iserver/accounts endpoint prior to cancelling an
// order. Use /iservers/account/orders endpoint to review open-order(s) and get latest
// order status.
func (c *Conn) Iserver_Account_AccountId_Order_OrderId_DELETE(
	// account id, or fa group if deleting a group order
	accountId string,
	// Customer order id, use /iservers/account/orders endpoint to query orderId.
	orderId string,
) (
	_200 *Iserver_Account_AccountId_Order_OrderId_DELETE_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/order/")
	_, _ = _u.WriteString(url.PathEscape(orderId))
	err = c.Do("DELETE", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Account_AccountId_Order_OrderId_DELETE_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Query strikes for Options/Warrants. For available contract months and exchanges use
// "/iserver/secdef/search"
func (c *Conn) Iserver_Secdef_Strikes_GET(
	// contract id
	conid string,
	// OPT/WAR
	sectype string,
	// contract month
	month string,
	// optional, default is SMART
	exchange string,
) (
	_200 *Iserver_Secdef_Strikes_GET_200,
	_500 *Iserver_Secdef_Strikes_GET_500,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/secdef/strikes")
	_, _ = _u.WriteString("?conid=")
	_, _ = _u.WriteString(url.PathEscape(conid))
	_, _ = _u.WriteString("&sectype=")
	_, _ = _u.WriteString(url.PathEscape(sectype))
	_, _ = _u.WriteString("&month=")
	_, _ = _u.WriteString(url.PathEscape(month))
	_, _ = _u.WriteString("&exchange=")
	_, _ = _u.WriteString(url.PathEscape(exchange))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Secdef_Strikes_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		case 500:
			_500 = &Iserver_Secdef_Strikes_GET_500{}
			if body != nil {
				return _500.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Please note, if alertId is 0, it will activate/deactivate all alerts
func (c *Conn) Iserver_Account_AccountId_Alert_Activate_POST(
	// account id
	accountId string,
	request *Iserver_Account_AccountId_Alert_Activate_POST_Request,
) (
	_200 *Iserver_Account_AccountId_Alert_Activate_POST_200,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/alert/activate")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Account_AccountId_Alert_Activate_POST_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Provides Contract Details of Futures, Options, Warrants, Cash and CFDs based on conid.
// To get the strike price for Options/Warrants use "/iserver/secdef/strikes" endpoint.
// Must call /secdef/search for the underlying contract first.
func (c *Conn) Iserver_Secdef_Info_GET(
	// underlying contract id
	conid string,
	// FUT/OPT/WAR/CASH/CFD
	sectype string,
	// contract month, only required for FUT/OPT/WAR in the format MMMYY, example JAN00
	month string,
	// optional, default is SMART
	exchange string,
	// optional, only required for OPT/WAR
	strike string,
	// C for call, P for put
	right string,
) (
	_200 Iserver_Secdef_Info_GET_200_List,
	_500 *Iserver_Secdef_Info_GET_500,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/secdef/info")
	_, _ = _u.WriteString("?conid=")
	_, _ = _u.WriteString(url.PathEscape(conid))
	_, _ = _u.WriteString("&sectype=")
	_, _ = _u.WriteString(url.PathEscape(sectype))
	_, _ = _u.WriteString("&month=")
	_, _ = _u.WriteString(url.PathEscape(month))
	_, _ = _u.WriteString("&exchange=")
	_, _ = _u.WriteString(url.PathEscape(exchange))
	_, _ = _u.WriteString("&strike=")
	_, _ = _u.WriteString(url.PathEscape(strike))
	_, _ = _u.WriteString("&right=")
	_, _ = _u.WriteString(url.PathEscape(right))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Iserver_Secdef_Info_GET_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		case 500:
			_500 = &Iserver_Secdef_Info_GET_500{}
			if body != nil {
				return _500.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Each login user only has one mobile trading assistant (MTA) alert with it's own unique
// tool id.
// The tool id cannot be changed. When modified a new order Id is generated. MTA alerts
// can not
// be created or deleted. If you call delete /iserver/account/:accountId/alert/:alertId,
// it will reset MTA to default. See [here](https://www.interactivebrokers.com/en/software/mobileiphonemobile/mobileiphone.htm#monitor/trading-assistant.htm)
// for more information on MTA alerts.
func (c *Conn) Iserver_Account_Mta_GET() (
	_200 *AlertResponse,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/mta")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &AlertResponse{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Logs the user out of the gateway session. Any further activity requires re-authentication.
func (c *Conn) Logout_POST() (
	_200 *Logout_POST_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/logout")
	err = c.Do("POST", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Logout_POST_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// If the gateway has not received any requests for several minutes an open session
// will automatically timeout. The tickle endpoint pings the server to prevent the session
// from ending.
func (c *Conn) Tickle_POST() (
	_200 *Tickle_POST_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/tickle")
	err = c.Do("POST", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Tickle_POST_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Cancel all market data request(s). To cancel market data for given conid, see /iserver/marketdata/{conid}/unsubscribe.
//
func (c *Conn) Iserver_Marketdata_Unsubscribeall_GET() (
	_200 *Iserver_Marketdata_Unsubscribeall_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/marketdata/unsubscribeall")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Marketdata_Unsubscribeall_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Used in tiered account structures (such as financial advisor and ibroker accounts)
// to return a list of sub-accounts for which the user can view position and  account-related
// information. This endpoint must be called prior to calling other  /portfolio endpoints
// for those subaccounts.  To query a list of accounts the user can trade, see /iserver/accounts.
func (c *Conn) Portfolio_Subaccounts_GET() (
	_200 *Account,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/portfolio/subaccounts")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Account{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

func (c *Conn) Fyi_Deliveryoptions_Email_PUT(
	// true/false
	enabled string,
) (
	_200 *Fyi_Deliveryoptions_Email_PUT_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/deliveryoptions/email")
	_, _ = _u.WriteString("?enabled=")
	_, _ = _u.WriteString(url.PathEscape(enabled))
	err = c.Do("PUT", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Fyi_Deliveryoptions_Email_PUT_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns an object of all positions matching the conid for all the selected accounts.
// For portfolio models the conid could be in more than one model, returning an array
// with the name of the model it belongs to. /portfolio/accounts or /portfolio/subaccounts
// must be called prior to this endpoint.
func (c *Conn) Portfolio_Positions_Conid_GET(
	// contract id
	conid int64,
) (
	_200 *Portfolio_Positions_Conid_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/portfolio/positions/")
	_, _ = _u.WriteString(url.PathEscape(strconv.FormatInt(int64(conid), 10)))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Portfolio_Positions_Conid_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns a summary of all account balances for the given accounts, if more than one
// account is passed, the result is consolidated.
func (c *Conn) Pa_Summary_POST(
	request *Pa_Summary_POST_Request,
) (
	_200 *Summary,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/pa/summary")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Summary{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Using the Contract Identifier get contract info. You can use this to prefill your
// order before you submit an order
func (c *Conn) Iserver_Contract_Conid_Info_GET(
	// contract id
	conid string,
) (
	_200 *Contract,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/contract/")
	_, _ = _u.WriteString(url.PathEscape(conid))
	_, _ = _u.WriteString("/info")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Contract{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns an object contains all stock contracts for given symbol(s)
func (c *Conn) Trsrv_Stocks_GET(
	// list of upper-sensitive symbols separated by comma
	symbols string,
) (
	_200 *Trsrv_Stocks_GET_200,
	_500 *Trsrv_Stocks_GET_500,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/trsrv/stocks")
	_, _ = _u.WriteString("?symbols=")
	_, _ = _u.WriteString(url.PathEscape(symbols))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Trsrv_Stocks_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		case 500:
			_500 = &Trsrv_Stocks_GET_500{}
			if body != nil {
				return _500.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Return the current choices of subscriptions, we can toggle the option
func (c *Conn) Fyi_Settings_GET() (
	_200 Fyi_Settings_GET_200_List,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/settings")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Fyi_Settings_GET_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns the total number of unread fyis
func (c *Conn) Fyi_Unreadnumber_GET() (
	_200 *Fyi_Unreadnumber_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/fyi/unreadnumber")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Fyi_Unreadnumber_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns information about margin, cash balances and other information related to
// specified account. See also /portfolio/{accountId}/ledger. /portfolio/accounts or
// /portfolio/subaccounts must be called prior to this endpoint.
func (c *Conn) Portfolio_AccountId_Summary_GET(
	// account id
	accountId string,
) (
	_200 *Portfolio_AccountId_Summary_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/portfolio/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/summary")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Portfolio_AccountId_Summary_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Information about the account's portfolio allocation by Asset Class, Industry and
// Category.  /portfolio/accounts or /portfolio/subaccounts must be called prior to
// this endpoint.
func (c *Conn) Portfolio_AccountId_Allocation_GET(
	// account id
	accountId string,
) (
	_200 Allocation_List,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/portfolio/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/allocation")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Allocation_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Get historical market Data for given conid, length of data is controlled by 'period'
// and 'bar'.
// Formatted as: min=minute, h=hour, d=day, w=week, m=month, y=year
// e.g. period =1y with bar =1w returns 52 data points (Max of 1000 data points supported).
// **Note**: There's a limit of 5 concurrent requests. Excessive requests will return
// a 'Too many requests' status 429 response.
func (c *Conn) Iserver_Marketdata_History_GET(
	// contract id
	conid string,
	// Exchange of the conid (e.g. ISLAND, NYSE, etc.). Default value is empty which corresponds
	// to primary exchange of the conid.
	exchange string,
	// available time period-- {1-30}min, {1-8}h, {1-1000}d, {1-792}w, {1-182}m, {1-15}y
	period string,
	// possible value-- 1min, 2min, 3min, 5min, 10min, 15min, 30min, 1h, 2h, 3h, 4h, 8h,
	// 1d, 1w, 1m
	bar string,
	// For contracts that support it, will determine if historical data includes outside
	// of regular trading hours.
	outsideRth bool,
) (
	_200 *HistoryData,
	_429 *Iserver_Marketdata_History_GET_429,
	_500 *SystemError,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/marketdata/history")
	_, _ = _u.WriteString("?conid=")
	_, _ = _u.WriteString(url.PathEscape(conid))
	_, _ = _u.WriteString("&exchange=")
	_, _ = _u.WriteString(url.PathEscape(exchange))
	_, _ = _u.WriteString("&period=")
	_, _ = _u.WriteString(url.PathEscape(period))
	_, _ = _u.WriteString("&bar=")
	_, _ = _u.WriteString(url.PathEscape(bar))
	_, _ = _u.WriteString("&outsideRth=")
	_, _ = _u.WriteString(url.PathEscape(strconv.FormatBool(outsideRth)))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &HistoryData{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		case 429:
			_429 = &Iserver_Marketdata_History_GET_429{}
			if body != nil {
				return _429.UnmarshalJSON(body)
			} else {
				return err
			}
		case 500:
			_500 = &SystemError{}
			if body != nil {
				return _500.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns trading related rules and info for contract
func (c *Conn) Iserver_Contract_Conid_InfoAndRules_GET(
	// IBKR contract identifier
	conid string,
	// Side of the market rules apply too. Set to true for Buy Orders, set to false for
	// Sell Orders
	isBuy bool,
) (
	_200 *Iserver_Contract_Conid_InfoAndRules_GET_200,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/contract/")
	_, _ = _u.WriteString(url.PathEscape(conid))
	_, _ = _u.WriteString("/info-and-rules")
	_, _ = _u.WriteString("?isBuy=")
	_, _ = _u.WriteString(url.PathEscape(strconv.FormatBool(isBuy)))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = &Iserver_Contract_Conid_InfoAndRules_GET_200{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Similar to /portfolio/{accountId}/allocation but returns a consolidated view of of
// all the accounts returned by /portfolio/accounts. /portfolio/accounts or /portfolio/subaccounts
// must be called prior to this endpoint.
func (c *Conn) Portfolio_Allocation_POST(
	request *Portfolio_Allocation_POST_Request,
) (
	_200 Allocation_List,
	err error,
) {
	var _r []byte
	if request != nil {
		_r, err = request.MarshalJSON()
	}
	if err != nil {
		return
	}
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/portfolio/allocation")
	err = c.DoRequest("POST", _u, _r, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Allocation_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// Returns a list of all positions matching the conid. For portfolio models the conid
// could be in more than one model, returning an array with the name of the model it
// belongs to.  /portfolio/accounts or /portfolio/subaccounts must be called prior to
// this endpoint.
func (c *Conn) Portfolio_AccountId_Position_Conid_GET(
	// account id
	accountId string,
	// contract id
	conid int64,
) (
	_200 Position_List,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/portfolio/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/position/")
	_, _ = _u.WriteString(url.PathEscape(strconv.FormatInt(int64(conid), 10)))
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Position_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}

// The response will contain both active and inactive alerts, but it won't have MTA
// alert
func (c *Conn) Iserver_Account_AccountId_Alerts_GET(
	// account id
	accountId string,
) (
	_200 Iserver_Account_AccountId_Alerts_GET_200_List,
	err error,
) {
	_u := bytebufferpool.Get()
	_, _ = _u.WriteString(c.host)
	_, _ = _u.WriteString("/iserver/account/")
	_, _ = _u.WriteString(url.PathEscape(accountId))
	_, _ = _u.WriteString("/alerts")
	err = c.Do("GET", _u, func(status int, body []byte, err error) error {
		if err != nil {
			return err
		}
		switch status {
		case 200:
			_200 = Iserver_Account_AccountId_Alerts_GET_200_List{}
			if body != nil {
				return _200.UnmarshalJSON(body)
			} else {
				return err
			}
		default:
			return StatusCodeError{Code: status}
		}
	})
	return
}
