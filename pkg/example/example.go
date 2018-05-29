package example

import (
	"fmt"

	"github.com/golang/glog"
)

type Example struct {
	Argument  string
	SuperFlag string
}

func (e *Example) Hello() error {
	glog.V(1).Infof("Hello %s with %q from -v 1", e.Argument, e.SuperFlag)
	return nil
}

func (e *Example) GetArgument() string {
	return e.Argument
}

func (e *Example) GetSuperFlag() (string, error) {
	if e.SuperFlag != "" {
		return e.SuperFlag, nil
	}
	err := fmt.Errorf("empty SuperFlag")
	glog.Errorf("Unexpected error: %v", err)
	return "", err
}
