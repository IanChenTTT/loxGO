package lox

import "fmt"

// errState recording program state
type errState struct {
	hadError bool
	s        string
}

// errReport  report current error occure
type errReport struct {
	line  int
	where string
	msg   string
	state errState
}

func (e *errState) Error() string {
	return e.s
}
func New(s string) error {
	return &errState{true, s}
}
func (eState *errState) erno(line int, msg string) {
	report(errReport{line, "", msg, *eState})
}
func report(ep errReport) {
	fmt.Println(fmt.Errorf("[line %d] Error:%s : %s", ep.line, ep.where, ep.msg))
	ep.state.hadError = true
}
