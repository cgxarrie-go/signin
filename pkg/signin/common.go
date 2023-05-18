package signin

type Item struct {
	ID     int       `json:"id"`
	SiteID int       `json:"site_id"`
	Space  ItemSpace `json:"space"`
}

type ItemSpace struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	Capacity    int        `json:"capacity"`
	Zones       []ItemZone `json:"zones"`
}

type ItemZone struct {
	ID       string `json:"id"`
	Name     string `json: "name"`
	Type     string `json:"type"`
	Capacity int    `json:"capacity"`
}
