package format

import "testing"

func TestPhone(t *testing.T) {
	source := "8 (985) 111-22-33"

	t.Log(PhoneLength(source, 12))
}
