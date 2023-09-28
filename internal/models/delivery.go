package models

import (
	"time"

	"hospital-app/config"
	"hospital-app/pkg"
)

type Delivery struct {
	BaseModel
	StaffID          string     `gorm:"not null;references:staffs(id)" json:"-"`
	PatientID        string     `gorm:"not null;references:patients(id)" json:"patient_id"`
	DeliveryDateTime *time.Time `gorm:"default:null" json:"delivery_date_time"`
	DeliveryMode     string     `gorm:"default:'vaginal'" json:"delivery_mode"`
	BabySex          string     `gorm:"not null" json:"baby_sex"`
	BabyWeight       string     `gorm:"default:'0'" json:"baby_weight"`
	Condition        string     `gorm:"default:null" json:"condition"`
	Note             string     `gorm:"default:null" json:"note"`
	Patient          *Patient   `json:"patient"`
}

// Delete function
func (m *Delivery) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

// Read function
func (m *Delivery) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ReadAll function
func (m *Delivery) ReadAll() (bool, []Delivery) {
	var Deliverys []Delivery
	if result := config.DB.Find(&Deliverys, &m); result.Error != nil {
		return false, Deliverys
	}
	return true, Deliverys
}

func (cg *Delivery) List(pagination pkg.Pagination, preloads []string, d *Delivery) (*pkg.Pagination, error) {
	var deliveries []*Delivery

	db := config.DB.Scopes(Paginate(deliveries, &pagination, config.DB))
	if len(preloads) > 0 {
		for _, v := range preloads {
			db.Preload(v)
		}
	}

	db.Find(&deliveries, d)
	pagination.Rows = deliveries

	return &pagination, nil
}
