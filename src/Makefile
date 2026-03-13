# Compiler and flags
CC = gcc
CFLAGS = -Wall -I.
LIBS = -lm
CHECK_LIBS = $(shell pkg-config --cflags --libs check)

# Source files
SRC = $(wildcard *.c)
MAIN_SRC = main.c
TEST_SRC = test_ourmath.c

# Build directories
BUILD_DIR = build
BIN = $(BUILD_DIR)/main
TEST_BIN = $(BUILD_DIR)/test_ourmath

# Default target
all: $(BIN)

# Build main program
$(BIN): $(SRC) ourmath.h
	@mkdir -p $(BUILD_DIR)
	$(CC) $(CFLAGS) -o $@ $(MAIN_SRC) $(LIBS)

# Build test program
$(TEST_BIN): $(TEST_SRC) ourmath.h
	@mkdir -p $(BUILD_DIR)
	$(CC) $(CFLAGS) -o $@ $(TEST_SRC) $(LIBS) $(CHECK_LIBS)

# Run main program
run: $(BIN)
	./$(BIN)

# Build and run tests
test: $(TEST_BIN)
	./$(TEST_BIN)

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)

# Help message
help:
	@echo "Makefile targets:"
	@echo "  all       - build the main program"
	@echo "  run       - run the main program"
	@echo "  test      - build and run Check unit tests"
	@echo "  clean     - remove build artifacts"
	@echo "  help      - show this message"