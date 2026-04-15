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
	cd $(GO_SRC_DIR) && go mod tidy && go build -o $(GO_BUILD_OUT)

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
			gtk4 \
			libadwaita \
			gdk-pixbuf2 \
			gobject-introspection \
			pkg-config \
			gcc \
			make \
			; \
	elif command -v apt >/dev/null; then \
		echo "📦 Using apt (Debian/Ubuntu)"; \
		sudo apt update; \
		sudo apt install -y \
			golang \
			check \
			libgtk-4-dev \
			libgtk-4-common \
			libadwaita-1-dev \
			libgdk-pixbuf-2.0-dev \
			libgirepository1.0-dev \
			pkg-config \
			gcc \
			make \
			; \
	elif command -v dnf >/dev/null; then \
		echo "📦 Using dnf (Fedora)"; \
		sudo dnf install -y \
			golang \
			check \
            check-devel \
			gtk4-devel \
			libadwaita-devel \
			gdk-pixbuf2-devel \
			gobject-introspection-devel \
			pkgconf-pkg-config \
			gcc \
			make \
			; \
	else \
		echo "⚠️ Unknown package manager. Please install manually:"; \
		echo "   - Go"; \
		echo "   - libcheck"; \
		echo "   - GTK4 development files (libgtk-4-dev or gtk4-devel)"; \
		echo "   - GDK Pixbuf development files"; \
		echo "   - pkg-config"; \
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
