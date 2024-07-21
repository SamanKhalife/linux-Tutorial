# iw

`iw` is a command-line tool used for configuring and managing wireless devices in Linux. It supports the modern nl80211 interface and is intended to replace older tools like `iwconfig`. The `iw` command provides more advanced and detailed control over wireless interfaces and supports the latest features of wireless hardware.

#### Basic Usage

To use `iw`, you typically need root privileges. The general syntax of the `iw` command is:

```sh
iw [OPTIONS] COMMAND [COMMAND_OPTIONS]
```

### Common Commands and Options

1. **List Devices**
   - List all wireless devices:
     ```sh
     iw dev
     ```
   - Example output:
     ```
     phy#0
        Interface wlan0
            ifindex 3
            wdev 0x1
            addr 00:11:22:33:44:55
            type managed
            txpower 20.00 dBm
     ```

2. **Link Status**
   - Show information about the current link:
     ```sh
     iw dev wlan0 link
     ```
   - Example output:
     ```
     Connected to 00:11:22:33:44:55 (on wlan0)
     SSID: MyNetwork
     freq: 2412
     RX: 1234567 bytes (2345 packets)
     TX: 987654 bytes (3456 packets)
     signal: -40 dBm
     tx bitrate: 54.0 MBit/s
     ```

3. **Scan for Networks**
   - Scan for available wireless networks:
     ```sh
     iw dev wlan0 scan
     ```
   - Example output:
     ```
     BSS 00:11:22:33:44:55(on wlan0)
        TSF: 123456789012 usec (0d, 12:34:56)
        freq: 2412
        beacon interval: 100 TUs
        capability: ESS (0x0421)
        signal: -40 dBm
        last seen: 50 ms ago
        SSID: MyNetwork
        Supported rates: 1.0 2.0 5.5 11.0 22.0 36.0 48.0 54.0
        DS Parameter set: channel 1
        RSN:     * Version: 1
                * Group cipher: CCMP
                * Pairwise ciphers: CCMP
                * Authentication suites: PSK
                * Capabilities: 16-PTKSA-RC 1-GTKSA-RC (0x0000)
     ```

4. **Set Interface Down**
   - Bring a wireless interface down:
     ```sh
     iw dev wlan0 set type managed
     ```
   - Example:
     ```sh
     ip link set wlan0 down
     ```

5. **Set Interface Up**
   - Bring a wireless interface up:
     ```sh
     iw dev wlan0 set type managed
     ```
   - Example:
     ```sh
     ip link set wlan0 up
     ```

6. **Join a Network**
   - Connect to a wireless network (requires WPA supplicant for encrypted networks):
     ```sh
     iw dev wlan0 connect MyNetwork
     ```
   - For encrypted networks, additional configuration with WPA supplicant or network manager is needed.

7. **Disconnect from a Network**
   - Disconnect from the current network:
     ```sh
     iw dev wlan0 disconnect
     ```

8. **Set Wireless Power Save Mode**
   - Enable power save mode:
     ```sh
     iw dev wlan0 set power_save on
     ```
   - Disable power save mode:
     ```sh
     iw dev wlan0 set power_save off
     ```

9. **Show Station Information**
   - Show information about connected stations (for access points):
     ```sh
     iw dev wlan0 station dump
     ```

10. **Change Wireless Interface Type**
    - Change the type of a wireless interface (e.g., from managed to monitor):
      ```sh
      iw dev wlan0 set type monitor
      ```

#### Example Configuration Workflow

1. **List Available Wireless Interfaces**
   ```sh
   iw dev
   ```

2. **Bring the Interface Down**
   ```sh
   ip link set wlan0 down
   ```

3. **Set Interface Type**
   ```sh
   iw dev wlan0 set type managed
   ```

4. **Bring the Interface Up**
   ```sh
   ip link set wlan0 up
   ```

5. **Scan for Available Networks**
   ```sh
   iw dev wlan0 scan
   ```

6. **Connect to a Network**
   ```sh
   iw dev wlan0 connect MyNetwork
   ```

### Conclusion

The `iw` tool is a powerful utility for managing and configuring wireless devices in Linux. It provides comprehensive control over wireless interfaces and supports advanced features. By mastering `iw`, administrators can effectively manage wireless connectivity and troubleshoot issues.
