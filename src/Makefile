LIBS = -lm
CFLAGS = -Wall
CC = gcc

BUILD_DIR = ./build

DOC_DIR = ../doc

SRC = $(wildcard *.c)

GO_SRC_DIR = .
GO_BUILD_OUT = $(BUILD_DIR)/app

STAGING_DIR = xpalacj00_xsobekd00_xstovim00_xvitekv00

# pkg-config helpers
CHECK_CFLAGS := $(shell pkg-config --cflags check 2>/dev/null)
CHECK_LIBS   := $(shell pkg-config --libs check 2>/dev/null)

# ---- OS detection ----
PKG_MANAGER := $(shell command -v pacman || command -v apt || command -v dnf)

# ---- Phony targets ----
.PHONY: all clean deps test run help doc stddev pack

# ---- Main target ----
all: stddev
	mkdir -p $(BUILD_DIR)
	cd $(GO_SRC_DIR) && go mod tidy && go build -o $(GO_BUILD_OUT)

pack: doc stddev
	$(eval CUR_DIR := $(shell pwd))
	$(eval TEMP_DIR := $(shell mktemp -d))
	mkdir -p $(TEMP_DIR)/$(STAGING_DIR)

# docs
	mkdir $(TEMP_DIR)/$(STAGING_DIR)/doc
	cp -r $(DOC_DIR)/* $(TEMP_DIR)/$(STAGING_DIR)/doc

# repo
	git clone git@git.sleepat.fit:SobekCo/Sobkulacka.git $(TEMP_DIR)/$(STAGING_DIR)/repo
	cd $(TEMP_DIR)/$(STAGING_DIR)/repo && git fetch --all && git pull --all

# installer
	mkdir $(TEMP_DIR)/$(STAGING_DIR)/install
	cp ../install.sh $(TEMP_DIR)/$(STAGING_DIR)/install.sh
	cp $(BUILD_DIR)/stddev $(TEMP_DIR)/$(STAGING_DIR)/stddev

	cd $(TEMP_DIR) && zip -r $(CUR_DIR)/$(STAGING_DIR).zip $(STAGING_DIR)

	rm -rf $(TEMP_DIR)

clean:
	rm -rf $(BUILD_DIR) $(DOC_DIR) data.txt perf.data perf.data.old

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

# ---- Stddev profiler ----

stddev:
	mkdir -p $(BUILD_DIR)
	$(CC) -o $(BUILD_DIR)/stddev \
		$(CFLAGS) -g -O2 \
		ourmath.c profiling/stddev.c \
		$(LIBS)

stddev_test: stddev
	sudo sysctl kernel.perf_event_paranoid=1
	awk 'BEGIN{srand(); for(i=0;i<10000000;i++) print rand()*1000}' > data.txt
	perf record -g $(BUILD_DIR)/stddev < data.txt
	perf report > ../profiling/perf_result.txt
	perf report
	rm -f data.txt perf.data perf.data.old

# ---- Run ----

run: all
	$(GO_BUILD_OUT)

help:
	@echo "Targets:"
	@echo "  all      - Build Desktop App and Profiling program"
	@echo "  deps     - Install system + Go dependencies"
	@echo "  test     - Run C tests"
	@echo "  run      - Run built app"
	@echo "  clean    - Remove build directory"
	@echo "  doc      - Generate doucmentation"
	@echo "  stddev   - Build only profiling program"
	@echo "  pack     - Packs project into format for project submission"
	@echo ""
	@echo "To build program on any Debian/Fedora/Arch based system, simply run these commands in order:"
	@echo "  - make deps"
	@echo "  - make all"
	@echo "On other linux platforms, please install all packages listed by the 'make deps' command and then use 'make all'"
	@echo "On non-linux platforms, you will also need to install the GTK4 platform"

# ---- Doc ----

doc:
	export PROJECT_NUMBER=$(git describe)
	doxygen Doxyfile
