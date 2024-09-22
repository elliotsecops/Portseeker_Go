#!/bin/bash

# List of all CVE .gz file URLs
urls=(
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-Modified.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-Recent.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2024.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2023.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2022.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2021.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2020.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2019.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2018.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2017.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2016.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2015.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2014.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2013.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2012.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2011.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2010.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2009.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2008.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2007.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2006.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2005.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2004.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2003.json.gz"
    "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-2002.json.gz"
)

# Function to download and unzip a file
download_and_unzip_file() {
    local url="$1"
    local file_name="${url##*/}"
    local json_file_name="${file_name%.gz}"
    echo "Downloading $file_name..."
    if wget -q "$url"; then
        echo "$file_name downloaded successfully."
        echo "Unzipping $file_name..."
        gunzip -f "$file_name"
        echo "$json_file_name unzipped successfully."
    else
        echo "Failed to download $file_name."
    fi
}

# Use coproc to create background processes for downloading and unzipping files
for url in "${urls[@]}"; do
    coproc download_and_unzip_file "$url"
done

# Wait for all background processes to finish
wait

echo "All files downloaded and unzipped!"