package project

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

func FMT(path string) error {
	err := os.Chdir(path)
	if err != nil {
		return errors.Wrap(err, "failed to navigate to dir to format")
	}

	cmd := exec.Command("gofmt", "-w", "-s", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, "failed to format dir")
	}

	return nil
}

func InitModule(path, module string) error {
	err := os.Chdir(path)
	if err != nil {
		return errors.Wrap(err, "failed to navigate to dir to initialize module")
	}

	cmd := exec.Command("go", "mod", "init", module)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, "failed to init go modules")
	}

	return nil
}

func Setup(path string) error {
	err := os.Chdir(path)
	if err != nil {
		return errors.Wrap(err, "failed to navigate to dir to run project setup")
	}

	cmd := exec.Command("make", "setup")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, "failed to run project setup")
	}

	return nil
}
