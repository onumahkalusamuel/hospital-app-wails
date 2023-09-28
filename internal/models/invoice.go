package models

import (
	"time"

	"hospital-app/config"
	"hospital-app/pkg"

	"gorm.io/gorm"
)

type Invoice struct {
	BaseModel
	DeletedAt      gorm.DeletedAt   `gorm:"index" json:"-"`
	StaffID        string           `gorm:"not null;references:staffs(id)" json:"-"`
	PatientID      string           `gorm:"not null;references:patients(id)" json:"patient_id"`
	Name           string           `gorm:"default:'Guest'" json:"name"`
	BillingAddress string           `gorm:"default:null" json:"billing_address"`
	InvoiceNumber  string           `gorm:"not null" json:"invoice_number"`
	InvoiceDate    *time.Time       `gorm:"not null" json:"invoice_date"`
	DueDate        *time.Time       `gorm:"not null" json:"due_date"`
	Details        []InvoiceDetails `gorm:"types:bytes;serializer:gob" json:"details"`
	Amount         uint             `gorm:"not null" json:"amount"`
	Subtotal       uint             `gorm:"not null" json:"sub_total"`
	Balance        uint             `gorm:"default:0" json:"balance"`
	Discount       uint             `gorm:"default:0" json:"discount"`
	Completed      uint             `gorm:"default:0" json:"completed"`
	Payments       []*Payment       `gorm:"constraint:onDelete:CASCADE" json:"payments"`
	Patient        *Patient         `json:"patient"`
}

type InvoiceDetails struct {
	Description string `json:"description"`
	Qty         uint   `json:"qty"`
	Price       uint   `json:"price"`
	Amount      uint   `json:"amount"`
}

func (m *Invoice) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

func (m *Invoice) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

func (m *Invoice) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Invoice) ReadAll() (bool, []Invoice) {
	var Invoices []Invoice
	if result := config.DB.Find(&Invoices, &m); result.Error != nil {
		return false, Invoices
	}
	return true, Invoices
}

func (cg *Invoice) List(pagination pkg.Pagination, i *Invoice) (*pkg.Pagination, error) {
	var invoices []*Invoice

	db := config.DB.Scopes(Paginate(invoices, &pagination, config.DB))
	db.Preload("Patient").Order("balance DESC, created_at ASC").Find(&invoices, i)
	pagination.Rows = invoices
	return &pagination, nil
}
