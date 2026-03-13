LIBS = -lm
CFLAGS = -Wall
CC = gcc

# Only include main program sources, not tests
SRC = main.c

all: $(SRC)
	mkdir -p ./build/
	$(CC) -o ./build/main $(CFLAGS) $(LIBS) $(SRC)

# Check library linking
CHECK_LIBS = -lcheck -pthread -lrt -lm -lsubunit

# Build and run tests
test: ./build/test_ourmath
	./build/test_ourmath

./build/test_ourmath: test_ourmath.c ourmath.h
	mkdir -p ./build/
	$(CC) $(CFLAGS) -o ./build/test_ourmath test_ourmath.c $(LIBS) $(CHECK_LIBS)

pack :

clean: 
	rm -rf build/

doc: all

run: all
	./build/main

stddev: all

help: