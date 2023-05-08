package models

import (
	"fmt"
	"time"
)

const (
	OfferMessageType               = "offer"
	BalanceMessageType             = "balance"
	InfoMessageType                = "info"
	XrateMessageType               = "xrate"
	BetslipMessageType             = "betslip"
	PmmMessageType                 = "pmm"
	SyncMessageType                = "sync"
	BetslipClosedMessageType       = "betslip_closed"
	OrderMessageType               = "order"
	BetMessageType                 = "bet"
	UnsubscribeEventMessageType    = "unsubscribe_event"
	SubscribeEventMessageType      = "subscribe_event"
	UnsubscribeAllEventMessageType = "unsubscribe_all_event"
)

// Offer universal
type Offer struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
	I int64   `json:"i"`
	N string  `json:"n"`
}
type Offers struct {
	EventID   int64     `json:"event_id"`
	Starts    time.Time `json:"starts"`
	Sport     string    `json:"sport"`
	WsReceive time.Time `json:"ws_receive"`
	OfferList []Offer   `json:"offer_list"`
}
type Timing struct {
	WsReceive       time.Time
	BeginOpen       time.Time
	EndOpen         time.Time
	StartJob        time.Time
	Complete        time.Time
	BeginPlaceFirst time.Time
	ConditionsOk    time.Time
	BeginFirstBet   time.Time
	EndFirstBet     time.Time
	StartCheck      time.Time
	BeginSecondBet  time.Time
	EndSecondBet    time.Time
	BeginStats      time.Time
	EndStats        time.Time
	BeforeSubmit    time.Time
}
type SubscribeEventRequest struct {
	CompetitionId int64     `json:"competition_id,omitempty"`
	Sport         string    `json:"sport,omitempty"`
	HomeID        int64     `json:"home_id"`
	AwayID        int64     `json:"away_id"`
	EventDate     time.Time `json:"event_date"`
}
type EventDB struct {
	Sport         string    `json:"sport"`
	Country       string    `json:"country"`
	CompetitionID int64     `json:"competition_id"`
	HomeID        int64     `json:"home_id"`
	AwayID        int64     `json:"away_id"`
	ID            int64     `json:"id"`
	EventDate     time.Time `json:"-"`
	Starts        time.Time `json:"starts"`
	Offline       bool      `json:"offline"`
}

func (e EventDB) EventID() string {
	return fmt.Sprintf("%s,%d,%d", e.EventDate.Format("2006-01-02"), e.HomeID, e.AwayID)
}

type EventWithScore struct {
	Sport         string    `json:"sport"`
	Country       string    `json:"country"`
	CompetitionID int64     `json:"competition_id"`
	HomeID        int64     `json:"home_id"`
	AwayID        int64     `json:"away_id"`
	ID            int64     `json:"id"`
	EventDate     time.Time `json:"-"`
	Starts        time.Time `json:"starts"`
	Offline       bool      `json:"offline"`
	TimeName      *string   `json:"time_name"`
	TimeMin       *int64    `json:"time_min"`
}

func (e EventWithScore) EventID() string {
	return fmt.Sprintf("%s,%d,%d", e.EventDate.Format("2006-01-02"), e.HomeID, e.AwayID)
}

type PmmMessage struct {
	BetslipId string `json:"betslip_id" mapstructure:"betslip_id"`
	Sport     string `json:"sport" mapstructure:"sport"`
	EventId   string `json:"event_id" mapstructure:"event_id"`
	Bookie    string `json:"bookie" mapstructure:"bookie"`
	Username  string `json:"username" mapstructure:"username"`
	BetType   string `json:"bet_type" mapstructure:"bet_type"`
	Status    struct {
		Code string `json:"code" mapstructure:"code"`
	} `json:"status" mapstructure:"status"`
	PriceList []struct {
		Effective struct {
			Price float64       `json:"price" mapstructure:"price"`
			Min   []interface{} `json:"min" mapstructure:"min"`
			Max   []interface{} `json:"max" mapstructure:"max"`
		} `json:"effective" mapstructure:"effective"`
	} `json:"price_list" mapstructure:"price_list"`
	Ts float64 `json:"ts"`
}

