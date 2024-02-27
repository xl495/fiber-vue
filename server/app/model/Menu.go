package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model `json:"gorm.Model"`
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

	// meta of the Menu
	// example: {"title":"system","icon":"system"}
	// required: false
	Meta Meta `gorm:"size:255" json:"meta"`
}

type Meta struct {
	// title of the Meta
	// example: system
	// required: true
	Title string `gorm:"size:255" json:"title"`

	// icon of the Meta
	// example: system
	// required: false
	Icon string `gorm:"size:255" json:"icon"`

	//	keepAlive of the Meta
	//	example: false
	//	required: false
	KeepAlive bool `gorm:"size:255" json:"keepAlive"`

	// hidden of the Menu
	// example: false
	// required: false
	Hidden bool `gorm:"size:255" json:"hidden"`
}
