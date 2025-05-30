package helpers

import (
	"bufio"
	"bytes"
	"database/sql/driver"
	"io"
	"math"
	"math/rand"
	"mime"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/ahargunyllib/thera-be/pkg/log"
)

func Contains(search string, words []string) bool {
	for _, word := range words {
		if search == word {
			return true
		}
	}

	return false
}

func ReadFile(filepath string, separator string) ([]string, error) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[HELPERS][ReadFiles] failed to open file")
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	results := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		results = append(results, line)
	}

	if err := scanner.Err(); err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[HELPERS][ReadFiles] failed to scan file")
		return nil, err
	}

	return results, nil
}

func CheckRowsAffected(rows int64, err error) error {
	if rows == 0 {
		return err
	}

	return nil
}

func GenerateRandomString(lenght int) string {
	alphaNumRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	randomRune := make([]rune, lenght)

	for i := 0; i < lenght; i++ {
		randomRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}

	return string(randomRune)
}

// Helper function to convert struct fields into a slice of interface{}
func StructToSlice(i interface{}) []interface{} {
	var result []interface{}
	val := reflect.ValueOf(i)

	for i := 0; i < val.NumField(); i++ {
		// if the field is of type entity, skip it
		if strings.Contains(val.Field(i).Type().String(), "entity") {
			continue
		}

		result = append(result, val.Field(i).Interface())
	}
	return result
}

// Helper function to convert slice of interface{} into a slice of driver.Value
func ConvertToDriverValue(values []interface{}) []driver.Value {
	driverValues := make([]driver.Value, len(values))
	for i, v := range values {
		driverValues[i] = driver.Value(v)
	}
	return driverValues
}

// Haversine calculates the distance between two latitude/longitude points in meters.
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Radius of Earth in meters
	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	deltaLat := (lat2 - lat1) * math.Pi / 180
	deltaLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func ToInt64Slice[T ~int64](values []T) []int64 {
	result := make([]int64, len(values))
	for i, v := range values {
		result[i] = int64(v)
	}
	return result
}

func FileToMultipartHeader(fieldName, filePath string) (*multipart.FileHeader, error) {
	// Open the local file
	file, err := os.Open(filePath)
	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[HELPERS][FileToMultipartHeader] failed to open file")
		return nil, err
	}
	defer file.Close()

	// Get the file extension and detect MIME type
	ext := strings.ToLower(filepath.Ext(filePath)) // e.g. .png
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream" // fallback
	}

	// Create a buffer and multipart writer
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Manually set headers for the file part
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="`+fieldName+`"; filename="`+filepath.Base(filePath)+`"`)
	h.Set("Content-Type", mimeType)

	// Create part with custom header
	part, err := writer.CreatePart(h)
	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[HELPERS][FileToMultipartHeader] failed to create form part with header")
		return nil, err
	}

	// Copy file contents into part
	if _, err := io.Copy(part, file); err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[HELPERS][FileToMultipartHeader] failed to copy file contents")
		return nil, err
	}
	writer.Close()

	// Parse back to get FileHeader
	reader := multipart.NewReader(&b, writer.Boundary())
	form, err := reader.ReadForm(10 << 20) // 10MB
	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[HELPERS][FileToMultipartHeader] failed to read form")
		return nil, err
	}

	files := form.File[fieldName]
	if len(files) == 0 {
		log.Error(log.CustomLogInfo{
			"error": "no files found in form",
		}, "[HELPERS][FileToMultipartHeader] no files found in form")
		return nil, os.ErrNotExist
	}
	return files[0], nil
}

func GetQRCodeFileNameFromURL(URL string) string {
	lastSlashIndex := len(URL) - 1
	for i := len(URL) - 1; i >= 0; i-- {
		if URL[i] == '/' {
			lastSlashIndex = i
			break
		}
	}

	return URL[lastSlashIndex+1:]
}
