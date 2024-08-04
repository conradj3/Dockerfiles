# Use Ubuntu 24.04 as the base image
FROM ubuntu:24.04

# Set the noninteractive frontend
ENV DEBIAN_FRONTEND=noninteractive

# Copy the JSON config into the container
COPY ./toolset/config.json /tmp/config.json
COPY ./scripts /tmp/scripts

# Install required packages
RUN apt-get update && \
    apt-get install -y \
    wget \
    jq \
    curl \
    && jq -r '.apt.vital_packages + .apt.common_packages + .apt.cmd_packages | .[]' /tmp/config.json | xargs apt-get install -y --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

# Install PowerShell
RUN sh ./tmp/scripts/Install-Pwsh.sh 

# Execute the PowerShell script to install modules
RUN apt-get update && \
    pwsh -File /tmp/scripts/Install-RequiredModules.ps1

# Install Azure CLI
RUN curl -sL https://aka.ms/InstallAzureCLIDeb | bash

# Clean up
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/config.json

# Set dialog for frontend
ENV DEBIAN_FRONTEND=dialog

# Set entrypoint
ENTRYPOINT [ "pwsh" ]
