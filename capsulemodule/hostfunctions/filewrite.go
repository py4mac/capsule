package hostfunctions

import (
	"errors"
	"github.com/bots-garden/capsule/commons"
	"strconv"
	_ "unsafe"
)

//export hostWriteFile
//go:linkname hostWriteFile
func hostWriteFile(filePathPtrPos uint32, size uint32, contentPtrPos uint32, contentSize uint32, retBuffPtrPos **byte, retBuffSize *int)

// WriteFile : Call host function: hostReadFile
// Pass a string as parameter
//Get a string from the host

func WriteFile(filePath string, content string) (string, error) {

	filePathPtrPos, size := getStringPtrPositionAndSize(filePath)
	contentPtrPos, contentSize := getStringPtrPositionAndSize(content)

	var buffPtr *byte
	var buffSize int

	hostWriteFile(filePathPtrPos, size, contentPtrPos, contentSize, &buffPtr, &buffSize)

	var resultStr = ""
	var err error
	valueStr := getStringResult(buffPtr, buffSize)

	// check the return value
	if commons.IsErrorString(valueStr) {
		errorMessage, errorCode := commons.GetErrorStringInfo(valueStr)
		if errorCode == 0 {
			err = errors.New(errorMessage)
		} else {
			err = errors.New(errorMessage + " (" + strconv.Itoa(errorCode) + ")")
		}

	} else {
		resultStr = valueStr
	}
	return resultStr, err

}
