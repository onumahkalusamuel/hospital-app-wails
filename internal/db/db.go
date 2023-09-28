package db

import (
	"log"

	"hospital-app/config"
	"hospital-app/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {

	var err error

	config.DB, err = gorm.Open(sqlite.Open(config.DatabaseFile), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	config.DB.AutoMigrate(
		&models.Delivery{},
		&models.Invoice{},
		&models.Patient{},
		&models.PatientHistory{},
		&models.Payment{},
		&models.Settings{},
		&models.Staff{},
	)
}

func InitialData() {
	for setting, value := range config.InitSettings {
		s := &models.Settings{Setting: setting}
		s.Read()
		if s.ID == "" {
			s = &models.Settings{Setting: setting, Value: value}
			s.Create()
		}
	}

	// add guest customer
	guest := &models.Patient{CardNo: "GUEST"}
	guest.Read()
	if guest.ID == "" {
		guest.Title = "Mr"
		guest.Lastname = "Guest"
		guest.Firstname = "Patient"
		guest.Address = "Address"
		guest.Sex = "Male"
		guest.Email = "guest@hospital.com"
		config.DB.Model(&models.Patient{}).Save(guest)
	}

}
