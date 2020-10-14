package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// DFA Representa un Automata Finito Determinista (AFD)
type DFA struct {
	Alphabet     []string     `json:"alphabet"`
	State        []string     `json:"state"`
	InitialState string       `json:"initialState"`
	FinalStates  []string     `json:"finalState"`
	Transitions  []Transition `json:"transitions"`
}

// Transition Representa una transicion en el AFD
type Transition struct {
	Init string `json:"init"`
	Alph string `json:"alph"`
	End  string `json:"end"`
}

// handleError Maneja errores
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// findString Toma un arreglo de string y busca un string en el. Si lo encuentra retorna su posicion y true, sino retorna -1 y false.
func findString(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// findTransition Toma un arreglo de Transition y busca un state en el. Si lo encuentra retorna su posicion y true, sino retorna -1 y false.
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

// isFinite construye de manera recursiva una cadena.
// Se debe pasar por copia el AFD y no por referencia.
func isFinite(state string, str string, dfa DFA) bool {
	// Validamos que state exista en state
	_, stateExist := findString(dfa.State, state)
	if !stateExist {
		panic(state + " not found in the state")
	}

	// Validamos las condiciones de parada.
	strLen := len(str)
	stateLen := len(dfa.State)
	if strLen > stateLen {
		return false
	}

	// Si corresponde a un estado final y no tiene transicion a el mismo, paramos y retornamos true.
	_, isFinal := findString(dfa.FinalStates, state)
	_, selfTransition := findSelfTransition(dfa.Transitions, state)
	if isFinal && !selfTransition {
		return true
	}

	// Validamos que exista al menos una transicion
	_, transitionExist := findTransition(dfa.Transitions, state)
	if !transitionExist {
		panic("Transition not found for state " + state)
	}

	// Si encontramos state en transitions llamamos recursivamente a isFinite con el siguiente estado y la nueva cadena.
	finite := true
	for _, transition := range dfa.Transitions {
		// Validamos que transition.Init exista en state
		_, initExist := findString(dfa.State, transition.Init)
		if !initExist {
			panic(transition.Init + " not found in the state")
		}

		// Validamos que transition.Alph exista en alphabet
		_, alphExist := findString(dfa.Alphabet, transition.Alph)
		if !alphExist {
			panic(transition.Alph + " not found in the alphabet")
		}

		if transition.Init != state {
			continue
		}

		// Construimos el nuevo string y realizamos la BFS para el siguiente estado.
		newStr := str + transition.Alph
		finite = finite && isFinite(transition.End, newStr, dfa)
	}

	return finite
}

// main Funcion principal del programa.
func main() {
	// Ingrese el nombre del archivo.
	fmt.Println("Enter the file name:")
	var fileName string
	fmt.Scanf("%s", &fileName)

	// Leemos el archivo.
	body, err := ioutil.ReadFile(fileName)
	handleError(err)

	// Solo con proposito de pruebas...
	// fmt.Println(string(body))

	// Mapeamos el archivo en una estructura.
	dfa := DFA{}
	err = json.Unmarshal([]byte(body), &dfa)
	handleError(err)

	// Solo con proposito de pruebas...
	// fmt.Println(dfa)

	// Validamos si el AFD es finito o infinito.
	finite := isFinite(dfa.InitialState, "", dfa)
	fmt.Println(finite)
}
