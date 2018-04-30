package notification

type AcceptedAlly struct {
	AllyID   int32   `yaml:"allyID"`
	CharID   int32   `yaml:"charID"`
	EnemyID  int32   `yaml:"enemyID"`
	IskValue float64 `yaml:"iskValue"`
	Time     int64   `yaml:"time"`
}

type AcceptedSurrender struct {
	CharID     int32   `yaml:"charID"`
	EntityID   int32   `yaml:"entityID"`
	IskValue   float64 `yaml:"iskValue"`
	OfferingID int32   `yaml:"offeringID"`
}

type AllMaintenanceBillMsg struct {
	AllianceID int32 `yaml:"allianceID"`
	DueDate    int64 `yaml:"dueDate"`
}

type AllWarCorpJoinedAllianceMsg struct {
	AllianceID int32 `yaml:"allianceID"`
	CorpID     int32 `yaml:"corpID"`
}

type AllWarDeclaredMsg struct {
	AgainstID    int32   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int32   `yaml:"declaredByID"`
	DelayHours   int32   `yaml:"delayHours"`
	HostileState int32   `yaml:"hostileState"`
}

type AllWarInvalidatedMsg struct {
	AgainstID    int32   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int32   `yaml:"declaredByID"`
	DelayHours   int32   `yaml:"delayHours"`
	HostileState int32   `yaml:"hostileState"`
}

type AllWarRetractedMsg struct {
	AgainstID    int32   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int32   `yaml:"declaredByID"`
	DelayHours   int32   `yaml:"delayHours"`
	HostileState int32   `yaml:"hostileState"`
}

type AllWarSurrenderMsg struct {
	AgainstID    int32   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int32   `yaml:"declaredByID"`
	DelayHours   int32   `yaml:"delayHours"`
	HostileState int32   `yaml:"hostileState"`
}

type AllianceCapitalChanged struct {
	AllianceID    int32 `yaml:"allianceID"`
	SolarSystemID int32 `yaml:"solarSystemID"`
}

type AllyContractCancelled struct {
	AggressorID  int32 `yaml:"aggressorID"`
	DefenderID   int32 `yaml:"defenderID"`
	TimeFinished int64 `yaml:"timeFinished"`
}

type AllyJoinedWarAggressorMsg struct {
	AllyID     int32 `yaml:"allyID"`
	DefenderID int32 `yaml:"defenderID"`
	StartTime  int64 `yaml:"startTime"`
}

type AllyJoinedWarAllyMsg struct {
	AggressorID int32 `yaml:"aggressorID"`
	AllyID      int32 `yaml:"allyID"`
	DefenderID  int32 `yaml:"defenderID"`
	StartTime   int64 `yaml:"startTime"`
}

type AllyJoinedWarDefenderMsg struct {
	AggressorID int32 `yaml:"aggressorID"`
	AllyID      int32 `yaml:"allyID"`
	StartTime   int64 `yaml:"startTime"`
}

type BillOutOfMoneyMsg struct {
	BillTypeID int32 `yaml:"billTypeID"`
	DueDate    int64 `yaml:"dueDate"`
}

type BillPaidCorpAllMsg struct {
	Amount  int32 `yaml:"amount"`
	DueDate int64 `yaml:"dueDate"`
}

type BountyClaimMsg struct {
	Amount float64 `yaml:"amount"`
	CharID int32   `yaml:"charID"`
}

type BountyESSShared struct {
	CharID   int32   `yaml:"charID"`
	MyIsk    float64 `yaml:"myIsk"`
	TotalIsk float64 `yaml:"totalIsk"`
}

type BountyESSTaken struct {
	CharID   int32   `yaml:"charID"`
	MyIsk    float64 `yaml:"myIsk"`
	TotalIsk float64 `yaml:"totalIsk"`
}

type BountyPlacedAlliance struct {
	Bounty         float64 `yaml:"bounty"`
	BountyPlacerID int32   `yaml:"bountyPlacerID"`
}

type BountyPlacedChar struct {
	Bounty         float64 `yaml:"bounty"`
	BountyPlacerID int32   `yaml:"bountyPlacerID"`
}

type BountyPlacedCorp struct {
	Bounty         float64 `yaml:"bounty"`
	BountyPlacerID int32   `yaml:"bountyPlacerID"`
}

type BountyYourBountyClaimed struct {
	Bounty   float64 `yaml:"bounty"`
	VictimID int32   `yaml:"victimID"`
}

type BuddyConnectContactAdd struct {
	Level   int32  `yaml:"level"`
	Message string `yaml:"message"`
}

type CharAppAcceptMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int32  `yaml:"charID"`
	CorpID          int32  `yaml:"corpID"`
}

type CharAppRejectMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int32  `yaml:"charID"`
	CorpID          int32  `yaml:"corpID"`
}

type CharAppWithdrawMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int32  `yaml:"charID"`
	CorpID          int32  `yaml:"corpID"`
}

type CharLeftCorpMsg struct {
	CharID int32 `yaml:"charID"`
	CorpID int32 `yaml:"corpID"`
}

