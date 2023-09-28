package models

import (
	"fmt"
	"strconv"
	"strings"

	"hospital-app/config"
	"hospital-app/pkg"

	"gorm.io/gorm"
)

type Patient struct {
	BaseModel
	DeletedAt          gorm.DeletedAt    `gorm:"index" json:"-"`
	StaffID            string            `gorm:"not null;references:staffs(id)" json:"-"`
	CardNo             string            `gorm:"default:null;unique" json:"card_no"`
	Title              string            `gorm:"default:null" json:"title"`
	Firstname          string            `gorm:"not null" json:"firstname"`
	Middlename         string            `gorm:"default:null" json:"middlename"`
	Lastname           string            `gorm:"not null" json:"lastname"`
	Phone              string            `gorm:"default:null" json:"phone"`
	Email              string            `gorm:"default:null" json:"email"`
	Sex                string            `gorm:"default:null" json:"sex"`
	MaritalStatus      string            `gorm:"default:'Single'" json:"marital_status"`
	Age                string            `gorm:"default:null" json:"age"`
	Address            string            `gorm:"default:null" json:"address"`
	Occupation         string            `gorm:"default:null" json:"occupation"`
	Anc                string            `gorm:"default:null" json:"anc"`
	CurrentAppointment string            `gorm:"default:null" json:"current_appointment"`
	CurrentStatus      string            `gorm:"default:null" json:"current_status"`
	Invoices           []*Invoice        `gorm:"constraint:onDelete:CASCADE" json:"invoices"`
	PatientHistory     []*PatientHistory `gorm:"constraint:onDelete:CASCADE" json:"patient_history"`
}

func (m *Patient) Create() error {
	return config.DB.Create(&m).Error
}

func (m *Patient) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

func (m *Patient) UpdateSingle(key string, value any) error {
	return config.DB.First(&m).Update(key, value).Error
}

func (m *Patient) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

func (m *Patient) Read() error {
	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Patient) ReadAll() (bool, []Patient) {
	var Patients []Patient
	if result := config.DB.Find(&Patients, &m); result.Error != nil {
		return false, Patients
	}
	return true, Patients
}

func (u *Patient) AfterCreate(tx *gorm.DB) (err error) {

	if u.CardNo != "" {
		return nil
	}

	s := Settings{Setting: "last_card_no"}
	s.Read()
	lastNo, _ := strconv.Atoi(s.Value)
	available := false
	nextCardNo := ""

	for available == false {
		nextCardNo = strings.TrimLeft(fmt.Sprintf("H%04v", lastNo+1), " ")

		// check existing
		var check []Patient
		tx.Model(u).Find(&check, "card_no='"+nextCardNo+"'")
		if len(check) > 0 {
			lastNo++
			continue
		}

		available = true
		tx.Model(u).Update("card_no", nextCardNo)
		tx.Model(s).Update("value", lastNo+1)
	}

	return
}

func (cg *Patient) List(pagination pkg.Pagination, preloads []string) (*pkg.Pagination, error) {
	var patients []*Patient

	db := config.DB.Scopes(Paginate(patients, &pagination, config.DB))
	if len(preloads) > 0 {
		for _, v := range preloads {
			db.Preload(v)
		}
	}

	if pagination.Query != "" {
		db.Where(
			"lastname LIKE '" + pagination.Query + "%' OR " +
				"firstname LIKE '" + pagination.Query + "%' OR " +
				"middlename LIKE '" + pagination.Query + "%' OR " +
				"card_no LIKE '%" + pagination.Query + "%' OR " +
				"phone LIKE '" + pagination.Query + "%' OR " +
				"email LIKE '" + pagination.Query + "%' OR " +
				"occupation LIKE '" + pagination.Query + "%'")
	}

	db.Find(&patients)

	pagination.Rows = patients

	return &pagination, nil
}
