.PHONY: all get-common

all: get-common

get-common:
	@echo "Getting LabStars/selpo-common package..."
	go get github.com/LabStars/selpo-common@latest

clean: 
	@echo "go mod clean cache..."
	go clean -modcache

set-hee: 
	@echo "setting up to moonintheroom..."
	git config user.name moonintheroom
	git config user.email moonintheroom@github.com

set-woo: 
	@echo "setting up to dbsehddn0901..."
	git config user.name dbsehdnd0901
	git config user.email dbsehddn0901@github.com

# go get github.com/LabStars/selpo-common 에러 시 
git-export:
	@echo "setting up to GOPRIVATE..."
	echo 'export GOPRIVATE=github.com/LabStars' >> ~/.zshrc
	source ~/.zshrc