type CharMedalMsg struct {
	CorpID  int32  `yaml:"corpID"`
	MedalID int32  `yaml:"medalID"`
	Reason  string `yaml:"reason"`
}

type CharTerminationMsg struct {
	CharID      int32   `yaml:"charID"`
	CorpID      int32   `yaml:"corpID"`
	RoleName    string  `yaml:"roleName"`
	RoleNameIDs []int32 `yaml:"roleNameIDs"`
	Security    float64 `yaml:"security"`
}

type CloneActivationMsg struct {
	CloneBought     int32 `yaml:"cloneBought"`
	CloneStationID  int32 `yaml:"cloneStationID"`
	CloneTypeID     int32 `yaml:"cloneTypeID"`
	CorpStationID   int32 `yaml:"corpStationID"`
	LastCloned      int64 `yaml:"lastCloned"`
	PodKillerID     int32 `yaml:"podKillerID"`
	SkillID         int32 `yaml:"skillID"`
	SkillPointsLost int32 `yaml:"skillPointsLost"`
}

type CloneActivationMsg2 struct {
	CloneStationID int32 `yaml:"cloneStationID"`
	CorpStationID  int32 `yaml:"corpStationID"`
	LastCloned     int64 `yaml:"lastCloned"`
	PodKillerID    int32 `yaml:"podKillerID"`
}

type CloneMovedMsg struct {
	CharsInCorpID int32 `yaml:"charsInCorpID"`
	CorpID        int32 `yaml:"corpID"`
	NewStationID  int32 `yaml:"newStationID"`
	StationID     int32 `yaml:"stationID"`
}

type CloneRevokedMsg2 struct {
	CorpID       int32 `yaml:"corpID"`
	NewStationID int32 `yaml:"newStationID"`
	StationID    int64 `yaml:"stationID"`
}

type ContactAdd struct {
	Level   int32  `yaml:"level"`
	Message string `yaml:"message"`
}

type ContactEdit struct {
	Level   float64 `yaml:"level"`
	Message string  `yaml:"message"`
}

type ContainerPasswordMsg struct {
	CharID        int32  `yaml:"charID"`
	Password      string `yaml:"password"`
	PasswordType  string `yaml:"passwordType"`
	SolarSystemID int32  `yaml:"solarSystemID"`
	StationID     int32  `yaml:"stationID"`
	TypeID        int32  `yaml:"typeID"`
}

type CorpAllBillMsg struct {
	Amount      float64 `yaml:"amount"`
	BillTypeID  int32   `yaml:"billTypeID"`
	CreditorID  int32   `yaml:"creditorID"`
	CurrentDate int64   `yaml:"currentDate"`
	DebtorID    int32   `yaml:"debtorID"`
	DueDate     int64   `yaml:"dueDate"`
	ExternalID  int32   `yaml:"externalID"`
	ExternalID2 int32   `yaml:"externalID2"`
}

type CorpAppAcceptMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int32  `yaml:"charID"`
	CorpID          int32  `yaml:"corpID"`
}

type CorpAppInvitedMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int32  `yaml:"charID"`
	CorpID          int32  `yaml:"corpID"`
	InvokingCharID  int32  `yaml:"invokingCharID"`
}

type CorpAppNewMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int32  `yaml:"charID"`
	CorpID          int32  `yaml:"corpID"`
}

type CorpAppRejectCustomMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int32  `yaml:"charID"`
	CorpID          int32  `yaml:"corpID"`
	CustomMessage   string `yaml:"customMessage"`
}

type CorpAppRejectMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int32  `yaml:"charID"`
	CorpID          int32  `yaml:"corpID"`
}

type CorpDividendMsg struct {
	Amount    float64 `yaml:"amount"`
	CorpID    int32   `yaml:"corpID"`
	IsMembers bool    `yaml:"isMembers"`
	Payout    float64 `yaml:"payout"`
}

type CorpFriendlyFireDisableTimerCompleted struct {
	CorpID int32 `yaml:"corpID"`
}

type CorpFriendlyFireDisableTimerStarted struct {
	CharID       int32 `yaml:"charID"`
	CorpID       int32 `yaml:"corpID"`
	TimeFinished int64 `yaml:"timeFinished"`
}

type CorpFriendlyFireEnableTimerCompleted struct {
	CorpID int32 `yaml:"corpID"`
}

type CorpFriendlyFireEnableTimerStarted struct {
	CharID       int32 `yaml:"charID"`
	CorpID       int32 `yaml:"corpID"`
	TimeFinished int64 `yaml:"timeFinished"`
}

type CorpKicked struct {
	CorpID int32 `yaml:"corpID"`
}

type CorpLiquidationMsg struct {
	Amount float64 `yaml:"amount"`
	CorpID int32   `yaml:"corpID"`
	Payout float64 `yaml:"payout"`
}

type CorpNewCEOMsg struct {
	CorpID   int32 `yaml:"corpID"`
	NewCeoID int32 `yaml:"newCeoID"`
	OldCeoID int32 `yaml:"oldCeoID"`
}

