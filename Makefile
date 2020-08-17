CC  = `which go`
SRC = src/main.go src/boostrap.go

EXEC =  gojvm

all:
	$(CC) build -o ./bin/$(EXEC) $(SRC)

vars:
	@echo "GO\t\t $(CC)"
	@echo "SOURCE FILES\t $(SRC)"