package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gg",
	Short: "Generates gitignore easily🌈",
	Long: `GG is a CLI tool to help developers generate 
the gitignore which is collected in 'toptal.com'
or create a custom template.`,
	Run: func(cmd *cobra.Command, args []string) {
		info := color.New(color.FgHiBlue).SprintFunc()
		option := color.New(color.Bold).SprintFunc()
		templates, _ := cmd.Flags().GetStringSlice("templates")
		file, _ := cmd.Flags().GetString("file")
		append, _ := cmd.Flags().GetBool("append")
		fmt.Printf("%s %s\n", info("Selected Template:"), option(strings.Join(templates, ", ")))
		fmt.Printf("%s %s\n", info("Selected File    :"), option(file))
		fmt.Printf("%s %s\n", info("Appending Mode   :"), option(strconv.FormatBool(append)))
		fmt.Println()
		process(templates, file, append)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gg.yaml)")

	rootCmd.Flags().StringSliceP("templates", "t", nil, "template name")
	rootCmd.Flags().StringP("file", "f", ".gitignore", "specify filename")
	rootCmd.Flags().BoolP("append", "a", false, "toggle append mode")
}

func initConfig() {
	versionInfo := color.New(color.FgHiYellow, color.Bold).SprintFunc()
	fmt.Printf("%s ⛄️\n\n", versionInfo("Gitignore Generator v0.0.1"))

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".gg")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func process(templates []string, file string, append bool) {
	content := ""
	fmt.Println("🍭 Generating .gitignore")
	fmt.Println()
	for _, t := range templates {
		fmt.Printf("Adding %s.gitignore ...\n", t)
		content += fetchGitIgnore(t)
	}
	saveFile(content, file, append)
	fmt.Println()
	fmt.Println("🍻 Done.")
}

func fetchGitIgnore(templateName string) string {
	API := fmt.Sprintf("https://www.toptal.com/developers/gitignore/api/%s", templateName)
	template := fmt.Sprintf("#### %s ####\n", templateName)

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
	template += string(body)
	template += "\n"
	return template
}

func saveFile(content string, filename string, append bool) {
	var file *os.File
	var err error
	if append {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

	} else {
		file, err = os.Create(filename)
		if err != nil {
			panic(err)
		}
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	if _, err := file.WriteString(content); err != nil {
		panic(err)
	}
	file.Sync()
}
