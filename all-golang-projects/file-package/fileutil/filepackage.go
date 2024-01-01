package file

import (
	"fmt"
	"os"
)

func Read(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}
