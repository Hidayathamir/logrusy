package logrusy_test

import (
	"os"

	"github.com/Hidayathamir/logrusy"
)

func Example_basic() {
	var log = logrusy.New()
	log.Formatter = new(logrusy.JSONFormatter)
	log.Formatter = new(logrusy.TextFormatter)                     //default
	log.Formatter.(*logrusy.TextFormatter).DisableColors = true    // remove colors
	log.Formatter.(*logrusy.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	log.Level = logrusy.TraceLevel
	log.Out = os.Stdout

	// file, err := os.OpenFile("logrusy.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	// 	log.Out = file
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }

	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logrusy.Entry)
			log.WithFields(logrusy.Fields{
				"omg":         true,
				"err_animal":  entry.Data["animal"],
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
				"number":      100,
			}).Error("The ice breaks!") // or use Fatal() to force the process to exit with a nonzero code
		}
	}()

	log.WithFields(logrusy.Fields{
		"animal": "walrus",
		"number": 0,
	}).Trace("Went to the beach")

	log.WithFields(logrusy.Fields{
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	log.WithFields(logrusy.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrusy.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrusy.Fields{
		"temperature": -4,
	}).Debug("Temperature changes")

	log.WithFields(logrusy.Fields{
		"animal": "orca",
		"size":   9009,
	}).Panic("It's over 9000!")

	// Output:
	// level=trace msg="Went to the beach" animal=walrus number=0
	// level=debug msg="Started observing beach" animal=walrus number=8
	// level=info msg="A group of walrus emerges from the ocean" animal=walrus size=10
	// level=warning msg="The group's number increased tremendously!" number=122 omg=true
	// level=debug msg="Temperature changes" temperature=-4
	// level=panic msg="It's over 9000!" animal=orca size=9009
	// level=error msg="The ice breaks!" err_animal=orca err_level=panic err_message="It's over 9000!" err_size=9009 number=100 omg=true
}
