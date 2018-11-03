// Package twofer or 2-fer is short for two for one. One for you and one for me.
package twofer

import (
	"bytes"
)

// ShareWith returns you if no name is given
func ShareWith(name string) string {

	buffer := bytes.Buffer{}
	you := "One for "
	me := ", one for me."

	if name == "" {
		name = "you"
	}

	buffer.WriteString(you)
	buffer.WriteString(name)
	buffer.WriteString(me)

	message := buffer.String()

	return message
}
