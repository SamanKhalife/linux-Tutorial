# GRUB
GRUB (GRand Unified Bootloader) is a widely used bootloader that allows users to boot multiple operating systems on a computer. It provides a flexible and powerful way to control the boot process, making it a crucial component in many Linux distributions. Here’s a detailed explanation of GRUB, its configuration, and usage:

### Key Features of GRUB

1. **Multiboot Support**: GRUB can boot various operating systems, including different versions of Linux, Windows, and other Unix-like systems.
2. **Extensible and Modular**: GRUB’s modular design allows for extensions and customizations, making it versatile.
3. **Command-Line Interface**: GRUB provides a command-line interface for advanced troubleshooting and configuration.
4. **Graphical Menu Interface**: GRUB can display a graphical menu to select from available boot entries.

### GRUB Configuration Files

#### grub.cfg

The primary configuration file for GRUB is `grub.cfg`, usually located in `/boot/grub` or `/boot/efi/EFI/<distro>` on UEFI systems.

```bash
/boot/grub/grub.cfg
```

This file is auto-generated and should not be edited directly. Instead, make changes to the source files that generate `grub.cfg` and regenerate it.

#### Source Configuration Files

- **/etc/default/grub**: Contains main configuration settings for GRUB.
- **/etc/grub.d/**: Contains scripts that generate `grub.cfg`.

### Configuring GRUB

#### /etc/default/grub

This file allows you to set various options for GRUB. Here are some common settings:

```bash
GRUB_DEFAULT=0
GRUB_TIMEOUT=5
GRUB_CMDLINE_LINUX_DEFAULT="quiet splash"
GRUB_CMDLINE_LINUX=""
```

- **GRUB_DEFAULT**: Sets the default menu entry (by index or name).
- **GRUB_TIMEOUT**: Sets the timeout before the default entry is booted.
- **GRUB_CMDLINE_LINUX_DEFAULT**: Adds kernel parameters for the default boot.
- **GRUB_CMDLINE_LINUX**: Adds kernel parameters for all boot entries.

#### /etc/grub.d/

This directory contains scripts used to create `grub.cfg`. Common scripts include:

- **00_header**: Sets up the initial configuration.
- **10_linux**: Adds entries for Linux kernels.
- **30_os-prober**: Adds entries for other operating systems.

### Updating GRUB Configuration

After modifying the source configuration files, update `grub.cfg` by running:

```bash
sudo update-grub
```

This command regenerates `grub.cfg` based on the scripts in `/etc/grub.d/` and settings in `/etc/default/grub`.

### GRUB Commands

#### Booting from the Command Line

In case of boot issues, GRUB’s command-line interface can be used to manually boot an operating system.

1. **Access GRUB Command Line**: Press `c` at the GRUB menu.
2. **Load the Kernel**:
   ```bash
   linux /vmlinuz-<kernel-version> root=/dev/sda1
   ```
3. **Load the Initial RAM Disk**:
   ```bash
   initrd /initrd.img-<kernel-version>
   ```
4. **Boot**:
   ```bash
   boot
   ```

#### Adding a Boot Entry

You can add a custom boot entry to `/etc/grub.d/40_custom`:

```bash
menuentry 'Custom Linux' {
    set root='hd0,1'
    linux /vmlinuz-<kernel-version> root=/dev/sda1
    initrd /initrd.img-<kernel-version>
}
```

Then update GRUB:

```bash
sudo update-grub
```

### Troubleshooting GRUB

#### Recovering GRUB

If GRUB is broken (e.g., after installing a new OS), you can recover it using a live CD/USB:

1. **Boot from Live CD/USB**.
2. **Mount the Root Filesystem**:
   ```bash
   sudo mount /dev/sda1 /mnt
   ```
3. **Mount Other Necessary Filesystems**:
   ```bash
   sudo mount --bind /dev /mnt/dev
   sudo mount --bind /proc /mnt/proc
   sudo mount --bind /sys /mnt/sys
   ```
4. **Chroot into the Mounted Filesystem**:
   ```bash
   sudo chroot /mnt
   ```
5. **Reinstall GRUB**:
   ```bash
   grub-install /dev/sda
   update-grub
   ```

#### Fixing Boot Entries

If an OS is not appearing in the GRUB menu, run:

```bash
sudo os-prober
sudo update-grub
```

### Conclusion

GRUB is a powerful and flexible bootloader that provides a robust way to manage the boot process for multiple operating systems. Understanding its configuration files, commands, and troubleshooting techniques is essential for Linux system administration and ensuring smooth system boots.
