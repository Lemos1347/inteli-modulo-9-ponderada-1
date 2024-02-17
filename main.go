package main

import (
	pub "ponderada-1/publisher"
	sub "ponderada-1/subscriber"
)

func main() {
	pub.PubWg.Add(1)
	go pub.PubMessage("sensors/solar_sensor")

	sub.SubWg.Add(1)
	go sub.RunSub("sensors/solar_sensor")

	sub.SubWg.Wait()
	pub.PubWg.Wait()
}
