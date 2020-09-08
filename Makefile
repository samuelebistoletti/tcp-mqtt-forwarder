CC=go

BUILD_DIR=builds/bin
DIST_DIR=builds/debian
DEPS_DIR=debian-deps

BIN=tcp-mqtt-forwarder
VERSION=$(shell cat VERSION)
BUILD_TIME=$(shell date +%FT%T%z)

ARCHS=amd64 386 arm

LDFLAGS=-ldflags='-s -w -X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}'

SOURCES=$(shell ls *.go)

.PHONY: all clean lint

all: clean
	@echo Building binaries for: $(ARCHS)
	@echo Building version: $(VERSION) \($(BUILD_TIME)\)
	@echo Building binaries ...
	@for arch in $(ARCHS) ; do 																																										 \
		if [ $$arch != "arm" ] ;                                                                                     \
			then GOARCH=$$arch $(CC) build $(LDFLAGS) -o $(BUILD_DIR)/$(VERSION)/$(BIN)_$$arch $(SOURCES);             \
			else GOARCH=$$arch GOARM=7 $(CC) build $(LDFLAGS) -o $(BUILD_DIR)/$(VERSION)/$(BIN)_$${arch}v7 $(SOURCES); \
		fi                                                                                                           \
	done

clean:
	@if [ -d $(BUILD_DIR) ] ; then rm -rf $(BUILD_DIR) ; fi

clean-dir:
	@if [ -d $(DIST_DIR) ] ; then rm -rf $(DIST_DIR) ; fi

deb: clean-dir all
	@echo Building deb packages for: $(ARCHS)
	@for arch in $(ARCHS) ; do                                                                                                                                             				  \
		if [ "$${arch}" = "386" ] ; then export dir_arch=i686 ; else export dir_arch=$${arch} ; fi 																					   && \
		if [ "$${arch}" = "arm" ] ; then export arch=$${arch}v7 && export dir_arch=armhf ; else export arch=$${arch} ; fi 															   && \
		$(MKDIR_P) $(DIST_DIR)/$(BIN)-$(VERSION)_$$dir_arch/ 																															   && \
		cp -r debian-template/* $(DIST_DIR)/$(BIN)-$(VERSION)_$$dir_arch/ 																														   && \
		find $(DIST_DIR)/$(BIN)-$(VERSION)_$$dir_arch/ -name '.gitignore' -exec rm {} \; 																							   && \
		cp $(BUILD_DIR)/$(VERSION)/$(BIN)_$${arch} $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/usr/bin/$(BIN) 																		   && \
		m4 -D DEB_ARCH=$$dir_arch -D DEB_VERSION=$(VERSION) -D DEB_PROJECT=$(BIN) $(DEPS_DIR)/systemd.m4 > $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/lib/systemd/system/$(BIN).service && \
		m4 -D DEB_ARCH=$$dir_arch -D DEB_VERSION=$(VERSION) -D DEB_PROJECT=$(BIN) $(DEPS_DIR)/conffiles.m4 > $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/DEBIAN/conffiles 			   && \
		m4 -D DEB_ARCH=$$dir_arch -D DEB_VERSION=$(VERSION) -D DEB_PROJECT=$(BIN) $(DEPS_DIR)/control.m4 > $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/DEBIAN/control 				   && \
		m4 -D DEB_ARCH=$$dir_arch -D DEB_VERSION=$(VERSION) -D DEB_PROJECT=$(BIN) $(DEPS_DIR)/postinst.m4 > $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/DEBIAN/postinst 				   && \
		m4 -D DEB_ARCH=$$dir_arch -D DEB_VERSION=$(VERSION) -D DEB_PROJECT=$(BIN) $(DEPS_DIR)/preinst.m4 > $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/DEBIAN/preinst 				   && \
		m4 -D DEB_ARCH=$$dir_arch -D DEB_VERSION=$(VERSION) -D DEB_PROJECT=$(BIN) $(DEPS_DIR)/prerm.m4 > $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/DEBIAN/prerm 					   && \
		m4 -D DEB_ARCH=$$dir_arch -D DEB_VERSION=$(VERSION) -D DEB_PROJECT=$(BIN) $(DEPS_DIR)/postrm.m4 > $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/DEBIAN/postrm 					   && \
		cp $(DEPS_DIR)/config.json $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/etc/$(BIN)/config.json                                                                  			   && \
		chmod -R 0755 $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/                                                                                                       			   && \
		chmod 0644 $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/DEBIAN/conffiles                                                                                          		       && \
		chmod 0644 $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/etc/$(BIN)/config.json                                                                                    		   && \
		fakeroot dpkg-deb --build $(DIST_DIR)/$(BIN)-$(VERSION)_$${dir_arch}/                                                                                              			   && \
		rm -r $(DIST_DIR)/$(BIN)-$(VERSION)_$$dir_arch/ 																												   				  \
	; done

MKDIR_P := mkdir -p