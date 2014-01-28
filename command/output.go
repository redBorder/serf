package command

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"fmt"
)

// Format some raw data for output. For better or worse, this currently forces
// the passed data object to implement fmt.Stringer, since it's pretty hard to
// implement a canonical *-to-string function.
func formatOutput(data interface{}, format string) ([]byte, error) {
	var out string

	switch format {

	case "json":
		jsonout, err := json.MarshalIndent(data.(fmt.Stringer), "", "  ")
		if err != nil {
			return nil, err
		}
		out = string(jsonout)

	case "xml":
		xmlout, err := xml.MarshalIndent(data.(fmt.Stringer), "", "  ")
		if err != nil {
			return nil, err
		}
		out = fmt.Sprintf("%s%s", xml.Header, xmlout)

	case "text":
		out = data.(fmt.Stringer).String()

	default:
		return nil, fmt.Errorf("Invalid output format \"%s\"", format)

	}
	return []byte(prepareOutput(out)), nil
}

// Apply some final formatting to make sure we don't end up with extra newlines
func prepareOutput(in string) string {
	return strings.TrimSpace(string(in))
}
