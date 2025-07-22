KEY = 1234ABCD$%^&(-+)
STR_INPUT = HelloWorld
FILES = test1.txt,test2.txt

# Build the binary
build:
	go build -o crypto-cli main.go

# Clean build artifacts
clean: 
	rm -f crypto-cli

# Encrypt a string
encrypt-string:
	go run main.go run --mode=encrypt --type=string --input="$(STR_INPUT)" --key="$(KEY)"

# Decrypting string
decrypt-string:
	go run main.go run --mode=encrypt --type=string --input="REPLACE_WITH_ENCRYPTED_STRING" --key="$(KEY)"

# Encrypting Files (non-concurrent)
encrypt-files:
	go run main.go run --mode=encrypt --type=file --input=$(FILES) --key="$(KEY)"

# Decrypting Files (non-concurrent)
decrypt-files:
	go run main.go run --mode=decrypt --type=file --input=$(FILES) --key="$(KEY)"

# Encrypting Files (non-concurrent)
encrypt-files-concurrent:
	go run main.go run --mode=encrypt --type=file --input=$(FILES) --key="$(KEY)" --concurrent=true

# Decrypting Files (non-concurrent)
decrypt-files-concurrent:
	go run main.go run --mode=decrypt --type=file --input=$(FILES) --key="$(KEY)" --concurrent=true

docker-build:
	docker-build -t crypto-cli .

docker-run-encrypt-string:
	docker run -rm crypto-cli run --mode=encrypt --type=string --input="docker Secret" --key="algo1234%^&*suck"