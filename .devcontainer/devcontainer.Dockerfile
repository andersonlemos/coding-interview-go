# syntax=docker/dockerfile:1
FROM golang:1.24-bookworm
# This Dockerfile is used to create a development container for the Go programming language.
# It is based on the official Golang image and installs additional tools and dependencies.
# The container is configured to run as a non-root user with sudo privileges.
# The container is built using the Dockerfile syntax version 1.
# The base image is the official Golang image with version 1.24 on Debian Bookworm.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Install dependencies
# The following environment variables are set to configure the installation of various packages.
# These packages include git, wget, openssl, openjdk-17-jre-headless, pre-commit, and hadolint.
# The versions of these packages are specified to ensure compatibility and stability.
# The DEBIAN_FRONTEND variable is set to noninteractive to suppress prompts during package installation.
# The apt-get update command is used to update the package lists for the Debian package manager.
# The apt-get install command is used to install the specified packages.
# The --no-install-recommends option is used to prevent the installation of recommended packages.
# The apt-get clean command is used to remove cached package files to reduce the size of the image.
# The rm -rf /var/lib/apt/lists/* command is used to remove the package lists to free up space.
# The echo command is used to add the non-root user to the sudoers file, allowing them to run commands as root without a password.
# The chmod command is used to set the permissions of the sudoers file to ensure that only the root user can modify it.
# The useradd command is used to create a new user with the specified username, UID, and GID.
# The groupadd command is used to create a new group with the specified GID.

ARG DEBIAN_FRONTEND=noninteractive
ARG GIT_VERSION=1:2.39.5-0+deb12u2
ARG WGET_VERSION=1.21.3-1+deb12u1
ARG OPEN_SSL_VERSION=3.0.15-1~deb12u1
ARG JDK_VERSION=17.0.14+7-1~deb12u1
ARG PRECOMMIT_VERSION=3.0.4-1
ARG HADOLINT_VERSION=v2.12.0
ARG SUDO_VERSION=1.9.13p3-1+deb12u1
ARG GNUPG2_VERSION=2.2.40-1.1
ARG GPG_AGENT_VERSION=2.2.40-1.1
ARG GPG_VERSION=2.2.40-1.1
ARG SOCAT_VERSION=1.7.4.4-2
ARG PIPX_VERSION=1.1.0-1
ARG COMMITIZEN_VERSION=4.4.1
#Install git, wget, openssl, openjdk-17-jre-headless, pre-commit
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    gpg=${GPG_VERSION} \
    gnupg2=${GNUPG2_VERSION} \
    gpg-agent=${GPG_AGENT_VERSION} \
    socat=${SOCAT_VERSION} \
    git=${GIT_VERSION} \
    wget=${WGET_VERSION} \
    openssl=${OPEN_SSL_VERSION} \
    openjdk-17-jre-headless=${JDK_VERSION} \
    pre-commit=${PRECOMMIT_VERSION} \
    pipx=${PIPX_VERSION} \
    sudo=${SUDO_VERSION} \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* 

RUN wget -q -O hadolint https://github.com/hadolint/hadolint/releases/download/${HADOLINT_VERSION}/hadolint-Linux-x86_64 \
    && chmod +x hadolint \ 
    && mv hadolint /usr/local/bin/ 

RUN  groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \ 
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

USER $USERNAME

RUN pipx ensurepath && pipx install --pip-args="--no-cache-dir" commitizen==${COMMITIZEN_VERSION}


