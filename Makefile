.PHONY: run
run:
	@read -p "Enter which day to execute: " day;\
	go run main.go $$day

.PHONY: test
test:
	go run main.go 10-animate

define TEMPLATE
package day$(day)

import (
	"github.com/rrichy/advent-of-code-2024/utils"
)

func Part$${part}() int {
	input := utils.ReadFile("day_${day}/input.txt")

	return 0
}
endef

define TEST_TEMPLATE
package day${day}

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	result := Part1()
	assert.Equal(t, result, 0)
}

func TestPart2(t *testing.T) {
	result := Part2()
	assert.Equal(t, result, 0)
}
endef

# eval echo "$TEMPLATE" > day_$$day/part_$$part.go ; \
# part=2; \
# eval echo "$$TEMPLATE" > day_$$day/part_$$part.go; \
# eval echo "$$TEST_TEMPLATE" > day_$$day/part$$part_test.go; \
# touch day_$$day/input day_$$day/poem

.PHONY: create
create:
	@read -p "Enter which day to create: " day;\
	mkdir -p day_$$day;
	TEMPLATE=$$(cat <<EOF \
package day\
\
import ( \
	"github.com/rrichy/advent-of-code-2024/utils" \
)\
\
func Part() int { \
	input := utils.ReadFile("day_/input.txt") \
\
	return 0\
}\
	EOF\
	); \
	echo "$$TEMPLATE"