type CorpNewsMsg struct {
	Body      string `yaml:"body"`
	CorpID    int32  `yaml:"corpID"`
	InEffect  int32  `yaml:"inEffect"`
	Parameter int32  `yaml:"parameter"`
	VoteType  int32  `yaml:"voteType"`
}

type CorpTaxChangeMsg struct {
	CorpID     int32   `yaml:"corpID"`
	NewTaxRate float64 `yaml:"newTaxRate"`
	OldTaxRate float64 `yaml:"oldTaxRate"`
}

type CorpVoteMsg struct {
	Body    string `yaml:"body"`
	Subject string `yaml:"subject"`
}

type CorpWarDeclaredMsg struct {
	AgainstID    int32   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int32   `yaml:"declaredByID"`
}

type CorpWarFightingLegalMsg struct {
	AgainstID    int32   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int32   `yaml:"declaredByID"`
}

type CorpWarInvalidatedMsg struct {
	AgainstID    int32   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int32   `yaml:"declaredByID"`
}

type CorpWarRetractedMsg struct {
	AgainstID    int32   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int32   `yaml:"declaredByID"`
}

type CorpWarSurrenderMsg struct {
	AgainstID    int32   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int32   `yaml:"declaredByID"`
}

type CustomsMsg struct {
	FactionID int32 `yaml:"factionID"`
	LostList  []struct {
		Fine     float64 `yaml:"fine"`
		Penalty  float64 `yaml:"penalty"`
		Quantity int32   `yaml:"quantity"`
		TypeID   int32   `yaml:"typeID"`
	} `yaml:"lostList"`
	SecurityLevel    float64 `yaml:"securityLevel"`
	ShouldAttack     int32   `yaml:"shouldAttack"`
	ShouldConfiscate int32   `yaml:"shouldConfiscate"`
	SolarSystemID    int32   `yaml:"solarSystemID"`
	StandingDivision float64 `yaml:"standingDivision"`
}

type DeclareWar struct {
	CharID     int32 `yaml:"charID"`
	DefenderID int32 `yaml:"defenderID"`
	EntityID   int32 `yaml:"entityID"`
}

type EntosisCaptureStarted struct {
	SolarSystemID   int32 `yaml:"solarSystemID"`
	StructureTypeID int32 `yaml:"structureTypeID"`
}

type FWAllianceWarningMsg struct {
	AllianceID       int32   `yaml:"allianceID"`
	CorpList         string  `yaml:"corpList"`
	FactionID        int32   `yaml:"factionID"`
	RequiredStanding float64 `yaml:"requiredStanding"`
}

type FWCharRankGainMsg struct {
	FactionID int32 `yaml:"factionID"`
	NewRank   int32 `yaml:"newRank"`
}

type FWCharRankLossMsg struct {
	FactionID int32 `yaml:"factionID"`
	NewRank   int32 `yaml:"newRank"`
}

type FWCorpJoinMsg struct {
	CorpID    int32 `yaml:"corpID"`
	FactionID int32 `yaml:"factionID"`
}

type FWCorpKickMsg struct {
	CorpID           int32   `yaml:"corpID"`
	CurrentStanding  float64 `yaml:"currentStanding"`
	FactionID        int32   `yaml:"factionID"`
	RequiredStanding float64 `yaml:"requiredStanding"`
}

type FWCorpLeaveMsg struct {
	CorpID    int32 `yaml:"corpID"`
	FactionID int32 `yaml:"factionID"`
}

type FWCorpWarningMsg struct {
	CorpID           int32   `yaml:"corpID"`
	CurrentStanding  float64 `yaml:"currentStanding"`
	FactionID        int32   `yaml:"factionID"`
	RequiredStanding float64 `yaml:"requiredStanding"`
}

type FacWarCorpJoinRequestMsg struct {
	CorpID    int32 `yaml:"corpID"`
	FactionID int32 `yaml:"factionID"`
}

type FacWarCorpJoinWithdrawMsg struct {
	CorpID    int32 `yaml:"corpID"`
	FactionID int32 `yaml:"factionID"`
}

type FacWarCorpLeaveRequestMsg struct {
	CorpID    int32 `yaml:"corpID"`
	FactionID int32 `yaml:"factionID"`
}

type FacWarCorpLeaveWithdrawMsg struct {
	CorpID    int32 `yaml:"corpID"`
	FactionID int32 `yaml:"factionID"`
}

type FacWarLPDisqualifiedEvent struct {
	Amount               int32 `yaml:"amount"`
	CharRefID            int32 `yaml:"charRefID"`
	CorpID               int32 `yaml:"corpID"`
	DisqualificationType int32 `yaml:"disqualificationType"`
	Event                int32 `yaml:"event"`
	ItemRefID            int32 `yaml:"itemRefID"`
	LocationID           int32 `yaml:"locationID"`
}

