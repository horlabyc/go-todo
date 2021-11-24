package todo

import "github.com/jinzhu/gorm"

const (
	PENDING    = "pending"
	IN_PROGRES = "inProgress"
	DONE       = "done"
)

type Todo struct {
	gorm.Model
	Name        string `gorm:"Not Null" json:"name"`
	Description string `json:"description"`
	Status      string `gorm:"Not Null" json:"status"`
}
