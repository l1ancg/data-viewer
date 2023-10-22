package pgk

type ConnectProps struct {
	Label string `json:"label"`
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Index int    `json:"index"`
}

type Connect interface {
	Query(ql string) ([]map[string]interface{}, error)
}
