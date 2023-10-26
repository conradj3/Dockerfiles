FROM ubuntu:20.04  
  
ENV DEBIAN_FRONTEND=noninteractive  
  
# Update and install required packages  
RUN apt-get update && \  
    apt-get install -y --no-install-recommends apt-transport-https curl lsb-release gnupg wget && \  
    rm -rf /var/lib/apt/lists/*  
  
# Install Azure CLI  
RUN mkdir -p /etc/apt/trusted.gpg.d/ && \  
    curl -sL https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > /etc/apt/trusted.gpg.d/microsoft.gpg && \  
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/trusted.gpg.d/microsoft.gpg] https://packages.microsoft.com/repos/azure-cli/ $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/azure-cli.list && \  
    apt-get update && \  
    apt-get install -y --no-install-recommends azure-cli && \  
    az extension add --name azure-devops  
  
# Install GitHub CLI  
RUN curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | gpg --dearmor | sudo tee /usr/share/keyrings/githubcli-archive-keyring.gpg > /dev/null && \  
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null && \  
    apt-get update && \  
    apt-get install -y --no-install-recommends gh  
  
# Install PowerShell  
RUN wget -q https://packages.microsoft.com/config/ubuntu/20.04/packages-microsoft-prod.deb && \  
    dpkg -i packages-microsoft-prod.deb && \  
    rm packages-microsoft-prod.deb && \  
    apt-get update && \  
    apt-get install -y --no-install-recommends powershell  
  
# Cleanup  
RUN apt-get clean && \  
    rm -rf /var/lib/apt/lists/*  
  
WORKDIR /root  
  
CMD [ "/bin/bash" ]  
