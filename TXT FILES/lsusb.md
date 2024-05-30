# lsusb

The `lsusb` command in Linux is used to display information about USB (Universal Serial Bus) buses in the system and the devices connected to them. It's a valuable tool for diagnosing and managing USB devices, providing detailed insights into the USB hardware connected to your system.

### Understanding `lsusb`

The `lsusb` command lists all the USB devices currently connected to your system, showing details about each device such as the vendor ID, product ID, and device description.

#### Basic Usage

To display a list of all USB devices, simply run:
```sh
lsusb
```

#### Example Output

Here’s an example of what the `lsusb` command might return:
```sh
Bus 002 Device 004: ID 8087:0024 Intel Corp. Integrated Rate Matching Hub
Bus 002 Device 003: ID 04f2:b2f4 Chicony Electronics Co., Ltd 
Bus 002 Device 002: ID 8087:0024 Intel Corp. Integrated Rate Matching Hub
Bus 002 Device 001: ID 1d6b:0002 Linux Foundation 2.0 root hub
Bus 001 Device 003: ID 0bda:58b0 Realtek Semiconductor Corp. 
Bus 001 Device 002: ID 1bcf:0005 Sunplus Innovation Technology Inc. Optical Mouse
Bus 001 Device 001: ID 1d6b:0002 Linux Foundation 2.0 root hub
```

#### Detailed Output

To get more detailed information about the USB devices, use the `-v` (verbose) option:
```sh
lsusb -v
```

#### Example Detailed Output

Here’s an example of detailed output for one device:
```sh
Bus 002 Device 004: ID 8087:0024 Intel Corp. Integrated Rate Matching Hub
Device Descriptor:
  bLength                18
  bDescriptorType         1
  bcdUSB               2.00
  bDeviceClass            9 Hub
  bDeviceSubClass         0 
  bDeviceProtocol         1 Full speed (or root) hub
  bMaxPacketSize0        64
  idVendor           0x8087 Intel Corp.
  idProduct          0x0024 Integrated Rate Matching Hub
  bcdDevice            0.00
  iManufacturer           0 
  iProduct                0 
  iSerial                 0 
  bNumConfigurations      1
```

#### Key Options

1. **-v**: Verbose output. Displays detailed information about each device.
   ```sh
   lsusb -v
   ```

2. **-t**: Show the hierarchical tree of devices.
   ```sh
   lsusb -t
   ```

3. **-s [bus]:[devnum]**: Show only devices in the specified bus and device number.
   ```sh
   lsusb -s 002:004
   ```

4. **-d [vendor]:[product]**: Show only devices with the specified vendor and product ID.
   ```sh
   lsusb -d 8087:0024
   ```

5. **-D [device]**: Show detailed information about the specified device file.
   ```sh
   lsusb -D /dev/bus/usb/002/004
   ```

### Practical Examples

1. **List All USB Devices**:
   ```sh
   lsusb
   ```
   This command lists all the currently connected USB devices.

2. **Display Detailed Information About All USB Devices**:
   ```sh
   sudo lsusb -v
   ```
   This command provides detailed information about all USB devices. Note that `sudo` may be required for full details.

3. **Show USB Device Tree**:
   ```sh
   lsusb -t
   ```
   This command shows a hierarchical tree of USB devices.

4. **Filter by Vendor and Product ID**:
   ```sh
   lsusb -d 8087:0024
   ```
   This command will list devices that match the specified vendor (8087) and product (0024) IDs.

5. **Detailed Information for a Specific Device**:
   ```sh
   lsusb -s 002:004
   ```
   This command displays information about the device on bus 002 with device number 004.

### Related Commands

1. **usb-devices**:
   Displays detailed information about USB devices. It's more human-readable than `lsusb -v`.
   ```sh
   usb-devices
   ```

2. **dmesg**:
   Displays system messages, including those related to USB device connections and disconnections.
   ```sh
   dmesg | grep -i usb
   ```

3. **udevadm**:
   Monitors udev events. Useful for diagnosing USB device issues.
   ```sh
   udevadm monitor --udev
   ```

### Conclusion

The `lsusb` command is a powerful tool for managing and diagnosing USB devices on Linux systems. It provides a quick way to list all USB devices and offers detailed information for troubleshooting and verifying hardware.


# help 

```
lsusb [options]

List USB devices.

Options:

-v, --verbose   Print detailed information about devices.
-t, --tree      Print a tree-like representation of the device hierarchy.
-h, --help       Show this help message.

For more information, see the lsusb man page.
```

