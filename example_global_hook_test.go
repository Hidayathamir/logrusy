package logrusy_test

import (
	"os"

	"github.com/Hidayathamir/logrusy"
)

var (
	mystring string
)

type GlobalHook struct {
}

func (h *GlobalHook) Levels() []logrusy.Level {
	return logrusy.AllLevels
}

func (h *GlobalHook) Fire(e *logrusy.Entry) error {
	e.Data["mystring"] = mystring
	return nil
}

func ExampleGlobalHook() {
	l := logrusy.New()
	l.Out = os.Stdout
	l.Formatter = &logrusy.TextFormatter{DisableTimestamp: true, DisableColors: true}
	l.AddHook(&GlobalHook{})
	mystring = "first value"
	l.Info("first log")
	mystring = "another value"
	l.Info("second log")
	// Output:
	// level=info msg="first log" mystring="first value"
	// level=info msg="second log" mystring="another value"
}
