package client

type Sso struct {
	Sso Client
}

type test struct {
	ID int
}

func (c *Client) IsLogin() error {
	body, _, err := c.do("GET", "/dockerapi/test", nil, false)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var t test
	err = json.Unmarshal(body, &t)
	fmt.Println(t.ID)
	return err
}

func main() {
	endpoint := "http://127.0.0.1:9000"
	c, err := NewClient(endpoint)
	c.IsLogin()
	fmt.Println(err)
	fmt.Println("Hello World!")
}
