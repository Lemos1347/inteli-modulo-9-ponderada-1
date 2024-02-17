package main

import (
	pub "ponderada-1/publisher"
	sub "ponderada-1/subscriber"
)

func main() {
	pub.PubWg.Add(1)
	go pub.PubMessage("test/topic")

	sub.SubWg.Add(1)
	go sub.RunSub("test/topic")

	sub.SubWg.Wait()
	pub.PubWg.Wait()
}
