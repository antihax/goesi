package goesi

// Factions resolves faction name to ID
var FactionsByName = map[string]int32{"Caldari": 500001, "Minmatar": 500002, "Amarr": 500003, "Gallente": 500004}

// Factions resolves faction ID to Name
var FactionsByID = map[int32]string{500001: "Caldari", 500002: "Minmatar", 500003: "Amarr", 500004: "Gallente"}

// FactionsAtWar resolves two enemy parties for each factionID
var FactionsAtWar = map[int32][]int32{
	500001: {500002, 500004}, // Caldari  : Minmatar, Gallente
	500003: {500002, 500004}, // Amarr    : Minmatar, Gallente
	500002: {500001, 500003}, // Minmatar : Caldari, Amarr
	500004: {500001, 500003}, // Gallente : Caldari, Amarr
}

// FactionAllies resolves friendly faction war IDs
var FactionAllies = map[int32]int32{
	500001: 500003,
	500003: 500001,
	500002: 500004,
	500004: 500002,
}
