package subcommands_test

import (
	//	"bazil.org/bazil/cliutil/positional"
	"bytes"
	"flag"
	"testing"

	"bazil.org/bazil/cliutil/subcommands"
	"bazil.org/bazil/cliutil/subcommands/test/calc"
	_ "bazil.org/bazil/cliutil/subcommands/test/calc/sum"
)

func TestUsage(t *testing.T) {
	result, err := subcommands.Parse(&calc.Calc, "calc", []string{"-help"})
	if err != flag.ErrHelp {
		t.Errorf("unexpected error message: %q", err.Error())
	}
	var buf bytes.Buffer
	result.UsageTo(&buf)
	if g, e := buf.String(), `Usage:
  calc [OPT..] COMMAND..

Options:
  -v=false: verbose output

Commands:
  sum    sum two numbers
`; g != e {
		t.Errorf("unexpected usage:\n%s", g)
		t.Logf("got: %q", g)
		t.Logf("exp: %q", e)
	}
}

func TestUsageSub(t *testing.T) {
	result, err := subcommands.Parse(&calc.Calc, "calc", []string{"sum", "-help"})
	if err != flag.ErrHelp {
		t.Errorf("unexpected error message: %q", err.Error())
	}
	var buf bytes.Buffer
	result.UsageTo(&buf)
	if g, e := buf.String(), `Usage:
  calc sum [OPT..] A B

Computes the sum of two integers, A and B.

Options:
  -frobnicate=false: frobnicate the qubbitz
`; g != e {
		t.Errorf("unexpected usage:\n%s", g)
		t.Logf("got: %q", g)
		t.Logf("exp: %q", e)
	}
}
