package encoder

func EncodeByte(b byte) ([]Nucleotide, error) {
	dna := make([]Nucleotide, 4)

	for i := 0; i < 4; i++ {
		shift := uint(6 - (i * 2))
		twoBits := (b >> shift) & 0x03

		n, err := ToNucleotide(twoBits)
		if err != nil {
			return nil, err
		}
		dna[i] = n
	}

	return dna, nil
}
