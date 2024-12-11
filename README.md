
# EnvBuddy

EnvBuddy is a simple and secure command-line tool for managing and encrypting `.env` files. It helps developers protect sensitive environment variables by encrypting `.env` files and easily decrypting them when necessary.

## Features

- Encrypt `.env` files to securely store sensitive information.
- Decrypt encrypted `.env` files.
- Supports AES encryption with a predefined key.
- Easy-to-use CLI with commands for encryption and decryption.

## Installation

### Prerequisites

- Go 1.18 or higher

### Clone the repository

```bash
git clone https://github.com/alexrosepizant/envbuddy.git
cd envbuddy
```

### Install dependencies

Run the following command to install the required Go modules:

```bash
go mod tidy
```

### Build the application

To build the executable, run:

```bash
go build -o envbuddy main.go
```

This will generate an executable file called `envbuddy`.

## Usage

After building the application, you can use the following commands.

### Encrypt a `.env` file

To encrypt a `.env` file, run:

```bash
./envbuddy encrypt <filename>
```

Example:

```bash
./envbuddy encrypt .env.sample
```

This will generate an encrypted file named `.env.sample.encrypted`.

### Decrypt a `.env` file

To decrypt an encrypted `.env` file, run:

```bash
./envbuddy decrypt <filename>
```

Example:

```bash
./envbuddy decrypt .env.sample.encrypted
```

This will generate a decrypted `.env.sample.decrypted` file.

## Key Management

The tool uses a fixed encryption key in the code (`myverystrongpasswordo32bitlength`). In production, consider using a more secure key management solution or allow the user to specify a key at runtime.

## Contributing

Contributions are welcome! Feel free to fork the repository, submit issues, or create pull requests.

### How to contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature-xyz`)
3. Commit your changes (`git commit -am 'Add feature xyz'`)
4. Push to the branch (`git push origin feature-xyz`)
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
