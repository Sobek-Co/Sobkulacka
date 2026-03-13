LIBS = -lm
CFLAGS = -Wall
CC = gcc

SRC = $(wildcard *.c)

all :
	mkdir -p ./build/
	$(CC) -o ./build/main $(CFLAGS) $(LIBS) ourmath.c main.c

pack :

clean : 
	rm -r build/

test :
	mkdir -p ./build/
	$(CC) -o ./build/test_program $(CFLAGS) ourmath.c test_ourmath.c -lcheck -lsubunit -pthread -lrt -lm
	./build/test_program

doc : all

run : all
	./build/main

stddev : all

help : 
