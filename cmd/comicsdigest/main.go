package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Local()
	nowStr := fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	title := "Comics " + nowStr
	comics(title)
	DeliveryInit()
	Deliver(title, HTMLToString())
}