type FacWarLPDisqualifiedKill struct {
	Amount               int32 `yaml:"amount"`
	CharRefID            int32 `yaml:"charRefID"`
	CorpID               int32 `yaml:"corpID"`
	DisqualificationType int32 `yaml:"disqualificationType"`
	Event                int32 `yaml:"event"`
	ItemRefID            int32 `yaml:"itemRefID"`
	LocationID           int32 `yaml:"locationID"`
}

type FacWarLPPayoutEvent struct {
	Amount     int32 `yaml:"amount"`
	CharRefID  int32 `yaml:"charRefID"`
	CorpID     int32 `yaml:"corpID"`
	Event      int32 `yaml:"event"`
	ItemRefID  int32 `yaml:"itemRefID"`
	LocationID int32 `yaml:"locationID"`
}

type FacWarLPPayoutKill struct {
	Amount     int32 `yaml:"amount"`
	CharRefID  int32 `yaml:"charRefID"`
	CorpID     int32 `yaml:"corpID"`
	Event      int32 `yaml:"event"`
	ItemRefID  int32 `yaml:"itemRefID"`
	LocationID int32 `yaml:"locationID"`
}

type GameTimeAdded struct {
}

type GameTimeReceived struct {
	Message      string `yaml:"message"`
	OfferID      int32  `yaml:"offerID"`
	Quantity     int32  `yaml:"quantity"`
	SenderCharID int32  `yaml:"senderCharID"`
}

type GameTimeSent struct {
	ReceiverCharID int32 `yaml:"receiverCharID"`
	SenderCharID   int32 `yaml:"senderCharID"`
}

type GiftReceived struct {
	Message      string `yaml:"message"`
	OfferID      int32  `yaml:"offerID"`
	Quantity     int32  `yaml:"quantity"`
	SenderCharID int32  `yaml:"senderCharID"`
}

type IncursionCompletedMsg struct {
	EmailMessageID float64   `yaml:"emailMessageId"`
	SolarSystemID  int32     `yaml:"solarSystemID"`
	TaleID         int32     `yaml:"taleID"`
	TopTen         [][]int32 `yaml:"topTen"`
}

type IndustryTeamAuctionLost struct {
	SolarSystemID int32             `yaml:"solarSystemID"`
	SystemBids    map[int32]float64 `yaml:"systemBids"`
	TeamNameInfo  []interface{}     `yaml:"teamNameInfo"`
	TotalIsk      float64           `yaml:"totalIsk"`
	YourAmount    float64           `yaml:"yourAmount"`
}

type InsuranceExpirationMsg struct {
	EndDate   int64  `yaml:"endDate"`
	ShipID    int32  `yaml:"shipID"`
	ShipName  string `yaml:"shipName"`
	StartDate int64  `yaml:"startDate"`
}

type InsuranceFirstShipMsg struct {
	IsHouseWarmingGift int32 `yaml:"isHouseWarmingGift"`
	ShipTypeID         int32 `yaml:"shipTypeID"`
}

type InsuranceInvalidatedMsg struct {
	EndDate   int64 `yaml:"endDate"`
	OwnerID   int32 `yaml:"ownerID"`
	Reason    int32 `yaml:"reason"`
	ShipID    int32 `yaml:"shipID"`
	StartDate int64 `yaml:"startDate"`
	TypeID    int32 `yaml:"typeID"`
}

type InsuranceIssuedMsg struct {
	EndDate   int64   `yaml:"endDate"`
	ItemID    int64   `yaml:"itemID"`
	Level     float64 `yaml:"level"`
	NumWeeks  int32   `yaml:"numWeeks"`
	ShipName  string  `yaml:"shipName"`
	StartDate int64   `yaml:"startDate"`
	TypeID    int32   `yaml:"typeID"`
}

type InsurancePayoutMsg struct {
	Amount float64 `yaml:"amount"`
	ItemID int32   `yaml:"itemID"`
	Payout int32   `yaml:"payout"`
}

type JumpCloneDeletedMsg1 struct {
	LocationID      int64   `yaml:"locationID"`
	LocationOwnerID int32   `yaml:"locationOwnerID"`
	OwnerID         int32   `yaml:"ownerID"`
	TypeIDs         []int32 `yaml:"typeIDs"`
}

type JumpCloneDeletedMsg2 struct {
	DestroyerID     int32   `yaml:"destroyerID"`
	LocationID      int32   `yaml:"locationID"`
	LocationOwnerID int32   `yaml:"locationOwnerID"`
	OwnerID         int32   `yaml:"ownerID"`
	TypeIDs         []int32 `yaml:"typeIDs"`
}

type KillReportFinalBlow struct {
	KillMailHash     string `yaml:"killMailHash"`
	KillMailID       int32  `yaml:"killMailID"`
	VictimID         int32  `yaml:"victimID"`
	VictimShipTypeID int32  `yaml:"victimShipTypeID"`
}

type KillReportVictim struct {
	KillMailHash     string `yaml:"killMailHash"`
	KillMailID       int32  `yaml:"killMailID"`
	VictimShipTypeID int32  `yaml:"victimShipTypeID"`
}

