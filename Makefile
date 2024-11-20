# Makefile

# Variables
TEMPL_DIR := internal/templates
GENERATED_DIR := generated
GO_FILES := $(wildcard *.go)
TEMPL_FILES := $(wildcard $(TEMPL_DIR)/*.templ)
GENERATED_FILES := $(patsubst $(TEMPL_DIR)/%.templ,$(GENERATED_DIR)/%_templ.go,$(TEMPL_FILES))

# Default target
all: generate run


tail:
	@echo "Compiling css"
	bun install \
	&& bun run build:css

# Generate Templ files
generate:
	@echo "Generating Templ files..."
	@templ generate \


# Move generated files to subdirectory
# move:
# 	echo "Moving generated files to $(GENERATED_DIR)..." \
# 	&& mkdir -p $(TEMPL_DIR)/$(GENERATED_DIR) \
# 	# && mv $(TEMPL_DIR)/*_templ.go $(TEMPL_DIR)/$(GENERATED_DIR)/ 2>/dev/null || true


run: 
	go run cmd/server/main.go


# Clean generated files

# Phony targets
.PHONY: all generate run
