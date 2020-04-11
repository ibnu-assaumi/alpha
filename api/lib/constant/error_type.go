package constant

const (
	// ErrorTypeBindParam : error type when bind parameter
	ErrorTypeBindParam string = "bind_param"
	// ErrorTypeValidateParam : error type when validate parameter
	ErrorTypeValidateParam string = "validate_param"
	// ErrorTypeNotExists : error type when record does not exists
	ErrorTypeNotExists string = "record_not_exists"
	// ErrorTypeExists : error type when record does exists
	ErrorTypeExists string = "record_already_exists"
	// ErrorTypeSQLQuery : error type when querying record
	ErrorTypeSQLQuery string = "sql_query"
	// ErrorTypeSQLInsert : error sql type when inserting record
	ErrorTypeSQLInsert string = "sql_insert"
	// ErrorTypeSQLInsertHistory : error sql type when inserting history record
	ErrorTypeSQLInsertHistory string = "sql_insert_history"
	// ErrorTypeSQLUpdate : error sql type when updating record
	ErrorTypeSQLUpdate string = "sql_update"
	// ErrorTypeSQLDelete : error sql type when deleting record
	ErrorTypeSQLDelete string = "sql_delete"
	// ErroTypeESQuery : error elasticsearch type when querying record
	ErroTypeESQuery string = "es_query"
)
