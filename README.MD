# Todos:
- output summary in the console (add this feature)
- Add -c to the parse command to change the default path of the config file

# Rain CSV Parser

The **Rain CSV Parser** is a command-line tool built in Golang to parse CSV files coming from different employers and standardize the data for processing. This tool provides a solution that follows best practices, is well-organized, bug-free, and scalable. It allows you to process CSV files individually, ensuring clean and organized data for further processing.

## Features

- Automatically interprets CSV files and harmonizes the data to conform with Rain's processing specifications.
- Detects data anomalies and presents an informative summary directly in the command line interface.- Efficiently processes CSV files, ensuring uniformity and compatibility with Rain's processing requirements.
- Identifies data inconsistencies and highlights potential issues, providing a comprehensive overview in the command line interface.
- Creates two distinct CSV files for easy data segregation: one for valid records and another for records with issues.
- Built using the robust Golang 1.19, leveraging Go modules for optimal package management and reliability.


## Installation

1. Navigate to the root directory of the project.

## Build

```bash
go build -o rcsv main.go
```

## Run
To parse a CSV file, use the following command:
```bash
./rcsv parse [filename.csv]
```
For example:
```bash
./rcsv parse input/roster1.csv
```


## How It Works
The Rain CSV Parser processes CSV files with varying structures. It dynamically adapts to column names and orders.
The tool uses a configuration file (config/config.json) to define expected column names and possible aliases for each column. This ensures flexibility and extensibility when dealing with different file formats.

## Architecture
The tool is built using the Cobra framework, Viper configuration management, and the Golang standard library.
The architecture is designed to handle CSV files with dynamic column orders and names.

### Key Benefits of Using a CLI Approach:
- **Automated Processing:** The CLI tool enables seamless automation of data processing tasks, streamlining operations and enhancing efficiency.
This is particularly advantageous when dealing with a high volume of CSV files.

- **Standardization:** Varied CSV file structures are managed through standardized processing rules defined in the configuration.
This ensures uniformity in processed data, simplifying downstream operations and analysis.

- **Scalability and Deployment:** Designed for deployment across diverse environments, the CLI tool's lightweight nature makes it suitable for both small and large-scale data processing requirements.

### Architecture Overview:
- `main.go`: The entry point that initializes the CLI interface and executes commands.
- `cmd/parse.go`: This file handles the core CSV processing logic. It reads input files, segregates valid and invalid records into separate CSV files, and dynamically maps columns based on the specified configuration.
  
By adopting a CLI-based approach, the Rain CSV Parser tool combines automation, standardization, and scalability to efficiently handle CSV files with varying structures and column orders.

## Unit Tests
The project includes comprehensive unit tests to ensure the reliability of the tool. These tests cover various scenarios, including valid and error cases. To run the tests:
```bash
go test -v ./...
```

## Future Enhancements
As requirements evolve, the tool can be enhanced with additional features:

- Support for a wider array of file formats beyond CSV.
- Enhanced error handling and logging capabilities.
- Add **Error Insights** by gathering statistics on the most common errors encountered during record processing and print the most frequent error message
- Show loading cursor while processing the file is being processed.
- Integration with Rain's broader data processing pipelines.


## Conclusion
The Rain CSV Parser is a powerful command-line tool that effectively parses and standardizes CSV files for Rain's processing. Its flexibility, scalability, and extensibility make it an invaluable asset for managing diverse data sources. The tool's architecture ensures seamless adaptation to future changes and improvements.