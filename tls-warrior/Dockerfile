# Use PowerShell Jammy-Jellyfish image
FROM mcr.microsoft.com/powershell:7.4-ubuntu-22.04

# Copy the JSON config and scripts into the container
COPY ./toolset/config.json /tmp/config.json
COPY ./scripts /tmp/scripts

# Set the non-interactive frontend
ENV DEBIAN_FRONTEND=noninteractive

# Install Libs and Powershell Modules
RUN apt-get update \
    && apt-get install -y --no-install-recommends jq \
    && jq -r '.apt.vital_packages + .apt.common_packages + .apt.cmd_packages | .[]' /tmp/config.json | xargs apt-get install -y --no-install-recommends \
    && pwsh -File /tmp/scripts/Install-RequiredModules.ps1 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Set dialog for frontend
ENV DEBIAN_FRONTEND=dialog

# Set entrypoint
ENTRYPOINT [ "pwsh" ]
