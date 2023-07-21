package model

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

//MYTUUID -> new datatype
type MYTUUID uuid.UUID

// StringToMYTYPE -> parse string to MYTUUID
func StringToMYTYPE(s string) (MYTUUID, error) {
	id, err := uuid.Parse(s)
	return MYTUUID(id), err
}

//String -> String Representation of Binary16
func (my MYTUUID) String() string {
	return uuid.UUID(my).String()
}

//GormDataType -> sets type to binary(16)
func (my MYTUUID) GormDataType() string {
	return "binary(16)"
}

func (my MYTUUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(my)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (my *MYTUUID) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*my = MYTUUID(s)
	return err
}

// Scan --> tells GORM how to receive from the database
func (my *MYTUUID) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = MYTUUID(parseByte)
	return err
}

// Value -> tells GORM how to save into the database
func (my MYTUUID) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}
