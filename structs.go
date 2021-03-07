package main

// Session is what connects you to the Tebex API
type Session struct {
	// Secret AKA X-Tebex-Secret is required to make requests to the API.
	Secret string
}

type ServerInfo struct {
	Account Account `json:"account"`
	Server  Server  `json:"server"`
}

type Account struct {
	ID       int    `json:"id"`
	Domain   string `json:"domain"`
	Name     string `json:"name"`
	Currency struct {
		Iso4217 string `json:"iso_4217"`
		Symbol  string `json:"symbol"`
	} `json:"currency"`
	OnlineMode bool   `json:"online_mode"`
	GameType   string `json:"game_type"`
	LogEvents  bool   `json:"log_events"`
}
type Server struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CommandQueue struct {
	Meta struct {
		ExecuteOffline bool `json:"execute_offline"`
		NextCheck      int  `json:"next_check"`
		More           bool `json:"more"`
	} `json:"meta"`
	Players `json:"players"`
}
type Players []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
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
