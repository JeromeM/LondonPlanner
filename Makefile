@PHONY: sanitizer

run_dev: sanitizer_dev
	uvicorn main:app --reload

run: sanitizer_build sanitizer
	uvicorn main:app

sanitizer_build: sanitizer_dev
	# ldflags :
	# -s Omit the symbol table and debug information
	# -w Omit the DWARF symbol table
	# These two parameter are used to reduce binary size
	@go build -o sanitizer/build/sanitize -ldflags="-s -w" -v sanitizer/main.go

sanitizer_dev:
	@go run sanitizer/main.go

sanitizer:
	@./sanitizer/build/sanitize