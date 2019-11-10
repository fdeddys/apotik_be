package database

import (
	"fmt"
	"log"
	dbmodels "distribution-system-be/models/dbModels"
	dto "distribution-system-be/models/dto"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
)

// GetMenus ...
func GetMenus(param dto.FilterMenu, offset int, limit int) ([]dbmodels.Menu, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var menus []dbmodels.Menu
	var total int
	var err error

	if offset == 0 && limit == 0 {
		err = db.Find(&menus).Error
		if err != nil {
			return menus, 0, err
		}
		return menus, 0, nil
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go asyncQueryMenus(db, offset, limit, &menus, param, errQuery)
	go asyncQueryCountMenus(db, &total, param, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		log.Println("errr resErrQuery -->", resErrQuery)
		return menus, 0, resErrQuery
	}

	if resErrCount != nil {
		log.Println("errr resErrCount -->", resErrCount)
		return menus, 0, resErrCount
	}
	return menus, total, err
}

// asyncQueryCountMenus ...
func asyncQueryCountMenus(db *gorm.DB, total *int, param dto.FilterMenu, resChan chan error) {
	var criteriaMenuName = "%"
	if strings.TrimSpace(param.MenuName) != "" {
		criteriaMenuName = param.MenuName + criteriaMenuName
	}

	err := db.Model(&dbmodels.Menu{}).Where("name ilike ? AND status = 1", criteriaMenuName).Count(&*total).Error

	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// asyncQueryMenus ...
func asyncQueryMenus(db *gorm.DB, offset int, limit int, menus *[]dbmodels.Menu, param dto.FilterMenu, resChan chan error) {

	var criteriaMenuName = "%"
	if strings.TrimSpace(param.MenuName) != "" {
		criteriaMenuName = param.MenuName + criteriaMenuName
	}

	err := db.Order("name ASC").Offset(offset).Limit(limit).Find(&menus, "name ilike ? AND status = 1", criteriaMenuName).Error
	if err != nil {
		resChan <- err
	}
	resChan <- nil
}

// GetUserMenus ...
func GetUserMenus(user string) ([]dbmodels.Menu, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var menus []dbmodels.Menu
	var err error

	err = db.Raw(`
		select d.id, d.name, d.description, link, icon, parent_id, d.status
		from m_users a join
		m_role_user b on (a.id = b.user_id) join
		m_role_menu c on(b.role_id = c.role_id) join
		m_menus d on(c.menu_id = d.id)
		where a.user_name = ? and d.status = 1
		group by d.id, a.user_name
	`, user).Scan(&menus).Error

	fmt.Println("Menus => ", menus)

	if err != nil {
		return menus, err
	}
	return menus, nil
}
