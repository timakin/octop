package client

import (
	"bufio"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func getToken() string {
	const tokenFile = "/var/tmp/op_token"
	const tokenMatcher = "([a-z0-9]{40})"
	_, err := os.Stat(tokenFile)
	var token string
	if err != nil {
		tokenScanning := true
		counter := 0
		r, _ := regexp.Compile(tokenMatcher)
		for tokenScanning && counter < 3 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("> Enter your access token: ")
			token, _ = reader.ReadString('\n')
			if token == "exit" {
				os.Exit(0)
			}
			if r.MatchString(token) {
				tokenScanning = false
			} else {
				fmt.Print("Token format is invalid!\n")
			}
			counter += 1
			if counter == 3 {
				fmt.Print("Input is invalid for 3 times. Recheck your access token.\n")
				fmt.Print("(cf. https://help.github.com/articles/creating-an-access-token-for-command-line-use/)\n")
				os.Exit(0)
			}
		}

		fout, err := os.Create(tokenFile)
		if err != nil {
			fmt.Print("[Error] Couldn't create token record file.\n")
			fmt.Print(err)
			os.Exit(0)
		}
		fout.WriteString(token)
		defer fout.Close()

		return token
	} else {
		f, err := ioutil.ReadFile(tokenFile)
		if err != nil {
			log.Fatal(err)
		}
		token := string(f)
		return token
	}
}

func newAuthenticatedClient() *http.Client {
	formattedToken := strings.Replace(getToken(), "\n", "", -1)
	TokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: formattedToken},
	)
	TokenClient := oauth2.NewClient(oauth2.NoContext, TokenSource)
	return TokenClient
}