type KillRightAvailable struct {
	CharID     int32   `yaml:"charID"`
	Price      float64 `yaml:"price"`
	ToEntityID int32   `yaml:"toEntityID"`
}

type KillRightAvailableOpen struct {
	CharID int32   `yaml:"charID"`
	Price  float64 `yaml:"price"`
}

type KillRightEarned struct {
	CharID int32 `yaml:"charID"`
}

type KillRightUnavailable struct {
	CharID     int32 `yaml:"charID"`
	ToEntityID int32 `yaml:"toEntityID"`
}

type KillRightUnavailableOpen struct {
	CharID int32 `yaml:"charID"`
}

type KillRightUsed struct {
	CharID int32 `yaml:"charID"`
}

type LocateCharMsg struct {
	AgentLocation struct {
		Region        int32 `yaml:"3"`
		Constellation int32 `yaml:"4"`
		SolarSystem   int32 `yaml:"5"`
		Station       int32 `yaml:"15"`
	} `yaml:"agentLocation"`
	CharacterID    int32 `yaml:"characterID"`
	MessageIndex   int32 `yaml:"messageIndex"`
	TargetLocation struct {
		Region        int32 `yaml:"3"`
		Constellation int32 `yaml:"4"`
		SolarSystem   int32 `yaml:"5"`
		Station       int32 `yaml:"15"`
	} `yaml:"targetLocation"`
}

type MadeWarMutual struct {
	CharID  int32 `yaml:"charID"`
	EnemyID int32 `yaml:"enemyID"`
}

type MercOfferedNegotiationMsg struct {
	AggressorID int32   `yaml:"aggressorID"`
	DefenderID  int32   `yaml:"defenderID"`
	IskValue    float64 `yaml:"iskValue"`
	MercID      int32   `yaml:"mercID"`
}

type MissionOfferExpirationMsg struct {
	Body            []string `yaml:"body"`
	Header          []string `yaml:"header"`
	MissionKeywords struct {
		ObjectiveDestinationID       int32 `yaml:"objectiveDestinationID"`
		ObjectiveDestinationSystemID int32 `yaml:"objectiveDestinationSystemID"`
		ObjectiveLocationID          int32 `yaml:"objectiveLocationID"`
		ObjectiveLocationSystemID    int32 `yaml:"objectiveLocationSystemID"`
		ObjectiveQuantity            int32 `yaml:"objectiveQuantity"`
		ObjectiveTypeID              int32 `yaml:"objectiveTypeID"`
		RewardQuantity               int32 `yaml:"rewardQuantity"`
		RewardTypeID                 int32 `yaml:"rewardTypeID"`
	} `yaml:"missionKeywords"`
}

type MoonminingAutomaticFracture struct {
	MoonID          int32             `yaml:"moonID"`
	MoonLink        string            `yaml:"moonLink"`
	OreVolumeByType map[int32]float64 `yaml:"oreVolumeByType"`
	SolarSystemID   int32             `yaml:"solarSystemID"`
	SolarSystemLink string            `yaml:"solarSystemLink"`
	StructureID     int64             `yaml:"structureID"`
	StructureLink   string            `yaml:"structureLink"`
	StructureName   string            `yaml:"structureName"`
	StructureTypeID int32             `yaml:"structureTypeID"`
}

type MoonminingExtractionCancelled struct {
	CancelledBy     int32  `yaml:"cancelledBy"`
	CancelledByLink string `yaml:"cancelledByLink"`
	MoonID          int32  `yaml:"moonID"`
	MoonLink        string `yaml:"moonLink"`
	SolarSystemID   int32  `yaml:"solarSystemID"`
	SolarSystemLink string `yaml:"solarSystemLink"`
	StructureID     int64  `yaml:"structureID"`
	StructureLink   string `yaml:"structureLink"`
	StructureName   string `yaml:"structureName"`
	StructureTypeID int32  `yaml:"structureTypeID"`
}

type MoonminingExtractionFinished struct {
	AutoTime        int64             `yaml:"autoTime"`
	MoonID          int32             `yaml:"moonID"`
	MoonLink        string            `yaml:"moonLink"`
	OreVolumeByType map[int32]float64 `yaml:"oreVolumeByType"`
	SolarSystemID   int32             `yaml:"solarSystemID"`
	SolarSystemLink string            `yaml:"solarSystemLink"`
	StructureID     int64             `yaml:"structureID"`
	StructureLink   string            `yaml:"structureLink"`
	StructureName   string            `yaml:"structureName"`
	StructureTypeID int32             `yaml:"structureTypeID"`
}

type MoonminingExtractionStarted struct {
	AutoTime        int64             `yaml:"autoTime"`
	MoonID          int32             `yaml:"moonID"`
	MoonLink        string            `yaml:"moonLink"`
	OreVolumeByType map[int32]float64 `yaml:"oreVolumeByType"`
	ReadyTime       int64             `yaml:"readyTime"`
	SolarSystemID   int32             `yaml:"solarSystemID"`
	SolarSystemLink string            `yaml:"solarSystemLink"`
	StartedBy       int32             `yaml:"startedBy"`
	StartedByLink   string            `yaml:"startedByLink"`
	StructureID     int64             `yaml:"structureID"`
	StructureLink   string            `yaml:"structureLink"`
	StructureName   string            `yaml:"structureName"`
	StructureTypeID int32             `yaml:"structureTypeID"`
}

