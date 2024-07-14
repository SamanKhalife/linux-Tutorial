# grub.cfg
`grub.cfg` is the main configuration file for GRUB 2 (GRand Unified Bootloader version 2), which is the default bootloader for most Linux distributions. This file contains the configuration directives that define the boot menu entries and the various boot parameters for the operating system.

### Overview of `grub.cfg`

#### Location

The `grub.cfg` file is typically located in the `/boot/grub/` or `/boot/grub2/` directory, depending on the Linux distribution and version.

### Purpose

The primary purpose of `grub.cfg` is to define the boot menu entries and the associated parameters for each entry. It allows users to select which operating system or kernel to boot into and pass specific parameters to the kernel at boot time.

### Structure and Basic Configuration

The `grub.cfg` file is automatically generated and should not be manually edited. Instead, users should modify the configuration files located in `/etc/default/grub` and the scripts in `/etc/grub.d/` and then regenerate `grub.cfg` using the `grub-mkconfig` command.

#### Basic Sections

1. **Global Settings**: Defines global parameters such as default boot entry, timeout, and appearance settings.
2. **Menu Entries**: Defines individual boot entries, specifying the kernel, initrd, and other boot parameters for each entry.

### Example of `grub.cfg`

Here is an example of a `grub.cfg` file:

```cfg
# Set the default boot entry
set default=0

# Set the boot menu timeout
set timeout=5

# Set the background image
if background_image /boot/grub/background.png; then
  set color_normal=light-gray/black
  set color_highlight=white/black
else
  set menu_color_normal=cyan/blue
  set menu_color_highlight=white/blue
fi

# Boot menu entries
menuentry 'Ubuntu, with Linux 5.4.0-42-generic' --class ubuntu --class gnu-linux --class gnu --class os {
  recordfail
  load_video
  gfxmode $linux_gfx_mode
  insmod gzio
  insmod part_msdos
  insmod ext2
  set root='hd0,msdos1'
  if [ x$feature_platform_search_hint = xy ]; then
    search --no-floppy --fs-uuid --set=root d1e2e1b0-51b1-45a4-8d11-7d153a6210b3
  else
    search --no-floppy --fs-uuid --set=root d1e2e1b0-51b1-45a4-8d11-7d153a6210b3
  fi
  linux   /boot/vmlinuz-5.4.0-42-generic root=UUID=d1e2e1b0-51b1-45a4-8d11-7d153a6210b3 ro  quiet splash $vt_handoff
  initrd  /boot/initrd.img-5.4.0-42-generic
}
```

### Key Directives

- **set default=0**: Sets the default boot entry (0 corresponds to the first entry).
- **set timeout=5**: Sets the timeout before the default entry is automatically booted.
- **menuentry 'Ubuntu, with Linux 5.4.0-42-generic'**: Defines a boot menu entry.
- **linux /boot/vmlinuz-5.4.0-42-generic**: Specifies the kernel to be booted.
- **initrd /boot/initrd.img-5.4.0-42-generic**: Specifies the initial RAM disk.

### Generating `grub.cfg`

To regenerate `grub.cfg`, use the following command:

```bash
sudo grub-mkconfig -o /boot/grub/grub.cfg
```

This command reads the configuration files from `/etc/default/grub` and the scripts in `/etc/grub.d/`, and then generates a new `grub.cfg` file based on these settings.

### Customizing GRUB Configuration

To customize the GRUB configuration:

1. **Edit `/etc/default/grub`**: Modify general settings like the default boot entry, timeout, and kernel parameters.

   Example:

   ```cfg
   GRUB_DEFAULT=0
   GRUB_TIMEOUT=5
   GRUB_CMDLINE_LINUX_DEFAULT="quiet splash"
   GRUB_CMDLINE_LINUX=""
   ```

2. **Edit scripts in `/etc/grub.d/`**: Add or modify specific menu entries and advanced settings.

3. **Regenerate `grub.cfg`**: Apply the changes by running:

   ```bash
   sudo grub-mkconfig -o /boot/grub/grub.cfg
   ```

### Conclusion

`grub.cfg` is an essential file for the GRUB 2 bootloader, containing configuration settings and boot menu entries. Properly managing and customizing `grub.cfg` through the appropriate configuration files and scripts ensures a smooth and flexible boot process for Linux systems.
