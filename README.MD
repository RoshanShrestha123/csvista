![logo](logo.png)

# CSVista

CSVista is a command-line tool written in Go that converts CSV files to JSON format and outputs the result in the terminal.

## Features

- Converts any CSV file to JSON format.
- Simple command-line interface.
- Supports specifying the path to the CSV file using a command-line flag (`--p=/filepath`).

## Installation

Make sure you have Go installed on your system. Then, you can install CSVista using `go get`:

```bash
go get github.com/RoshanShrestha123/csvista
```

Alternatively, you can clone the repository and build it manually:

```bash
git clone https://github.com/RoshanShrestha123/csvista
cd csvista
go build
```

Usage
To convert a CSV file to JSON, run CSVista from the terminal with the path to your CSV file specified using the --p flag:

```bash

csvista --p=/path/to/your/file.csv
```
Example:

```bash
csvista --p=/Users/username/documents/data.csv
```
This will print the JSON output to the terminal.

Command-line Options
--p=/filepath: Specify the path to the CSV file that you want to convert to JSON.


##### TODO
- [ ] export the JSON to .json file 