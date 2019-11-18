package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

const fileContent = `println("Hi);
println("This");
println("is")
println("Just");
println("a");
println("Generated");
println("Content");`

// Point ...
type Point struct {
	X int
	Y int
}

// ReadFromFile 51 columns x 7 lines of 1, will ignore any char except 1, will ignore char outside col [0-50] and row [0-6]
func ReadFromFile(filePath string) ([]Point, error) {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	dataStr := string(data)

	x, y := 0, 0
	points := []Point{}
	for c := range dataStr {
		char := dataStr[c]
		if char == '\n' || x >= panelWidth {
			x = 0
			y++
			continue
		} else if char == '1' {
			p := Point{
				X: x,
				Y: y,
			}
			points = append(points, p)
		} else if y >= panelHeight {
			break
		}
		x++
	}

	return points, nil
}

// GenerateFile TODO: actually check how github implement the intensity level
func GenerateFile(basePath string, idx int) (string, error) {
	filePath := fmt.Sprintf("%s/file-%d.txt", basePath, idx)

	err := ioutil.WriteFile(filePath, []byte(fileContent), os.FileMode(0755))

	return filePath, err
}
