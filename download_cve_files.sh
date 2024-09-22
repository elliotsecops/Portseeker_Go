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

# Directory to save downloaded files
download_dir="cve_json_files"

# Create the directory if it doesn't exist
mkdir -p "$download_dir"

# Function to download a file
download_file() {
    local url="$1"
    local file_name="${url##*/}"
    echo "Downloading $file_name..."
    if wget -q -P "$download_dir" "$url"; then
        echo "$file_name downloaded successfully."
    else
        echo "Failed to download $file_name."
    fi
}

# Export the function so it can be used by xargs
export -f download_file

# Use xargs to download files in parallel
printf '%s\n' "${urls[@]}" | xargs -P 5 -n 1 -I {} bash -c 'download_file "$@"' _ {}

echo "All files downloaded!"