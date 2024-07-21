# mkisofs

`mkisofs` is a command-line utility used to create ISO 9660 filesystem images, which can be used for CD-ROMs, DVDs, and other optical media. The utility is often used for creating ISO images from a directory tree, allowing users to bundle files and directories into an ISO image that can be burned to an optical disc or used as a virtual filesystem.

### Key Features

- **Create ISO Images**: Generates ISO 9660 filesystem images from files and directories.
- **Customize Filesystem**: Allows customization of various aspects of the ISO filesystem, such as volume labels and file attributes.
- **Compatibility**: Supports different ISO standards and extensions like Rock Ridge (for UNIX-like systems) and Joliet (for Windows).

### Basic Usage

The general syntax for `mkisofs` is:

```sh
mkisofs [options] -o output.iso input_directory
```

- **`[options]`**: Various command-line options to customize the ISO image.
- **`-o output.iso`**: Specifies the name of the output ISO file.
- **`input_directory`**: The directory whose contents will be included in the ISO image.

### Common Options

- **`-o <file>`**: Specifies the output ISO file name.
- **`-J`**: Enables Joliet extensions, which allow longer filenames and additional character sets.
- **`-R`**: Enables Rock Ridge extensions for UNIX-like systems, supporting long filenames and file permissions.
- **`-V <volume_label>`**: Sets the volume label for the ISO image.
- **`-b <boot_image>`**: Specifies a boot image for creating a bootable ISO.
- **`-no-emul-boot`**: Indicates no emulation of boot image; used with `-b` for creating bootable ISOs.
- **`-iso-level <level>`**: Sets the ISO 9660 level (e.g., 1 for standard, 2 for extensions, 3 for further extensions).
- **`-allow-lowercase`**: Allows lowercase characters in filenames (useful for cross-platform compatibility).
- **`-d`**: Allows the use of filenames with non-printable characters.

### Examples

#### Create a Basic ISO Image

```sh
mkisofs -o mydisk.iso /path/to/directory
```

This command creates an ISO image named `mydisk.iso` from the contents of `/path/to/directory`.

#### Create an ISO Image with Joliet and Rock Ridge Extensions

```sh
mkisofs -o mydisk.iso -J -R /path/to/directory
```

This command includes both Joliet and Rock Ridge extensions, allowing for better compatibility with both Windows and UNIX-like systems.

#### Create a Bootable ISO Image

```sh
mkisofs -o bootable.iso -b boot.img -no-emul-boot /path/to/directory
```

This command creates a bootable ISO image named `bootable.iso` using `boot.img` as the boot image.

### Advanced Usage

- **Multi-Session ISOs**: `mkisofs` does not handle multi-session ISOs. Use tools like `cdrecord` or `genisoimage` for creating multi-session images.
- **Hybrid Images**: For creating hybrid ISOs that work as both ISO images and bootable USBs, consider using `genisoimage` with specific options.

### Comparison with `genisoimage`

`mkisofs` is essentially a part of the `genisoimage` package. While `mkisofs` is the traditional tool, `genisoimage` is often used as it incorporates improvements and additional features. The syntax and options for `genisoimage` are similar to `mkisofs`.

### Summary

`mkisofs` is a versatile tool for creating ISO 9660 filesystem images from directories. By leveraging its various options, users can create standard ISO images, bootable ISOs, and images with extended file attributes for better cross-platform compatibility. For most use cases involving ISO image creation, `mkisofs` provides a powerful and flexible solution.
