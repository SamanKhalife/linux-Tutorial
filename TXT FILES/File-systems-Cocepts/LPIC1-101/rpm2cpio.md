# rpm2cpio

The `rpm2cpio` command is a utility that converts an RPM (Red Hat Package Manager) package to a CPIO archive. This can be particularly useful when you want to extract the contents of an RPM package without installing it on your system. The extracted contents can then be examined or used as needed.

### Basic Usage

The basic syntax for `rpm2cpio` is:

```sh
rpm2cpio package_name.rpm | cpio -idmv
```

This command pipeline converts the RPM package to a CPIO archive and then extracts the archive using the `cpio` command.

### Example Usage

#### Extracting an RPM Package

1. **Download or have the RPM package**: Ensure you have the RPM package file you want to extract. For example, `example-1.0-1.x86_64.rpm`.

2. **Convert and Extract**: Use the following command to convert the RPM package to a CPIO archive and extract it:

    ```sh
    rpm2cpio example-1.0-1.x86_64.rpm | cpio -idmv
    ```

    - **`-i`**: Extract files from the archive.
    - **`-d`**: Create leading directories where needed.
    - **`-m`**: Retain previous file modification times.
    - **`-v`**: Verbosely list files processed.

#### Breaking Down the Command

1. **Convert RPM to CPIO**: The `rpm2cpio` command reads the RPM package and converts it to a CPIO archive format.

2. **Extract CPIO Archive**: The `cpio` command reads the CPIO archive from the standard input and extracts the files.

### Practical Example

Suppose you have an RPM package named `example-1.0-1.x86_64.rpm` and you want to extract its contents:

1. **Place the RPM file in a directory** where you want to extract its contents, or navigate to the directory where the RPM is located.

2. **Run the extraction command**:

    ```sh
    rpm2cpio example-1.0-1.x86_64.rpm | cpio -idmv
    ```

After running this command, you will see a list of files being extracted, and the contents of the RPM package will be extracted to the current directory.

### Additional Notes

- **Inspecting Contents**: You can inspect the contents of an RPM package before extracting by using the `rpm` command with the `-qlp` option:

    ```sh
    rpm -qlp package_name.rpm
    ```

    This command lists the files contained in the RPM package without extracting them.

- **Creating a Directory for Extraction**: To keep things organized, you might want to create a separate directory for extraction:

    ```sh
    mkdir extracted_files
    cd extracted_files
    rpm2cpio ../package_name.rpm | cpio -idmv
    ```

    This will extract the contents into the `extracted_files` directory.

### Conclusion

The `rpm2cpio` command is a useful tool for extracting the contents of RPM packages without installing them. By converting the RPM package to a CPIO archive and then using the `cpio` command to extract it, you can easily access the files within the package.
