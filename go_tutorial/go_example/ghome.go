package main

import (
	"github.com/evalphobia/google-home-client-go/googlehome"
)

func main() {
	cli, err := googlehome.NewClientWithConfig(googlehome.Config{
		Hostname: "192.168.3.4",
		Lang:     "en",
		Accent:   "GB",
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// Speak text on Google Home.
	cli.Notify("Hello")

	// Change language
	cli.SetLang("ja")
	cli.Notify("こんにちは、グーグル。")

	// Or set language in Notify()
	//cli.Notify("你好、Google。", "zh")
}