# efibootmgr
`efibootmgr` is a command-line utility for managing UEFI boot entries. It allows users to create, modify, and delete boot entries, as well as change the boot order for UEFI-based systems. This tool is essential for managing the boot process on systems with UEFI firmware.

### Basic Usage

The basic syntax for `efibootmgr` is:

```bash
efibootmgr [options]
```

### Common Options

- `-v, --verbose`: Increase verbosity, showing detailed information about the boot entries.
- `-b ####, --bootnum ####`: Specify the boot entry number to modify (e.g., `0000`).
- `-B, --delete-bootnum`: Delete the specified boot entry.
- `-c, --create`: Create a new boot entry.
- `-n ####, --bootnext ####`: Set the boot entry for the next boot only.
- `-o ####,####,..., --bootorder ####,####,...`: Set the boot order.
- `-t ####, --timeout ####`: Set the boot manager timeout (in seconds).

### Examples

#### Viewing Current Boot Entries

To list all current boot entries and their details:

```bash
efibootmgr -v
```

This command displays all UEFI boot entries, including the boot order, boot entries, and the associated details.

#### Creating a New Boot Entry

To create a new boot entry for a Linux kernel:

```bash
efibootmgr --create --disk /dev/sda --part 1 --label "MyLinux" --loader /vmlinuz-5.4.0-26-generic
```

- `--disk /dev/sda`: Specifies the disk where the boot loader is located.
- `--part 1`: Specifies the partition number where the boot loader is located.
- `--label "MyLinux"`: Provides a label for the boot entry.
- `--loader /vmlinuz-5.4.0-26-generic`: Specifies the path to the boot loader.

#### Deleting a Boot Entry

To delete a specific boot entry (e.g., `Boot0003`):

```bash
efibootmgr --delete-bootnum --bootnum 0003
```

#### Setting the Boot Order

To change the boot order, you can specify a new sequence of boot entry numbers:

```bash
efibootmgr --bootorder 0001,0002,0003
```

This command sets the boot order to prioritize `Boot0001`, followed by `Boot0002`, and then `Boot0003`.

#### Setting the Next Boot Entry

To set a specific boot entry to be used for the next boot only:

```bash
efibootmgr --bootnext 0002
```

This command will set `Boot0002` as the boot entry for the next boot. After that boot, the system will revert to the regular boot order.

### Troubleshooting

#### Common Issues

1. **Permission Denied**: Ensure you run `efibootmgr` with root privileges, typically by using `sudo`.
2. **No EFI Variables Found**: This usually indicates that the system is not booted in UEFI mode. Ensure your system is in UEFI mode and not legacy BIOS mode.
3. **EFI Variables Are Read-Only**: Some systems may lock the EFI variables, preventing modifications. This can often be resolved by changing settings in the UEFI firmware or ensuring the correct boot mode.

### Conclusion

`efibootmgr` is a powerful utility for managing UEFI boot entries and the boot order on UEFI-based systems. It's particularly useful for managing multi-boot environments, recovering from boot issues, and configuring new boot entries for custom kernels or operating systems. Understanding how to use `efibootmgr` effectively can significantly enhance your ability to control the boot process on modern systems.
