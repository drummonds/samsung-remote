// This is in transition from being an example to a remote application
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/drummonds/samsung-remote/ws"
)

func main() {
	tv_ip := os.Getenv("TV_IP")
	remote := ws.NewRemote(ws.SamsungRemoteConfig{
		BaseUrl: fmt.Sprintf("wss://%s:8002", tv_ip),
		Name:    "drummonds",
		// Token:   "My_TOKEN",   // You do want to have a token and set the TV
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	token, err := remote.ConnectWithContext(nil, ctx)
	if err != nil {
		panic(fmt.Errorf("problem with remote %e", err))
	}

	fmt.Printf("%s\n", token)

	remote.SendKey("KEY_VOLDOWN")

	remote.OpenBrowser("")
	remote.Move(50, 10)
	remote.SendText([]byte(`www.bbc.co.uk`))
	apps, err := remote.GetInstalledApps()
	if err != nil {
		panic(fmt.Errorf("problem with remote %e", err))
	}
	fmt.Printf("%#v\n", apps)

	for _, app := range apps {
		if app.Name == "YouTube" {
			remote.StartApp(app.Id)
		}
	}

	status, err := remote.GetAppStatus("111299001912")
	if err != nil {
		panic(fmt.Errorf("problem with remote %e", err))
	}
	fmt.Printf("%#v", status)
}
