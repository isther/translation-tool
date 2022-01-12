package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/isther/translation-tool/client"
	"github.com/isther/translation-tool/config"
	"github.com/spf13/cobra"
)

var (
	version = "v1.0.0"

	rootCmd = &cobra.Command{
		Use:     "tsl",
		Version: version,
		Short:   "A cmd tool for translate",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				client.PrintDaily()
				return
			}

			if config.Instance.AppKey == "" {
				color.Red("请先完成配置(~/.translate/config)\n")
			}

			text := ""
			for _, v := range args {
				text = strings.Join([]string{text, v}, " ")
			}

			trans(text)
		},
	}
)

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return nil
}

func trans(text string) {
	c := client.NewRequest(text)
	result, err := c.Post()
	if err != nil {
		color.Red("Error: %v", err)
		return
	}

	if result.ErrorCode != "0" {
		color.Red("Error: %v\nCode: %v", result.ErrorCode)
		return
	}

	result.PrintQuery()
	result.PrintTranslate()
	result.PrintWeb()
}
