package encoder

import (
	"errors"
)

type Nucleotide byte

const (
	A Nucleotide = 'A'
	C Nucleotide = 'C'
	T Nucleotide = 'T'
	G Nucleotide = 'G'
)

func ToNucleotide(b byte) (Nucleotide, error) {
	switch b {
	case 0:
		return A, nil
	case 1:
		return C, nil
	case 2:
		return T, nil
	case 3:
		return G, nil
	default:
		return 0, errors.New("Value of bits is invalid for DNA")
	}
}

func FromNucleotide(n Nucleotide) (byte, error) {
	switch n {
	case A:
		return 0, nil
	case C:
		return 1, nil
	case G:
		return 2, nil
	case T:
		return 3, nil
	default:
		return 0, errors.New("Invalid Nucleotide")
	}
}
