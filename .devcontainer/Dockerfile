FROM mcr.microsoft.com/vscode/devcontainers/go:0-1-bullseye AS Go
# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/cpp/.devcontainer/base.Dockerfile

# [Choice] Debian / Ubuntu version (use Debian 11, Ubuntu 18.04/22.04 on local arm64/Apple Silicon): debian-11, debian-10, ubuntu-22.04, ubuntu-20.04, ubuntu-18.04
ARG VARIANT="bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/cpp:0-bullseye

# [Optional] Install CMake version different from what base image has already installed. 
# CMake reinstall choices: none, 3.21.5, 3.22.2, or versions from https://cmake.org/download/
ARG REINSTALL_CMAKE_VERSION_FROM_SOURCE="none"

# Optionally install the cmake for vcpkg
COPY ./reinstall-cmake.sh /tmp/
RUN if [ "${REINSTALL_CMAKE_VERSION_FROM_SOURCE}" != "none" ]; then \
    chmod +x /tmp/reinstall-cmake.sh && /tmp/reinstall-cmake.sh ${REINSTALL_CMAKE_VERSION_FROM_SOURCE}; \
    fi \
    && rm -f /tmp/reinstall-cmake.sh

# [Optional] Uncomment this section to install additional vcpkg ports.
# RUN su vscode -c "${VCPKG_ROOT}/vcpkg install"
# [Optional] Uncomment this section to install additional packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends wget unzip sqlite3

RUN cd /tmp && wget https://github.com/protocolbuffers/protobuf/releases/download/v21.6/protobuf-cpp-3.21.6.zip && unzip protobuf-cpp-3.21.6.zip -d /opt 
RUN cd /opt/protobuf-3.21.6 && cmake . && cmake --build . --parallel 10
RUN cd /opt/protobuf-3.21.6 && cmake --install .

COPY --from=Go /go /go
COPY --from=Go /usr/local/go /usr/local/go
ENV GOROOT=/usr/local/go \
    GOPATH=/go
ENV PATH=${GOPATH}/bin:${GOROOT}/bin:${PATH}
RUN chown -R vscode ${GOPATH} ${GOROOT}

USER vscode
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest