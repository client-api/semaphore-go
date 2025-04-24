
# Clean up

$generatedFilesPath = "./.openapi-generator/FILES"

if (Test-Path -Path $generatedFilesPath) {
    $generatedFiles = Get-Content -Path $generatedFilesPath

    foreach ($generatedFile in $generatedFiles) {
                $generatedFile = $generatedFile.Trim()

        if ($generatedFile -in '.gitignore', '.openapi-generator-ignore', 'README.md', 'api/openapi.yaml') {
            continue
        }

        if (Test-Path -Path $generatedFile) {
            Remove-Item -Path $generatedFile -Force
        }
    }
}

# Download generator

if (-Not (Test-Path -Path "./.openapi-generator/openapi-generator-cli.jar")) {
    $xmlString = Invoke-WebRequest -Uri "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/maven-metadata.xml" -UseBasicParsing

    if (-Not $xmlString) {
        Write-Host "Hiba történt az XML adat lekérésekor."
        exit
    }

    $xml = [xml]$xmlString.Content
    $latestVersion = $xml.metadata.versioning.latest
    $latestVersionUrl = "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$latestVersion/openapi-generator-cli-$latestVersion.jar"

    Invoke-WebRequest -Uri $latestVersionUrl -OutFile "./.openapi-generator/openapi-generator-cli.jar"
}

# Check Java

$output = & java -version 2>&1
if ($output -match "java version" -or $output -match "openjdk version") {
    Write-Host "Java version: $output"
} else {
    Write-Host "Java is not found."
    exit 1
}

# Check Generator Version

$output = & java -jar .\.openapi-generator\openapi-generator-cli.jar version 2>&1
if ($output -match "$latestVersion") {
    Write-Host "OpenAPI Generator version: $output"
} else {
    Write-Host "Unable to verify OpenAPI Generator version"
    exit 1
}

# Generate

java -jar ./.openapi-generator/openapi-generator-cli.jar generate `
    --input-spec "https://raw.githubusercontent.com/client-api/specs-store/refs/heads/main/semaphore/versions/2.13.0.yaml" `
    --generator-name "go" `
    --output ./ `
    --enable-post-process-file `
    --git-host "github.com" `
    --git-repo-id "semaphore-go" `
    --git-user-id "client-api" `
    --package-name "semaphore" `
    --additional-properties "disallowAdditionalPropertiesIfNotPresent=true,enumClassPrefix=true,generateInterfaces=true,generateMarshalJSON=true,generateUnmarshalJSON=true,hideGenerationTimestamp=false,isGoSubmodule=false,packageName=semaphore-go,packageVersion=1.0.0,prependFormOrBodyParameters=false,structPrefix=false,useDefaultValuesForRequiredVars=false,useOneOfDiscriminatorLookup=false,withAWSV4Signature=false,withGoMod=true,withXml=false"

if ($LASTEXITCODE -eq 0) {
    Write-Host "A generálás sikeresen lefutott."
} else {
    Write-Host "Hiba történt a generálás közben."
}