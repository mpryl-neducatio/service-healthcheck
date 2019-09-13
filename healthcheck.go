package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
	"os"
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

func waitForServices(services []Service, waitInterval int) {
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

func isHealthy(host string, port int) bool {
	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	return err == nil
}

func sendHealthCheckExitCode(services []Service) {
	for _, service := range services {
		if !isHealthy(service.Host, service.Port) {
			os.Exit(1)
		}
	}
	os.Exit(0)
}

func readCmdLineFlags() (string, int, string){
	configFileFlag := flag.String("f", "services.yml", "Config file path")
	waitIntervalFlag := flag.Int("i", 5, "Wait interval in seconds")
	modeFlag := flag.String("m", "wait", "Health check mode (wait|exit_code)")
	flag.Parse()

	return *configFileFlag, *waitIntervalFlag, *modeFlag
}

func main() {
	configFilePath, waitInterval, mode := readCmdLineFlags()
	fmt.Printf("Reading services config from file: %s\n", configFilePath)
	services := readConfig(configFilePath)
	if mode == "wait" {
		waitForServices(services, waitInterval)
	} else if mode == "exit_code" {
		sendHealthCheckExitCode(services)
	} else {
		log.Fatalf("Incorrect mode %s", mode)
	}
}
