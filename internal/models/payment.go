package models

import (
	"hospital-app/config"

	"gorm.io/gorm"
)

type Payment struct {
	BaseModel
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	StaffID          string         `gorm:"not null;references:staffs(id)" json:"-"`
	InvoiceID        string         `gorm:"not null;references:invoices(id)" json:"invoice_id"`
	Note             string         `gorm:"default:null" json:"note"`
	AmountPaid       uint           `gorm:"default:0" json:"amount_paid"`
	Balance          uint           `gorm:"default:0" json:"balance"`
	PaymentReference string         `gorm:"default:null" json:"payment_reference"`
	Status           string         `gorm:"not null;default:'paid'" json:"status"`
	PaymentProof     string         `gorm:"default:null" json:"payment_proof"`
}

func (m *Payment) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

func (m *Payment) UpdateSingle(key string, value string) error {
	return config.DB.First(&m).Update(key, value).Error
}

func (m *Payment) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

func (m *Payment) Read() error {
	return config.DB.First(&m, &m).Error
}

func (m *Payment) ReadAll() (bool, []Payment) {
	var payments []Payment
	if result := config.DB.Find(&payments, &m); result.Error != nil {
		return false, payments
	}
	return true, payments
}
