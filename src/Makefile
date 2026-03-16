LIBS = -lm
CFLAGS = -Wall
CC = gcc

BUILD_DIR = ./build

CHECK_VERSION = 0.15.2
CHECK_PREFIX = $(BUILD_DIR)/check
CHECK_SRC_DIR = $(BUILD_DIR)/check-source
CHECK_TAR = $(CHECK_SRC_DIR)/check-$(CHECK_VERSION).tar.gz
CHECK_EXTRACT_DIR = $(CHECK_SRC_DIR)/check-$(CHECK_VERSION)

SRC = $(wildcard *.c)

all :
	mkdir -p $(BUILD_DIR)/
	$(CC) -o $(BUILD_DIR)/main $(CFLAGS) $(LIBS) ourmath.c main.c

pack :

clean : 
	rm -rf $(BUILD_DIR)

check-deps :
	@if [ ! -f $(CHECK_PREFIX)/lib/libcheck.a ] && [ ! -f $(CHECK_PREFIX)/lib/libcheck.so ]; then \
		mkdir -p $(CHECK_SRC_DIR) $(CHECK_PREFIX); \
		if [ ! -f $(CHECK_EXTRACT_DIR)/configure ]; then \
			if [ ! -f $(CHECK_TAR) ]; then \
				wget -O $(CHECK_TAR) https://github.com/libcheck/check/releases/download/$(CHECK_VERSION)/check-$(CHECK_VERSION).tar.gz; \
			fi; \
			tar -xzf $(CHECK_TAR) -C $(CHECK_SRC_DIR); \
		fi; \
		cd $(CHECK_EXTRACT_DIR) && ./configure --prefix=$(PWD)/$(CHECK_PREFIX) --disable-subunit; \
		$(MAKE) -C $(CHECK_EXTRACT_DIR); \
		$(MAKE) -C $(CHECK_EXTRACT_DIR) install; \
	fi

check_install :
	@if ! pkg-config --exists check; then \
		echo "Check library not found. Building locally in $(CHECK_PREFIX)..."; \
		$(MAKE) check-deps; \
	fi

test : check-deps
	mkdir -p $(BUILD_DIR)/
	$(CC) -o $(BUILD_DIR)/test_program $(CFLAGS) -I$(CHECK_PREFIX)/include ourmath.c test_ourmath.c -L$(CHECK_PREFIX)/lib -lcheck -pthread -lrt -lm
	LD_LIBRARY_PATH=$(CHECK_PREFIX)/lib $(BUILD_DIR)/test_program

doc : all

run : all
	$(BUILD_DIR)/main

stddev : all

help :
