# buildah
FROM podmanbuildbase AS buildah
ARG FUSEOVERLAYFS_VERSION=2cbd1c4a2d3ab06c4e39276363c74c6a9c62c0fe
ARG BUILDAH_VERSION=v1.8.3
RUN set -eux; \
	git clone https://github.com/containers/fuse-overlayfs /fuse-overlayfs; \
	cd /fuse-overlayfs; \
	git checkout "${FUSEOVERLAYFS_VERSION}"; \
	sh autogen.sh; \
	LIBS="-ldl" LDFLAGS="-static" ./configure --prefix /usr; \
	make; \
	make install; \
	fuse-overlayfs --help >/dev/null
RUN git clone --branch ${BUILDAH_VERSION} https://github.com/containers/buildah $GOPATH/src/github.com/containers/buildah
WORKDIR $GOPATH/src/github.com/containers/buildah
RUN make static && mv buildah.static /usr/local/bin/buildah

