package jsonschema

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"

	"github.com/Bhinneka/alpha/api/lib/errorutil"
	"github.com/Bhinneka/alpha/api/lib/path"
)

// Validate : validate parameter with json schema
func Validate(fileName string, param interface{}) error {

	filePath := fmt.Sprintf("%s/service/jsonschema/%s", path.Root(), fileName)
	schema := gojsonschema.NewReferenceLoader(fmt.Sprintf("file:///%s", filePath))
	document := gojsonschema.NewGoLoader(param)
	schemaResult, err := gojsonschema.Validate(schema, document)
	if err != nil {
		return err
	}

	errs := errorutil.NewError(fileName)
	if !schemaResult.Valid() {
		for _, schemaError := range schemaResult.Errors() {
			errs.AppendMessage(schemaError.String())
		}
	}

	if errs.HasError() {
		return errs
	}

	return nil
}
