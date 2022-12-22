package log

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/evilsocket/islazy/tui"
)

type MessageCallback = func(verbosity Verbosity, message string)

var (
	// Level represents the current verbosity level of the logging system.
	Level = INFO
	// Output represents the log output file path if filled or, if empty, stdout.
	Output = ""
	// NoEffects disables all effects and colors if set to true.
	NoEffects = false
	// OnFatal represents the callback/action to execute on Fatal messages.
	OnFatal = ExitOnFatal
	// A custom callback to execute for every log message.
	Callback = dummyCallback

	lock        = &sync.Mutex{}
	currMessage = ""
	currLevel   = INFO
	writer      = os.Stdout

	reEffects = []*regexp.Regexp{
		regexp.MustCompile("\x033\\[\\d+m"),
		regexp.MustCompile("\\\\e\\[\\d+m"),
		regexp.MustCompile("\x1b\\[\\d+m"),
	}
)

// Open initializes the logging system.
func Open() (err error) {
	if Output != "" {
		writer, err = os.OpenFile(Output, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	}
	return
}

// Close finalizes the logging system.
func Close() {
	if writer != os.Stdout {
		writer.Close()
	}
}

func dummyCallback(verbosity Verbosity, message string) {

}

func removeEffects(s string) string {
	for _, re := range reEffects {
		s = re.ReplaceAllString(s, "")
	}
	return s
}

func emit(v Verbosity, s string) {
	// remove all effects if found
	plain := removeEffects(s)

	Callback(v, plain)

	if NoEffects {
		s = plain
	}
	fmt.Fprintf(writer, "%s\n", s)
}

func do(v Verbosity, format string, args ...interface{}) {
	lock.Lock()
	defer lock.Unlock()

	if Level > v {
		return
	}

	logLine := Format
	currLevel = v
	currMessage = format
	if args != nil {
		currMessage = fmt.Sprintf(format, args...)
	}
	// process token -> callback
	for token, cb := range Tokens {
		logLine = strings.Replace(logLine, token, cb(), -1)
	}
	// process token -> effect
	for token, effect := range Effects {
		logLine = strings.Replace(logLine, token, effect, -1)
	}
	// make sure an user error does not screw the log
	if tui.HasEffect(logLine) && !strings.HasSuffix(logLine, tui.RESET) {
		logLine += tui.RESET
	}

	emit(v, logLine)
}

// Raw emits a message without format to the logs.
func Raw(format string, args ...interface{}) {
	lock.Lock()
	defer lock.Unlock()

	currMessage = fmt.Sprintf(format, args...)
	emit(INFO, currMessage)
}

// Debug emits a debug message.
func Debug(format string, args ...interface{}) {
	do(DEBUG, format, args...)
}

// Info emits an informative message.
func Info(format string, args ...interface{}) {
	do(INFO, format, args...)
}

// Important emits an important informative message.
func Important(format string, args ...interface{}) {
	do(IMPORTANT, format, args...)
}

// Warning emits a warning message.
func Warning(format string, args ...interface{}) {
	do(WARNING, format, args...)
}

// Error emits an error message.
func Error(format string, args ...interface{}) {
	do(ERROR, format, args...)
}

// Fatal emits a fatal error message and calls the log.OnFatal callback.
func Fatal(format string, args ...interface{}) {
	do(FATAL, format, args...)
	OnFatal()
}
