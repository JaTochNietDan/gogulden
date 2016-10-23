package gogulden

// Help will display a list of commands when the parameter cmd is left as an
// empty string. When cmd is not an empty string, help will provide information
// about the provided command, if it exists.
func (c *Client) Help(cmd string) (string, error) {
	var result string
	err := c.runCommand(&result, "help", cmd)
	if err != nil {
		return "", err
	}
	return result, nil
}
