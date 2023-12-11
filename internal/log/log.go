package log

import (
	"os"
)

func Info(message string) {
	os.Stdout.WriteString(message + "\n")
}

func Error(message string) {
	os.Stderr.WriteString(message + "\n")

}
