# sha256sum

The `sha256sum` command in Unix and Linux is used to compute and check SHA-256 hash values. SHA-256 is part of the SHA-2 (Secure Hash Algorithm 2) family, which generates a 256-bit (32-byte) hash value, commonly used to verify data integrity and authenticate files.

### Basic Usage

The basic syntax for the `sha256sum` command is:

```sh
sha256sum [OPTION]... [FILE]...
```

- **`FILE`**: The file(s) for which to compute the SHA-256 checksum. If no file is specified, `sha256sum` reads from standard input.

### Examples

#### Calculating SHA-256 Checksum

To calculate the SHA-256 checksum of a file:

```sh
sha256sum file.txt
```

This command outputs the SHA-256 hash value followed by the file name.

Example output:

```
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855  file.txt
```

#### Verifying SHA-256 Checksum

To verify the integrity of a file using a previously generated checksum:

1. Create a checksum file:

```sh
sha256sum file.txt > file.txt.sha256
```

2. Verify the file against the checksum:

```sh
sha256sum -c file.txt.sha256
```

This command checks the hash value in `file.txt.sha256` against the hash value of `file.txt` and reports if the file is unchanged.

Example output:

```
file.txt: OK
```

### Options

#### Reading Checksum from File

To read and verify checksums from a file:

```sh
sha256sum -c checksums.sha256
```

The `checksums.sha256` file should contain lines with the format `<checksum>  <filename>`.

#### Quiet Mode

To run in quiet mode and only display a message for files that do not match:

```sh
sha256sum -c -q file.txt.sha256
```

#### Status Mode

To output only the status for verification without printing the detailed results:

```sh
sha256sum -c --status file.txt.sha256
```

This command returns an exit status of `0` if all files are verified correctly, `1` if any files are mismatched, and `2` for other errors.

#### Binary Mode

To read files in binary mode:

```sh
sha256sum --binary file.txt
```

#### Text Mode

To read files in text mode (default):

```sh
sha256sum --text file.txt
```

### Practical Use Cases

#### Verifying Downloaded Files

When downloading software or large files, it is common to verify the integrity of the downloaded files using the provided SHA-256 checksum:

```sh
sha256sum -c software_download.sha256
```

#### Creating Checksums for Multiple Files

To generate SHA-256 checksums for all files in a directory:

```sh
sha256sum * > all_files.sha256
```

#### Checking Data Integrity

To ensure data integrity when transferring files between systems, compute and verify checksums before and after transfer:

On the source system:

```sh
sha256sum datafile > datafile.sha256
```

On the destination system:

```sh
sha256sum -c datafile.sha256
```

### Summary

The `sha256sum` command is a robust tool for computing and verifying SHA-256 hash values, ensuring the integrity and authenticity of files. By using `sha256sum`, you can safeguard against data corruption and verify that files have not been tampered with. Understanding its options and use cases can significantly enhance your data management and security practices. 
