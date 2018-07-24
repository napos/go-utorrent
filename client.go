package utorrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
)

type Client struct {
	API      string
	Username string
	Password string

	token      string
	user_agent *http.Client
}

func (c *Client) setToken() error {
	httpResp, err := c.get("/token.html", nil)
	if err != nil {
		return fmt.Errorf("Error retrieving token: %s", err.Error())
	}

	defer httpResp.Body.Close()

	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		return fmt.Errorf("Error reading token: %s", err.Error())
	}

	if httpResp.StatusCode != 200 {
		if httpResp.StatusCode == 401 {
			return fmt.Errorf("Unable to create client: Incorrect Username or Password")
		}
		return fmt.Errorf("Error %d Unable to create client: %s", httpResp.StatusCode, body)
	}

	re := regexp.MustCompile("(?:<div.*>)(.*)(?:</div>)")
	c.token = re.FindStringSubmatch(string(body))[1]

	return nil
}

func NewClient(c *Client) (*Client, error) {
	cookieJar, _ := cookiejar.New(nil)

	c.user_agent = &http.Client{
		CheckRedirect: nil,
		Jar:           cookieJar,
	}

	if c.API == "" {
		c.API = "http://localhost:8085/gui"
	}

	err := c.setToken()
	if err != nil {
		return nil, err
	}

	return c, nil
}
