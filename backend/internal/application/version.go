package application

type Version struct {
	Id      int    `json:"id"  gorm:"primarykey"`
	Version string `json:"version"`
}

// todo: version management with table 'version'
