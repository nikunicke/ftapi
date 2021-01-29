NAME = ftapi

all:
	go build -o ${NAME} ./cmd/ftapi

fclean:
	rm -f $(NAME)