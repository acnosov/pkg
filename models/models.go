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
	SendTime  time.Time
}
type Timing struct {
	WsReceive       time.Time
	BeforeProcess   time.Time
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
	BeginExit       time.Time
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

func (e *EventDB) EventID() string {
	return fmt.Sprintf("%s,%d,%d", e.EventDate.Format("2006-01-02"), e.HomeID, e.AwayID)
}

type EventKey struct {
	Sport     string    `json:"sport"`
	EventDate time.Time `json:"event_date"`
	HomeID    int64     `json:"home_id"`
	AwayID    int64     `json:"away_id"`
}

func (e *EventDB) Key() EventKey {
	return EventKey{Sport: e.Sport, EventDate: e.EventDate, HomeID: e.HomeID, AwayID: e.AwayID}
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
	ScoreHome     *int64    `json:"score_home"`
	ScoreAway     *int64    `json:"score_away"`
}

func (e EventWithScore) EventID() string {
	return fmt.Sprintf("%s,%d,%d", e.EventDate.Format("2006-01-02"), e.HomeID, e.AwayID)
}

type BetslipData struct {
	BetslipMessage BetslipMessage
	PmmMap         map[string]PmmMessage
	Ts             float64
}

type PmmMessage struct {
	BetslipId string `json:"betslip_id"`
	Sport     string `json:"sport"`
	EventId   string `json:"event_id"`
	Bookie    string `json:"bookie"`
	Username  string `json:"username"`
	BetType   string `json:"bet_type"`
	Status    struct {
		Code string `json:"code"`
	} `json:"status"`
	PriceList []PriceList `json:"price_list"`
	Ts        float64     `json:"ts"`
	SendTime  time.Time
}

type PriceList struct {
	Effective Effective `json:"effective"`
}

type Effective struct {
	Price float64       `json:"price"`
	Min   []interface{} `json:"min"`
	Max   []interface{} `json:"max"`
}

type BetslipMessage struct {
	BetslipId          string  `json:"betslip_id"`
	Sport              string  `json:"sport"`
	EventId            string  `json:"event_id"`
	BetType            string  `json:"bet_type"`
	BetTypeTemplate    string  `json:"bet_type_template"`
	BetTypeDescription string  `json:"bet_type_description"`
	ExpiryTs           float64 `json:"expiry_ts"`
	IsOpen             bool    `json:"is_open"`
	CloseReason        string  `json:"close_reason"`
	Accounts           []struct {
		Bookie   string `json:"bookie"`
		Username string `json:"username"`
		BetType  string `json:"bet_type"`
	} `json:"accounts"`
	MultipleAccounts      bool     `json:"multiple_accounts"`
	EquivalentBets        bool     `json:"equivalent_bets"`
	EquivalentBetsBookies []string `json:"equivalent_bets_bookies"`
	WantBookies           []string `json:"want_bookies"`
	BookiesWithOffers     []string `json:"bookies_with_offers"`
	CustomerUsername      string   `json:"customer_username"`
	CustomerCcy           string   `json:"customer_ccy"`
	BetslipType           string   `json:"betslip_type"`
	Ts                    float64  `json:"ts"`
	SendTime              time.Time
}
type BalanceMessage struct {
	Balance   []interface{} `json:"balance"`
	OpenStake []interface{} `json:"open_stake"`
	Ts        float64       `json:"ts"`
	UserID    int64         `json:"user_id"`
	SendTime  time.Time
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
	Ccy      string  ` json:"ccy,omitempty"`
	Rate     float64 `json:"rate,omitempty"`
	Ts       float64 `json:"ts,omitempty"`
	SendTime time.Time
}
type Result struct {
	HtHome int `json:"ht_home"`
	HtAway int `json:"ht_away"`
	FtHome int `json:"ft_home"`
	FtAway int `json:"ft_away"`
}
type EventInfo struct {
	EventId            string `json:"event_id"`
	HomeId             int    `json:"home_id"`
	HomeTeam           string `json:"home_team"`
	AwayId             int    `json:"away_id"`
	AwayTeam           string `json:"away_team"`
	CompetitionId      int    `json:"competition_id"`
	CompetitionName    string `json:"competition_name"`
	CompetitionCountry string `json:"competition_country"`
	StartTime          string `json:"start_time"`
	Date               string `json:"date"`
	Result             Result `json:"result,omitempty"`
}

type OrderData struct {
	OrderId            int64         `json:"order_id"`
	OrderType          string        `json:"order_type"`
	BetType            string        `json:"bet_type"`
	BetTypeTemplate    string        `json:"bet_type_template"`
	BetTypeDescription string        `json:"bet_type_description"`
	Sport              string        `json:"sport"`
	Placer             string        `json:"placer"`
	WantPrice          float64       `json:"want_price"`
	WantStake          []interface{} `json:"want_stake"`
	CcyRate            float64       `json:"ccy_rate"`
	PlacementTime      string        `json:"placement_time"`
	ExpiryTime         string        `json:"expiry_time"`
	Closed             bool          `json:"closed"`
	CloseReason        string        `json:"close_reason"`
	EventInfo          EventInfo     `json:"event_info"`
	UserData           string        `json:"user_data"`
	Status             string        `json:"status"`
	KeepOpenIr         bool          `json:"keep_open_ir"`
	ExchangeMode       string        `json:"exchange_mode"`
	Price              float64       `json:"price"`
	Stake              []interface{} `json:"stake"`
	ProfitLoss         []interface{} `json:"profit_loss"`
	Bets               []BetMessage  `json:"bets"`
	BetBarValues       struct {
		Success    []interface{} `json:"success"`
		Inprogress []interface{} `json:"inprogress"`
		Danger     []interface{} `json:"danger"`
		Unplaced   []interface{} `json:"unplaced"`
	} `json:"bet_bar_values"`
	Ts            float64  `json:"ts"`
	BetBookieList []string `json:"bet_bookie_list"`
	SendTime      time.Time
}

type InfoMessage struct {
	QueueSize        int     `json:"queue_size"`
	QueueSizeMax     int     `json:"queue_size_max"`
	RegisteredEvents int     `json:"registered_events"`
	MaxQueueSize     int     `json:"max_queue_size"`
	Ts               float64 `json:"ts"`
	SendTime         time.Time
}

type BetMessage struct {
	BetId      int64         `json:"bet_id"`
	BetType    string        `json:"bet_type"`
	Bookie     string        `json:"bookie"`
	CcyRate    float64       `json:"ccy_rate"`
	EventId    string        `json:"event_id"`
	GotPrice   float64       `json:"got_price"`
	GotStake   []interface{} `json:"got_stake"`
	OrderId    int           `json:"order_id"`
	ProfitLoss []interface{} `json:"profit_loss"`
	Reconciled bool          `json:"reconciled"`
	Sport      string        `json:"sport"`
	Status     struct {
		Code string `json:"code"`
	} `json:"status"`
	Username     string        `json:"username"`
	WantPrice    float64       `json:"want_price"`
	WantStake    []interface{} `json:"want_stake"`
	ExchangeRole any           `json:"exchange_role"`
}

type BetslipClosedMessage struct {
	BetslipId   string `json:"betslip_id"`
	CloseReason string `json:"close_reason"`
	SendTime    time.Time
}

type SyncMessage struct {
	Token string `json:"token" `
}
