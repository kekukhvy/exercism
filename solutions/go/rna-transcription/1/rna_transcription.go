package strand

func ToRNA(dna string) string {

	var rna = make([]rune, len(dna))

	for i, c := range dna {
		switch c {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		}
	}

	return string(rna)
}