type BetslipData struct {
	BetslipMessage BetslipMessage
	PmmMap         map[string]PmmMessage
	Ts             float64
}

type BetslipMessage struct {
	BetslipId          string  `json:"betslip_id" mapstructure:"betslip_id"`
	Sport              string  `json:"sport" mapstructure:"sport"`
	EventId            string  `json:"event_id" mapstructure:"event_id"`
	BetType            string  `json:"bet_type" mapstructure:"bet_type"`
	BetTypeTemplate    string  `json:"bet_type_template" mapstructure:"bet_type_template"`
	BetTypeDescription string  `json:"bet_type_description" mapstructure:"bet_type_description"`
	ExpiryTs           float64 `json:"expiry_ts" mapstructure:"expiry_ts"`
	IsOpen             bool    `json:"is_open" mapstructure:"is_open"`
	CloseReason        string  `json:"close_reason" mapstructure:"close_reason"`
	Accounts           []struct {
		Bookie   string `json:"bookie" mapstructure:"bookie"`
		Username string `json:"username" mapstructure:"username"`
		BetType  string `json:"bet_type" mapstructure:"bet_type"`
	} `json:"accounts" mapstructure:"accounts"`
	MultipleAccounts      bool     `json:"multiple_accounts" mapstructure:"multiple_accounts"`
	EquivalentBets        bool     `json:"equivalent_bets" mapstructure:"equivalent_bets"`
	EquivalentBetsBookies []string `json:"equivalent_bets_bookies" mapstructure:"equivalent_bets_bookies"`
	WantBookies           []string `json:"want_bookies" mapstructure:"want_bookies"`
	BookiesWithOffers     []string `json:"bookies_with_offers" mapstructure:"bookies_with_offers"`
	CustomerUsername      string   `json:"customer_username" mapstructure:"customer_username"`
	CustomerCcy           string   `json:"customer_ccy" mapstructure:"customer_ccy"`
	BetslipType           string   `json:"betslip_type" mapstructure:"betslip_type"`
	Ts                    float64  `json:"ts"`
}
type BalanceMessage struct {
	Balance   []interface{} `json:"balance" mapstructure:"balance"`
	OpenStake []interface{} `json:"open_stake" mapstructure:"open_stake"`
	Ts        float64       `json:"ts"`
	UserID    int64         `json:"user_id"`
}

