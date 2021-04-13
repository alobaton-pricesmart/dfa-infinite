package dfa

// Transition represents a transaction in the DFA
type Transition struct {
	Init string `json:"init"`
	Alph string `json:"alph"`
	End  string `json:"end"`
}

// findTransition look for an state in a transition slice. Returns the position and true if found, -1 and false otherwise.
func findTransition(slice []Transition, state string) (int, bool) {
	for i, item := range slice {
		if item.Init == state {
			return i, true
		}
	}
	return -1, false
}

// findTransition Toma un arreglo de Transition y busca un state que tenga un transition a el mismo. Si lo encuentra retorna su posicion y true, sino retorna -1 y false.
func findSelfTransition(slice []Transition, state string) (int, bool) {
	for i, item := range slice {
		if item.Init == state && item.End == state {
			return i, true
		}
	}
	return -1, false
}
