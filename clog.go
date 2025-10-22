package clog

import (
	"log"

	"github.com/tobiashort/cfmt-go"
)

func Print(v ...any) {
	log.Print(cfmt.Sprint(v...))
}

func Printf(format string, v ...any) {
	log.Printf(cfmt.Sprintf(format, v...))
}

func Println(v ...any) {
	log.Println(cfmt.Sprintln(v...))
}

func Fatal(v ...any) {
	log.Fatal(cfmt.Sprint(v...))
}

func Fatalf(format string, v ...any) {
	log.Fatalf(cfmt.Sprintf(format, v...))
}

func Fatalln(v ...any) {
	log.Fatalln(cfmt.Sprintln(v...))
}

func Panic(v ...any) {
	log.Panic(cfmt.Sprint(v...))
}

func Panicf(format string, v ...any) {
	log.Panicf(cfmt.Sprintf(format, v...))
}

func Panicln(v ...any) {
	log.Panicln(cfmt.Sprintln(v...))
}
