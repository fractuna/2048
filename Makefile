all: run

BIN=2048
MAIN=src/main.go src/state.go src/utils.go src/keys.go src/item.go src/config.go src/data.go
FLATPAK_PACKAGE=go2048
FLATPAK_MANIFEST="org.flatpak.2048.yml"
BUILD_MODE=x11

run:
	tools/go-bindata -o src/data.go resources/ # DEBUG
	go run -tags ${BUILD_MODE} ${MAIN}

build:
	tools/go-bindata -o src/data.go resources/
	go build -v -tags ${BUILD_MODE} -o ${BIN} ${MAIN}

# Flatpak environment setup (TODO: Add)
#
flatpak-run:
	flatpak run org.flatpak.go2048

flatpak:
	flatpak-builder --user --force-clean --install-deps-from=flathub --repo=repo --install ${FLATPAK_PACKAGE} ${FLATPAK_MANIFEST}

flatpak-clean:
	rm -rf ${FLATPAK_PACKAGE}
	rm -rf repo
	rm -rf .flatpak-builder

clean:
	rm -rf ${BIN}
