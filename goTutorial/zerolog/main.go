package main

import (
	// "log"
	// "os"

	"net"
	"os"

	// "github.com/docker/docker/daemon/logger"
	"github.com/rs/zerolog"
	// "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	//? example 1
	// logger := zerolog.New(os.Stdout)
	// logger.Info().Msg("Hello Wold")
	/*
		output akan seperti ini :
		{"level":"info","message":"Hello World!"}
	*/
	//? example2
	// logger := zerolog.New(os.Stdout)
	// logger.Info().Msg("Hello Wold")
	// logger.Debug().Msg("Hello Wold")
	// logger.Error().Msg("Hello Wold")
	// logger.Fatal().Msg("Hello Wold")
	// logger.Panic().Msg("Hello Wold")
	//! import errors
	/*
		output akan seperti ini :
		{"level":"info","message":"Hello World!"}
		{"level":"Debug","message":"Hello World!"}
		{"level":"Error","message":"Hello World!"}
		{"level":"Fatal","message":"Hello World!"}
		{"level":"Panic","message":"Hello World!"}
		exit status 1
	*/
	//? example 3
	// zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// logger := zerolog.New(os.Stdout)
	// logger.Info().Msg("this is info message")
	// logger.Debug().Msg("this is info message")
	// logger.Error().Msg("this is info message")
	// logger.Warn().Msg("this is info message")
	/*
		output akan seperti ini :
		{"level":"error","message":"this is info message"}
	*/
	//? example 4
	// logger := zerolog.New(os.Stdout)
	// logger.Info().Str("name", "John Doe").Msg("")
	//! import errors
	/*
		output akan seperti ini :
		{"level":"info","name":"John Doe"}
	*/
	//? example 5
	// err := errors.New("Something went wrong")

	// logger := zerolog.New(os.Stdout)
	// logger.Error().Err(err).Msg("")

	//! import errors
	/*
		{"level":"error","error":"Something went wrong"}
	*/
	//? example 6

	// log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	// log.Info().Msg("Hello world")

	// import error
	/*
		{"level":"info","time":"2023-10-13T10:33:47+07:00","message":"Hello world"}
	*/
	//? example 7
	// log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	// log.Info().Msg("Hello world")
	//! import error
	/*
		{"level":"info","time":"2023-10-13T10:59:31+07:00","caller":"C:/Users/muhammad.rahman/Desktop/ArifDokumentasi/Dokuments/Golang/praktek/zerolog/main.go:74","message":"Hello world"}
	*/
	//? example 8 - masuk ke File
	// file, err := os.OpenFile("file.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// logger := zerolog.New(file).With().Timestamp().Logger()
	// logger.Info().Msg("I am logging to a file")
	//! membuat file dengan nama file.log
	//? example 9 - masuk ke file dengan rotasi
	// go get gopkg.in/natefinch/lumberjack.v2
	// lumberjackLogger := &lumberjack.Logger{
	// 	Filename:   "./file.log",
	// 	MaxSize:    1,
	// 	MaxBackups: 3,
	// 	MaxAge:     28,
	// 	Compress:   true,
	// }

	// logger := zerolog.New(lumberjackLogger).With().Timestamp().Logger()
	// loop := 0
	// for {
	// 	logger.Info().Msg("Iam logging to a file with rotation")
	// 	if loop == 1000000 {
	// 		break
	// 	}
	// 	loop++
	// }
	//? example 10 - masuk ke server jarak jauh
	con, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer con.Close()
	logger := zerolog.New(con).With().Timestamp().Logger()
	logger.Info().Msg("I am logging to a remote server")
	os.Exit(0)
}
