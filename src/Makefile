LIBS = -lm
CFLAGS = -Wall
CC = gcc

BUILD_DIR = ./build

SRC = $(wildcard *.c)

GO_SRC_DIR = .
GO_BUILD_OUT = $(BUILD_DIR)/app

# pkg-config helpers
CHECK_CFLAGS := $(shell pkg-config --cflags check 2>/dev/null)
CHECK_LIBS   := $(shell pkg-config --libs check 2>/dev/null)

# ---- OS detection ----
PKG_MANAGER := $(shell command -v pacman || command -v apt || command -v dnf)

# ---- Main target ----
all:
	mkdir -p $(BUILD_DIR)
	cd $(GO_SRC_DIR) && go build -o $(GO_BUILD_OUT)

clean:
	rm -rf $(BUILD_DIR)

# ---- Dependencies ----

deps:
	@echo "🔍 Installing dependencies..."
	@if command -v pacman >/dev/null; then \
		echo "📦 Using pacman (Arch-based)"; \
		sudo pacman -Sy --needed \
			go \
			check \
			vulkan-devel \
			libx11 \
			libxkbcommon \
			wayland \
			mesa \
			; \
	elif command -v apt >/dev/null; then \
		echo "📦 Using apt (Debian/Ubuntu)"; \
		sudo apt update; \
		sudo apt install -y \
			golang \
			check \
			libvulkan-dev \
			libx11-dev \
			libxkbcommon-dev \
			libwayland-dev \
			mesa-common-dev \
			; \
	elif command -v dnf >/dev/null; then \
		echo "📦 Using dnf (Fedora)"; \
		sudo dnf install -y \
			golang \
			check \
			vulkan-devel \
			libX11-devel \
			libxkbcommon-devel \
			wayland-devel \
			mesa-libGL-devel \
			; \
	else \
		echo "⚠️ Unknown package manager. Please install manually:"; \
		echo "   - Go"; \
		echo "   - libcheck"; \
		echo "   - Vulkan headers (vulkan.h)"; \
		echo "   - X11 / Wayland dev libraries"; \
	fi

	@echo "📦 Fetching Go dependencies..."
	cd $(GO_SRC_DIR) && go mod tidy

# ---- Testing ----

check_install:
	@if ! pkg-config --exists check; then \
		echo "❌ libcheck not found. Run 'make deps'"; \
		exit 1; \
	fi

test: check_install
	mkdir -p $(BUILD_DIR)
	$(CC) -o $(BUILD_DIR)/test_program \
		$(CFLAGS) $(CHECK_CFLAGS) \
		ourmath.c test/test_ourmath.c \
		$(CHECK_LIBS) -pthread -lrt -lm
	$(BUILD_DIR)/test_program

# ---- Run ----

run: all
	$(GO_BUILD_OUT)

help:
	@echo "Targets:"
	@echo "  all      - Build Go app"
	@echo "  deps     - Install system + Go dependencies"
	@echo "  test     - Run C tests"
	@echo "  run      - Run built app"
	@echo "  clean    - Remove build directory"
