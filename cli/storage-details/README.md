# Storage Details CLI Tool

Welcome to the **Storage Details CLI Tool**! This tool is built with Go, utilizing the powerful Cobra and Viper libraries. It provides detailed information about the storage usage on your system.

## Features

- **View Storage Details**: Get a comprehensive overview of your storage usage.
- **Filter by File Type**: Filter storage details by specific file types.
- **Sort by Size**: Sort files and directories by size.
- **Configuration**: Easily configurable using Viper.

## Installation

To install the CLI tool, use the following command:

```sh
go install github.com/yourusername/storage-details-cli@latest
```

## Usage

Here are some common commands you can use with the Storage Details CLI Tool:

### View Storage Details

```sh
storage-details view
```

### Filter by File Type

```sh
storage-details filter --type=jpg
```

### Sort by Size

```sh
storage-details sort --order=desc
```

## Configuration

The tool uses Viper for configuration management. You can create a configuration file named `config.yaml` in the following directory:

```sh
$HOME/.config/storage-details/
```

### Example Configuration

```yaml
storage:
    path: "/path/to/your/storage"
    exclude:
        - "node_modules"
        - ".git"
```

### Setup
1. For intialize run `~/go/bin/cobra-cli init`
2. To create custom command run `~/go/bin/cobra-cli add filesTors`
3. Run it using `go run . <command>`. Example: `go run . filesTors`


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or feedback, please open an issue on the [GitHub repository](https://github.com/yourusername/storage-details-cli).

Thank you for using the Storage Details CLI Tool!
