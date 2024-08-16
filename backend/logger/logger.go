package logger 


import (
	"log/slog"
	"os"
)


func GetLogger() *slog.Logger{
	var jsonHandler *slog.JSONHandler = slog.NewJSONHandler(os.Stderr, nil)
	var myslog *slog.Logger = slog.New(jsonHandler)

	return myslog

}