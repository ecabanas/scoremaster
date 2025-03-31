BIN		=		main

MAIN	=		main.go
RM		=		rm -rf 

###########################################

build:
		go build -o "$(BIN)"

run: build
		go run main.go

clean:		
		$(RM) ./logs.txt
		$(RM) ./scoremaster.db

test:
		go test ./test -v

.PHONY: build run clean test