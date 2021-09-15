package ed25519

import (
	"testing"
)

//==================================== Helper =============================================

type StaticRand struct { id int }
func (sr *StaticRand) Read(x []byte) (int, error) { return sr.id, nil }

func TestCrypto(t *testing.T) {
	
}
