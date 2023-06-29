package config

import (
	"log"
	"os"
	"path"
)

func EnsureBaseDir(fpath string) error {
	baseDir := path.Dir(fpath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}

func LogInto(filePath string) log.Logger {
	var InfoLogger *log.Logger
	err := EnsureBaseDir(filePath)
	if err != nil {
		log.Fatal("Could not create log folders")
	}
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	InfoLogger.SetOutput(os.Stdout)
	return *InfoLogger
}
