package number

import (
	"encoding/json"
	"strings"

	"github.com/opctl/opctl/sdks/go/data/coerce"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/op/params/param/formats"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
)

// Validate validates a value against a number parameter
func Validate(
	value *model.Value,
	constraints map[string]interface{},
) []error {
	// register custom format checkers
	gojsonschema.FormatCheckers.Add("integer", formats.IntegerFormatChecker{})

	valueAsNumber, err := coerce.ToNumber(value)
	if err != nil {
		return []error{err}
	}

	// guard no constraints
	if constraints != nil {
		errs := []error{}

		valueJSONBytes, err := json.Marshal(valueAsNumber.Number)
		if err != nil {
			// handle syntax errors specially
			return append(
				errs,
				errors.Wrap(err, "error validating parameter"),
			)
		}

		result, err := gojsonschema.Validate(
			gojsonschema.NewGoLoader(constraints),
			gojsonschema.NewBytesLoader(valueJSONBytes),
		)
		if err != nil {
			// handle syntax errors specially
			return append(
				errs,
				errors.Wrap(err, "error validating parameter"),
			)
		}

		for _, errString := range result.Errors() {
			// enum validation errors include `(root) ` prefix we don't want
			errs = append(errs, errors.New(strings.TrimPrefix(errString.Description(), "(root) ")))
		}

		return errs
	}

	return nil
}
