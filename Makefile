NAME = ftapi

all:
	go build -o ${NAME} ./cmd/ftapi

get-started: examples/getStarted.go
	go build -o get-started ./examples

fclean:
	rm -f $(NAME)