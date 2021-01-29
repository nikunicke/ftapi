NAME = ftAPI

all:
	go build -o ${NAME} ./cmd/ftAPI

fclean:
	rm -f $(NAME)