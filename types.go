package lobby

const GameId = "DontStarveTogether"

// ExplicitPlatforms means that platforms could be used as query params in klei api.
var ExplicitPlatforms = []string{Steam.String(), PSN.String(), Rail.String(), XBOne.String(), Switch.String()}

const (
	Steam Platform = 1
	PSN   Platform = 2
	// Rail is alias of WeGame, only serve at ap-east-1
	Rail  Platform = 4
	XBOne Platform = 16
	// PS4Official can not be use in api query params
	PS4Official Platform = 19
	Switch      Platform = 32
)

// Platform represents dst server platform, it may be updated by klei in the future
type Platform uint

func (p Platform) String() string {
	switch p {
	case 1:
		return "Steam"
	case 2:
		return "PSN"
	case 4:
		return "Rail"
	case 16:
		return "XBOne"
	case 19:
		return "PS4Official"
	case 32:
		return "Switch"
	}
	return "unknown platform"
}

// Region represents dst lobby server region, it may be updated by klei in the future
const (
	UsEast1     = "us-east-1"
	EuCentral   = "eu-central-1"
	ApSoutheast = "ap-southeast-1"
	ApEast      = "ap-east-1"
)

// Lobby server urls, may be updated by klei in the future
const (
	LobbyRegionURL  = `https://lobby-v2-cdn.klei.com/regioncapabilities-v2.json`
	LobbyServersURL = `https://lobby-v2-cdn.klei.com/{{.region}}-{{.platform}}.json.gz`
	LobbyDetailsURL = "https://lobby-v2-{{.region}}.klei.com/lobby/read"
)

type Regions struct {
	Regions []struct {
		Region string `json:"Region"`
	} `json:"LobbyRegions"`
}

// Server includes all the information about single dst server
type Server struct {
	// network options
	Guid  string `json:"guid"`
	RowId string `json:"__rowId"`
	// only at steam platform
	SteamId   string   `json:"steamid"`
	SteamRoom string   `json:"steamroom"`
	Session   string   `json:"session"`
	Address   string   `json:"__addr"`
	Port      int      `json:"port"`
	Host      string   `json:"host"`
	Platform  Platform `json:"platform"`

	ClanOnly bool `json:"clanonly"`
	LanOnly  bool `json:"lanonly"`

	// second shard
	Secondaries map[string]Secondaries

	// game options
	Name     string `json:"name"`
	GameMode string `json:"mode"`
	Intent   string `json:"intent"`
	Season   string `json:"season"`
	TagStr   string `json:"tags"`
	Tags     []string
	Version  int `json:"v"`
	// max players allowed
	MaxConnections int `json:"maxconnections"`
	// online players number
	Connected int `json:"connected"`

	Mod             bool `json:"mods"`
	Pvp             bool `json:"pvp"`
	HasPassword     bool `json:"password"`
	IsDedicated     bool `json:"dedicated"`
	ClientHosted    bool `json:"clienthosted"`
	AllowNewPlayers bool `json:"allownewplayers"`
	ServerPaused    bool `json:"serverpaused"`
	FriendOnly      bool `json:"fo"`
}

// Secondaries represents the secondaries shard among dst servers
type Secondaries struct {
	Id      string `json:"id"`
	SteamId string `json:"steamid"`
	Address string `json:"__addr"`
	Port    int    `json:"port"`
}

type Servers struct {
	List []Server `json:"GET"`
}

type Player struct {
	Name    string
	Prefab  string
	SteamId string
	// hex color code
	Colour string
	// shard level
	Level int
}

type Mod struct {
	Id       string
	Name     string
	Version1 string
	Version2 string
	Enabled  bool
}

// ServerDetails includes some details information
type ServerDetails struct {
	// repeat options
	Server

	Tick          int  `json:"tick"`
	ClientModsOff bool `json:"clientmodsoff"`
	Nat           int  `json:"nat"`

	// raw lua script data
	Data          string `json:"data"`
	WorldGen      string `json:"worldgen"`
	OnlinePlayers string `json:"players"`
	Mods          []any  `json:"mods_info"`

	// parsed lua data
	Details struct {
		Day                int
		DayElapsedInSeason int
		DaysLeftInSeason   int
		Players            []Player
		ModsInfo           []Mod
	}
}
