package services

import (
	"fiber-vue/app/database"
	"fiber-vue/app/model"
)

func GetMenuList() ([]model.Menu, error) {
	var menus []model.Menu

	if err := database.DB.Find(&menus).Error; err != nil {
		return nil, err
	}

	menuMap := make(map[uint][]model.Menu)

	for i := range menus {
		if menus[i].ParentId == 0 {
			menuMap[0] = append(menuMap[0], menus[i])
		} else {
			menuMap[uint(menus[i].ParentId)] = append(menuMap[uint(menus[i].ParentId)], menus[i])
		}
	}

	for i := range menus {
		if children, ok := menuMap[menus[i].ID]; ok {
			menus[i].Children = children
		}
	}

	return menus, nil
}

func CreateMenu(menu model.Menu) (model.Menu, error) {

	ok := database.DB.Create(&menu)

	if ok.Error != nil {
		return model.Menu{}, ok.Error
	}

	return menu, nil

}

func UpdateMenu(menuId int, menu model.Menu) (model.Menu, error) {

	ok := database.DB.Model(&model.Menu{}).Where("id = ?", menuId).Updates(menu)

	if ok.Error != nil {
		return model.Menu{}, ok.Error
	}

	return menu, nil

}

func RemoveMenu(menuId int) error {

	ok := database.DB.Where("id = ?", menuId).Delete(&model.Menu{})

	if ok.Error != nil {
		return ok.Error
	}

	return nil
}
