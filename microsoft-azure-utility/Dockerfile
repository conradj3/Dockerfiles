# Use an official base image
FROM ubuntu:20.04

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies and tools
RUN apt-get update && \
    apt-get install -y apt-transport-https curl lsb-release gnupg && \
    rm -rf /var/lib/apt/lists/* 

# Install Azure CLI
RUN mkdir -p /etc/apt/keyrings && \
    curl -sLS https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor | tee /etc/apt/keyrings/microsoft.gpg > /dev/null && \
    chmod go+r /etc/apt/keyrings/microsoft.gpg  && \
    AZ_DIST=$(lsb_release -cs) && \
    echo "deb [arch=`dpkg --print-architecture` signed-by=/etc/apt/keyrings/microsoft.gpg] https://packages.microsoft.com/repos/azure-cli/ $AZ_DIST main" | tee /etc/apt/sources.list.d/azure-cli.list
    
RUN apt-get update && \
    apt-get install azure-cli

# Install Azure DevOps CLI
RUN az extension add --name azure-devops

# Install GitHub CLI
RUN type -p curl >/dev/null || (sudo apt update && sudo apt install curl -y) && \
    curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg && \
    chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null && \
    apt update && \
    apt install gh -y
# Download the Microsoft repository keys
RUN wget -q https://packages.microsoft.com/config/ubuntu/$VERSION_ID/packages-microsoft-prod.deb

# Register the Microsoft repository keys
RUN dpkg -i packages-microsoft-prod.deb

# Delete the the Microsoft repository keys file
RUN rm packages-microsoft-prod.deb

# Update the list of packages after we added packages.microsoft.com
RUN sudo apt-get update

###################################
# Install PowerShell
RUN apt-get install powershell -y 
# Set working directory
WORKDIR /root

# Define default command
CMD [ "/bin/bash" ]
