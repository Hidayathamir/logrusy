//go:build !windows
// +build !windows

package logrusy_test

import (
	"log/syslog"
	"os"

	"github.com/Hidayathamir/logrusy"
	slhooks "github.com/Hidayathamir/logrusy/hooks/syslog"
)

// An example on how to use a hook
func Example_hook() {
	var log = logrusy.New()
	log.Formatter = new(logrusy.TextFormatter)                     // default
	log.Formatter.(*logrusy.TextFormatter).DisableColors = true    // remove colors
	log.Formatter.(*logrusy.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	if sl, err := slhooks.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, ""); err == nil {
		log.Hooks.Add(sl)
	}
	log.Out = os.Stdout

	log.WithFields(logrusy.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrusy.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrusy.Fields{
		"omg":    true,
		"number": 100,
	}).Error("The ice breaks!")

	// Output:
	// level=info msg="A group of walrus emerges from the ocean" animal=walrus size=10
	// level=warning msg="The group's number increased tremendously!" number=122 omg=true
	// level=error msg="The ice breaks!" number=100 omg=true
}