type MoonminingLaserFired struct {
	FiredBy         int32             `yaml:"firedBy"`
	FiredByLink     string            `yaml:"firedByLink"`
	MoonID          int32             `yaml:"moonID"`
	MoonLink        string            `yaml:"moonLink"`
	OreVolumeByType map[int32]float64 `yaml:"oreVolumeByType"`
	SolarSystemID   int32             `yaml:"solarSystemID"`
	SolarSystemLink string            `yaml:"solarSystemLink"`
	StructureID     int64             `yaml:"structureID"`
	StructureLink   string            `yaml:"structureLink"`
	StructureName   string            `yaml:"structureName"`
	StructureTypeID int32             `yaml:"structureTypeID"`
}

type NPCStandingsGained [][]float64

type NPCStandingsLost [][]float64

type OfferedSurrender struct {
	CharID    int32   `yaml:"charID"`
	EntityID  int32   `yaml:"entityID"`
	IskValue  float64 `yaml:"iskValue"`
	OfferedID int32   `yaml:"offeredID"`
}

type OfferedToAlly struct {
	CharID     int32   `yaml:"charID"`
	DefenderID int32   `yaml:"defenderID"`
	EnemyID    int32   `yaml:"enemyID"`
	IskValue   float64 `yaml:"iskValue"`
}

type OldLscMessages struct {
	Subject string `yaml:"subject"`
	Body    string `yaml:"body"`
}

type OperationFinished struct {
	OperationID int32 `yaml:"operationID"`
	Rewards     struct {
		Isk int32 `yaml:"isk"`
	} `yaml:"rewards"`
}

type OrbitalAttacked struct {
	AggressorAllianceID int32   `yaml:"aggressorAllianceID"`
	AggressorCorpID     int32   `yaml:"aggressorCorpID"`
	AggressorID         int32   `yaml:"aggressorID"`
	PlanetID            int32   `yaml:"planetID"`
	PlanetTypeID        int32   `yaml:"planetTypeID"`
	ShieldLevel         float64 `yaml:"shieldLevel"`
	SolarSystemID       int32   `yaml:"solarSystemID"`
	TypeID              int32   `yaml:"typeID"`
}

type OrbitalReinforced struct {
	AggressorAllianceID int32 `yaml:"aggressorAllianceID"`
	AggressorCorpID     int32 `yaml:"aggressorCorpID"`
	AggressorID         int32 `yaml:"aggressorID"`
	PlanetID            int32 `yaml:"planetID"`
	PlanetTypeID        int32 `yaml:"planetTypeID"`
	ReinforceExitTime   int64 `yaml:"reinforceExitTime"`
	SolarSystemID       int32 `yaml:"solarSystemID"`
	TypeID              int32 `yaml:"typeID"`
}

type OwnershipTransferred struct {
	CharacterLinkData       []interface{} `yaml:"characterLinkData"`
	CharacterName           string        `yaml:"characterName"`
	FromCorporationLinkData []interface{} `yaml:"fromCorporationLinkData"`
	FromCorporationName     string        `yaml:"fromCorporationName"`
	SolarSystemLinkData     []interface{} `yaml:"solarSystemLinkData"`
	SolarSystemName         string        `yaml:"solarSystemName"`
	StructureLinkData       []interface{} `yaml:"structureLinkData"`
	StructureName           string        `yaml:"structureName"`
	ToCorporationLinkData   []interface{} `yaml:"toCorporationLinkData"`
	ToCorporationName       string        `yaml:"toCorporationName"`
}

type ReimbursementMsg struct {
	AddCloneInfo  int32 `yaml:"addCloneInfo"`
	ShipTypeID    int32 `yaml:"shipTypeID"`
	SolarSystemID int32 `yaml:"solarSystemID"`
	StationID     int32 `yaml:"stationID"`
}

type ResearchMissionAvailableMsg struct {
}

type RetractsWar struct {
	CharID  int32 `yaml:"charID"`
	EnemyID int32 `yaml:"enemyID"`
}

type SeasonalChallengeCompleted struct {
	ChallengeNameID int32 `yaml:"challenge_name_id"`
	MaxProgress     int32 `yaml:"max_progress"`
	PointsAwarded   int32 `yaml:"points_awarded"`
	ProgressText    int32 `yaml:"progress_text"`
}

type SovAllClaimAquiredMsg struct {
	AllianceID    int32 `yaml:"allianceID"`
	CorpID        int32 `yaml:"corpID"`
	SolarSystemID int32 `yaml:"solarSystemID"`
}

type SovAllClaimLostMsg struct {
	AllianceID    int32 `yaml:"allianceID"`
	CorpID        int32 `yaml:"corpID"`
	SolarSystemID int32 `yaml:"solarSystemID"`
}

