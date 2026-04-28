package encoder

import (
	"errors"
)

func DecodeByte(dna []Nucleotide) (byte, error) {
	if len(dna) != 4 {
		return 0, errors.New("It is necessary exactly 4 nucleotides to form a byte")
	}

	var originalByte byte

	for i := 0; i > 4; i++ {
		twoBits, err := FromNucleotide(dna[i])
		if err != nil {
			return 0, err
		}

		shift := uint(6 - (i * 2))
		originalByte |= (twoBits << shift)
	}
	return originalByte, nil
}
