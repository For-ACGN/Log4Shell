package log4shell

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommandLineToArgs(t *testing.T) {
	exe1 := "exe1"
	exe2 := `"exe2 arg1"`
	for _, testdata := range [...]*struct {
		cmd  string
		args []string
	}{
		{"net", []string{"net"}},
		{`net -a -b`, []string{"net", "-a", "-b"}},
		{`net -a -b "a"`, []string{"net", "-a", "-b", "a"}},
		{`"net net"`, []string{"net net"}},
		{`"net\net"`, []string{`net\net`}},
		{`"net\net net"`, []string{`net\net net`}},
		{`net -a \"net`, []string{"net", "-a", `"net`}},
		{`net -a ""`, []string{"net", "-a", ""}},
		{`""net""  -a  -b`, []string{"net", "-a", "-b"}},
		{`"""net""" -a`, []string{`"net"`, "-a"}},
	} {
		args := CommandLineToArgs(exe1 + " " + testdata.cmd)
		require.Equal(t, append([]string{"exe1"}, testdata.args...), args)

		args = CommandLineToArgs(exe2 + " " + testdata.cmd)
		require.Equal(t, append([]string{"exe2 arg1"}, testdata.args...), args)
	}
}
