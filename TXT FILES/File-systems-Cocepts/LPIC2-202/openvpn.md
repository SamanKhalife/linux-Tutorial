# OpenVPN

OpenVPN is an open-source software application that implements virtual private network (VPN) techniques for creating secure point-to-point or site-to-site connections. It can be used to connect remote users or sites securely over the internet. This guide covers the installation, configuration, and use of OpenVPN.

### Installation

#### On Debian/Ubuntu:

```bash
sudo apt update
sudo apt install openvpn
```

#### On CentOS/RHEL:

```bash
sudo yum install epel-release
sudo yum install openvpn
```

### Configuration

#### Server Configuration

1. **Generate Server Certificates and Keys**

   OpenVPN uses the OpenSSL library to provide encryption, and you need to set up a Certificate Authority (CA) to generate server and client certificates and keys.

   ```bash
   wget https://github.com/OpenVPN/easy-rsa/releases/download/v3.0.8/EasyRSA-3.0.8.tgz
   tar xvf EasyRSA-3.0.8.tgz
   cd EasyRSA-3.0.8/
   ./easyrsa init-pki
   ./easyrsa build-ca nopass
   ./easyrsa gen-req server nopass
   ./easyrsa sign-req server server
   ./easyrsa gen-dh
   openvpn --genkey --secret ta.key
   ```

2. **Configure the Server**

   Create the server configuration file, usually located at `/etc/openvpn/server.conf`:

   ```bash
   sudo nano /etc/openvpn/server.conf
   ```

   Example configuration:

   ```conf
   port 1194
   proto udp
   dev tun
   ca ca.crt
   cert server.crt
   key server.key
   dh dh.pem
   server 10.8.0.0 255.255.255.0
   ifconfig-pool-persist ipp.txt
   push "redirect-gateway def1 bypass-dhcp"
   push "dhcp-option DNS 8.8.8.8"
   push "dhcp-option DNS 8.8.4.4"
   keepalive 10 120
   cipher AES-256-CBC
   user nobody
   group nogroup
   persist-key
   persist-tun
   status openvpn-status.log
   log-append /var/log/openvpn.log
   verb 3
   ```

3. **Start and Enable OpenVPN Service**

   ```bash
   sudo systemctl start openvpn@server
   sudo systemctl enable openvpn@server
   ```

#### Client Configuration

1. **Generate Client Certificates and Keys**

   ```bash
   ./easyrsa gen-req client1 nopass
   ./easyrsa sign-req client client1
   ```

2. **Create the Client Configuration File**

   Create the client configuration file, usually located at `/etc/openvpn/client/client1.ovpn`:

   ```bash
   sudo nano /etc/openvpn/client/client1.ovpn
   ```

   Example configuration:

   ```conf
   client
   dev tun
   proto udp
   remote your_server_ip 1194
   resolv-retry infinite
   nobind
   user nobody
   group nogroup
   persist-key
   persist-tun
   ca ca.crt
   cert client1.crt
   key client1.key
   remote-cert-tls server
   cipher AES-256-CBC
   verb 3
   ```

3. **Transfer Configuration and Certificates to Client**

   Transfer the configuration file (`client1.ovpn`), CA certificate (`ca.crt`), client certificate (`client1.crt`), and client key (`client1.key`) to the client machine.

4. **Start the OpenVPN Client**

   On the client machine, start OpenVPN with the configuration file:

   ```bash
   sudo openvpn --config /path/to/client1.ovpn
   ```

### Management and Usage

#### OpenVPN Management Commands

- **Start OpenVPN Service**
  ```bash
  sudo systemctl start openvpn@server
  ```

- **Stop OpenVPN Service**
  ```bash
  sudo systemctl stop openvpn@server
  ```

- **Enable OpenVPN Service at Boot**
  ```bash
  sudo systemctl enable openvpn@server
  ```

- **Disable OpenVPN Service at Boot**
  ```bash
  sudo systemctl disable openvpn@server
  ```

- **Check OpenVPN Service Status**
  ```bash
  sudo systemctl status openvpn@server
  ```

### Security Considerations

- **Regularly Update OpenVPN**: Ensure that OpenVPN and all dependencies are regularly updated to the latest version to benefit from security patches and improvements.
- **Use Strong Encryption**: Use strong encryption methods such as AES-256-CBC.
- **Firewall Configuration**: Ensure that the necessary ports (e.g., 1194) are open on your firewall to allow OpenVPN traffic.
- **Secure Certificates and Keys**: Protect the CA, server, and client certificates and keys. Never share private keys.

### Conclusion

OpenVPN is a versatile and secure tool for setting up VPN connections. By following the installation and configuration steps, you can create a robust VPN solution for connecting remote users or sites securely. Regular maintenance and security practices will ensure the VPN remains secure and reliable.
