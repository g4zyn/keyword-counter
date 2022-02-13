package result

// Summary
type Summary map[string]int

type Result struct {
	CorpusName string
	Summary    Summary
}

type Store interface{}

func New(corpus string, summary Summary) {
	_ = &Result{
		CorpusName: corpus,
		Summary:    summary,
	}
}
