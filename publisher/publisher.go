package publisher

import (
	"fmt"
	"math/rand"
	"ponderada-1/publisher/sensors"
	"sync"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var PubWg = sync.WaitGroup{}

func randSleep() {
	sleepTime := rand.Intn(5) + 1
	time.Sleep(time.Duration(sleepTime) * time.Second)
}

func PubMessage(topic string) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		solarReading, err := sensors.GenerateReading("publisher/sensors/dados_sensor_radiacao_solar.csv")
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
