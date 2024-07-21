# iwlist

The `iwlist` command is a part of the `wireless-tools` package on Linux systems and is used for querying wireless network interfaces for various information. It's often used in conjunction with the `iwconfig` command, which is used to configure wireless network interfaces.

#### Basic Usage

To use `iwlist`, you typically need to specify the wireless interface you want to query and the type of information you want to retrieve.

#### Common Commands and Options

##### 1. **Scan for Wireless Networks**

- **Command:**
  ```sh
  iwlist <interface> scan
  ```
- **Description:**
  Scans for available wireless networks and displays information about each network.
- **Example:**
  ```sh
  iwlist wlan0 scan
  ```
- **Example Output:**
  ```
  wlan0     Scan completed :
            Cell 01 - Address: 00:11:22:33:44:55
                      ESSID:"NetworkName"
                      Mode:Master
                      Frequency:2.412 GHz (Channel 1)
                      Quality=70/70  Signal level=-30 dBm  
                      Encryption key:on
                      IE: IEEE 802.11i/WPA2 Version 1
                      ...
  ```

##### 2. **Show Wireless Interface Information**

- **Command:**
  ```sh
  iwlist <interface> <info_type>
  ```
- **Description:**
  Displays various types of information about the wireless interface. Common info types include:
  - `scan`: Shows available networks (already covered).
  - `freq`: Displays the available frequencies.
  - `power`: Shows power settings.
- **Example:**
  ```sh
  iwlist wlan0 freq
  ```
- **Example Output:**
  ```
  wlan0     2.412 GHz
            2.417 GHz
            2.422 GHz
            ...
  ```

##### 3. **Display Detailed Wireless Network Information**

- **Command:**
  ```sh
  iwlist <interface> scan | grep -i <keyword>
  ```
- **Description:**
  Allows filtering the scan results to find specific information.
- **Example:**
  ```sh
  iwlist wlan0 scan | grep -i "SSID"
  ```
- **Example Output:**
  ```
  SSID:"NetworkName"
  SSID:"AnotherNetwork"
  ```

##### 4. **Display Signal Strength and Quality**

- **Command:**
  ```sh
  iwlist <interface> scan | grep -E "Quality|Signal level"
  ```
- **Description:**
  Filters the scan results to display only the signal strength and quality of networks.
- **Example:**
  ```sh
  iwlist wlan0 scan | grep -E "Quality|Signal level"
  ```
- **Example Output:**
  ```
  Quality=70/70  Signal level=-30 dBm  
  Quality=40/70  Signal level=-60 dBm
  ```

#### Important Concepts

- **ESSID**: Extended Service Set Identifier. It is the name of the wireless network.
- **Mode**: Indicates the mode of operation, such as Master (AP) or Managed (Client).
- **Frequency**: The frequency on which the network operates.
- **Quality**: Indicates the signal quality of the network, typically a percentage.
- **Signal Level**: The strength of the signal received from the network, measured in dBm (decibels relative to a milliwatt).
- **Encryption Key**: Indicates whether the network is secured with encryption.

#### Example Scenarios

1. **Finding Available Networks:**
   - Use `iwlist` to scan for and list all available wireless networks in the vicinity.
   
2. **Checking Network Strength:**
   - Evaluate the signal strength and quality of nearby networks to determine the best one to connect to.

3. **Determining Frequency Bands:**
   - Check which frequency bands are in use by scanning for available frequencies.

4. **Troubleshooting Wireless Connections:**
   - Analyze scan results to troubleshoot connectivity issues or find out why a network might not be appearing.

### Conclusion

The `iwlist` command is a powerful tool for querying wireless network interfaces and gathering detailed information about available networks and their characteristics. By understanding and using `iwlist`, network administrators and users can effectively manage and troubleshoot wireless network connections.
