package benchtest

import (
	"fmt"
	"testing"
)

func WriteSprintf() string {
	return fmt.Sprintf("%s%s", "$", "my_value")
}

func WriteAdd() string {
	return "$" + "my_value"
}

func BenchmarkWriteSprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WriteSprintf()
	}
}

func BenchmarkWriteAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WriteAdd()
	}
}
