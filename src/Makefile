LIBS = -lm
CFLAGS = -Wall
CC = gcc

CHECK_VERSION = 0.15.2
CHECK_PREFIX = ./build/check
CHECK_SRC_DIR = ./build/check-source
CHECK_TAR = $(CHECK_SRC_DIR)/check-$(CHECK_VERSION).tar.gz
CHECK_EXTRACT_DIR = $(CHECK_SRC_DIR)/check-$(CHECK_VERSION)

SRC = $(wildcard *.c)

all :
	mkdir -p ./build/
	$(CC) -o ./build/main $(CFLAGS) $(LIBS) ourmath.c main.c

pack :

clean : 
	rm -rf build/

check-deps:
	mkdir -p $(CHECK_SRC_DIR) $(CHECK_PREFIX)
	@if [ ! -f $(CHECK_EXTRACT_DIR)/configure ]; then \
		if [ ! -f $(CHECK_TAR) ]; then \
			wget -O $(CHECK_TAR) https://github.com/libcheck/check/releases/download/$(CHECK_VERSION)/check-$(CHECK_VERSION).tar.gz; \
		fi; \
		tar -xzf $(CHECK_TAR) -C $(CHECK_SRC_DIR); \
	fi
	cd $(CHECK_EXTRACT_DIR) && ./configure --prefix=$(PWD)/$(CHECK_PREFIX) && make && make install

check_install :
	@if ! pkg-config --exists check; then \
		echo "Check library not found. Building locally in $(CHECK_PREFIX)..."; \
		$(MAKE) check-deps; \
	fi

test : check-deps
	mkdir -p ./build/
	$(CC) -o ./build/test_program $(CFLAGS) -I$(CHECK_PREFIX)/include ourmath.c test_ourmath.c -L$(CHECK_PREFIX)/lib -lcheck -lsubunit -pthread -lrt -lm
	LD_LIBRARY_PATH=$(CHECK_PREFIX)/lib ./build/test_program

doc : all

run : all
	./build/main

stddev : all

help :