package main

import (
	"fmt"
	"log"
	"os"
)

const MAIN_FILE_TEMPLATE = `package main

import _ "embed"

//go:embed input
var input string

func main() {
	Part1()
	Part2()
}
`

const TEST_FILE_TEMPLATE = `package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	result := Part1()
	assert.Equal(t, result, 1473620)
}

func TestPart2(t *testing.T) {
	result := Part2()
	assert.Equal(t, result, 902620)
}
`

func getPartFileTemplate(part int) string {
	return fmt.Sprintf(`package main

import (
	"log"
	"time"
)

func Part%d() int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	return 0
}
`, part)
}

func directoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func main() {
	args := os.Args[1:]

	day := "day_" + args[0]
	if !directoryExists(day) {
		err := os.Mkdir(day, 0755)
		if err != nil {
			log.Fatal(err)
		}

		// Create input file
		err = os.WriteFile(day+"/input", []byte(""), 0644)
		if err != nil {
			log.Fatal(err)
		}

		// Create main.go
		err = os.WriteFile(day+"/main.go", []byte(MAIN_FILE_TEMPLATE), 0644)
		if err != nil {
			log.Fatal(err)
		}

		// Create test.go
		err = os.WriteFile(day+"/"+day+"_test.go", []byte(TEST_FILE_TEMPLATE), 0644)
		if err != nil {
			log.Fatal(err)
		}

		// Create part_1.go
		err = os.WriteFile(day+"/part_1.go", []byte(getPartFileTemplate(1)), 0644)
		if err != nil {
			log.Fatal(err)
		}

		// Create part_2.go
		err = os.WriteFile(day+"/part_2.go", []byte(getPartFileTemplate(2)), 0644)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal("Directory already exists")
	}
}
