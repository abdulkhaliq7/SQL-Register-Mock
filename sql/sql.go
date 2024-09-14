package sql

import "fmt"

type driver interface {
	Open(name string)
}

var drivers = make(map[string]driver)

func Register(name string, driver driver) {

	drivers[name] = driver
}

func Open(name string) {

	driveri, ok := drivers[name]

	if !ok {
		fmt.Println("Unknown Driver", driveri)
		return
	}

	driveri.Open(name)

}