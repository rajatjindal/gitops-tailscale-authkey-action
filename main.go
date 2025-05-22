package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
	"tailscale.com/client/tailscale/v2"
)

func main() {
	ctx := context.Background()

	oauthId := os.Getenv("TS_OAUTH_ID")
	oauthSecret := os.Getenv("TS_OAUTH_SECRET")
	tailnet := os.Getenv("TS_TAILNET")
	tags := strings.Split(os.Getenv("TS_AUTHKEY_TAGS"), ",")
	reusable := os.Getenv("TS_AUTHKEY_REUSABLE") == "true"
	preauthorized := os.Getenv("TS_AUTHKEY_PREAUTHORIZED") == "true"
	outputFile := os.Getenv("TS_AUTHKEY_FILE")
	if outputFile == "" {
		outputFile = "tailscale-authkey.txt"
	}

	fmt.Printf(`generating authkey with input params:
tags: %s
reusable: %t
preauthorized: %t
outputFile: %s
\n`, tags, reusable, preauthorized, outputFile)

	oauthConfig := &clientcredentials.Config{
		ClientID:     oauthId,
		ClientSecret: oauthSecret,
		TokenURL:     "https://api.tailscale.com/api/v2/oauth/token",
	}
	client := oauthConfig.Client(ctx)

	tclient := tailscale.Client{
		Tailnet: tailnet,
		HTTP:    client,
	}

	create := struct {
		Reusable      bool     `json:"reusable"`
		Ephemeral     bool     `json:"ephemeral"`
		Tags          []string `json:"tags"`
		Preauthorized bool     `json:"preauthorized"`
	}{
		Reusable:      reusable,
		Preauthorized: preauthorized,
		Tags:          tags,
	}

	devices := struct {
		Create struct {
			Reusable      bool     `json:"reusable"`
			Ephemeral     bool     `json:"ephemeral"`
			Tags          []string `json:"tags"`
			Preauthorized bool     `json:"preauthorized"`
		} `json:"create"`
	}{
		Create: create,
	}

	authkey, err := tclient.Keys().CreateAuthKey(ctx, tailscale.CreateKeyRequest{
		Capabilities: tailscale.KeyCapabilities{
			Devices: devices,
		},
	})
	if err != nil {
		fmt.Printf("failed to create authkey. Error: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(outputFile, []byte(authkey.Key), 0600)
	if err != nil {
		fmt.Printf("failed to save authkey file. Error: %v\n", err)
		os.Exit(1)
	}
}
