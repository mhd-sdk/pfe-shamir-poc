package shamir

import (
	"github.com/corvus-ch/shamir"
)

type ShamirPart struct {
	Key byte   `json:"key"`
	Val []byte `json:"val"`
}

// SplitSeed splits a secret into a number of `parts`, with a `threshold` of parts needed to reconstruct the secret
func SplitSeed(secret []byte, parts int, threshold int) (splitResult []ShamirPart, err error) {
	shamirSplit, err := shamir.Split(secret, parts, threshold)
	if err != nil {
		return nil, err
	}
	splitResult = make([]ShamirPart, 0)
	for Key, Val := range shamirSplit {
		splitResult = append(splitResult, ShamirPart{Key, Val})
	}
	return splitResult, nil
}

// ReconstructSeed reconstructs a secret from a number of `parts`
func ReconstructSeed(parts []ShamirPart) (reconstructed []byte, err error) {
	shamirCombine := make(map[byte][]byte)
	for _, part := range parts {
		shamirCombine[part.Key] = part.Val
	}
	reconstructed, err = shamir.Combine(shamirCombine)
	if err != nil {
		return nil, err
	}
	return reconstructed, nil
}
