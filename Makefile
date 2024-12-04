.PHONY: run
run:
	@read -p "Enter which day to execute: " day;\
	go run main.go $$day