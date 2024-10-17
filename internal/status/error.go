package status

import "errors"

var err error

func DbConnectionError() error {
	err = errors.New("error connecting to DB")
	return err
}

func PaginationDataValidationError() error {
	err = errors.New("incorrect pagination data")
	return err
}

func ExecutingQueryError(info string) error {
	err = errors.New("error executing query " + info)
	return err
}

func DataConversionError() error {
	err = errors.New("incorrect data conversion")
	return err
}
