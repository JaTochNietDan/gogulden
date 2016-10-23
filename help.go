package gogulden

func (c *Client) Help(cmd string) (string, error) {
	var result string
	err := c.runCommand(&result, "help", cmd)
	if err != nil {
		return "", err
	}
	return result, nil
}
