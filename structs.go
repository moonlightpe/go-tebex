package main

import (
	"time"
)

// Session ...
type Session struct {
	// Secret AKA X-Tebex-Secret is required to make every request to the API.
	Secret string
}

// ServerInfo tells you information about your webstore.
type ServerInfo struct {
	// Account ...
	Account struct {
		ID         int    `json:"id"`
		Domain     string `json:"domain"`
		Name       string `json:"name"`
		Currency   `json:"currency"`
		OnlineMode bool   `json:"online_mode"`
		GameType   string `json:"game_type"`
		LogEvents  bool   `json:"log_events"`
	} `json:"account"`
	// Server ...
	Server struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"server"`
}

// Currency ...
type Currency struct {
	Iso4217 string `json:"iso_4217"`
	Symbol  string `json:"symbol"`
}

type CommandQueue struct {
	Meta struct {
		ExecuteOffline bool `json:"execute_offline"`
		NextCheck      int  `json:"next_check"`
		More           bool `json:"more"`
	} `json:"meta"`
	Players `json:"players"`
}
type Player []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type Players = Player

type OfflineQueue struct {
	Meta struct {
		Limited bool `json:"limited"`
	} `json:"meta"`
	Commands []struct {
		ID         int    `json:"id"`
		Command    string `json:"command"`
		Payment    int    `json:"payment"`
		Package    int    `json:"package"`
		Conditions struct {
			Delay int `json:"delay"`
		} `json:"conditions"`
		Player Player `json:"player"`
	} `json:"commands"`
}

type OnlineQueue struct {
	Commands []struct {
		ID         int    `json:"id"`
		Command    string `json:"command"`
		Payment    int    `json:"payment"`
		Package    int    `json:"package"`
		Conditions struct {
			Delay int `json:"delay"`
			Slots int `json:"slots"`
		} `json:"conditions"`
	} `json:"commands"`
}

type Listings struct {
	Categories []struct {
		ID                int           `json:"id"`
		Order             int           `json:"order"`
		Name              string        `json:"name"`
		OnlySubcategories bool          `json:"only_subcategories"`
		Subcategories     []interface{} `json:"subcategories"`
		Packages          []struct {
			ID    int    `json:"id"`
			Order int    `json:"order"`
			Name  string `json:"name"`
			Price string `json:"price"`
			Image bool   `json:"image"`
			Sale  struct {
				Active   bool   `json:"active"`
				Discount string `json:"discount"`
			} `json:"sale"`
		} `json:"packages"`
	} `json:"categories"`
}

type Package struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Image        bool   `json:"image"`
	Price        int    `json:"price"`
	ExpiryLength int    `json:"expiry_length"`
	ExpiryPeriod string `json:"expiry_period"`
	Type         string `json:"type"`
	Category     struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	GlobalLimit       int    `json:"global_limit"`
	GlobalLimitPeriod string `json:"global_limit_period"`
	UserLimit         int    `json:"user_limit"`
	UserLimitPeriod   string `json:"user_limit_period"`
	Servers           []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"servers"`
	RequiredPackages []interface{} `json:"required_packages"`
	RequireAny       bool          `json:"require_any"`
	CreateGiftcard   bool          `json:"create_giftcard"`
	ShowUntil        bool          `json:"show_until"`
	GuiItem          string        `json:"gui_item"`
	Disabled         bool          `json:"disabled"`
	DisableQuantity  bool          `json:"disable_quantity"`
	CustomPrice      bool          `json:"custom_price"`
	ChooseServer     bool          `json:"choose_server"`
	LimitExpires     bool          `json:"limit_expires"`
	InheritCommands  bool          `json:"inherit_commands"`
	VariableGiftcard bool          `json:"variable_giftcard"`
}

// Packages are a list of package
type Packages = []Package

type CommunityGoal struct {
	ID            int         `json:"id"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	Account       int         `json:"account"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Image         string      `json:"image"`
	Target        string      `json:"target"`
	Current       string      `json:"current"`
	Repeatable    int         `json:"repeatable"`
	LastAchieved  interface{} `json:"last_achieved"`
	TimesAchieved int         `json:"times_achieved"`
	Status        string      `json:"status"`
	Sale          int         `json:"sale"`
}

type CommunityGoals = []CommunityGoal

type AllPayments []struct {
	ID       int    `json:"id"`
	Amount   string `json:"amount"`
	Date     string `json:"date"`
	Currency struct {
		Iso4217 string `json:"iso_4217"`
		Symbol  string `json:"symbol"`
	} `json:"currency"`
	Player `json:"player"`
}

