package utils

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	// Layout standard date format layout
	Layout       = "2006-01-02 15:04:05"        // Date Format: YYYY-MM-DD HH:MM:SS
	LayoutWithMS = "2006-01-02 15:04:05.999999" // Date Format with microseconds: YYYY-MM-DD HH:MM:SS.MCS
)

// IsValidEmail checks if the provided email string is valid
func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// ParseAndConvertUUID parses a UUID from a string and returns uuid.NullUUID with validity.
func ParseAndConvertUUID(id string) (uuid.NullUUID, error) {
	if id != "" {
		parsedID, err := uuid.Parse(id)
		if err != nil {
			return uuid.NullUUID{}, fmt.Errorf("invalid UUID: %w", err)
		}
		return uuid.NullUUID{UUID: parsedID, Valid: true}, nil
	}
	return uuid.NullUUID{Valid: false}, nil
}

// ParseAndConvertUUIDPointer converts a *string to a uuid.NullUUID.
func ParseAndConvertUUIDPointer(id *string) (uuid.NullUUID, error) {
	if id != nil && *id != "" {
		parsedID, err := uuid.Parse(*id)
		if err != nil {
			return uuid.NullUUID{}, err
		}
		return uuid.NullUUID{UUID: parsedID, Valid: true}, nil
	}
	return uuid.NullUUID{Valid: false}, nil
}

// ConvertUUIDToNullUUID converts uuid.UUID to a uuid.NullUUID.
func ConvertUUIDToNullUUID(id uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{
		UUID:  id,
		Valid: true,
	}
}

// ResolveAndFallbackUUID resolves the provided UUID pointer or falls back to the existing valid UUID.
func ResolveAndFallbackUUID(providedID *string, currentID uuid.NullUUID) string {
	if providedID != nil && *providedID != "" {
		return *providedID
	} else if currentID.Valid {
		return currentID.UUID.String()
	}
	return ""
}

// UuidToString converts a uuid.NullUUID to a *string. It returns nil if the UUID is not valid.
func UuidToString(u uuid.NullUUID) *string {
	if u.Valid {
		str := u.UUID.String()
		return &str
	}
	return nil
}

// Int64ToNullInt64 converts a pointer to int64 into sql.NullInt64. If the pointer is nil, it returns an invalid NullInt64.
func Int64ToNullInt64(i *int64) sql.NullInt64 {
	if i != nil {
		return sql.NullInt64{Int64: *i, Valid: true}
	}
	return sql.NullInt64{Valid: false}
}

// NullInt64ToNullInt32 converts sql.NullInt64 to sql.NullInt32.
func NullInt64ToNullInt32(input sql.NullInt64) sql.NullInt32 {
	if input.Valid {
		return sql.NullInt32{Int32: int32(input.Int64), Valid: true}
	}
	return sql.NullInt32{Valid: false}
}

// NullInt32ToNullInt64 converts a sql.NullInt32 to sql.NullInt64.
func NullInt32ToNullInt64(input sql.NullInt32) sql.NullInt64 {
	if input.Valid {
		return sql.NullInt64{
			Int64: int64(input.Int32),
			Valid: true,
		}
	}
	return sql.NullInt64{Valid: false}
}

// NullInt32ToInt64Pointer converts a sql.NullInt32 to a pointer to int64, returning nil if the value is not valid.
func NullInt32ToInt64Pointer(input sql.NullInt32) *int64 {
	if input.Valid {
		int64Value := int64(input.Int32)
		return &int64Value
	}
	return nil
}

// NullStringToString converts a sql.NullString to a string, returning the value if valid, or an empty string if not.
func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// ToNullString converts a string to a sql.NullString.
func ToNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  value != "",
	}
}

// ToNullBool converts a boolean to a sql.NullBool.
func ToNullBool(value bool) sql.NullBool {
	return sql.NullBool{
		Bool:  value,
		Valid: true, // Always valid, as a bool can't be null in Go
	}
}

// ToFormattedDateString converts sql.NullTime to a formatted date string.
func ToFormattedDateString(nullTime sql.NullTime, layout string) string {
	if nullTime.Valid {
		return nullTime.Time.Format(layout)
	}
	return ""
}

// ParseWithFallbackTime parses a time string if provided, otherwise falls back to the existing value.
func ParseWithFallbackTime(newVal *string, existingVal sql.NullTime) (sql.NullTime, error) {
	if newVal != nil {
		if *newVal == "" {
			// If the input string is empty, return an invalid NullTime to set the date to NULL.
			return sql.NullTime{Valid: false}, nil
		}
		parsedDate, err := time.Parse(LayoutWithMS, *newVal)
		if err != nil {
			return sql.NullTime{}, err
		}
		return sql.NullTime{Time: parsedDate, Valid: true}, nil
	}
	return existingVal, nil
}

// FallbackNullString returns newVal if it's not empty; otherwise, it falls back to existingVal.
func FallbackNullString(newVal, existingVal string) string {
	if newVal != "" {
		return newVal
	}
	return existingVal
}

// FallbackNullBool returns newVal if provided; otherwise, it falls back to existingVal.
func FallbackNullBool(newVal, existingVal bool) bool {
	return newVal
}

// FallbackStringPointerToNullString takes a string pointer and falls back to the current sql.NullString if the pointer is nil or empty.
func FallbackStringPointerToNullString(newVal *string, existingVal sql.NullString) sql.NullString {
	if newVal != nil && *newVal != "" {
		return sql.NullString{
			String: *newVal,
			Valid:  true,
		}
	}
	return existingVal
}

