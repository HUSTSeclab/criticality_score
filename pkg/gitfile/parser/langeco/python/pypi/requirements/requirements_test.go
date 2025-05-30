package requirements

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, _ := os.Open("")
	defer file.Close()
	data, _ := io.ReadAll(file)
	t.Run("Parse requirements.txt", func(t *testing.T) {
		pkg, deps, err := Parse(string(data))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(*pkg)
		for index, dep := range *deps {
			fmt.Println(index, dep)
		}
	})
}
