package config

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/denisbrodbeck/machineid"
	"gorm.io/gorm"
)

const (
	UpdatePath = "./updates/"
	JwtSecret  = "secret"
	AppKey     = "hospital"
	ServerPort = "8788"
)

var (
	DB             *gorm.DB
	HashedID, err  = machineid.ProtectedID(AppKey)
	ActivationCode = ""
)

// initial db settings
var (
	InitSettings = map[string]string{
		"installation_code": HashedID,
		"activation_code":   "",
		"last_card_no":      "1",
		"hospital_name":     "",
		"hospital_address":  "",
		"hospital_phone":    "",
		"hospital_email":    "",
		"hospital_logo":     "",
	}
)

const (
	ADMIN_ROLE = iota + 1
	DOCTOR_ROLE
	NURSE_ROLE
)

type BodyStructure map[string]string

var AppDataFolder = fmt.Sprintf("%v/%v", os.Getenv("APPDATA"), "HCMS")
var DatabaseFile = fmt.Sprintf("%v/%v", AppDataFolder, "hospital.appdb")
var FilesFolder = fmt.Sprintf("%v/%v", AppDataFolder, "files")

var SERVER_HANDLE = &http.Server{
	Addr: net.JoinHostPort("0.0.0.0", ServerPort),
}