// StringPointerToNullString converts a *string to sql.NullString.
func StringPointerToNullString(value *string) sql.NullString {
	if value != nil && *value != "" {
		return sql.NullString{String: *value, Valid: true}
	}
	return sql.NullString{Valid: false}
}

// StringToNullString converts a string to sql.NullString.
func StringToNullString(input string) sql.NullString {
	return sql.NullString{
		String: input,
		Valid:  input != "",
	}
}

func StringToInt32(s string) int32 {
	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	} else {
		return int32(i64)
	}
}

func StringToBool(s string) bool {
	value, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return value
}

func ValidateFileSize(s int64) bool {
	byteToMB := math.Round(float64(s) / (1 << 20))
	if byteToMB > 10 {
		return false
	}
	return true
}

func ValidateFileType(imageFile string) bool {
	imageByte, err := ReadImageFile(imageFile)
	if err != nil {
		return false
	}
	fType := http.DetectContentType(imageByte)
	if fType == "image/jpeg" || fType == "image/png" {
		return true
	}
	return false
}

// FallbackInt64PointerToNullInt64 takes a double pointer to int64 and falls back to the current sql.NullInt64 if the pointer is nil.
func FallbackInt64PointerToNullInt64(newVal **int64, existingVal sql.NullInt64) sql.NullInt64 {
	if newVal != nil && *newVal != nil {
		return sql.NullInt64{
			Int64: **newVal,
			Valid: true,
		}
	}
	if newVal == nil {
		return sql.NullInt64{Valid: false}
	}
	return existingVal
}

// FallbackStringPointerToString returns newVal if it's not nil; otherwise, it falls back to existingVal.
func FallbackStringPointerToString(newVal *string, existingVal string) string {
	if newVal != nil && *newVal != "" {
		return *newVal
	}
	return existingVal
}

// FallbackBoolPointerToBool takes a bool pointer and falls back to the current bool if the pointer is nil.
func FallbackBoolPointerToBool(newVal *bool, existingVal bool) bool {
	if newVal != nil {
		return *newVal
	}
	return existingVal
}

// FallbackInt32Pointer returns the dereferenced value of the pointer if it is not nil; otherwise, it returns the fallback value.
func FallbackInt32Pointer(ptr *int32, existingVal int32) int32 {
	if ptr != nil {
		return *ptr
	}
	return existingVal
}

func FallbackInt32PointerToSqlNullInt32(ptr *int32, existingVal sql.NullInt32) sql.NullInt32 {
	if ptr != nil {
		return sql.NullInt32{
			Int32: *ptr,
			Valid: true,
		}
	}
	return existingVal
}

// NullableString converts a sql.NullString to a *string.
func NullableString(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

// NullableInt32 converts a sql.NullInt32 to a *int32.
func NullableInt32(ni sql.NullInt32) *int32 {
	if ni.Valid {
		return &ni.Int32
	}
	return nil
}

// NullableBool converts a sql.NullBool to a *bool.
func NullableBool(nb sql.NullBool) *bool {
	if nb.Valid {
		return &nb.Bool
	}
	return nil
}

// Int32ToNullInt32 converts an *int32 to sql.NullInt32.
func Int32ToNullInt32(i *int32) sql.NullInt32 {
	if i != nil {
		return sql.NullInt32{Int32: *i, Valid: true}
	}
	return sql.NullInt32{Valid: false}
}

// BoolToNullBool converts a *bool to sql.NullBool.
func BoolToNullBool(b *bool) sql.NullBool {
	if b != nil {
		return sql.NullBool{Bool: *b, Valid: true}
	}
	return sql.NullBool{Valid: false}
}

// GenerateRandomString generates a random string of the specified length.
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// ReadCSVFile reads a CSV file and returns a slice of slices of strings.
func ReadCSVFile(filePath string) ([][]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// ConvertValueToPointer returns a pointer to the given value.
func ConvertValueToPointer[T any](v T) *T {
	return &v
}

func IsPresignedURL(url string) bool {
	return strings.Contains(url, "X-Amz-Date") && strings.Contains(url, "X-Amz-Expires")
}

// ParseAndConvertNullUUID parses a null UUID from a string and returns uuid.NullUUID with validity.
func ParseAndConvertNullUUID(id string) (uuid.NullUUID, error) {
	if id != "" {
		parsedID, err := uuid.Parse(id)
		if err != nil {
			return uuid.NullUUID{}, fmt.Errorf("invalid UUID: %w", err)
		}
		return uuid.NullUUID{UUID: parsedID, Valid: true}, nil
	}
	return uuid.NullUUID{Valid: false}, nil
}

// ConvertStringToNullUUID converts UUID as a string into a uuid.NullUUID.
func ConvertStringToNullUUID(id string) uuid.NullUUID {
	if len(id) == 0 {
		return uuid.NullUUID{}
	}

	uuidIn, err := ParseAndConvertNullUUID(id)
	if err != nil {
		fmt.Printf("ERROR: Unable to convert, %s, %v", uuidIn, err)
	}

	return uuidIn
}

func ReadImageFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	fileBytes, err := io.ReadAll(f)
	if err != nil {
		return nil, nil
	}

	return fileBytes, err
}

func RemoveS3EnvKey(key string) string {
	if i := strings.IndexByte(key, '/'); i != -1 {
		key = key[i+1:]
	}
	return key
}
