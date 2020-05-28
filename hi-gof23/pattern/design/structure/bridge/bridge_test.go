package bridge

import "testing"

func TestDemoBridge(t *testing.T) {
	auto := Auto{}
	manual := Manual{}
	amt := AMT{}

	bmw := BMWCar{}
	bmw.SetTransmission(auto)
	bmw.run()

	bmw.SetTransmission(manual)
	bmw.run()

	bmw.SetTransmission(amt)
	bmw.run()

	benz := BenZCar{}
	benz.SetTransmission(auto)
	benz.run()
}
