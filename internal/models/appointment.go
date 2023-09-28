package models

import (
	"time"

	"hospital-app/config"
)

type Appointment struct {
	BaseModel
	StaffID   string     `gorm:"not null;references:staffs(id)" json:"-"`
	PatientID string     `gorm:"not null;references:patients(id)" json:"patient_id"`
	DateTime  *time.Time `gorm:"default:null" json:"date_time"`
	Type      string     `gorm:"not null" json:"type"`
	Note      string     `gorm:"default:null" json:"note"`
}

// Delete function
func (m *Appointment) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

// Read function
func (m *Appointment) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ReadAll function
func (m *Appointment) ReadAll() (bool, []Appointment) {
	var Appointments []Appointment
	if result := config.DB.Find(&Appointments, &m); result.Error != nil {
		return false, Appointments
	}
	return true, Appointments
}
