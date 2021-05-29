package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show available templates.",
	Long:  `This command can show available templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		template, _ := cmd.Flags().GetString("template")
		list(template)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("template", "t", "all", "Show default/custom/all templates.")
}

func list(template string) {
	defaultInfo := color.New(color.FgHiBlue, color.Bold).SprintFunc()
	customInfo := color.New(color.FgHiRed, color.Bold).SprintFunc()
	switch template {
	case "all":
		fmt.Printf("🍻 %s\n", defaultInfo("Default Template"))
		getToptalList()
		fmt.Print("\n\n")
		fmt.Printf("🥂 %s\n", customInfo("Custom Template"))
		getCustomList()
	case "default":
		fmt.Printf("🍻 %s\n", defaultInfo("Default Template"))
		getToptalList()
	case "custom":
		fmt.Printf("🥂 %s\n", customInfo("Custom Template"))
		getCustomList()
	default:
		fmt.Println("Wrong")
	}
}

func getToptalList() {
	API := "https://www.toptal.com/developers/gitignore/api/list"
	resp, http_err := http.Get(API)
	if http_err != nil {
		fmt.Println("Error: An error occurs in the http connection.")
		fmt.Println(http_err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		fmt.Println("Error: An error occurs in I/O.")
		fmt.Println(read_err)
		os.Exit(1)
	}

	types := strings.Split(strings.Join(strings.Split(string(body), "\n"), ","), ",")
	for _, t := range types {
		fmt.Println("➢", t)
	}
}

func getCustomList() {
	fmt.Println("No custom templates.")
}
