//go:build !windows && !nacl && !plan9
// +build !windows,!nacl,!plan9

package syslog

import (
	"fmt"
	"log/syslog"
	"os"

	"github.com/Hidayathamir/logrusy"
)

// SyslogHook to send logs via syslog.
type SyslogHook struct {
	Writer        *syslog.Writer
	SyslogNetwork string
	SyslogRaddr   string
}

// Creates a hook to be added to an instance of logger. This is called with
// `hook, err := NewSyslogHook("udp", "localhost:514", syslog.LOG_DEBUG, "")`
// `if err == nil { log.Hooks.Add(hook) }`
func NewSyslogHook(network, raddr string, priority syslog.Priority, tag string) (*SyslogHook, error) {
	w, err := syslog.Dial(network, raddr, priority, tag)
	return &SyslogHook{w, network, raddr}, err
}

func (hook *SyslogHook) Fire(entry *logrusy.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	switch entry.Level {
	case logrusy.PanicLevel:
		return hook.Writer.Crit(line)
	case logrusy.FatalLevel:
		return hook.Writer.Crit(line)
	case logrusy.ErrorLevel:
		return hook.Writer.Err(line)
	case logrusy.WarnLevel:
		return hook.Writer.Warning(line)
	case logrusy.InfoLevel:
		return hook.Writer.Info(line)
	case logrusy.DebugLevel, logrusy.TraceLevel:
		return hook.Writer.Debug(line)
	default:
		return nil
	}
}

func (hook *SyslogHook) Levels() []logrusy.Level {
	return logrusy.AllLevels
}
