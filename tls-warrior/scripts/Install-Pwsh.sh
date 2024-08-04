# Install Required Dependency Lib
wget https://mirror.it.ubc.ca/ubuntu/pool/main/i/icu/libicu72_72.1-3ubuntu3_amd64.deb 
dpkg -i libicu72_72.1-3ubuntu3_amd64.deb 

# Install PowerShell 7.4.4
wget https://github.com/PowerShell/PowerShell/releases/download/v7.4.4/powershell_7.4.4-1.deb_amd64.deb 
dpkg -i powershell_7.4.4-1.deb_amd64.deb 

# Clean Up
rm libicu72_72.1-3ubuntu3_amd64.deb 
rm powershell_7.4.4-1.deb_amd64.deb 