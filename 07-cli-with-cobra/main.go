package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCommand *cobra.Command

func init() {
	helloCommand = &cobra.Command{
		Use:   "hello",
		Short: "Print the hello world",
		Run:   sayHello,
	}

	helloCommand.Flags().StringP("name", "n", "World", "Who to say hello to")
	helloCommand.MarkFlagRequired("name")
	helloCommand.Flags().StringP("language", "l", "en", "Which language to say hello in")
}

func sayHello(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	greeting := "Hello"
	language, _ := cmd.Flags().GetString("language")

	switch language {
	case "en":
		greeting = "Hello"
	case "es":
		greeting = "Hola"
	case "fr":
		greeting = "Bonjour"
	case "de":
		greeting = "Hallo"
	}

	fmt.Printf("%s %s!\n", greeting, name)
}

func main() {
	helloCommand.Execute()
}
