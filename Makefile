.PHONY: run
run:
	@read -p "Enter which day to execute: " day;\
	cd day_$$day; \
	GOOS=windows go run main.go part_1.go part_2.go

.PHONY: run2
run2:
	@read -p "Enter which day to execute: " day;\
	cd day_$$day; \
	GOOS=windows go run main.go part_1.go part_2.go animation.go

.PHONY: generate-day
generate-day:
	@read -p "Enter which day to generate: " day;\
	go run tools/template.go $$day