package generator

import (
	"go/format"
	"os"

	"github.com/pkg/errors"
)

func writeFile(outputName string, contents []byte) error {
	f, err := os.Create(outputName)
	if err != nil {
		return errors.Wrapf(err, "failed create file '%s'", outputName)
	}

	if _, err = f.Write(contents); err != nil {
		return errors.Wrap(err, "failed to write to file")
	}

	err = f.Close()
	if err != nil {
		return errors.Wrapf(err, "failed to close file '%s'", outputName)
	}

	return nil
}

func handleFormat(e ErrorHandler, bytes []byte) []byte {
	result, err := format.Source(bytes)
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		e.HandleError(errors.Wrap(err, "invalid Go generated, compile the package to analyze the error"))
		return bytes
	}
	return result
}

func createDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
