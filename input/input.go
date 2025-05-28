package input

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("Couldn't get input")
	}
	return strings.TrimSpace(strings.TrimSuffix(input, "\n")), nil
}
