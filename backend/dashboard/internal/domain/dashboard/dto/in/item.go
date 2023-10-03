package dashboards_dto_in

type Item struct {
	DashId      string `json:"dash_id"`
	ItemType    int    `json:"item_type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DataQueries any    `json:"data_queries"`
	Options     any    `json:"options"`
}