type SovCommandNodeEventStarted struct {
	CampaignEventType int32 `yaml:"campaignEventType"`
	ConstellationID   int32 `yaml:"constellationID"`
	SolarSystemID     int32 `yaml:"solarSystemID"`
}

type SovStationEnteredFreeport struct {
	Freeportexittime int64 `yaml:"freeportexittime"`
	SolarSystemID    int32 `yaml:"solarSystemID"`
	StructureTypeID  int32 `yaml:"structureTypeID"`
}

type SovStructureDestroyed struct {
	SolarSystemID   int32 `yaml:"solarSystemID"`
	StructureTypeID int32 `yaml:"structureTypeID"`
}

type SovStructureReinforced struct {
	CampaignEventType int32 `yaml:"campaignEventType"`
	DecloakTime       int64 `yaml:"decloakTime"`
	SolarSystemID     int32 `yaml:"solarSystemID"`
}

type SovStructureSelfDestructCancel struct {
	CharID          int32 `yaml:"charID"`
	SolarSystemID   int32 `yaml:"solarSystemID"`
	StructureTypeID int32 `yaml:"structureTypeID"`
}

type SovStructureSelfDestructFinished struct {
	SolarSystemID   int32 `yaml:"solarSystemID"`
	StructureTypeID int32 `yaml:"structureTypeID"`
}

type SovStructureSelfDestructRequested struct {
	CharID          int32  `yaml:"charID"`
	CorpName        string `yaml:"corpName"`
	DestructTime    int64  `yaml:"destructTime"`
	SolarSystemID   int32  `yaml:"solarSystemID"`
	StructureTypeID int32  `yaml:"structureTypeID"`
}

type SovereigntyIHDamageMsg struct {
	AggressorAllianceID int32   `yaml:"aggressorAllianceID"`
	AggressorCorpID     int32   `yaml:"aggressorCorpID"`
	AggressorID         int32   `yaml:"aggressorID"`
	ArmorValue          float64 `yaml:"armorValue"`
	HullValue           float64 `yaml:"hullValue"`
	ShieldValue         float64 `yaml:"shieldValue"`
	SolarSystemID       int32   `yaml:"solarSystemID"`
}

type SovereigntySBUDamageMsg struct {
	AggressorAllianceID int32   `yaml:"aggressorAllianceID"`
	AggressorCorpID     int32   `yaml:"aggressorCorpID"`
	AggressorID         int32   `yaml:"aggressorID"`
	ArmorValue          float64 `yaml:"armorValue"`
	HullValue           float64 `yaml:"hullValue"`
	ShieldValue         float64 `yaml:"shieldValue"`
	SolarSystemID       int32   `yaml:"solarSystemID"`
}

type SovereigntyTCUDamageMsg struct {
	AggressorAllianceID int32   `yaml:"aggressorAllianceID"`
	AggressorCorpID     int32   `yaml:"aggressorCorpID"`
	AggressorID         int32   `yaml:"aggressorID"`
	ArmorValue          float64 `yaml:"armorValue"`
	HullValue           float64 `yaml:"hullValue"`
	ShieldValue         float64 `yaml:"shieldValue"`
	SolarSystemID       int32   `yaml:"solarSystemID"`
}

type StationServiceDisabled struct {
	SolarSystemID   int32 `yaml:"solarSystemID"`
	StructureTypeID int32 `yaml:"structureTypeID"`
}

type StationServiceEnabled struct {
	SolarSystemID   int32 `yaml:"solarSystemID"`
	StructureTypeID int32 `yaml:"structureTypeID"`
}

type StructureAnchoring struct {
	OwnerCorpLinkData     []interface{} `yaml:"ownerCorpLinkData"`
	OwnerCorpName         string        `yaml:"ownerCorpName"`
	SolarsystemID         int32         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int32         `yaml:"structureTypeID"`
	TimeLeft              int64         `yaml:"timeLeft"`
	VulnerableTime        int64         `yaml:"vulnerableTime"`
}

type StructureDestroyed struct {
	OwnerCorpLinkData     []interface{} `yaml:"ownerCorpLinkData"`
	OwnerCorpName         string        `yaml:"ownerCorpName"`
	SolarsystemID         int32         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int32         `yaml:"structureTypeID"`
}

type StructureFuelAlert struct {
	ListOfTypesAndQty     [][]int32     `yaml:"listOfTypesAndQty"`
	SolarsystemID         int32         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int32         `yaml:"structureTypeID"`
}

type StructureItemsDelivered struct {
	CharID                int32         `yaml:"charID"`
	ListOfTypesAndQty     [][]int32     `yaml:"listOfTypesAndQty"`
	SolarsystemID         int32         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int32         `yaml:"structureTypeID"`
}

type StructureLostArmor struct {
	SolarsystemID         int32         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int32         `yaml:"structureTypeID"`
	TimeLeft              int64         `yaml:"timeLeft"`
	Timestamp             int64         `yaml:"timestamp"`
	VulnerableTime        int64         `yaml:"vulnerableTime"`
}

