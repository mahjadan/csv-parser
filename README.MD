# Rain CSV Processing Tool

The **Rain CSV Parser Tool** is a powerful command-line utility crafted in Golang, designed to streamline the parsing and standardization of CSV files received from diverse employers. With a keen focus on accuracy, efficiency, and adaptability, this tool transforms raw data into a structured format, ensuring seamless integration into Rain's processing pipeline.

### Key Features:

- Automated Harmonization: Effortlessly transforms CSV data to match Rain's processing requirements, regardless of its original structure.

- Anomaly Detection: Quickly identifies and reports data anomalies, offering insightful summaries directly in the command line.

- Efficient Processing: Ensures uniformity and compatibility with Rain's standards, streamlining data processing.

- Smart Segmentation: Separates data into distinct valid and issue-containing CSV files, simplifying further analysis.

- Golang Ecosystem Compatibility: Developed in Golang with support for Go modules, the Rain CSV Processing Tool embraces the latest language features, ensuring reliability, maintainability, and future-proofing.


## Installation

1. Navigate to the root directory of the project.
2. Ensure you have Go 1.20 or later installed.
## Building the Tool
To build the Rain CSV Processing Tool, run the following command in your terminal:
```bash
go build -o rcsv main.go
```

## Usage
To utilize the Rain CSV Processing Tool for parsing CSV files, follow these steps:

1. Run the command below, replacing [filename.csv] with the path to the CSV file you want to parse:

    ```bash
    ./rcsv parse [filename.csv]
    ```
    For example, to process a CSV file named roster1.csv located in the input directory:
    ```bash
    ./rcsv parse input/roster1.csv
    ```
    The tool will initiate the parsing process, creating separate valid and invalid CSV files. These files will be stored in the output directory within the project, providing a clear distinction between correctly processed data and records with issues.



For further insights into the available command options, the tool offers a built-in help feature. You can access this feature by running the following command:
```bash
./rcsv parse -h
```
This command will provide you with information on how to change the **configuration file path** and the **output directory** for processed CSV files. 

## Using Cobra and Viper
The Rain CSV Processing Tool leverages the Cobra and Viper packages for building the command-line interface and handling configurations. These packages are excellent choices for this project due to their robustness, ease of use, and popularity within the Go community.

Cobra simplifies the creation of powerful command-line interfaces by providing a simple and intuitive framework. Viper, on the other hand, enhances configuration management by supporting various configuration formats and providing a clean way to bind configuration values to structs.

By using Cobra and Viper, the Rain CSV Processing Tool achieves a well-structured and user-friendly command-line experience while ensuring flexible and efficient configuration management.

## How It Works
The Rain CSV Processing Tool processes CSV files with varying structures. It dynamically adapts to column names and orders.

The tool uses a configuration file (config/config.json) to define expected column names and possible aliases for each column, as well as what will be the output columns/format. This ensures flexibility and extensibility when dealing with different file formats.

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
As the project evolves, the tool can be further improved with the following enhancements:

- **Expanded File Format Support:** Extend support to encompass a broader range of file formats, moving beyond CSV.
- **Enhanced Validation:** Incorporate additional validation mechanisms to align with evolving company requirements.
- **Robust Error Handling:** Strengthen error handling and introduce comprehensive logging capabilities for improved troubleshooting.
- **Error Insights:** Introduce an insightful feature that compiles statistics on commonly encountered errors during record processing, highlighting the most frequent error messages.
- **Enhanced User Experience:** Enhance user experience by implementing a loading cursor in the terminal to indicate ongoing file processing.
- **Seamless Data Integration:** Seamlessly integrate the tool with Rain's wider data processing pipelines to streamline overall data workflows.


## Conclusion
The Rain CSV Parser is a powerful command-line tool that effectively parses and standardizes CSV files for Rain's processing. Its flexibility, scalability, and extensibility make it an invaluable asset for managing diverse data sources. The tool's architecture ensures seamless adaptation to future changes and improvements.