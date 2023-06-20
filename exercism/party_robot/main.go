package main

import (
	"fmt"
)

const ()

func main() {
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s", name)
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, table_number int, name_seatmate string, direction string, distance float64) string {
	return fmt.Sprintf(`Welcome to my party, %s!
You have been assigned to table %d. Your table is %s, exactly %f meters from here.
You will be sitting next to %s`, name, table_number, direction, distance, name_seatmate)
}
