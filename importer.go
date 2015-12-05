package main

type transaction struct {
	date        string
	posted      string
	code        string
	description string
	amount      string
	cost        string
	total       string
	note        string
}

type Importer interface {
	Import() (transaction, error)
}

func importAll(in Importer) []Block {
	var blocks []Block

	for t, err := in.Import(); err == nil; {
		_ = t
		lines := []string{}

		b := Block{lines: lines}
		blocks = append(blocks, b)
	}

	return blocks
}
