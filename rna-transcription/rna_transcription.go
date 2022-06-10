package strand

var DNAtoRNA = map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

func ToRNA(dna string) (out string) {
	for _, nucleotide := range dna {
		out += string(DNAtoRNA[nucleotide])
	}
	return
}