type User struct {
	CreatedAt time.Time  `json:"-"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	Host      string     `json:"host"`
	SessionID string     `json:"-"`
	ID        int64      `json:"id"`
	Active    bool       `json:"active"`
	Currency  string     `json:"currency"`
}
type XRateMessage struct {
	Ccy  string  `mapstructure:"ccy" json:"ccy,omitempty"`
	Rate float64 `mapstructure:"rate" json:"rate,omitempty"`
	Ts   float64 `json:"ts,omitempty"`
}
type Result struct {
	HtHome int `json:"ht_home"`
	HtAway int `json:"ht_away"`
	FtHome int `json:"ft_home"`
	FtAway int `json:"ft_away"`
}
type EventInfo struct {
	EventId            string `json:"event_id" mapstructure:"event_id"`
	HomeId             int    `json:"home_id" mapstructure:"home_id"`
	HomeTeam           string `json:"home_team" mapstructure:"home_team"`
	AwayId             int    `json:"away_id" mapstructure:"away_id"`
	AwayTeam           string `json:"away_team" mapstructure:"away_team"`
	CompetitionId      int    `json:"competition_id" mapstructure:"competition_id"`
	CompetitionName    string `json:"competition_name" mapstructure:"competition_name"`
	CompetitionCountry string `json:"competition_country" mapstructure:"competition_country"`
	StartTime          string `json:"start_time" mapstructure:"start_time"`
	Date               string `json:"date" mapstructure:"date"`
	Result             Result `json:"result,omitempty"`
}

type OrderData struct {
	OrderId            int64         `json:"order_id" mapstructure:"order_id"`
	OrderType          string        `json:"order_type" mapstructure:"order_type"`
	BetType            string        `json:"bet_type" mapstructure:"bet_type"`
	BetTypeTemplate    string        `json:"bet_type_template" mapstructure:"bet_type_template"`
	BetTypeDescription string        `json:"bet_type_description" mapstructure:"bet_type_description"`
	Sport              string        `json:"sport" mapstructure:"sport"`
	Placer             string        `json:"placer" mapstructure:"placer"`
	WantPrice          float64       `json:"want_price" mapstructure:"want_price"`
	WantStake          []interface{} `json:"want_stake" mapstructure:"want_stake"`
	CcyRate            float64       `json:"ccy_rate" mapstructure:"ccy_rate"`
	PlacementTime      string        `json:"placement_time" mapstructure:"placement_time"`
	ExpiryTime         string        `json:"expiry_time" mapstructure:"expiry_time"`
	Closed             bool          `json:"closed" mapstructure:"closed"`
	CloseReason        string        `json:"close_reason" mapstructure:"close_reason"`
	EventInfo          EventInfo     `json:"event_info" mapstructure:"event_info"`
	UserData           string        `json:"user_data" mapstructure:"user_data"`
	Status             string        `json:"status" mapstructure:"status"`
	KeepOpenIr         bool          `json:"keep_open_ir" mapstructure:"keep_open_ir"`
	ExchangeMode       string        `json:"exchange_mode" mapstructure:"exchange_mode"`
	Price              float64       `json:"price" mapstructure:"price"`
	Stake              []interface{} `json:"stake" mapstructure:"stake"`
	ProfitLoss         []interface{} `json:"profit_loss" mapstructure:"profit_loss"`
	Bets               []BetMessage  `json:"bets" mapstructure:"bets"`
	BetBarValues       struct {
		Success    []interface{} `json:"success" mapstructure:"success"`
		Inprogress []interface{} `json:"inprogress" mapstructure:"inprogress"`
		Danger     []interface{} `json:"danger" mapstructure:"danger"`
		Unplaced   []interface{} `json:"unplaced" mapstructure:"unplaced"`
	} `json:"bet_bar_values" mapstructure:"bet_bar_values"`
	Ts            float64  `json:"ts"`
	BetBookieList []string `json:"bet_bookie_list"`
}

type InfoMessage struct {
	QueueSize        int     `mapstructure:"queue_size" json:"queue_size"`
	QueueSizeMax     int     `mapstructure:"queue_size_max" json:"queue_size_max"`
	RegisteredEvents int     `mapstructure:"registered_events" json:"registered_events"`
	MaxQueueSize     int     `mapstructure:"max_queue_size" json:"max_queue_size"`
	Ts               float64 `json:"ts"`
}

type BetMessage struct {
	BetId      int64         `json:"bet_id" mapstructure:"bet_id"`
	BetType    string        `json:"bet_type" mapstructure:"bet_type"`
	Bookie     string        `json:"bookie" mapstructure:"bookie"`
	CcyRate    float64       `json:"ccy_rate" mapstructure:"ccy_rate"`
	EventId    string        `json:"event_id" mapstructure:"event_id"`
	GotPrice   float64       `json:"got_price" mapstructure:"got_price"`
	GotStake   []interface{} `json:"got_stake" mapstructure:"got_stake"`
	OrderId    int           `json:"order_id" mapstructure:"order_id"`
	ProfitLoss []interface{} `json:"profit_loss" mapstructure:"profit_loss"`
	Reconciled bool          `json:"reconciled" mapstructure:"reconciled"`
	Sport      string        `json:"sport" mapstructure:"sport"`
	Status     struct {
		Code string `json:"code" mapstructure:"code"`
	} `json:"status" mapstructure:"status"`
	Username  string        `json:"username" mapstructure:"username"`
	WantPrice float64       `json:"want_price" mapstructure:"want_price"`
	WantStake []interface{} `json:"want_stake" mapstructure:"want_stake"`
}
