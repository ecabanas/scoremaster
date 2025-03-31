BIN		=		main

MAIN	=		main.go
RM		=		rm -rf 

###########################################

build:
		go build -o "$(BIN)"

run: build
		go run main.go

remove:		
		$(RM) ./logs.txt
		$(RM) ./scoremaster.db

.PHONY: build run remove