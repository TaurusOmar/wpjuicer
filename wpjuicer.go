package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {

	fmt.Println(`
	 __    __          __          _                   
	/ / /\ \ \ _ __    \ \  _   _ (_)  ___   ___  _ __ 
	\ \/  \/ /| '_ \    \ \| | | || | / __| / _ \| '__|
	 \  /\  / | |_) |/\_/ /| |_| || || (__ |  __/| |   
	  \/  \/  | .__/ \___/  \__,_||_| \___| \___||_|   
                  |_|                                      
					  @TaurusOmar_

`)
	url := getUserInput("Enter the URL or host: ")

	detectWordPressVersion(url)
	detectWordPressUsers(url)

	directories := []string{
		"/wp-admin.php",
		"/wp-config.php",
		"/wp-content/uploads",
		"/Wp-load",
		"/wp-signup.php",
		"/Wp-json",
		"/wp-includes/",
		"/wp-login.php",
		"/wp-links-opml.php",
		"/wp-activate.php",
		"/wp-blog-header.php",
		"/wp-cron.php",
		"/wp-links.php",
		"/wp-mail.php",
		"/xmlrpc.php",
		"/wp-settings.php",
		"/wp-trackback.php",
		"/wp-signup.php",
		"/admin-bar.php",
		"/wp-json/wp/v2/users",
		"/wp-json/wp/v2/plugins",
		"/wp-json/wp/v2/themes",
		"/wp-json/wp/v2/comments",
		"/wordpress.tmp",
		"/wordpress.rar",
		"/wordpress.7z",
		"/wordpress.sql",
		"/wordpress.sql.tar.gz",
		"/wordpress.sql.zip",
		"/wordpress.tar.gz",
		"/wordpress.txt.gz",
		"/wordpress.old",
		"/wordpress.zip",
		"/wp.tmp",
		"/wp.rar",
		"/wp.7z",
		"/wp.sql",
		"/wp.sql.tar.gz",
		"/wp.sql.zip",
		"/wp.tar.gz",
		"/wp.txt.gz",
		"/wp.old",
		"/wp.zip",
		"/wp-config.php",
		"/wp-config.php-bak",
		"/wp-config.php.bak",
		"/wp-config.php.new",
		"/wp-config.php.old",
		"/wp-config.php?aam-media=1",
		"/wp-config.php_Old",
		"/wp-config.php_bak",
		"/wp-config.php_new",
		"/wp-config.php_old",
		"/wp-config.php~",
		"/wp-config.bak",
		"/wp-config.bak1",
		"/logs",
		"/.htaccess",
		"/.htaccess-dev",
		"/.htaccess-local",
		"/.htaccess-marco",
		"/.htaccess.BAK",
		"/.htaccess.bak",
		"/.htaccess.bak1",
		"/.htaccess.inc",
		"/.htaccess.old",
		"/.htaccess.orig",
		"/.htaccess.sample",
		"/.htaccess.save",
		"/.htaccess.txt",
		"/.htaccess/",
		"/.htaccess_extra",
		"/.htaccess_orig",
		"/.htaccess_sc",
		"/.htaccessBAK",
		"/.htaccessOLD",
		"/.htaccessOLD2",
		"/.htaccess~",
		"/.htgroup",
		"/.htpasswd",
		"/.htpasswd-old",
		"/.htpasswd.bak",
		"/.htpasswd.inc",
		"/.htpasswd/",
		"/.htpasswd_test",
		"/.htpasswds",
		"/.htpasswrd",
		"/#.htaccess#",
		"/secure/.htaccess",
		"/.htaccess.swp",
	}

	fmt.Println("Directories:")
	for _, directory := range directories {
		fullURL := url + directory
		resp, err := http.Get(fullURL)
		if err != nil {
			fmt.Printf("\033[31mError accessing %s: %s\033[0m\n", fullURL, err)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("\033[32mDirectory %s exists: %s\033[0m\n", directory, fullURL)
		} else {
			fmt.Printf("\033[31mCannot access directory %s: %s\033[0m\n", directory, fullURL)
		}
	}
}

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	// Ensure the URL starts with http:// or https://
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
		input = "http://" + input
	}
	return input
}

func detectWordPressVersion(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("\033[31mError accessing %s: %s\033[0m\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("\033[31mError reading response body: %s\033[0m\n", err)
		return
	}

	re := regexp.MustCompile(`(?i)content="WordPress (\d+\.\d+(\.\d+)?)`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) > 1 {
		fmt.Printf("\033[33mDetected WordPress version: %s\033[0m\n", matches[1])
	} else {
		fmt.Println("\033[31mCould not detect WordPress version.\033[0m")
	}
}

func detectWordPressUsers(url string) {
	resp, err := http.Get(url + "/wp-json/wp/v2/users")
	if err != nil {
		fmt.Printf("\033[31mError accessing %s: %s\033[0m\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("\033[31mCould not fetch user list.\033[0m")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("\033[31mError reading response body: %s\033[0m\n", err)
		return
	}

	re := regexp.MustCompile(`"name":"([^"]+)"`)
	matches := re.FindAllStringSubmatch(string(body), -1)
	if len(matches) > 0 {
		fmt.Println("Detected users:")
		for _, match := range matches {
			fmt.Printf("\033[33m- %s\033[0m\n", match[1])
		}
	} else {
		fmt.Println("\033[31mCould not detect users.\033[0m")
	}
}
