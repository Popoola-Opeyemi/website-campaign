package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/Popoola-Opeyemi/sprintBackend/models"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"golang.org/x/crypto/bcrypt"
)

// ImageData ...
type ImageData struct {
	Data string `json:"data"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// HashPassword return a bcrypt hash of the input string
// and a non-nil error if an error occurs
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares the hash with a string
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// RequestIP ...
func RequestIP(req *http.Request) string {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return ""
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return ""
	}
	return userIP.String()
}

// Base64Encode ...
func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

// Base64Decode ...
func Base64Decode(src string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(src)
	return string(res), err
}

// Capitalize ...
func Capitalize(val *string) string {
	Capitalized := strings.Title(strings.ToLower(*val))
	return Capitalized
}

// ToString ...
func ToString(val *int) string {
	return strconv.Itoa(*val)
}

// ToNum ...
func ToNum(str string) (int, error) {
	return strconv.Atoi(str)
}

// SessionFromEchoContext ...
func SessionFromEchoContext(c echo.Context, name string) (*sessions.Session, error) {
	sess, err := session.Get(name, c)
	if err != nil {
		return nil, errors.New("session has expired")
	}
	return sess, err
}

// GetRoleNumFromEchoSession ...
func GetRoleNumFromEchoSession(c echo.Context, name string) (rn int64, err error) {

	sess, err := SessionFromEchoContext(c, name)
	if err != nil {
		return 0, err
	}
	val, ok := sess.Values["role"]
	if !ok {
		return 0, errors.New("session is invalid")
	}
	rn = val.(int64)
	return
}

// GenerateID function generates id of a student by appending to the current count in the database ...
func GenerateID(settings *models.Settings) (regID string) {
	count := settings.RegistrationCount

	if count <= 9 {
		regID = fmt.Sprintf("LS00%d", count)
	}

	if count >= 10 && count <= 99 {
		regID = fmt.Sprintf("LS0%d", count)
	}

	if count >= 100 {
		regID = fmt.Sprintf("LS%d", count)
	}

	settings.RegistrationCount++
	return
}

// SaveImageData save a data uri to a file and return a url where the file can be reached
// func SaveImageData(env *Enviroment, imgData json.RawMessage) (image *ImageData, err error) {
// 	saver := &AssetSave{}
// 	saver.Init(env)

// 	asset := &Asset{}
// 	image = &ImageData{}

// 	// convert images in json
// 	if len(imgData) > 0 {
// 		err = json.Unmarshal(imgData, image)
// 		if err != nil {
// 			return
// 		}

// 		if !strings.HasPrefix(image.Data, "data:") {
// 			// not a data uri
// 			return
// 		}

// 		// extract data uri and save to disk
// 		asset, err = saver.Save(&image.Data)
// 		if err != nil {
// 			return
// 		}

// 		image.Data = asset.FileURL
// 	}

// 	return
// }
