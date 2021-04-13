package dfa

import "dfa-infinite/pkg/strutil"

// DFA Representa un Automata Finito Determinista (AFD)
type DFA struct {
	Alphabet     []string     `json:"alphabet"`
	State        []string     `json:"state"`
	InitialState string       `json:"initialState"`
	FinalStates  []string     `json:"finalState"`
	Transitions  []Transition `json:"transitions"`
}

// isFinite construye de manera recursiva una cadena.
// Se debe pasar por copia el AFD y no por referencia.
func (d DFA) IsFinite(state string, str string) bool {
	// Validamos que state exista en state
	_, stateExist := strutil.Find(d.State, state)
	if !stateExist {
		panic(state + " not found in the state")
	}

	// Validamos las condiciones de parada.
	strLen := len(str)
	stateLen := len(d.State)
	if strLen > stateLen {
		return false
	}

	// Si corresponde a un estado final y no tiene transicion a el mismo, paramos y retornamos true.
	_, isFinal := strutil.Find(d.FinalStates, state)
	_, selfTransition := findSelfTransition(d.Transitions, state)
	if isFinal && !selfTransition {
		return true
	}

	// Validamos que exista al menos una transicion
	_, transitionExist := findTransition(d.Transitions, state)
	if !transitionExist {
		panic("Transition not found for state " + state)
	}

	// Si encontramos state en transitions llamamos recursivamente a isFinite con el siguiente estado y la nueva cadena.
	finite := true
	for _, transition := range d.Transitions {
		// Validamos que transition.Init exista en state
		_, initExist := strutil.Find(d.State, transition.Init)
		if !initExist {
			panic(transition.Init + " not found in the state")
		}

		// Validamos que transition.Alph exista en alphabet
		_, alphExist := strutil.Find(d.Alphabet, transition.Alph)
		if !alphExist {
			panic(transition.Alph + " not found in the alphabet")
		}

		if transition.Init != state {
			continue
		}

		// Construimos el nuevo string y realizamos la BFS para el siguiente estado.
		newStr := str + transition.Alph
		finite = finite && d.IsFinite(transition.End, newStr)
	}

	return finite
}
