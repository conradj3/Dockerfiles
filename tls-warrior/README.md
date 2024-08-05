# Jammy Jellyfish PowerShell Utility Container

A utility container with preloaded Modules for troubleshooting TLS and Network inside Kubernetes or from Docker Engine.

## Lib Packages

## Vital Packages

1. **curl**: A command-line tool for transferring data using various protocols (HTTP, FTP, etc.). It's commonly used for downloading files or interacting with APIs.
1. **tar**: A tool for creating and extracting archives. It supports various compression methods like gzip and bzip2.
1. **unzip**: A utility for extracting files from ZIP archives.
1. **wget**: A command-line utility for downloading files from the web. Supports HTTP, HTTPS, and FTP protocols.
1. **jq**: A lightweight and flexible command-line JSON processor, used for parsing and manipulating JSON data.

1. **ca-certificates**: Common CA certificates. Ensures that the system can verify the authenticity of SSL connections.

### Common Packages

1. **dnsutils**: A collection of utilities for querying DNS (Domain Name System) servers, such as `dig`, `nslookup`, and `host`.
1. **dpkg**: The Debian package manager, which installs, removes, and manages Debian packages.
1. **iproute2**: A collection of utilities for controlling networking, including `ip` for network interface configuration and routing.
1. **iputils-ping**: A set of network utilities, including `ping` for checking the reachability of hosts on a network.
1. **openssh-client**: A suite of secure networking utilities based on the SSH protocol, including `ssh`, `scp`, and `sftp` for secure communication and file transfer.

### Command Packages

1. **coreutils**: A package containing the basic file, shell, and text manipulation utilities of the GNU operating system, such as `ls`, `cat`, `cp`, and `mv`.
1. **file**: A utility for determining the type of a file based on its content and metadata.
1. **findutils**: A collection of tools for searching and finding files, including `find`, `xargs`, and `locate`.
1. **ftp**: The File Transfer Protocol client, used for transferring files to and from remote hosts.
1. **net-tools**: A collection of programs for controlling the network subsystem of the Linux kernel, including `ifconfig`, `netstat`, `route`, and `arp`.
1. **nmap**: A network exploration and security auditing tool. It’s used for network discovery and vulnerability scanning.
1. **netcat-openbsd**: A versatile networking tool, also known as `nc`, used for reading from and writing to network connections.
1. **ssh**: A secure shell client used for logging into and executing commands on remote machines securely.
1. **sshpass**: A non-interactive SSH password authentication tool, used to automate SSH login with a password.
1. **telnet**: A command-line tool for interacting with remote systems using the Telnet protocol, often used for debugging and testing network services.
1. **time**: A command used to measure the duration of execution of a program or command.
1. **zip**: A utility for creating ZIP archives, which compresses files into a single archive for easier distribution and storage.

## Powershell Packages

### Azure Modules

The following Azure module is required:

1. **az**: Installs 11.5.0 of the entirety of Az.* of Powershell Modules

### PowerShell Modules

The following PowerShell modules are required:

1. **Microsoft.Graph**: The Microsoft Graph PowerShell SDK is a collection of modules that contain cmdlets for calling Microsoft Graph. Microsoft Graph is the gateway to data and intelligence in Microsoft 365.

1. **Pester**: Pester is a testing framework for PowerShell, used for writing and running tests to ensure your PowerShell code performs as expected.

1. **PSScriptAnalyzer**: PSScriptAnalyzer is a static code checker for PowerShell modules and scripts, providing best practices and rule-based analysis to improve script quality.

1. **PSTCPIp**: The PSTCPIp module provides cmdlets for TCP/IP related operations, allowing for network-related scripting and automation tasks.

1. **PSGraphQL**: PSGraphQL is a PowerShell module for interacting with GraphQL APIs, enabling the execution of queries and mutations against GraphQL endpoints.

1. **PSJsonWebToken**: PSJsonWebToken provides cmdlets for creating and validating JSON Web Tokens (JWT), often used for authentication and secure information exchange.

1. **PowerShellGet**: PowerShellGet is a module with cmdlets for discovering, installing, updating, and publishing PowerShell modules and scripts.

1. **PSReadLine**: PSReadLine enhances the PowerShell command-line editing experience with features like syntax highlighting, multi-line editing, and customizable key bindings.
