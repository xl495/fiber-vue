package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	// parent_id of the Menu
	// example: 0
	// required: false
	ParentId int `gorm:"size:255" json:"parentId,default:0"`

	// path of the Menu
	// example: /system/menu
	// required: true
	Path string `gorm:"size:255" json:"path"`

	// component of the Menu
	// example: system/menu/index
	// required: true
	Component string `gorm:"size:255" json:"component"`

	// name of the Menu
	// example: 菜单管理
	// required: true
	Name string `gorm:"size:255" json:"name"`

	// sort of the Menu
	// example: 1
	// required: false
	Sort int `gorm:"size:255" json:"sort"`

	// Mate of the Menu
	// example: {"title":"system","icon":"system"}
	// required: false
	Mate json.RawMessage `gorm:"type:longtext" json:"mate"`

	Children []Menu `gorm:"-" json:"children"`
}

type Mate struct {
	// title of the Mate
	// example: system
	// required: true
	Title string `gorm:"size:255" json:"title"`

	// icon of the Mate
	// example: system
	// required: false
	Icon string `gorm:"size:255" json:"icon"`

	//	keepAlive of the Mate
	//	example: false
	//	required: false
	KeepAlive bool `gorm:"size:255" json:"keepAlive"`

	// hidden of the Menu
	// example: false
	// required: false
	Hidden bool `gorm:"size:255" json:"hidden"`
}

func (m Mate) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *Mate) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	return json.Unmarshal(bytes, &m)
}
