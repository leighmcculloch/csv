# csv

Extract columns from CSV files.

## Install

### Binary (Linux; macOS; Windows)

Download and install the binary from the [releases](https://github.com/leighmcculloch/csv/releases) page.

### From Source

```
go get 4d63.com/csv
```

## Usage

```
cat file.csv | csv '{{index . 3}}'
```
