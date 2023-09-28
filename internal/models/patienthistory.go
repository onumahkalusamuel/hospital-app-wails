package models

import (
	"time"

	"hospital-app/config"
	"hospital-app/pkg"
)

type PatientHistory struct {
	BaseModel
	StaffID   string         `gorm:"not null;references:staffs(id)" json:"-"`
	PatientID string         `gorm:"not null;references:patients(id)" json:"patient_id"`
	Subject   string         `gorm:"default:null" json:"subject"`
	Date      *time.Time     `gorm:"default:null" json:"date_time"`
	Details   HistoryDetails `gorm:"serializer:json" json:"details"`
	Patient   *Patient       `json:"patient"`
}

type HistoryDetails struct {
	General     DetailsInner     `json:"general"`
	Appointment AppointmentInner `json:"appointment"`
	Admission   AdmissionInner   `json:"admission"`
	Discharge   DetailsInner     `json:"discharge"`
	Diagnosis   DetailsInner     `json:"diagnosis"`
	Examination DetailsInner     `json:"examination"`
	TestResult  DetailsInner     `json:"testresult"`
	Treatment   DetailsInner     `json:"treatment"`
}

type DetailsInner struct {
	Note     string `json:"note"`
	Document string `json:"document"`
}

type AdmissionInner struct {
	DetailsInner
	WardNumber string `json:"ward_number"`
	RoomNumber string `json:"room_number"`
}

type AppointmentInner struct {
	DetailsInner
	AppointmentType string `json:"appointment_type"`
}

func (m *PatientHistory) Create() error {
	return config.DB.Create(&m).Error
}

func (m *PatientHistory) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

func (m *PatientHistory) UpdateSingle(key string, value any) error {
	return config.DB.First(&m).Update(key, value).Error
}

func (m *PatientHistory) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

func (m *PatientHistory) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *PatientHistory) ReadAll() (bool, []PatientHistory) {
	var PatientHistorys []PatientHistory
	if result := config.DB.Find(&PatientHistorys, &m); result.Error != nil {
		return false, PatientHistorys
	}
	return true, PatientHistorys
}

func (m *PatientHistory) List(pagination pkg.Pagination, ph *PatientHistory) (*pkg.Pagination, error) {
	var records []*PatientHistory
	config.DB.Scopes(Paginate(ph, &pagination, config.DB)).Find(&records, ph)
	pagination.Rows = records
	return &pagination, nil
}
