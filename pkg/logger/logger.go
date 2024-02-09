package logger

import (
	"fmt"
	"time"
)

const (
	Fatal = 0
	Panic = 1
	Warn  = 2
	Info  = 3
)

type LogMessage struct {
	Lvl  int
	Tag  string
	Msg  string
	Date string
}

type Logger struct {
	Tag string
}

func (l *Logger) Log(msg string, lvl int) LogMessage {
	m := LogMessage{
		Lvl:  lvl,
		Tag:  l.Tag,
		Msg:  msg,
		Date: time.Now().String(),
	}

	log := fmt.Sprintf("[%d][%s]: \t %s \t %s", m.Lvl, l.Tag, msg, time.Now().String())
	fmt.Println(log)
	return m
}

func (l *Logger) Error(msg string) {
	log := fmt.Sprintf("[%s]: \t %s \t %s", l.Tag, msg, time.Now().String())
	fmt.Println(log)
}

func GetLogger(tag string) *Logger {
	return &Logger{Tag: tag}
}
