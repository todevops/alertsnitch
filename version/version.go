package version

import "fmt"

// Name is the application name
const Name = "Alertsnitch"

// Version is the application Version
const Version = "v0.0.1"

// Date is the built date and time
const Date = "2021-03-06"

// GetVersion returns the version as a string
func GetVersion() string {
	return fmt.Sprintf("%s \nVersion: %s \nDate: %s", Name, Version, Date)
}
