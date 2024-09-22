# Portseeker_Go

## Tabla de Contenidos

1. [Introducción](#introducción)
2. [Características](#características)
3. [Instalación](#instalación)
4. [Uso](#uso)
   - [Bandera de Línea de Comando](#bandera-de-línea-de-comando)
   - [Ejecutando el Escaneo](#ejecutando-el-escaneo)
   - [Exportando Resultados](#exportando-resultados)
5. [Estructura del Proyecto](#estructura-del-proyecto)
6. [Dependencias](#dependencias)
7. [Contribuyendo](#contribuyendo)

## Introducción

**Portseeker_Go** es una aplicación de Go diseñada para escanear un hostname y un rango de puertos especificados utilizando Nmap, coincidir los puertos abiertos con vulnerabilidades conocidas (CVEs) y mostrar un resumen de los resultados del escaneo. La aplicación también proporciona una opción para exportar los resultados en formatos PDF, CSV o JSON. Esta herramienta es particularmente útil para profesionales de la seguridad y administradores de sistemas que necesitan evaluar rápidamente la postura de seguridad de sus sistemas.

## Características

- **Integración de Nmap:** Utiliza Nmap para escanear el hostname y el rango de puertos especificados.
- **Coincidencia de CVE:** Coincide los puertos abiertos con vulnerabilidades conocidas (CVEs) a partir de un archivo JSON proporcionado.
- **Tablero de Resumen:** Proporciona un resumen de los resultados del escaneo, incluyendo el total de puertos abiertos, el total de vulnerabilidades y un desglose por gravedad.
- **Opciones de Exportación:** Permite a los usuarios exportar los resultados del escaneo en formatos PDF, CSV o JSON.
- **Interacción del Usuario:** Pregunta al usuario después del escaneo si desean exportar los resultados y en qué formato.

## Instalación

### Prerrequisitos

- **Go (Golang):** Asegúrese de tener Go instalado en su sistema. Puede descargarlo desde [aquí](https://golang.org/dl/).
- **Nmap:** Asegúrese de que Nmap esté instalado en su sistema. Puede descargarlo desde [aquí](https://nmap.org/download.html).

### Pasos

1. **Clonar el Repositorio:**

   ```sh
   git clone https://github.com/elliotsecops/Portseeker_Go.git
   cd Portseeker_Go
   ```
2. ## Descargando Archivos CVE

Para descargar los archivos CVE más recientes, debe utilizar el script `download_cve_files.sh` proporcionado. Este script descargará los archivos JSON CVE necesarios del NVD (National Vulnerability Database) y los guardará en el directorio `cve_json_files`.

#### Ejecutando el Script

1. **Navegar al Directorio `cve_json_files`:**
   ```bash
   cd cve_json_files
   ```
Este comando cambia el directorio actual al directorio `cve_json_files`. Este es el directorio donde el script descargará y descomprimirá los archivos JSON CVE.

2. **Hacer el Script Ejecutable:**
   ```bash
   chmod +x download_cve_files.sh
   ```

El comando `chmod +x` hace que el script `download_cve_files.sh` sea ejecutable. Esto es necesario porque, por defecto, los archivos descargados de un repositorio o creados en su sistema pueden no tener el permiso de ejecución. La bandera `+x` agrega el permiso de ejecución al script, lo que permite que se ejecute como un programa.

3. **Ejecutar el Script:**
   ```bash
   ./download_cve_files.sh
   ```
Al ejecutar este comando, el script:
    - Descargará los archivos JSON CVE especificados del NVD (National Vulnerability Database).
    - Descomprimirá los archivos `.gz` descargados para extraer los archivos `.json`.
    - Guardará los archivos `.json` descomprimidos en el directorio `cve_json_files`.

Siguiendo estos pasos, asegura que los archivos CVE se descarguen y se coloquen en el directorio correcto, listos para su uso por la aplicación `Portseeker_Go`.

3. **Inicializar el Módulo de Go:**

   ```sh
   go mod init Portseeker_Go
   ```

4. **Instalar Dependencias:**

   ```sh
   go mod tidy
   ```

5. **Compilar el Proyecto:**

   ```sh
   go build -o portseeker cmd/main.go
   ```

## Uso

### Bandera de Línea de Comando

La aplicación admite las siguientes banderas de línea de comando:

- `--host`: El hostname a escanear (por defecto: `127.0.0.1`).
- `--start`: Comienzo del rango de puertos (por defecto: `1`).
- `--end`: Final del rango de puertos (por defecto: `1024`).
- `--cve`: Ruta al archivo JSON CVE (por defecto: `cve_json_files/nvdcve-1.1-2024.json`).

### Ejecutando el Escaneo

Para ejecutar el escaneo, utilice el siguiente comando:

```sh
go run cmd/main.go --host=127.0.0.1 --start=1 --end=1024 --cve=cve_json_files/nvdcve-1.1-2024.json
```
Al ejecutar este comando, la aplicación `Portseeker_Go` escaneará el hostname y el rango de puertos especificados, coincidirá los puertos abiertos con vulnerabilidades conocidas a partir del archivo JSON CVE proporcionado y mostrará un resumen de los resultados del escaneo.

- **`go run cmd/main.go`**: Compila y ejecuta la aplicación `Portseeker_Go`.
- **`--host=127.0.0.1`**: Especifica el hostname o la dirección IP a escanear (localhost en este caso).
- **`--start=1`**: Especifica el número de puerto de inicio para el escaneo (puerto 1).
- **`--end=1024`**: Especifica el número de puerto final para el escaneo (puerto 1024).
- **`--cve=cve_json_files/nvdcve-1.1-2024.json`**: Especifica la ruta al archivo JSON CVE que contiene las vulnerabilidades conocidas.

### Exportando Resultados

Después de que el escaneo se haya completado, el programa preguntará al usuario si desean exportar los resultados. Si el usuario elige exportar los resultados, se le pedirá que seleccione el formato (JSON, CSV o PDF) y que proporcione un nombre de archivo.

Ejemplo de interacción:

```sh
¿Desea exportar los resultados? (sí/no): sí
¿En qué formato desea exportar los resultados? (json/csv/pdf): pdf
Introduzca el nombre de archivo para la exportación: vulnerabilidades
Se exportaron correctamente los resultados a PDF.
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

- **cmd/main.go:** El punto de entrada de la aplicación. Define las banderas de línea de comando para el hostname, el puerto de inicio, el puerto final y el archivo CVE. Carga los datos CVE, escanea el hostname y el rango de puertos especificados utilizando Nmap y muestra el tablero de resumen.
- **internal/cve/cve.go:** Contiene funciones para cargar los datos CVE a partir de un archivo JSON.
- **internal/scanner/scanner.go:** Contiene funciones para ejecutar Nmap, analizar la salida y coincidir los CVEs con los puertos abiertos.
- **cve_json_files/:** Directorio que contiene los archivos JSON CVE.
- **pkg/cve/cve.go:** Contiene funciones para cargar los datos CVE a partir de un archivo JSON.
- **pkg/exporter/exporter.go:** Contiene funciones para exportar los resultados del escaneo a formatos JSON, CSV y PDF.
- **pkg/scanner/scanner.go:** Contiene funciones para ejecutar Nmap, analizar la salida y coincidir los CVEs con los puertos abiertos.
- **go.mod:** Archivo de módulo de Go que lista las dependencias del proyecto.

## Dependencias

El proyecto depende de las siguientes dependencias:

- **Nmap:** Para escanear el hostname y el rango de puertos especificados.
- **GoFpdf:** Para exportar los resultados del escaneo al formato PDF.

Estas dependencias se administran utilizando los módulos de Go. El archivo `go.mod` lista todas las dependencias requeridas.

## Contribuciones

Las contribuciones son bienvenidas. Si desea contribuir a este proyecto, siga los siguientes pasos:

1. **Forkear el Repositorio:** Haga clic en el botón "Fork" en la esquina superior derecha de la página del repositorio.
2. **Clonar el Fork:** Clone el repositorio forked en su máquina local.
3. **Crear una Nueva Rama:** Cree una nueva rama para sus cambios.
4. **Realizar Cambios:** Realice los cambios necesarios en el proyecto.
5. **Confirmar Cambios:** Confirme sus cambios con un mensaje de confirmación descriptivo.
6. **Enviar Cambios:** Envíe sus cambios al repositorio forked.
7. **Crear una Solicitud de Extracción:** Vaya al repositorio original y cree una solicitud de extracción a partir de la rama forked.

Asegúrese de que su código cumpla con los estándares de programacion existentes e incluya las pruebas apropiadas.
 
---

¡Gracias por usar **Portseeker_Go**! Si tiene alguna pregunta o necesita más ayuda, no dudes en abrir un issue.

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
2. ## Downloading CVE Files

To download the latest CVE files you have to use the provided `download_cve_files.sh` script. This script will download the necessary CVE JSON files from the NVD (National Vulnerability Database) and save them in the `cve_json_files` directory.

#### Running the Script

1. **Navigate to the `cve_json_files` Directory:**
   ```bash
   cd cve_json_files
   ```
 This command changes the current directory to the `cve_json_files` directory. This is where the script will download and unzip the CVE JSON files.

2. **Make the Script Executable:**
   ```bash
   chmod +x download_cve_files.sh
   ```

   - The `chmod +x` command makes the `download_cve_files.sh` script executable. This is necessary because, by default, files downloaded from a repository or created on your system may not have the executable permission. The `+x` flag adds the execute permission to the script, allowing it to be run as a program.

3. **Run the Script:**
   ```bash
   ./download_cve_files.sh
   ```
When you run this command, the script will:
    - Download the specified CVE JSON files from the NVD (National Vulnerability Database).
     - Unzip the downloaded `.gz` files to extract the `.json` files.
     - Save the unzipped `.json` files in the `cve_json_files` directory.

By following these steps, you ensure that the CVE files are downloaded and placed in the correct directory, ready for use by the `Portseeker_Go` application.

3. **Initialize the Go Module:**

   ```sh
   go mod init Portseeker_Go
   ```

4. **Install Dependencies:**

   ```sh
   go mod tidy
   ```

5. **Build the Project:**

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
By running this command, the Portseeker_Go application will scan the specified host and port range, match the open ports with known vulnerabilities from the provided CVE JSON file, and display a summary of the scan results.

- **`go run cmd/main.go`**: Compiles and runs the `Portseeker_Go` application.
- **`--host=127.0.0.1`**: Specifies the hostname or IP address to scan (localhost in this case).
- **`--start=1`**: Specifies the starting port number for the scan (port 1).
- **`--end=1024`**: Specifies the ending port number for the scan (port 1024).
- **`--cve=cve_json_files/nvdcve-1.1-2024.json`**: Specifies the path to the CVE JSON file that contains the known vulnerabilities.

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
 


### Running the Script

To run the script, navigate to the project directory and execute the following command:

```sh
bash download_cve_files.sh
```

This will download the latest CVE files and save them in the `cve_json_files` directory.

---

Thank you for using **Portseeker_Go**! If you have any questions or need further assistance, please feel free to open an issue on the GitHub repository.
