package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "now [timestamp]\n  now [YYYY-MM-DD]\n  now [YYYY-MM-DD H:i]\n  now [YYYY-MM-DD H:i:s]",
	DisableFlagsInUseLine: true,
	Short:                 "A simple cli tool to convert unix timestamps or human readable dates.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println(time.Now().Unix())
		}
		if len(args) >= 1 {
			ts, err := strconv.ParseInt(args[0], 10, 64)
			if err == nil {
				switch len(args[0]) {
				case 16:
					fmt.Println(time.Unix(0, ts*int64(time.Microsecond)).Format("2006-01-02 15:04:05"))
				case 13:
					fmt.Println(time.Unix(0, ts*int64(time.Millisecond)).Format("2006-01-02 15:04:05"))
				case 10:
					fmt.Println(time.Unix(0, ts*int64(time.Second)).Format("2006-01-02 15:04:05"))
				default:
					return errors.New("timestamps must be 10, 13, or 16 integers.")
				}
			} else {
				input := strings.Join(args, " ")
				defaultLayout := []string{
					"2006-01-02",
					"2006-01-02 15:04",
					"2006-01-02 15:04:05",
				}
				for _, l := range defaultLayout {
					utime, err := time.ParseInLocation(l, input, time.Local)
					if err == nil {
						fmt.Println(utime.Unix())
						os.Exit(0)
					}
				}
				fmt.Println("Date format error, for example:")
				for _, l := range defaultLayout {
					fmt.Println("now " + l)
				}
			}
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
