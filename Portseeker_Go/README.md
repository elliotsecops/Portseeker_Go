**Portseeker_Go**

## Introducción

**Portseeker_Go** es una aplicación de Go diseñada para escanear un hostname y un rango de puertos especificados utilizando Nmap, coincidir los puertos abiertos con vulnerabilidades conocidas (CVEs) y mostrar una resumen de los resultados de la escaneo. La aplicación también proporciona la opción de exportar los resultados en formatos PDF, CSV o JSON.

## Características

- **Integración con Nmap:** Utiliza Nmap para escanear el hostname y el rango de puertos especificados.
- **Coincidencia de CVE:** Coincide los puertos abiertos con vulnerabilidades conocidas (CVEs) de un archivo JSON proporcionado.
- **Panel de resumen:** Proporciona un resumen de los resultados de la escaneo, incluyendo el número total de puertos abiertos, el número total de vulnerabilidades y una desglose por severidad.
- **Opciones de exportación:** Permite a los usuarios exportar los resultados de la escaneo en formatos PDF, CSV o JSON.
- **Interacción con el usuario:** Pide al usuario después de la escaneo si desea exportar los resultados y en qué formato.

## Instalación

### Requisitos previos

- **Go (Golang):** Asegúrese de tener Go instalado en su sistema. Puede descargarlo desde [aquí](https://golang.org/dl/).
- **Nmap:** Asegúrese de tener Nmap instalado en su sistema. Puede descargarlo desde [aquí](https://nmap.org/download.html).

### Pasos

1. **Clonar el repositorio:**

```sh
git clone https://github.com/elliotsecops/Portseeker_Go.git
cd Portseeker_Go
```

2. **Iniciar el módulo Go:**

```sh
go mod init Portseeker_Go
```

3. **Instalar dependencias:**

```sh
go mod tidy
```

4. **Compilar el proyecto:**

```sh
go build -o portseeker cmd/main.go
```

## Uso

### Bandejas de línea de comandos

La aplicación admite las siguientes bandejas de línea de comandos:

- `--host`: El hostname a escanear (por defecto: `127.0.0.1`).
- `--start`: El inicio del rango de puertos (por defecto: `1`).
- `--end`: El fin del rango de puertos (por defecto: `1024`).
- `--cve`: La ruta al archivo JSON de CVE (por defecto: `cve_json_files/nvdcve-1.1-2024.json`).

### Ejecución del escaneo

Para ejecutar el escaneo, use el siguiente comando:

```sh
go run cmd/main.go --host=127.0.0.1 --start=1 --end=1024 --cve=cve_json_files/nvdcve-1.1-2024.json
```

### Exportación de resultados

Después de que el escaneo se ha completado, el programa preguntará al usuario si desea exportar los resultados. Si el usuario elige exportar los resultados, se le pedirá que seleccione el formato (JSON, CSV o PDF) y proporcione un nombre de archivo.

Ejemplo de interacción:

```sh
¿Desea exportar los resultados? (sí/no): sí
En qué formato desea exportar los resultados? (json/csv/pdf): pdf
Ingrese el nombre del archivo para exportar: vulns
Resultados exportados con éxito al PDF.
```

## Estructura del proyecto

El proyecto está organizado en los siguientes directorios y archivos:

```sh
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

### Descripción de directorios y archivos

- **cmd/main.go:** Punto de entrada de la aplicación. Define bandejas de línea de comandos para el hostname, el inicio del rango de puertos, el fin del rango de puertos y el archivo JSON de CVE. Carga los datos de CVE, escanea el hostname y el rango de puertos utilizando Nmap y muestra el panel de resumen.
- **internal/cve/cve.go:** Contiene funciones para cargar datos de CVE de un archivo JSON.
- **internal/scanner/scanner.go:** Contiene funciones para ejecutar Nmap, analizar el resultado y coincidir CVE con puertos abiertos.
- **cve_json_files/:** Directorio que contiene archivos JSON de CVE.
- **pkg/cve/cve.go:** Contiene funciones para cargar datos de CVE de un archivo JSON.
- **pkg/exporter/exporter.go:** Contiene funciones para exportar resultados de escaneo a formatos JSON, CSV y PDF.
- **pkg/scanner/scanner.go:** Contiene funciones para ejecutar Nmap, analizar el resultado y coincidir CVE con puertos abiertos.
- **go.mod:** Archivo de módulo Go que lista las dependencias del proyecto.

## Dependencias

El proyecto depende de las siguientes dependencias:

- **Nmap:** Para escanear el hostname y el rango de puertos especificados.
- **GoFpdf:** Para exportar resultados de escaneo a formato PDF.

Estas dependencias se gestionan mediante módulos Go. El archivo `go.mod` lista las dependencias requeridas.

## Contribuciones

Las contribuciones son bienvenidas! Si desea contribuir a este proyecto solo haz un branch y envía un pull request.

Asegúrese de que su código cumpla con los estándares de codificación existentes y incluya pruebas adecuadas.

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

---

Thank you for using **Portseeker_Go**! If you have any questions or need further assistance, please feel free to open an issue on the GitHub repository.