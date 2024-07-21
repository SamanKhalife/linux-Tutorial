# UEFI shell

The **UEFI Shell** is a command-line interface provided by the UEFI (Unified Extensible Firmware Interface) firmware environment. It allows users to interact directly with the system’s firmware and perform various tasks such as managing files, running scripts, and configuring the system. The UEFI Shell is particularly useful for low-level system management and troubleshooting.

### Overview

- **Purpose**: The UEFI Shell provides a command-line environment within the UEFI firmware that can be used for various tasks like executing shell commands, running scripts, and managing the UEFI environment.
- **Environment**: It operates independently of the operating system and is accessed during the system's boot process, usually through a special UEFI menu or boot option.

### Accessing the UEFI Shell

1. **Direct Boot**:
   - Some UEFI firmware implementations include the UEFI Shell as an option in the firmware settings. You can boot into it directly from the UEFI firmware setup menu.

2. **Bootable Media**:
   - You can also create a bootable USB drive with the UEFI Shell installed. This is useful for systems that do not have a built-in UEFI Shell.

   - **Creating Bootable Media**:
     - Obtain a UEFI Shell binary (e.g., `Shell.efi`).
     - Copy it to the EFI System Partition (ESP) on a USB drive.
     - Boot from the USB drive and select the UEFI Shell.

### Common Commands

Here are some common UEFI Shell commands:

1. **`help`**:
   - Displays a list of available commands and their usage.

   ```sh
   help
   ```

2. **`ls`**:
   - Lists files and directories in the current directory.

   ```sh
   ls
   ```

3. **`cd`**:
   - Changes the current directory.

   ```sh
   cd <directory>
   ```

4. **`cp`**:
   - Copies files from one location to another.

   ```sh
   cp <source> <destination>
   ```

5. **`rm`**:
   - Removes files or directories.

   ```sh
   rm <file_or_directory>
   ```

6. **`mkdir`**:
   - Creates a new directory.

   ```sh
   mkdir <directory>
   ```

7. **`echo`**:
   - Displays a message or variable value.

   ```sh
   echo "Hello, UEFI Shell!"
   ```

8. **`map`**:
   - Lists all mapped file systems, including the EFI System Partition.

   ```sh
   map
   ```

9. **`load`**:
   - Loads a UEFI application or driver.

   ```sh
   load <filename>
   ```

10. **`unload`**:
    - Unloads a UEFI application or driver.

    ```sh
    unload <filename>
    ```

11. **`bcfg`**:
    - Manages the boot configuration, including adding or removing boot entries.

    ```sh
    bcfg boot dump
    bcfg boot add 1 fs0:\EFI\BOOT\BOOTX64.EFI "My Boot Entry"
    ```

### Scripting in UEFI Shell

UEFI Shell supports scripting, which allows you to automate tasks using shell scripts. These scripts have a `.nsh` extension and are similar to batch files in Windows or shell scripts in Unix-like systems.

**Example `.nsh` Script**:

```sh
@echo off
echo "Starting UEFI Shell script..."
cd fs0:
mkdir MyDirectory
cp MyFile.efi MyDirectory
echo "Script completed."
```

### Common Use Cases

1. **Troubleshooting**:
   - UEFI Shell can be used to diagnose and troubleshoot boot issues or firmware problems, such as testing hardware or examining system settings.

2. **Firmware Updates**:
   - Use the UEFI Shell to apply firmware updates or patches by running UEFI applications provided by the motherboard or system manufacturer.

3. **File Management**:
   - Manage files on the EFI System Partition (ESP), such as copying boot loaders or configuration files.

4. **Boot Configuration**:
   - Modify or create UEFI boot entries to change the boot order or add new boot options.

### Example Commands

1. **List File Systems**:

   ```sh
   map
   ```

2. **Change Directory and List Contents**:

   ```sh
   cd fs0:
   ls
   ```

3. **Add a New Boot Entry**:

   ```sh
   bcfg boot add 1 fs0:\EFI\BOOT\BOOTX64.EFI "New Boot Entry"
   ```

4. **Run a UEFI Application**:

   ```sh
   fs0:\EFI\BOOT\BOOTX64.EFI
   ```

### Summary

- **Purpose**: Provides a command-line environment within the UEFI firmware for system management and configuration.
- **Access**: Can be accessed via UEFI firmware setup or bootable media.
- **Commands**: Includes file management, boot configuration, and scripting.
- **Use Cases**: Troubleshooting, firmware updates, file management, and boot configuration.

The UEFI Shell is a powerful tool for advanced users and administrators who need to interact with the system at a low level, especially when dealing with firmware and boot issues.
