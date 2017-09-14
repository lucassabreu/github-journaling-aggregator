// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/lucassabreu/github-journaling-aggregator/formatter"
	"github.com/lucassabreu/github-journaling-aggregator/githubclient"
	"github.com/lucassabreu/github-journaling-aggregator/report"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	_         = iota
	TODAY int = 1 << (10 * iota)
	YESTERDAY
	LAST_WEEK
	DAYS
	DATE
)

var (
	cfgFile string

	dateFilterType int = TODAY
	today          bool
	yesterday      bool
	lastWeek       bool
	days           int
	date           string

	formatterType string

	username string
	token    string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "github-journaling-aggregator <username>",
	Short: "Create a simple report using your activity feed at GitHub",
	Long: `Create a simple report using your activity feed at GitHub.

	Will receive a username, access token and a beginning date and generate a report based on the users activity feed on GitHub`,
	Args: validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		username = args[0]
		var beginningDate time.Time = time.Now()

		switch dateFilterType {
		case YESTERDAY:
			beginningDate = time.Now().AddDate(0, 0, -1)
		case LAST_WEEK:
			beginningDate = time.Now().AddDate(0, 0, -1*int(time.Now().Weekday()))
		case DAYS:
			beginningDate = time.Now().AddDate(0, 0, days*-1)
		case DATE:
			var err error
			beginningDate, err = time.Parse("2006-01-02", date)
			if err != nil {
				log.Fatal(err)
			}
		}

		y, m, d := beginningDate.Date()
		beginningDate = time.Date(y, m, d, 0, 0, 0, 0, time.Local)

		client := githubclient.NewGithubClient(username, token, nil)
		r := report.New(client, username, beginningDate)

		f, err := getFormatter()
		if err != nil {
			log.Fatal(err)
		}

		r.AttachFormatter(f)
		r.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const (
	FORMAT_DEFAULT    = FORMAT_RAW
	FORMAT_RAW        = "raw"
	FORMAT_TABLE      = "table"
	FORMAT_MD         = "md"
	FORMAT_GROUP_LINE = "groupline"
	FORMAT_CSV        = "csv"
)

var formats = []string{
	FORMAT_RAW,
	FORMAT_TABLE,
	FORMAT_MD,
	FORMAT_GROUP_LINE,
	FORMAT_CSV,
}

func getFormatter() (f report.Formatter, err error) {
	switch formatterType {
	case FORMAT_CSV:
		t := formatter.NewCSV(os.Stdout)
		f = &t
	case FORMAT_GROUP_LINE:
		t := formatter.NewGroupLineTable(os.Stdout)
		f = &t
	case FORMAT_MD:
		t := formatter.NewMDTable(os.Stdout)
		f = &t
	case FORMAT_TABLE:
		t := formatter.NewTable(os.Stdout)
		f = &t
	case FORMAT_RAW:
		r := formatter.NewRaw(os.Stdout)
		f = &r
	default:
		err = fmt.Errorf("Format %s is not valid !", formatterType)
	}
	return
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&formatterType, "format", "f", FORMAT_DEFAULT, fmt.Sprintf(
		"how the events should be displayed, the options are: %s",
		strings.Join(formats, ", "),
	))
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.github-journaling-aggregator.yaml)")
	RootCmd.PersistentFlags().StringVar(&token, "token", "", "github access token (or user password), if not set $GITHUB_TOKEN will be used")
	RootCmd.PersistentFlags().BoolVarP(&today, "today", "t", false, "use today as beginning date (default)")
	RootCmd.PersistentFlags().BoolVarP(&yesterday, "yesterday", "y", false, "use yesterday as beginning date")
	RootCmd.PersistentFlags().BoolVarP(&lastWeek, "last-week", "w", false, "use the last sunday as beginning date")
	RootCmd.PersistentFlags().IntVarP(&days, "days", "d", 0, "use today as beginning date")
	RootCmd.PersistentFlags().StringVar(&date, "date", "", "set a beginning date (format 2017-12-31)")
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if token == "" {
		return fmt.Errorf("token must be informmed or GITHUB_TOKEN environment var set")
	}

	timeParamCount := 0

	if today {
		timeParamCount++
		dateFilterType = TODAY
	}

	if yesterday {
		timeParamCount++
		dateFilterType = YESTERDAY
	}

	if lastWeek {
		timeParamCount++
		dateFilterType = LAST_WEEK
	}

	if days > 0 {
		timeParamCount++
		dateFilterType = DAYS
	}

	if date != "" {
		timeParamCount++
		dateFilterType = DATE
	}

	if timeParamCount > 1 {
		return fmt.Errorf("can't mix the beginning flags")
	}

	if len(args) != 1 || strings.TrimSpace(args[0]) == "" {
		return fmt.Errorf("username is required")
	}
	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".github-journaling-aggregator" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".github-journaling-aggregator")
	}

	if token == "" {
		var ok bool
		if token, ok = os.LookupEnv("GITHUB_TOKEN"); !ok {
			token = ""
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
