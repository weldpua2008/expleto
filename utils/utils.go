package utils

import (
	// "bytes"
	// "encoding/xml"
	"fmt"
	// "io/ioutil"
	"path/filepath"
	"runtime"
	// "strconv"
	// "strings"
	// "text/template"
	// "time"
)

func FormatError(err error) error {
	pc, fn, line, _ := runtime.Caller(1)
	return fmt.Errorf("%v\nTRACE: %s[%s:%d]", err, runtime.FuncForPC(pc).Name(), filepath.Base(fn), line)
}