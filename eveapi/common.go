package eveapi

// Common structures

type imageList struct {
	I32x32 struct {
		Href string `json:"href,omitempty"`
	} `json:"32x32,omitempty"`
	I64x64 struct {
		Href string `json:"href,omitempty"`
	} `json:"64x64,omitempty"`
	I128x128 struct {
		Href string `json:"href,omitempty"`
	} `json:"128x128,omitempty"`
	I256x256 struct {
		Href string `json:"href,omitempty"`
	} `json:"256x256,omitempty"`
}

// Corporation or alliance in most structures.
type entityReference struct {
	idHref
	Name  string    `json:"name"`
	IsNPC bool      `json:"isNPC"`
	Logo  imageList `json:"logo"`
}

type characterReference struct {
	idHref
	Name        string          `json:"name"`
	Corporation entityReference `json:"corporation,omitempty"`
	Alliance    entityReference `json:"alliance,omitempty"`
	IsNPC       bool            `json:"isNPC"`
	Capsuleer   struct {
		Href string
	}
	Portrait imageList `json:"portrait"`
}

// Killmail references
type itemReference struct {
	idHref
	Name string `json:"name"`
	Icon struct {
		Href string `json:"href"`
	} `json:"icon"`
}

type simpleHref struct {
	Href string `json:"href"`
}

// ID & Href references
type idHref struct {
	Href string `json:"href"`
	ID   int64  `json:"id"`
}
