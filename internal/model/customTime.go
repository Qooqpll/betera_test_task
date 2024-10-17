package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type CustomTime time.Time

// MarshalJSON implements the json.Marshaler interface.
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(ct).Format("2006-01-02"))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// Удаляем кавычки
	str := string(b[1 : len(b)-1])
	// Пытаемся распарсить дату в формате YYYY-MM-DD
	if len(str) == 10 { // YYYY-MM-DD
		parsedTime, err := time.Parse("2006-01-02", str)
		if err != nil {
			return fmt.Errorf("failed to parse date: %w", err)
		}
		*ct = CustomTime(parsedTime)
		return nil
	}
	// Если формат не совпадает, возвращаем ошибку
	return fmt.Errorf("unsupported date format: %s", str)
}

// Value implements the driver.Valuer interface.
func (ct CustomTime) Value() (driver.Value, error) {
	return time.Time(ct).Format("2006-01-02"), nil
}

// Scan implements the sql.Scanner interface.
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		*ct = CustomTime(time.Time{})
		return nil
	}
	t, err := time.Parse("2006-01-02", value.(string))
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}
