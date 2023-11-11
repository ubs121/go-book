package debug

import (
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	Info := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info.Println("Hello from logger !")
}
