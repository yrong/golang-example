package main

import (
	"fmt"
	"log"
	"strings"
	"errors"
	u "github.com/go-goodies/go_utils"
)
// enums indicating number of characters for that type of word
// ex: a TINY word has 4 or fewer characters
const (
	TEENINY WordSize = 1
	SMALL   WordSize = 4 << iota
	MEDIUM                      // assigned 8 from iota
	LARGE                       // assigned 16 from iota
	XLARGE  WordSize = 32000
)

type WordSize int

func (ws WordSize) String() string {
	var s string
	if ws&TEENINY == TEENINY {
		s = "TEENINY"
	}
	return s
}

// ChainLink allows us to chain function/method calls.  It also keeps
// data internal to ChainLink, avoiding the side effect of mutated data.
type ChainLink struct {
	Data []string
}

func (v *ChainLink)Value() []string {
	return v.Data
}

// stringFunc is a first-class method, used as a parameter to _map
type stringFunc func(s string) (result string)

// _map uses stringFunc to modify (up-case) each string in the slice
func (v *ChainLink)_map(fn stringFunc) *ChainLink {
	var mapped []string
	orig := *v
	for _, s := range orig.Data {
		mapped = append(mapped, fn(s))  // first-class function
	}
	v.Data = mapped
	return v
}

// _filter uses embedded logic to filter the slice of strings
// Note: We could have chosen to use a first-class function
func (v *ChainLink)_filter(max WordSize) *ChainLink {
	filtered := []string{}
	orig := *v
	for _, s := range orig.Data {
		if len(s) <= int(max) {             // embedded logic
			filtered = append(filtered, s)
		}
	}
	v.Data = filtered
	return v
}


func main() {
	nums := []string{
		"tiny",
		"marathon",
		"philanthropinist",
		"supercalifragilisticexpialidocious"}

	data := ChainLink{nums};
	orig_data := data.Value()
	fmt.Printf("unfiltered: %#v\n", data.Value())

	filtered := data._filter(MEDIUM)
	fmt.Printf("filtered: %#v\n", filtered)

	fmt.Printf("filtered and mapped (MEDIUM sized words): %#v\n",
		filtered._map(strings.ToUpper).Value())

	data = ChainLink{nums}
	fmt.Printf("filtered and mapped (MEDIUM sized words): %#v\n",
		data._filter(MEDIUM)._map(strings.ToUpper).Value())

	data = ChainLink{nums}
	fmt.Printf("filtered twice and mapped (SMALL sized words): %#v\n",
		data._filter(MEDIUM)._map(strings.ToUpper)._filter(SMALL).Value())

	data = ChainLink{nums}
	val := data._map(strings.ToUpper)._filter(XLARGE).Value()
	fmt.Printf("mapped and filtered (XLARGE sized words): %#v\n", val)

	// heredoc with interpoloation
	constants := `
** Constants ***
SMALL: %d
MEDIUM: %d
LARGE: %d
XLARGE: %d
`
	fmt.Printf(constants, SMALL, MEDIUM, LARGE, XLARGE)
	fmt.Printf("TEENINY: %s\n\n", TEENINY)

	fmt.Printf("Join(nums, \"|\")     : %v\n", u.Join(nums, "|"))
	fmt.Printf("Join(orig_data, \"|\"): %v\n", u.Join(orig_data, "|"))
	fmt.Printf("Join(data, \"|\")     : %v\n\n", u.Join(data.Value(), "|"))

	if u.Join(nums, "|") == u.Join(orig_data, "|") {
		fmt.Println("No Side Effects!")
	} else {
		log.Print(errors.New("WARNING - Side Effects!"))

	}
}