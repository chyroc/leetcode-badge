all: build

build:
	CGO_ENABLED=0 go build -o ./bin/leetcode-badge github.com/Chyroc/leetcode-badge

linux_build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/leetcode-badge github.com/Chyroc/leetcode-badge
