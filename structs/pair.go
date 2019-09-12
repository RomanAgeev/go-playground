package structs

type Pair struct {
	Fst interface{}
	Snd interface{}
}

func NewPair(fst interface{}, snd interface{}) *Pair {
	return &Pair{
		Fst: fst,
		Snd: snd,
	}
}
