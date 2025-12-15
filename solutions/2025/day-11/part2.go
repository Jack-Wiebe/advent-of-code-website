package main

import (
	"advent-of-code-website/solutions/utils"
	"strings"
)

func Part2() (int,error) {

	scanner, err := utils.ReadFile("")
	if err != nil {
			return -1 , err
	}

	serverMap := make(map[string] []string)


	for scanner.Scan(){

		line := scanner.Text()

		strArr := strings.Split(line, ": ")
		key:= strArr[0]
		values:=strings.Split(strArr[1], " ")

		serverMap[key] = values

	}

	rack := ServerRack{serverMap, map[Device]int{}}
	count := rack.recursiveCaching(Device{name:"svr", mask: 0})

	return count, nil
}

const (
	NONE = 0
	FFT  = 1 << 0
	DAC  = 1 << 1
	BOTH = FFT | DAC
)

type Device struct {
	name string
	mask int
}

type ServerRack struct {
	serverMap map[string][]string
	cache     map[Device]int
}

func (rack *ServerRack) recursiveCaching(current Device) (int) {

	if current.name == "out" {
		if(current.mask == BOTH){
			return 1
		}else{
			return 0
		}
	}

	if current.name == "fft" {
		current.mask |= FFT
	} else if current.name == "dac" {
		current.mask |= DAC
	}

	sum:=0
	for _, next := range rack.serverMap[current.name] {
		d := Device{next, current.mask}
		count, cached := rack.cache[d]
		if !cached {
			count = rack.recursiveCaching(d)
			rack.cache[d] = count
		}
		sum += count
	}
	return sum

}