type PaginatedPayments struct {
	Total       int    `json:"total"`
	PerPage     int    `json:"per_page"`
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	NextPageURL string `json:"next_page_url"`
	PrevPageURL string `json:"prev_page_url"`
	From        int    `json:"from"`
	To          int    `json:"to"`
	Data        []struct {
		ID      int       `json:"id"`
		Amount  string    `json:"amount"`
		Date    time.Time `json:"date"`
		Gateway struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"gateway"`
		Status   string `json:"status"`
		Currency `json:"currency"`
		Player   `json:"player"`
		Packages []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"packages"`
	} `json:"data"`
}

type Payment struct {
	ID       int    `json:"id"`
	Amount   string `json:"amount"`
	Status   string `json:"status"`
	Date     string `json:"date"`
	Currency struct {
		Iso4217 string `json:"iso_4217"`
		Symbol  string `json:"symbol"`
	} `json:"currency"`
	Player struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		UUID string `json:"uuid"`
	} `json:"player"`
	Packages []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"packages"`
}

type CheckoutURL struct {
	URL     string `json:"url"`
	Expires string `json:"expires"`
}

type GiftCard struct {
	Data struct {
		ID      int    `json:"id"`
		Code    string `json:"code"`
		Balance struct {
			Starting  string `json:"starting"`
			Remaining string `json:"remaining"`
			Currency  string `json:"currency"`
		} `json:"balance"`
		Note string `json:"note"`
		Void bool   `json:"void"`
	} `json:"data"`
}

type GiftCards = []GiftCard

type Coupon struct {
	Data struct {
		ID        int    `json:"id"`
		Code      string `json:"code"`
		Effective struct {
			Type       string        `json:"type"`
			Packages   []interface{} `json:"packages"`
			Categories []interface{} `json:"categories"`
		} `json:"effective"`
		Discount struct {
			Type       string `json:"type"`
			Percentage int    `json:"percentage"`
			Value      int    `json:"value"`
		} `json:"discount"`
		Expire struct {
			RedeemUnlimited string `json:"redeem_unlimited"`
			ExpireNever     string `json:"expire_never"`
			Limit           int    `json:"limit"`
			Date            string `json:"date"`
		} `json:"expire"`
		BasketType string `json:"basket_type"`
		StartDate  string `json:"start_date"`
		UserLimit  int    `json:"user_limit"`
		Minimum    int    `json:"minimum"`
		Username   string `json:"username"`
		Note       string `json:"note"`
	} `json:"data"`
}

type AllCoupons struct {
	Pagination struct {
		Totalresults int         `json:"totalResults"`
		Currentpage  int         `json:"currentPage"`
		Lastpage     int         `json:"lastPage"`
		Previous     interface{} `json:"previous"`
		Next         string      `json:"next"`
	} `json:"pagination"`
	Data []Coupon `json:"data"`
}

type Ban struct {
	ID           int       `json:"id"`
	Time         time.Time `json:"time"`
	IP           string    `json:"ip"`
	PaymentEmail string    `json:"payment_email"`
	Reason       string    `json:"reason"`
	User         struct {
		Ign  string `json:"ign"`
		UUID string `json:"uuid"`
	} `json:"user"`
}

type Bans = []Ban

type Sales struct {
	Data []struct {
		ID        int `json:"id"`
		Effective struct {
			Type       string        `json:"type"`
			Packages   []int         `json:"packages"`
			Categories []interface{} `json:"categories"`
		} `json:"effective"`
		Discount struct {
			Type       string `json:"type"`
			Percentage int    `json:"percentage"`
			Value      int    `json:"value"`
		} `json:"discount"`
		Start  int `json:"start"`
		Expire int `json:"expire"`
		Order  int `json:"order"`
	} `json:"data"`
}

type Lookup struct {
	Player struct {
		ID               string `json:"id"`
		CreatedAt        string `json:"created_at"`
		UpdatedAt        string `json:"updated_at"`
		CacheExpire      string `json:"cache_expire"`
		Username         string `json:"username"`
		Meta             string `json:"meta"`
		PluginUsernameID int    `json:"plugin_username_id"`
	} `json:"player"`
	Bancount       int `json:"banCount"`
	Chargebackrate int `json:"chargebackRate"`
	Payments       []struct {
		TxnID    string  `json:"txn_id"`
		Time     int     `json:"time"`
		Price    float64 `json:"price"`
		Currency string  `json:"currency"`
		Status   int     `json:"status"`
	} `json:"payments"`
	Purchasetotals struct {
		Usd float64 `json:"USD"`
		Gbp float64 `json:"GBP"`
	} `json:"purchaseTotals"`
}
