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

# ---- Phony targets ----
.PHONY: all clean deps test run help doc

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
		sudo pacman -Sy --needed --noconfirm \
			go \
			check \
			vulkan-devel \
			libx11 \
			libxkbcommon \
			libxkbcommon-x11 \
			libxcb \
			libxcb-util \
			libxcb-keysyms \
			libxcb-xfixes \
			libxcursor \
			libxfixes \
			wayland \
			mesa \
			lcov \
			pkgconf \
			gcc \
			make \
			; \
	elif command -v apt >/dev/null; then \
		echo "📦 Using apt (Debian/Ubuntu)"; \
		sudo apt update; \
		sudo apt install -y \
			golang \
			check \
			libegl-dev \
      libgles2-mesa-dev \
			libvulkan-dev \
			libx11-dev \
			libxkbcommon-dev \
			libxkbcommon-x11-dev \
			libxcb1-dev \
			libxcb-xkb-dev \
			libxcb-xfixes0-dev \
			libx11-xcb-dev \
			libxcursor-dev \
			libxfixes-dev \
			libwayland-dev \
			mesa-common-dev \
			lcov \
			pkg-config \
			gcc \
			make \
			; \
	elif command -v dnf >/dev/null; then \
		echo "📦 Using dnf (Fedora)"; \
		sudo dnf install -y \
			golang \
			check \
			vulkan-devel \
			libX11-devel \
			libxkbcommon-devel \
			libxkbcommon-x11-devel \
			libxcb-devel \
			libxcb-devel \
			libxcb-devel \
			libXcursor-devel \
			libXfixes-devel \
			wayland-devel \
			mesa-libGL-devel \
			lcov \
			pkgconf-pkg-config \
			gcc \
			make \
			; \
	else \
		echo "⚠️ Unknown package manager. Please install manually:"; \
		echo "   - Go"; \
		echo "   - libcheck"; \
		echo "   - Vulkan headers (vulkan.h)"; \
		echo "   - X11 / Wayland dev libraries"; \
	fi

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

# ---- Doc ----

doc:
	export PROJECT_NUMBER=$(git describe)
	doxygen Doxyfile
