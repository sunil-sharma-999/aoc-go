SRC_DIR := ./src
BUILD_DIR := ./dist

# Find all main.go files in src subdirectories
MAIN_FILES := $(shell find $(SRC_DIR) -name main.go)

# Create a build rule for each main.go file
.PHONY: all
all: $(MAIN_FILES:$(SRC_DIR)/%/main.go=$(BUILD_DIR)/%.exe)

$(BUILD_DIR)/%.exe: $(SRC_DIR)/%/main.go
	@echo "Building $<"
	mkdir -p $(dir $@)
	go build -o $@ $<

.PHONY: run
run:
	@year=$(word 2, $(MAKECMDGOALS)); \
	day=$(word 3, $(MAKECMDGOALS)); \
	args=$(filter-out $@,$(MAKECMDGOALS)); \
	$(BUILD_DIR)/$$year/day$$day.exe $$args

# Prevent make from interpreting year and day as targets
%:
	@: