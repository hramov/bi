package data_sourse_dto_in

type QueryOptions struct {
	Source string `json:"source"`
	Query  string `json:"query"`
	Params []any  `json:"params"`
}
