#!/bin/bash

if [ -f "plog_test.go" ]; then
	echo "plog test go file already exist!"
	exit 0
fi

echo "generating new test go file..."

echo """
package plog_test

import (
	\"fmt\"
	\"testing\"

	\"github.com/devproje/plog\"
	\"github.com/devproje/plog/level\"
	\"github.com/devproje/plog/log\"
	\"github.com/devproje/plog/utils\"
)

var testStr = \"Hello, World!\"

func TestLogging(t *testing.T) {
	log.SetLevel(level.Trace)
	log.SetTimestamp(false)
	print()

	log.SetTimestamp(true)
	print()

	log.SetTimeFormat(utils.TimeFormatter(\"YYYY-MM-DD hh:mm:ss a ZZZ\"))
	print()
}

func TestCustomLogging(t *testing.T) {
	logger := plog.New()

	logger.Info(testStr)
}

func print() {
	log.Error(fmt.Sprintf(\"%s\n\", testStr))
	log.Warn(fmt.Sprintf(\"%s\n\", testStr))
	log.Info(fmt.Sprintf(\"%s\n\", testStr))
	log.Debug(fmt.Sprintf(\"%s\n\", testStr))
	log.Trace(fmt.Sprintf(\"%s\n\", testStr))
	fmt.Printf(\"n\")

	log.Errorf(\"%s\n\", testStr)
	log.Warnf(\"%s\n\", testStr)
	log.Infof(\"%s\n\", testStr)
	log.Debugf(\"%s\n\", testStr)
	log.Tracef(\"%s\n\", testStr)
	fmt.Printf(\"\n\")

	log.Errorln(testStr)
	log.Warnln(testStr)
	log.Infoln(testStr)
	log.Debugln(testStr)
	log.Traceln(testStr)

	fmt.Printf(\"\n\n\")
}
""" > plog_test.go

echo "test go file generating complete!"