package global

import "fmt"

// errState recording program state
type ErrState struct {
	HadError bool
	S        string
}

// errReport  report current error occure
type ErrReport struct {
	line  int
	where string
	msg   string
	state ErrState
}

func (e *ErrState) Error() string {
	return e.S
}
func New(s string) error {
	return &ErrState{true, s}
}
func (eState *ErrState) Erno(line int, msg string) {
	report(ErrReport{line, "", msg, *eState})
}
func report(ep ErrReport) {
	fmt.Println(fmt.Errorf("[line %d] Error:%s : %s", ep.line, ep.where, ep.msg))
	ep.state.HadError = true
}
