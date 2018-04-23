package auorm

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// Raw sql查询返回[]map[string]string类型
func Raw(db *gorm.DB, sqlQuery string, sqlValues ...interface{}) (result []map[string]string, err error) {
	rows, err := db.Raw(sqlQuery, sqlValues...).Rows()
	if err != nil {
		return
	}
	cols, err := rows.Columns()
	if err != nil {
		return
	}
	values := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}
		var value string
		resultC := map[string]string{}
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			resultC[cols[i]] = value
		}
		result = append(result, resultC)
	}
	return
}
