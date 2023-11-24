package shamir

import (
	"fmt"
	"testing"
)

func TestShamir(t *testing.T) {
	mySecret := []byte("This is a secret")
	split, err := SplitSeed(mySecret, 5, 3)
	if err != nil {
		t.Fatalf("Error splitting secret: %s", err.Error())
	}
	first := split[0]
	second := split[1]
	third := split[2]

	group := []ShamirPart{third, first, second}

	reconstructed, err := ReconstructSeed(group)
	if err != nil {
		t.Fatalf("Error reconstructing secret: %s", err.Error())
	}
	fmt.Println(string(reconstructed))
}
