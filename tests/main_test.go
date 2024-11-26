package tests

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
     // Get the directory of the current file
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        log.Fatal("Unable to get the current file path")
    }

    // Load .env from the project root
    projectRoot := filepath.Join(filepath.Dir(filename), "..")
    err := godotenv.Load(filepath.Join(projectRoot, ".env"))
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    os.Exit(m.Run())
}
