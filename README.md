# Portseeker_Go

## Tabla de Contenidos

1. [Introducción](#introducción)
2. [Características](#características)
3. [Instalación](#instalación)
4. [Uso](#uso)
   - [Bandera de Línea de Comandos](#bandera-de-línea-de-comandos)
   - [Ejecutar el Escaneo](#ejecutar-el-escaneo)
   - [Exportar Resultados](#exportar-resultados)
5. [Estructura del Proyecto](#estructura-del-proyecto)
6. [Dependencias](#dependencias)
7. [Contribuir](#contribuir)
8. [Licencia](#licencia)
9. [Descargar Archivos CVE](#descargar-archivos-cve)

## Introducción

**Portseeker_Go** es una aplicación en Go diseñada para escanear un hostname y rango de puertos especificados utilizando Nmap, hacer coincidir los puertos abiertos con vulnerabilidades conocidas (CVEs) y mostrar un resumen de los resultados del escaneo. La aplicación también proporciona una opción para exportar los resultados en formatos PDF, CSV o JSON. Esta herramienta es particularmente útil para profesionales de seguridad y administradores de sistemas que necesitan evaluar rápidamente la postura de seguridad de sus sistemas.

## Características

- **Integración con Nmap:** Utiliza Nmap para escanear el hostname y rango de puertos especificados.
- **Coincidencia de CVEs:** Hace coincidir los puertos abiertos con vulnerabilidades conocidas (CVEs) desde un archivo JSON proporcionado.
- **Tablero de Resumen:** Proporciona un resumen de los resultados del escaneo, incluyendo el total de puertos abiertos, el total de vulnerabilidades y un desglose por severidad.
- **Opciones de Exportación:** Permite a los usuarios exportar los resultados del escaneo en formatos PDF, CSV o JSON.
- **Interacción del Usuario:** Solicita al usuario después del escaneo si desea exportar los resultados y en qué formato.

## Instalación

### Requisitos Previos

- **Go (Golang):** Asegúrate de tener Go instalado en tu sistema. Puedes descargarlo desde [aquí](https://golang.org/dl/).
- **Nmap:** Asegúrate de tener Nmap instalado en tu sistema. Puedes descargarlo desde [aquí](https://nmap.org/download.html).

### Pasos

1. **Clonar el Repositorio:**

   ```sh
   git clone https://github.com/elliotsecops/Portseeker_Go.git
   cd Portseeker_Go
   ```

2. **Inicializar el Módulo Go:**

   ```sh
   go mod init Portseeker_Go
   ```

3. **Instalar Dependencias:**

   ```sh
   go mod tidy
   ```

4. **Compilar el Proyecto:**

   ```sh
   go build -o portseeker cmd/main.go
   ```

## Uso

### Bandera de Línea de Comandos

La aplicación soporta las siguientes banderas de línea de comandos:

- `--host`: El hostname a escanear (por defecto: `127.0.0.1`).
- `--start`: Inicio del rango de puertos (por defecto: `1`).
- `--end`: Fin del rango de puertos (por defecto: `1024`).
- `--cve`: Ruta al archivo JSON de CVE (por defecto: `cve_json_files/nvdcve-1.1-2024.json`).

### Ejecutar el Escaneo

Para ejecutar el escaneo, usa el siguiente comando:

```sh
go run cmd/main.go --host=127.0.0.1 --start=1 --end=1024 --cve=cve_json_files/nvdcve-1.1-2024.json
```

### Exportar Resultados

Después de que el escaneo termine, el programa solicitará al usuario si desea exportar los resultados. Si el usuario elige exportar los resultados, se le pedirá que seleccione el formato (JSON, CSV o PDF) y proporcione un nombre de archivo.

Ejemplo de interacción:

```sh
Do you want to export the results? (yes/no): yes
In which format do you want to export the results? (json/csv/pdf): pdf
Enter the filename for export: vulns
Results exported to PDF successfully.
```

## Estructura del Proyecto

El proyecto está organizado en los siguientes directorios y archivos:

```
Portseeker_Go/
├── cmd/
│   └── main.go
├── internal/
│   ├── cve/
│   │   └── cve.go
│   └── scanner/
│       └── scanner.go
├── cve_json_files/
│   ├── nvdcve-1.1-2002.json
│   ├── nvdcve-1.1-2003.json
│   ├── nvdcve-1.1-2004.json
│   ├── nvdcve-1.1-2005.json
│   ├── nvdcve-1.1-2006.json
│   ├── nvdcve-1.1-2007.json
│   ├── nvdcve-1.1-2008.json
│   ├── nvdcve-1.1-2009.json
│   ├── nvdcve-1.1-2010.json
│   ├── nvdcve-1.1-2011.json
│   ├── nvdcve-1.1-2012.json
│   ├── nvdcve-1.1-2013.json
│   ├── nvdcve-1.1-2014.json
│   ├── nvdcve-1.1-2015.json
│   ├── nvdcve-1.1-2016.json
│   ├── nvdcve-1.1-2017.json
│   ├── nvdcve-1.1-2018.json
│   ├── nvdcve-1.1-2019.json
│   ├── nvdcve-1.1-2020.json
│   ├── nvdcve-1.1-2021.json
│   ├── nvdcve-1.1-2022.json
│   ├── nvdcve-1.1-2023.json
│   ├── nvdcve-1.1-2024.json
│   ├── nvdcve-1.1-Modified.json
│   └── nvdcve-1.1-Recent.json
├── pkg/
│   ├── cve/
│   │   └── cve.go
│   ├── exporter/
│   │   └── exporter.go
│   └── scanner/
│       └── scanner.go
└── go.mod
```

### Descripción de Directorios y Archivos

- **cmd/main.go:** El punto de entrada de la aplicación. Define las banderas de línea de comandos para el hostname, puerto de inicio, puerto de fin y archivo CVE. Carga los datos de CVE, escanea el hostname y rango de puertos especificados utilizando Nmap, y muestra el tablero de resumen.
- **internal/cve/cve.go:** Contiene funciones para cargar datos de CVE desde un archivo JSON.
- **internal/scanner/scanner.go:** Contiene funciones para ejecutar Nmap, analizar la salida y hacer coincidir CVEs con puertos abiertos.
- **cve_json_files/:** Directorio que contiene archivos JSON de CVE.
- **pkg/cve/cve.go:** Contiene funciones para cargar datos de CVE desde un archivo JSON.
- **pkg/exporter/exporter.go:** Contiene funciones para exportar resultados de escaneo a formatos JSON, CSV y PDF.
- **pkg/scanner/scanner.go:** Contiene funciones para ejecutar Nmap, analizar la salida y hacer coincidir CVEs con puertos abiertos.
- **go.mod:** Archivo de módulo Go que lista las dependencias del proyecto.

## Dependencias

El proyecto depende de las siguientes dependencias:

- **Nmap:** Para escanear el hostname y rango de puertos especificados.
- **GoFpdf:** Para exportar resultados de escaneo a formato PDF.

Estas dependencias son gestionadas utilizando módulos Go. El archivo `go.mod` lista todas las dependencias requeridas.

## Descargar Archivos CVE

Para descargar los archivos CVE más recientes, puedes usar el script `download_cve_files.sh` proporcionado. Este script descargará los archivos JSON de CVE necesarios desde el NVD (National Vulnerability Database) y los guardará en el directorio `cve_json_files`.

### Script `download_cve_files.sh`

```bash
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
```

### Ejecutar el Script

Para ejecutar el script, navega al directorio del proyecto y ejecuta el siguiente comando:

```sh
bash download_cve_files.sh
```

Esto descargará los archivos CVE más recientes y los guardará en el directorio `cve_json_files`.

---

¡Gracias por usar **Portseeker_Go**! Si tienes alguna pregunta o necesitas más ayuda, por favor abre un issue en el repositorio de GitHub.

---

# Portseeker_Go

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
   - [Command-Line Flags](#command-line-flags)
   - [Running the Scan](#running-the-scan)
   - [Exporting Results](#exporting-results)
5. [Project Structure](#project-structure)
6. [Dependencies](#dependencies)
7. [Contributing](#contributing)
8. [License](#license)
9. [Downloading CVE Files](#downloading-cve-files)

## Introduction

**Portseeker_Go** is a Go application designed to scan a specified hostname and port range using Nmap, match the open ports with known vulnerabilities (CVEs), and display a summary of the scan results. The application also provides an option to export the results in PDF, CSV, or JSON formats. This tool is particularly useful for security professionals and system administrators who need to quickly assess the security posture of their systems.

## Features

- **Nmap Integration:** Utilizes Nmap to scan the specified hostname and port range.
- **CVE Matching:** Matches open ports with known vulnerabilities (CVEs) from a provided JSON file.
- **Summary Dashboard:** Provides a summary of the scan results, including total open ports, total vulnerabilities, and a breakdown by severity.
- **Export Options:** Allows users to export the scan results in PDF, CSV, or JSON formats.
- **User Interaction:** Prompts the user after the scan to ask if they want to export the results and in which format.

## Installation

### Prerequisites

- **Go (Golang):** Ensure you have Go installed on your system. You can download it from [here](https://golang.org/dl/).
- **Nmap:** Ensure Nmap is installed on your system. You can download it from [here](https://nmap.org/download.html).

### Steps

1. **Clone the Repository:**

   ```sh
   git clone https://github.com/elliotsecops/Portseeker_Go.git
   cd Portseeker_Go
   ```

2. **Initialize the Go Module:**

   ```sh
   go mod init Portseeker_Go
   ```

3. **Install Dependencies:**

   ```sh
   go mod tidy
   ```

4. **Build the Project:**

   ```sh
   go build -o portseeker cmd/main.go
   ```

## Usage

### Command-Line Flags

The application supports the following command-line flags:

- `--host`: The hostname to scan (default: `127.0.0.1`).
- `--start`: Start of the port range (default: `1`).
- `--end`: End of the port range (default: `1024`).
- `--cve`: Path to the CVE JSON file (default: `cve_json_files/nvdcve-1.1-2024.json`).

### Running the Scan

To run the scan, use the following command:

```sh
go run cmd/main.go --host=127.0.0.1 --start=1 --end=1024 --cve=cve_json_files/nvdcve-1.1-2024.json
```

### Exporting Results

After the scan is finished, the program will prompt the user if they want to export the results. If the user chooses to export the results, they will be asked to select the format (JSON, CSV, or PDF) and provide a filename.

Example interaction:

```sh
Do you want to export the results? (yes/no): yes
In which format do you want to export the results? (json/csv/pdf): pdf
Enter the filename for export: vulns
Results exported to PDF successfully.
```

## Project Structure

The project is organized into the following directories and files:

```
Portseeker_Go/
├── cmd/
│   └── main.go
├── internal/
│   ├── cve/
│   │   └── cve.go
│   └── scanner/
│       └── scanner.go
├── cve_json_files/
│   ├── nvdcve-1.1-2002.json
│   ├── nvdcve-1.1-2003.json
│   ├── nvdcve-1.1-2004.json
│   ├── nvdcve-1.1-2005.json
│   ├── nvdcve-1.1-2006.json
│   ├── nvdcve-1.1-2007.json
│   ├── nvdcve-1.1-2008.json
│   ├── nvdcve-1.1-2009.json
│   ├── nvdcve-1.1-2010.json
│   ├── nvdcve-1.1-2011.json
│   ├── nvdcve-1.1-2012.json
│   ├── nvdcve-1.1-2013.json
│   ├── nvdcve-1.1-2014.json
│   ├── nvdcve-1.1-2015.json
│   ├── nvdcve-1.1-2016.json
│   ├── nvdcve-1.1-2017.json
│   ├── nvdcve-1.1-2018.json
│   ├── nvdcve-1.1-2019.json
│   ├── nvdcve-1.1-2020.json
│   ├── nvdcve-1.1-2021.json
│   ├── nvdcve-1.1-2022.json
│   ├── nvdcve-1.1-2023.json
│   ├── nvdcve-1.1-2024.json
│   ├── nvdcve-1.1-Modified.json
│   └── nvdcve-1.1-Recent.json
├── pkg/
│   ├── cve/
│   │   └── cve.go
│   ├── exporter/
│   │   └── exporter.go
│   └── scanner/
│       └── scanner.go
└── go.mod
```

### Directory and File Descriptions

- **cmd/main.go:** The entry point of the application. It defines command-line flags for the hostname, start port, end port, and CVE file. It loads the CVE data, scans the specified hostname and port range using Nmap, and displays the summary dashboard.
- **internal/cve/cve.go:** Contains functions to load CVE data from a JSON file.
- **internal/scanner/scanner.go:** Contains functions to run Nmap, parse the output, and match CVEs with open ports.
- **cve_json_files/:** Directory containing CVE JSON files.
- **pkg/cve/cve.go:** Contains functions to load CVE data from a JSON file.
- **pkg/exporter/exporter.go:** Contains functions to export scan results to JSON, CSV, and PDF formats.
- **pkg/scanner/scanner.go:** Contains functions to run Nmap, parse the output, and match CVEs with open ports.
- **go.mod:** Go module file that lists the project dependencies.

## Dependencies

The project relies on the following dependencies:

- **Nmap:** For scanning the specified hostname and port range.
- **GoFpdf:** For exporting scan results to PDF format.

These dependencies are managed using Go modules. The `go.mod` file lists all the required dependencies.

## Contributing

Contributions are welcome! If you would like to contribute to this project, please follow these steps:

1. **Fork the Repository:** Click the "Fork" button on the top-right corner of the repository page.
2. **Clone Your Fork:** Clone the forked repository to your local machine.
3. **Create a New Branch:** Create a new branch for your changes.
4. **Make Changes:** Make the necessary changes to the project.
5. **Commit Changes:** Commit your changes with a descriptive commit message.
6. **Push Changes:** Push your changes to your forked repository.
7. **Create a Pull Request:** Go to the original repository and create a pull request from your forked branch.

Please ensure that your code adheres to the existing coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Downloading CVE Files

To download the latest CVE files, you can use the provided `download_cve_files.sh` script. This script will download the necessary CVE JSON files from the NVD (National Vulnerability Database) and save them in the `cve_json_files` directory.

### `download_cve_files.sh` Script

```bash
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
```

### Running the Script

To run the script, navigate to the project directory and execute the following command:

```sh
bash download_cve_files.sh
```

This will download the latest CVE files and save them in the `cve_json_files` directory.

---

Thank you for using **Portseeker_Go**! If you have any questions or need further assistance, please feel free to open an issue on the GitHub repository.
