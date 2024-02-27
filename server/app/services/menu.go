package services

import (
	"fiber-vue/app/database"
	"fiber-vue/app/model"
)

func GetMenuList() ([]model.Menu, error) {

	var menus []model.Menu

	ok := database.DB.Find(&menus)

	if ok.Error != nil {
		return nil, ok.Error
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
