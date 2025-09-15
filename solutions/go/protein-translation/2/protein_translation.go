package protein

import (
	"errors"
)

var ErrStop = errors.New("stop processing")
var ErrInvalidBase = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	var codons = []string{}

	for i := 0; i <= 2; i++ {

		start := i * 3
		end := (i * 3) + 3
		v, err := FromCodon(rna[start:end])
		if err != nil {
			if errors.Is(err, ErrInvalidBase) {
				return nil, ErrInvalidBase
			} else {
				break
			}
		}

		codons = append(codons, v)
	}

	return codons, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
