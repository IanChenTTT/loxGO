package global

import (
	"fmt"
	t "github.com/IanChenTTT/loxGO/internal/lox/token"
)

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

// Erno will just have line and msg
func (eState *ErrState) Erno(line int, msg string) {
	report(ErrReport{
		line:  line,
		where: "",
		msg:   msg,
		state: *eState,
	})
}

// ErnoDetail will list all the struct parameter
func (eState *ErrState) ErnoDetail(line int, where string, msg string) {
	report(ErrReport{
		line:  line,
		where: where,
		msg:   msg,
		state: *eState,
	})
}
func (eState *ErrState) ErnoToken(tok t.Token, msg string) {
	err := ErrReport{
		line:  tok.Line,
		msg:   msg,
		where: "at: " + tok.Lexemes,
		state: *eState,
	}
	if tok.Types == t.EOF {
		err.where = " at end"
		report(err)
		return
	}
	report(err)
}
func report(ep ErrReport) {
	fmt.Println(fmt.Errorf("[line %d] Error:%s : %s", ep.line, ep.where, ep.msg))
	ep.state.HadError = true
}
