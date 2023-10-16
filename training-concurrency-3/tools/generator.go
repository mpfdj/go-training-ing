package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var states = []string{"INFO", "WARNING", "ERROR", "NOTICE", "DEBUG", "TRACE", "FATAL", "CRITICAL", "ALERT", "EMERGENCY", "PANIC"}
var descriptions = []string{
	"Application started successfully",
	"Low disk space on drive C:",
	"Failed to connect to database",
	"User JohnD logged in",
	"High memory usage detected",
	"Timeout while connecting to server",
	"Scheduled maintenance started",
	"Unusual network activity detected",
	"Disk write error on drive D:",
	"User JaneD logged out",
	"Server reboot required",
	"Application update available",
	"Critical system error",
	"New user registered",
	"Database backup completed",
	"SSL certificate expired",
	"Invalid login attempt",
	"File not found",
	"Server overloaded",
	"DNS resolution failed",
	"Invalid request received",
	"Unauthorized access attempt",
	"Service unavailable",
	"Invalid configuration detected",
	"System shutdown initiated",
	"Unexpected system restart",
	"Invalid input received",
	"File upload failed",
	"Invalid file format",
	"Invalid parameter provided",
}

func main() {
	names := []string{"log.txt"}

	for _, name := range names {
		file, err := os.Create(name)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		for i := 0; i < 50000; i++ {
			timestamp := time.Now().
				Add(-time.Duration(r.Intn(720)) * time.Hour).
				Format("2006-01-02 15:04:05")
			state := states[r.Intn(len(states))]

			var description string

			for j := 0; j < 100; j++ {
				description += descriptions[r.Intn(len(descriptions))]

				if j < 100 {
					description += " "
				}
			}

			_, err = file.Write([]byte(fmt.Sprintf("%s %s %s\n", timestamp, state, description)))
			if err != nil {
				panic(err)
			}
		}
	}
}
