package main

import (
	"fmt"
	"math/rand"
	"os"
	"publisher/sensors"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func randSleep() {
	sleepTime := rand.Intn(5) + 1
	time.Sleep(time.Duration(sleepTime) * time.Second)
}

func pubMessage(topic string, csvPath string) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		solarReading, err := sensors.GenerateReading(csvPath)
		if err == nil {
			randSleep()
			fmt.Printf("\033[33mDado lido: %s  \033[0m\n", solarReading)
			token := client.Publish(topic, 1, false, solarReading)
			token.Wait()
		} else {
			fmt.Printf("\033[31m%s\033[0m\n", err.Error())
			break
		}
	}
	fmt.Println("\033[35mPublisher encerrado! \033[0m")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\033[31mMissing csv path\033[0m")
		os.Exit(1)
	}
	pubMessage("sensors/solar_sensor", os.Args[1])
}
