package main

import (
	"fmt"
	"time"
)

const ()

func main() {
	fmt.Println(AnniversaryDate())
}

func Schedule(date string) time.Time {
	t, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t
}

func HasPassed(date string) bool {
	t, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	return t.Before(time.Now())
}

func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 19:04:05", date)

	if err != nil {
		panic(err)
	}

	return (t.Hour() >= 12 || t.Hour() < 16)
}

func Description(date string) string {
	t, err := time.Parse("1/02/2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("You have an appointment at %s", t.Format("Monday, January 02, 2006, at 15:04"))

}

func AnniversaryDate() time.Time {
	t, _ := time.Parse("2006-01-02 00:00:00", fmt.Sprintf("%d-09-15 00:00:00", time.Now().Year()))

	return t
}
