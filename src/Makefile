LIBS = -lm
CFLAGS = -Wall
CC = gcc

SRC = $(wildcard *.c)

all : $(SRC)
	mkdir -p ./build/
	$(CC) -o ./build/main $(CFLAGS) $(LIBS) $(SRC)

pack :

clean : 
	rm -r build/

test : all

doc : all

run : all
	./build/main

stddev : all

help : 
