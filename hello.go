package main

import (
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/client"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"time"
)

func main() {
	helloGoSimple()
	helloGoAdvanced()
	helloGoAttributes()
}

func helloGoSimple() {
	fmt.Println("-- Running Hello Go Simple --")
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

	if err = unlaunchClient.AwaitUntilReady(5 * time.Second); err != nil {
		fmt.Printf("Unlaunch Client isn't ready %s\n", err)
	}

	variation := unlaunchClient.Variation(flagKey, "user123", nil)
	fmt.Printf("The variation for feature is: %s\n", variation)

	unlaunchClient.Shutdown()

	fmt.Printf("Client shutdown successfully? %v\n\n", unlaunchClient.IsShutdown())
}

func helloGoAdvanced() {
	fmt.Println("-- Running Hello Go Advanced which passes attributes and prints evaluation reason --")
	const apiKey = "prod-server-7f05812f-cd5a-41dd-b79b-6f30d84cd335"
	const flagKey = "mfa-feature"

	cfg := client.DefaultConfig()
	cfg.LoggerConfig = &logger.LogOptions{
		Level:    "INFO",
		Colorful: true,
	}

	factory, _ := client.NewUnlaunchClientFactory(apiKey, cfg)

	unlaunchClient := factory.Client()

	if err := unlaunchClient.AwaitUntilReady(5 * time.Second); err != nil {
		fmt.Printf("Unlaunch Client isn't ready %s\n", err)
	}

	// Let's use the Feature(...) function to get both variation and evaluation reason
	// Let's pass some attributes
	attr := make(map[string]interface{})
	attr["registered"] = true

	feature := unlaunchClient.Feature(flagKey, "user123", attr)
	fmt.Printf("The variation for feature is: %s. Evaluation reason is: %s\n",
		feature.Variation, feature.EvaluationReason)

	unlaunchClient.Shutdown()

	fmt.Printf("Client shutdown successfully? %v\n\n", unlaunchClient.IsShutdown())
}

func helloGoAttributes() {
	fmt.Println("-- Running Hello Go Attributes which passes attributes and prints evaluation reason --")
	const apiKey = "prod-server-7f05812f-cd5a-41dd-b79b-6f30d84cd335"
	const flagKey = "mfa-feature"

	cfg := client.DefaultConfig()
	cfg.LoggerConfig = &logger.LogOptions{
		Level:    "INFO",
		Colorful: true,
	}

	factory, _ := client.NewUnlaunchClientFactory(apiKey, cfg)

	unlaunchClient := factory.Client()

	if err := unlaunchClient.AwaitUntilReady(5 * time.Second); err != nil {
		fmt.Printf("Unlaunch Client isn't ready %s\n", err)
	}

	// Let's pass some attributes
	attr := make(map[string]interface{}) // map to store all attributes

	attr["registered"] = true
	attr["device"] = "iphone"
	attr["age"] = 30
	attr["startDate"] = time.Now().UTC().Unix()

	// let's make a "Set" to store some user Id
	// We use map because "Go" doesn't have a set type. Only keys are used. The value could be anything and is ignored
	userIDs := make(map[string]bool)
	userIDs["user123@gmail.com"] = true
	userIDs["userabc@gmail.com"] = true

	feature := unlaunchClient.Feature(flagKey, "user123", attr)
	fmt.Printf("The variation for feature is: %s\n", feature.Variation)

	unlaunchClient.Shutdown()

	fmt.Printf("Client shutdown successfully? %v\n\n", unlaunchClient.IsShutdown())
}

func helloGoConfiguration() {
	fmt.Println("-- Running Hello Go Attributes which passes attributes and prints evaluation reason --")
	const apiKey = "prod-server-7f05812f-cd5a-41dd-b79b-6f30d84cd335"
	const flagKey = "mfa-feature"

	cfg := client.DefaultConfig()
	// Customize default client
	cfg.PollingInterval = 15 * time.Second // How often flags are fetched
	cfg.HTTPTimeout = 3 * time.Second
	cfg.MetricsFlushInterval = 30 * time.Second // How often metrics are sent
	cfg.MetricsQueueSize = 500

	factory, _ := client.NewUnlaunchClientFactory(apiKey, cfg)

	unlaunchClient := factory.Client()

	if err := unlaunchClient.AwaitUntilReady(5 * time.Second); err != nil {
		fmt.Printf("Unlaunch Client isn't ready %s\n", err)
	}

	// Let's pass some attributes
	attr := make(map[string]interface{}) // map to store all attributes

	attr["registered"] = true
	attr["device"] = "iphone"
	attr["age"] = 30
	attr["startDate"] = time.Now().UTC().Unix()

	// let's make a "Set" to store some user Id
	// We use map because "Go" doesn't have a set type. Only keys are used. The value could be anything and is ignored
	userIDs := make(map[string]bool)
	userIDs["user123@gmail.com"] = true
	userIDs["userabc@gmail.com"] = true

	feature := unlaunchClient.Feature(flagKey, "user123", attr)
	fmt.Printf("The variation for feature is: %s\n", feature.Variation)

	unlaunchClient.Shutdown()

	fmt.Printf("Client shutdown successfully? %v\n\n", unlaunchClient.IsShutdown())
}