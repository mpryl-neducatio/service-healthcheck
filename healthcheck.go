package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"gopkg.in/yaml.v2"
	"time"
)

type Service struct {
	Name string
	Host string
	Port int
}

func readConfig(configFilePath string) (services []Service) {
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &services)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return
}

func waitForService(name string, host string, port int, c chan int, waitInterval int) {
	counter := 0
	for {
		_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			time.Sleep(time.Duration(waitInterval) * 1000 * time.Millisecond)
			counter += waitInterval
			fmt.Printf("Waiting for %s for %d seconds...\n", name, counter)
		} else {
			fmt.Printf("%s is ready after %d seconds.\n", name, counter)
			c <- 1
			break
		}
	}
}

func readCmdLineFlags() (string, int){
	configFileFlag := flag.String("f", "services.yml", "Config file path")
	waitIntervalFlag := flag.Int("i", 5, "Wait interval in seconds")
	flag.Parse()

	return *configFileFlag, *waitIntervalFlag
}

func main() {
	configFilePath, waitInterval := readCmdLineFlags()
	services := readConfig(configFilePath)
	counter := 0
	c := make(chan int, len(services))
	for _, service := range services {
		go waitForService(service.Name, service.Host, service.Port, c, waitInterval)
	}
	for i := range c {
		counter += i
		if counter == len(services) {
			close(c)
		}
	}
}
