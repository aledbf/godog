package godog

import (
	"fmt"

	"github.com/DATA-DOG/godog/gherkin"
)

type registeredFormatter struct {
	name        string
	fmt         Formatter
	description string
}

var formatters []*registeredFormatter

// RegisterFormatter registers a feature suite output
// Formatter as the name and descriptiongiven.
// Formatter is used to represent suite output
func RegisterFormatter(name, description string, f Formatter) {
	formatters = append(formatters, &registeredFormatter{
		name:        name,
		fmt:         f,
		description: description,
	})
}

// Formatter is an interface for feature runner
// output summary presentation.
//
// New formatters may be created to represent
// suite results in different ways. These new
// formatters needs to be registered with a
// RegisterFormatter function call
type Formatter interface {
	Node(interface{})
	Failed(*gherkin.Step, *StepDef, error)
	Passed(*gherkin.Step, *StepDef)
	Skipped(*gherkin.Step)
	Undefined(*gherkin.Step)
	Summary()
}

// failed represents a failed step data structure
// with all necessary references
type failed struct {
	step *gherkin.Step
	def  *StepDef
	err  error
}

func (f failed) line() string {
	var tok *gherkin.Token
	var ft *gherkin.Feature
	if f.step.Scenario != nil {
		tok = f.step.Scenario.Token
		ft = f.step.Scenario.Feature
	} else {
		tok = f.step.Background.Token
		ft = f.step.Background.Feature
	}
	return fmt.Sprintf("%s:%d", ft.Path, tok.Line)
}

// passed represents a successful step data structure
// with all necessary references
type passed struct {
	step *gherkin.Step
	def  *StepDef
}

// skipped represents a skipped step data structure
// with all necessary references
type skipped struct {
	step *gherkin.Step
}

// undefined represents a pending step data structure
// with all necessary references
type undefined struct {
	step *gherkin.Step
}