package postgresql

import (
	"fmt"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
)

func setDBCallback(db *gorm.DB) {
	// callback to capture query
	db.Callback().Create().After("gorm:create").Register("report_sentry:after_create", reportToSentry)
	db.Callback().Query().After("gorm:query").Register("report_sentry:after_query", reportToSentry)
	db.Callback().Delete().After("gorm:delete").Register("report_sentry:after_delete", reportToSentry)
	db.Callback().Update().After("gorm:update").Register("report_sentry:after_update", reportToSentry)
	db.Callback().RowQuery().After("gorm:row_query").Register("report_sentry:after_row_query", reportToSentry)
}

func reportToSentry(scope *gorm.Scope) {
	sqlValues := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(scope.Value)), ","), "[]")
	sentry.CaptureMessage(
		fmt.Sprintf("SQL: %s\nVALUES: %v", scope.SQL, sqlValues),
	)
}
