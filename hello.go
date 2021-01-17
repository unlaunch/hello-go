package main

import (
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/client"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"time"
)

func main() {
	const apiKey = "prod-server-7f05812f-cd5a-41dd-b79b-6f30d84cd335"
	const flagKey = "mfa-feature"

	cfg := client.DefaultConfig()
	cfg.LoggerConfig = &logger.LogOptions{
		Level:    "INFO",
		Colorful: true,
	}
	factory, err := client.NewUnlaunchClientFactory(apiKey, cfg)

	if err != nil {
		fmt.Printf("Unable to initialize Unlaunch Client because there was an error %s\n", err)
		return
	}

	unlaunchClient := factory.Client()

	if err = unlaunchClient.BlockUntilReady(5 * time.Second); err != nil {
		fmt.Printf("Unlaunch Client isn't ready %s\n", err)
	}

	variation := unlaunchClient.Variation(flagKey, "user123", nil)
	fmt.Printf("The variation for feature is: %s\n", variation)

	// Let's pass some attributes
	attr := make(map[string]interface{})
	attr["registered"] = true

	variation = unlaunchClient.Variation(flagKey, "user123", attr)
	fmt.Printf("The variation for feature with registered=true is: %s\n", variation)

	unlaunchClient.Shutdown()

	fmt.Printf("Client shutdown successfully? %v", unlaunchClient.IsShutdown())
}
