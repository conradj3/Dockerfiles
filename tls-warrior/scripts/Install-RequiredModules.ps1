# Read the config file
$config = Get-Content -Raw -Path /tmp/config.json | ConvertFrom-Json

# Install PowerShell modules
foreach ($module in $config.powershellModules) {
    Install-Module -Name $module.name -Scope AllUsers -Force
}

# Install specific Azure CLI version
Install-Module -Name Az -RequiredVersion $config.azureModules[0].versions[0] -Scope AllUsers -Force

# Set the default PowerShell version to the specified one
Install-Package -Name PowerShell -RequiredVersion $config.pwsh.version -Scope AllUsers -Force