type StructureLostShields struct {
	SolarsystemID         int32         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int32         `yaml:"structureTypeID"`
	TimeLeft              int64         `yaml:"timeLeft"`
	Timestamp             int64         `yaml:"timestamp"`
	VulnerableTime        int64         `yaml:"vulnerableTime"`
}

type StructureOnline struct {
	SolarsystemID         int32         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int32         `yaml:"structureTypeID"`
}

type StructureServicesOffline struct {
	ListOfServiceModuleIDs []int32       `yaml:"listOfServiceModuleIDs"`
	SolarsystemID          int32         `yaml:"solarsystemID"`
	StructureID            int64         `yaml:"structureID"`
	StructureShowInfoData  []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID        int32         `yaml:"structureTypeID"`
}

type StructureUnanchoring struct {
	OwnerCorpLinkData     []interface{} `yaml:"ownerCorpLinkData"`
	OwnerCorpName         string        `yaml:"ownerCorpName"`
	SolarsystemID         int32         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int32         `yaml:"structureTypeID"`
	TimeLeft              int64         `yaml:"timeLeft"`
}

type StructureUnderAttack struct {
	AllianceID            int32         `yaml:"allianceID"`
	AllianceLinkData      []interface{} `yaml:"allianceLinkData"`
	AllianceName          string        `yaml:"allianceName"`
	ArmorPercentage       float64       `yaml:"armorPercentage"`
	CharID                int32         `yaml:"charID"`
	CorpLinkData          []interface{} `yaml:"corpLinkData"`
	CorpName              string        `yaml:"corpName"`
	HullPercentage        float64       `yaml:"hullPercentage"`
	ShieldPercentage      float64       `yaml:"shieldPercentage"`
	SolarsystemID         int32         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int32         `yaml:"structureTypeID"`
}

type StructureWentHighPower struct {
	SolarsystemID         int32         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int32         `yaml:"structureTypeID"`
}

type StructureWentLowPower struct {
	SolarsystemID         int32         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int32         `yaml:"structureTypeID"`
}

type StructuresReinforcementChanged struct {
	AllStructureInfo [][]interface{} `yaml:"allStructureInfo"`
	Hour             int32           `yaml:"hour"`
	NumStructures    int32           `yaml:"numStructures"`
	Timestamp        int64           `yaml:"timestamp"`
	Weekday          int32           `yaml:"weekday"`
}

type TowerAlertMsg struct {
	AggressorAllianceID int32   `yaml:"aggressorAllianceID"`
	AggressorCorpID     int32   `yaml:"aggressorCorpID"`
	AggressorID         int32   `yaml:"aggressorID"`
	ArmorValue          float64 `yaml:"armorValue"`
	HullValue           float64 `yaml:"hullValue"`
	MoonID              int32   `yaml:"moonID"`
	ShieldValue         float64 `yaml:"shieldValue"`
	SolarSystemID       int32   `yaml:"solarSystemID"`
	TypeID              int32   `yaml:"typeID"`
}

type TowerResourceAlertMsg struct {
	AllianceID    int32 `yaml:"allianceID"`
	CorpID        int32 `yaml:"corpID"`
	MoonID        int32 `yaml:"moonID"`
	SolarSystemID int32 `yaml:"solarSystemID"`
	TypeID        int32 `yaml:"typeID"`
	Wants         []struct {
		Quantity int32 `yaml:"quantity"`
		TypeID   int32 `yaml:"typeID"`
	} `yaml:"wants"`
}

type WarAllyOfferDeclinedMsg struct {
	AggressorID int32 `yaml:"aggressorID"`
	AllyID      int32 `yaml:"allyID"`
	CharID      int32 `yaml:"charID"`
	DefenderID  int32 `yaml:"defenderID"`
}

type WarSurrenderDeclinedMsg struct {
	IskValue float64 `yaml:"iskValue"`
	OwnerID  int32   `yaml:"ownerID"`
}

type WarSurrenderOfferMsg struct {
	IskValue         float64 `yaml:"iskValue"`
	OwnerID1         int32   `yaml:"ownerID1"`
	OwnerID2         int32   `yaml:"ownerID2"`
	WarNegotiationID int32   `yaml:"warNegotiationID"`
}

type NotificationTypeMoonminingExtractionStarted struct {
	AutoTime        int64             `yaml:"autoTime"`
	MoonID          int32             `yaml:"moonID"`
	MoonLink        string            `yaml:"moonLink"`
	OreVolumeByType map[int32]float64 `yaml:"oreVolumeByType"`
	ReadyTime       int64             `yaml:"readyTime"`
	SolarSystemID   int32             `yaml:"solarSystemID"`
	SolarSystemLink string            `yaml:"solarSystemLink"`
	StartedBy       int32             `yaml:"startedBy"`
	StartedByLink   string            `yaml:"startedByLink"`
	StructureID     int64             `yaml:"structureID"`
	StructureLink   string            `yaml:"structureLink"`
	StructureName   string            `yaml:"structureName"`
	StructureTypeID int32             `yaml:"structureTypeID"`
}
