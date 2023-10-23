package src

//type Column struct {
//	Name      string  `json:"name"`
//	Label     string  `json:"label"`
//	DataType  string  `json:"dataType"`
//	OrderBy   string  `json:"orderBy"`
//	Display   bool    `json:"display"`
//	Condition bool    `json:"condition"`
//	Convert   Convert `json:"convert"`
//}
//
//type Convert struct {
//	Key   string `json:"key"`
//	Value string `json:"value"`
//}

type Connect interface {
	Init(data string) (*MySQLClient, error)
	Destroy(ql string)
	Query(ql string) ([]map[string]interface{}, error)
}
