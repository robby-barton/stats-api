package espn

const teamInfoUrl = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/teams?limit=1000"

type TeamInfoESPN struct {
	Sports []Sport `json:"sports"`
}

type Sport struct {
	Id      int64    `json:"id,string"`
	Leagues []League `json:"leagues"`
	Name    string   `json:"name"`
	Slug    string   `json:"slug"`
}

type League struct {
	Abbreviation string     `json:"abbreviation"`
	Id           int64      `json:"id,string"`
	Name         string     `json:"name"`
	ShortName    string     `json:"shortName"`
	Slug         string     `json:"slug"`
	Teams        []TeamWrap `json:"teams"`
	Year         int64      `json:"year"`
}

type TeamWrap struct {
	Team TeamInfo `json:"team"`
}

type TeamInfo struct {
	Abbreviation     string `json:"abbreviation"`
	AltColor         string `json:"alternateColor"`
	Color            string `json:"color"`
	DisplayName      string `json:"displayName"`
	Id               int64  `json:"id,string"`
	IsActive         bool   `json:"isActive"`
	IsAllStar        bool   `json:"isAllStar"`
	Links            []Link `json:"links"`
	Location         string `json:"location"`
	Logos            []Logo `json:"logos"`
	Name             string `json:"name"`
	Nickname         string `json:"nickname"`
	ShortDisplayName string `json:"shortDisplayName"`
	Slug             string `json:"slug"`
}

type Link struct {
	Href       string   `json:"href"`
	IsExternal bool     `json:"isExternal"`
	IsHidden   bool     `json:"isHidden"`
	IsPremium  bool     `json:"isPremium"`
	Language   string   `json:"language"`
	Rel        []string `json:"rel"`
	ShortText  string   `json:"shortText"`
	Text       string   `json:"text"`
}

type Logo struct {
	Alt    string   `json:"alt"`
	Height int64    `json:"height"`
	Href   string   `json:"href"`
	Rel    []string `json:"rel"`
	Width  int64    `json:"width"`
}
