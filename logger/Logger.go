package logger

import (
	"encoding/json"
	"log/slog"
	"os"
)

type InfoLog struct {
	Method  string      `json:"method"`
	Data    interface{} `json:"data"`
}

type ErrorLog struct {
	Method  string      `json:"method"`
	Data    interface{} `json:"data"`
}

var (
	logger *slog.Logger
)

func init() {
	log := slog.NewJSONHandler(os.Stdout, nil)
	logger = slog.New(log)
}

func Info(method string, message string, data interface{}) {

	log := InfoLog{
		Method: method,
		Data: data,
	}

	ConvertToJson, _ := json.Marshal(log)

	logger.Info(message,"Info",string(ConvertToJson))
}

func Error(method string, message string, data interface{}) {

	log := ErrorLog{
		Method: method,
		Data: data,
	}

	ConvertToJson, _ := json.Marshal(log)

	logger.Error(message,"Error",string(ConvertToJson))
}