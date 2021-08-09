// +build gofuzzbeta

package sub

import "testing"

func Fuzz(f *testing.F) {
	f.Fuzz(func(t *testing.T, data string) {
		if data == "this is a test" {
			panic(data)
		}
	})
}
