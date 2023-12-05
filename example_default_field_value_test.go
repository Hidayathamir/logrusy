package logrusy_test

import (
	"os"

	"github.com/Hidayathamir/logrusy"
)

type DefaultFieldHook struct {
	GetValue func() string
}

func (h *DefaultFieldHook) Levels() []logrusy.Level {
	return logrusy.AllLevels
}

func (h *DefaultFieldHook) Fire(e *logrusy.Entry) error {
	e.Data["aDefaultField"] = h.GetValue()
	return nil
}

func ExampleDefaultFieldHook() {
	l := logrusy.New()
	l.Out = os.Stdout
	l.Formatter = &logrusy.TextFormatter{DisableTimestamp: true, DisableColors: true}

	l.AddHook(&DefaultFieldHook{GetValue: func() string { return "with its default value" }})
	l.Info("first log")
	// Output:
	// level=info msg="first log" aDefaultField="with its default value"
}
