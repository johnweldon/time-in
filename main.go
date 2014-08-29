package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var MainCmd = &cobra.Command{
	Use:   "time-in",
	Short: "time-in shows the current time in some pre-configured locations",
	Run:   timeIn,
}

var DefaultShow = []string{
	"Local",
	"UTC",
}

func main() {
	viper.SetConfigName(".time-in")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	viper.ReadInConfig()
	MainCmd.Execute()
}

func timeIn(cmd *cobra.Command, args []string) {
	Show := DefaultShow
	if zones := viper.GetStringSlice("timezones"); len(zones) > 0 {
		Show = zones
	}
	for _, z := range Show {
		loc, err := time.LoadLocation(z)
		if err != nil {
			continue
		}
		t := time.Now()
		in := t.In(loc)
		fmt.Fprintf(os.Stdout, "%20s %s\n", z, in.Format("2006-01-02 15:04:05 -0700 MST"))
	}
}
