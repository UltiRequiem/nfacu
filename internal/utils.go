package internal

import "fmt"

// Return a valid XML line
func ParseLine(key, val string) string {
	return fmt.Sprintf(`    <add key="%s" value="%s" />`, key, val)
}
