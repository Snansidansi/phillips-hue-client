.PHONY: default build run clean translate translate-update

default: run

build:
	@mkdir -p build
	@go build -o ./build/phillips-hue-desktop-client

run: build
	@LANG=en_US.UTF-8 ./build/phillips-hue-desktop-client

clean:
	@rm -r build

translate:
	@fyne translate translations/en.json
	@fyne translate translations/de.json

translate-update:
	@fyne translate --update translations/en.json
	@fyne translate --update translations/de.json
