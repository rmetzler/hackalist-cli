package cmd

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

const API_URL = "http://www.hackalist.org/api/1.0/2019/01.json"

func listHackathons() {

	resp, err := http.Get(API_URL)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	respBody, err:= ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}

var hackathonCmd = &cobra.Command{
	Use: "hackathons",
	Short: "Lists all the hackathons available",
	Run: func(cmd *cobra.Command, args []string) {
		listHackathons()
	},
}

func init() {
	rootCmd.AddCommand(hackathonCmd)
}