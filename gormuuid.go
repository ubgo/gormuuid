package gormuuid

import (
	"database/sql/driver"
	"strings"

	"github.com/google/uuid"
	"github.com/ubgo/gouuid"
)

type UUIDArray []uuid.UUID

func (uidarray UUIDArray) GormDataType() string {
	return "uuid[]"
}

func (a UUIDArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	params := "{" + strings.Join(gouuid.ToStringSlice(a), ",") + "}"
	return params, nil
}

// Do not delete - Why we did not use this
// GormValue does not accepts the pointer type value so in order to support i had to create two different types
// So i decied to used the Value() method above
// func (uidarray UUIDArray) GormValue(ctx context.Context, db *gorm.DB) (expr clause.Expr) {
// 	params := "{" + strings.Join(ArrayToString(uidarray), ",") + "}"
// 	return clause.Expr{
// 		SQL: "?::uuid[]",
// 		Vars: []interface{}{params},
// 	}
// }

// Scan implements the sql.Scanner interface
func (uidarray *UUIDArray) Scan(v interface{}) error {
	val := ""
	switch s := v.(type) {
	case []byte:
		val = string(s)
	}
	parsed := gouuid.PgStringArrayToUUIDSlide(val)
	*uidarray = parsed
	return nil
}
