package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/deckarep/gosx-notifier"
)

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

	fmt.Println("Mode : ", *environment)
	if *environment == "development" {
		pomodoroTime = time.Second * time.Duration(*pomodoroTimeFlag)
		shortBreakTime = time.Second * time.Duration(*shortBreakTimeFlag)
		longBreakTime = time.Second * time.Duration(*longBreakTimeFlag)
	} else {
		pomodoroTime = time.Minute * time.Duration(*pomodoroTimeFlag)
		shortBreakTime = time.Minute * time.Duration(*shortBreakTimeFlag)
		longBreakTime = time.Minute * time.Duration(*longBreakTimeFlag)
	}

	showNotification("Pomodoro starts now")

	for i := 1; i <= pomodoroCount; i++ {
		fmt.Printf("\npomodoro %d started\n", i)
		pomodoroEndsAt := time.Now().Add(pomodoroTime)
		pomodoroTicker := time.NewTicker(time.Second * 1)
		pomodoroTimer := time.NewTimer(pomodoroTime)
		go PrintTick(pomodoroEndsAt, pomodoroTicker)
		<-pomodoroTimer.C
		pomodoroTicker.Stop()

		if i%4 == 0 {
			fmt.Printf("\n----- long break -----\n")
			showNotification("long break time")
			longBreakEndsAt := time.Now().Add(longBreakTime)
			longBreakTicker := time.NewTicker(time.Second * 1)
			longBreakTimer := time.NewTimer(longBreakTime)
			go PrintTick(longBreakEndsAt, longBreakTicker)
			<-longBreakTimer.C
			longBreakTicker.Stop()
			showNotification("long break ends")
		} else {
			showNotification("short break time")
			fmt.Printf("\n----- short break -----\n")
			shortBreakEndsAt := time.Now().Add(shortBreakTime)
			shortBreakTicker := time.NewTicker(time.Second * 1)
			shortBreakTimer := time.NewTimer(shortBreakTime)
			go PrintTick(shortBreakEndsAt, shortBreakTicker)
			<-shortBreakTimer.C
			shortBreakTicker.Stop()
			showNotification("short break ends")
		}

	}
}

func PrintTick(et time.Time, tick *time.Ticker) {
	for t := range tick.C {
		fmt.Printf("\rDue time : %s", fmtDuration(et.Sub(t)))
	}
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d", m, s)
}

func showNotification(message string) {

	note := gosxnotifier.NewNotification(message)
	note.Title = "Pomodoro notification"
	note.Sound = gosxnotifier.Glass

	//Optionally, an app icon (10.9+ ONLY)
	note.AppIcon = "./icon/pom.png"

	//Optionally, a content image (10.9+ ONLY)
	// note.ContentImage = "./icon/pom.png"

	err := note.Push()

	//If necessary, check error
	if err != nil {
		log.Println("Error while pushing the notifications!")
	}

}
