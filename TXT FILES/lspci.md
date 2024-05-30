# lspci

The `lspci` command in Linux is used to display information about all PCI (Peripheral Component Interconnect) buses in the system and all devices connected to them. It is a powerful tool for diagnosing hardware issues and for obtaining detailed information about PCI devices.

### Understanding `lspci`

The `lspci` command provides details about PCI devices, which include network cards, sound cards, USB controllers, and other hardware connected to the PCI bus.

#### Basic Usage

To display a list of all PCI devices, simply run:
```sh
lspci
```

#### Example Output

Here’s an example of what the `lspci` command might return:
```sh
00:00.0 Host bridge: Intel Corporation 6th-9th Gen Core Processor PCIe Controller (rev 07)
00:02.0 VGA compatible controller: Intel Corporation HD Graphics 530 (rev 06)
00:14.0 USB controller: Intel Corporation 100 Series/C230 Series Chipset Family USB 3.0 xHCI Controller (rev 31)
00:16.0 Communication controller: Intel Corporation 100 Series/C230 Series Chipset Family MEI Controller #1 (rev 31)
00:1f.2 SATA controller: Intel Corporation 100 Series/C230 Series Chipset Family SATA AHCI Controller (rev 31)
```

#### Detailed Output

To get more detailed information about the PCI devices, use the `-v` (verbose) option:
```sh
lspci -v
```

#### Example Detailed Output

Here’s an example of detailed output:
```sh
00:02.0 VGA compatible controller: Intel Corporation HD Graphics 530 (rev 06) (prog-if 00 [VGA controller])
    Subsystem: Dell Device 06b9
    Flags: bus master, fast devsel, latency 0, IRQ 128
    Memory at d0000000 (64-bit, non-prefetchable) [size=16M]
    Memory at c0000000 (64-bit, prefetchable) [size=256M]
    I/O ports at e000 [size=64]
    Expansion ROM at 000c0000 [virtual] [disabled] [size=128K]
    Capabilities: <access denied>
    Kernel driver in use: i915
    Kernel modules: i915
```

#### Key Options

1. **-v**: Verbose output. Displays detailed information about each device.
   ```sh
   lspci -v
   ```

2. **-vv**: Very verbose output. Displays even more detailed information.
   ```sh
   lspci -vv
   ```

3. **-k**: Show kernel drivers handling each device.
   ```sh
   lspci -k
   ```

4. **-t**: Show a tree view of the devices.
   ```sh
   lspci -t
   ```

5. **-nn**: Show numeric IDs for devices.
   ```sh
   lspci -nn
   ```

6. **-s [bus:device.function]**: Show only devices in the specified slot.
   ```sh
   lspci -s 00:02.0
   ```

7. **-d [vendor:device]**: Show only devices with the specified vendor and device ID.
   ```sh
   lspci -d 8086:1237
   ```

### Practical Examples

1. **Display Information About a Specific Device**:
   ```sh
   lspci -s 00:14.0
   ```
   This command will display information about the device located at bus 00, device 14, function 0.

2. **Show Kernel Drivers Handling Devices**:
   ```sh
   lspci -k
   ```
   This command will show which kernel drivers are being used for each device.

3. **List All Devices with Vendor and Device IDs**:
   ```sh
   lspci -nn
   ```
   This command will display all devices with their numeric vendor and device IDs.

4. **Show PCI Device Tree**:
   ```sh
   lspci -t
   ```
   This command displays a tree view of the devices, showing the hierarchy and relationships between them.

5. **Filter by Vendor and Device ID**:
   ```sh
   lspci -d 8086:1237
   ```
   This command will list devices that match the specified vendor (8086) and device (1237) IDs.

### Related Commands

1. **lspci -vv**: For very verbose output, which includes detailed information about each device's configuration space.
   ```sh
   lspci -vv
   ```

2. **lspci -xxx**: To display the raw hexadecimal dump of the PCI configuration space.
   ```sh
   lspci -xxx
   ```

3. **lspci -b**: To display the IRQ numbers in decimal instead of hexadecimal.
   ```sh
   lspci -b
   ```

4. **lshw**: Another command that provides detailed information about hardware.
   ```sh
   sudo lshw -short
   ```

### Conclusion

The `lspci` command is a vital tool for system administrators and users who need to gather detailed information about the PCI devices on their Linux systems. It is particularly useful for diagnosing hardware issues, verifying device drivers, and understanding the hardware configuration.
