package main

import (
	"fmt"
	"time"

	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

func sendNotification(msg string) {

	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/pom.png",
		AppName:     "Pomodoro",
	})

	notify.Push("Notification", msg, "/Users/nikunjshukla/pom.png", notificator.UR_CRITICAL)
}

func main() {
	pomodoroTime := time.Second * 15
	shortBreakTime := time.Second * 5
	longBreakTime := time.Second * 10
	// pomodoroCount := 0
	fmt.Println("Pomodoro starts now")
	for i := 0; i < 3; i++ {
		time.Sleep(pomodoroTime)
		sendNotification("short break time")
		time.Sleep(shortBreakTime)
	}
	time.Sleep(pomodoroTime)
	sendNotification("long break time")
	time.Sleep(longBreakTime)

}

// tick every 25 mins for 4 times
// after every 25 mins tick - sleep 5 mins
