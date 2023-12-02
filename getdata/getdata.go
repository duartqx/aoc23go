package getdata

import (
	"bufio"
	"os"
)

func GetInputChannel(filename string) (<-chan string, error) {

	ch := make(chan string)

	file, err := os.Open(filename)
	if err != nil {
		return ch, err
	}

	scan := bufio.NewScanner(file)

	go func() {

		defer close(ch)
		defer file.Close()

		for scan.Scan() {
			ch <- scan.Text()
		}
	}()

	return ch, nil
}
