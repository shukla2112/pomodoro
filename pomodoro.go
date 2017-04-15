package main

import (
	"flag"
	"time"

	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

func sendNotification(msg string) {

	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/pom.png",
		AppName:     "Pomodoro",
	})

	notify.Push("", msg, "/Users/nikunjshukla/Desktop/pom.png", notificator.UR_CRITICAL)
}

func main() {

	var pomodoroCount int
	var pomodoroTime, shortBreakTime, longBreakTime time.Duration
	//flags
	pomodoroTimeFlag := flag.Int("pomodorotime", 25, "time for each pomodoro (in mins)")
	shortBreakTimeFlag := flag.Int("short", 5, "time for short break (in mins)")
	longBreakTimeFlag := flag.Int("long", 15, "time for long break (in mins)")
	environment := flag.String("env", "development", "runtime environment")
	flag.IntVar(&pomodoroCount, "count", 100, "number of pomodro sessions")

	flag.Parse()

	if *environment == "development" {
		pomodoroTime = time.Second * time.Duration(*pomodoroTimeFlag)
		shortBreakTime = time.Second * time.Duration(*shortBreakTimeFlag)
		longBreakTime = time.Second * time.Duration(*longBreakTimeFlag)
	} else {
		pomodoroTime = time.Minute * time.Duration(*pomodoroTimeFlag)
		shortBreakTime = time.Minute * time.Duration(*shortBreakTimeFlag)
		longBreakTime = time.Minute * time.Duration(*longBreakTimeFlag)
	}

	sendNotification("Pomodoro starts now")
	for i := 1; i <= pomodoroCount; i++ {
		time.Sleep(pomodoroTime)
		if i%4 == 0 {
			sendNotification("long break time")
			time.Sleep(longBreakTime)
			sendNotification("long break ends")
		}
		sendNotification("short break time")
		time.Sleep(shortBreakTime)
		sendNotification("short break ends")
	}
}
