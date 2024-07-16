# openvas-mkcert

The `openvas-mkcert` command is used to create SSL certificates for securing communications between OpenVAS components. This includes generating a new Certificate Authority (CA) certificate and server certificate, which are essential for encrypting data transmitted during vulnerability scans and other operations.

#### Basic Usage

To generate SSL certificates for OpenVAS, use the `openvas-mkcert` command:

```bash
sudo openvas-mkcert
```

This command will guide you through the process of creating a new CA certificate and server certificate. 

#### Steps

1. **Run the Command**:
    ```bash
    sudo openvas-mkcert
    ```

2. **Interactive Prompts**:
    The command will prompt you to enter various details for the certificates. These details include:

    - **Country Name**: The two-letter ISO code for your country.
    - **State or Province Name**: The full name of your state or province.
    - **Locality Name**: The name of your city or locality.
    - **Organization Name**: The full name of your organization.
    - **Organizational Unit Name**: The name of your organizational unit (e.g., IT Department).
    - **Common Name**: The fully qualified domain name (FQDN) of your OpenVAS server.
    - **Email Address**: An email address associated with the certificate.

    Example prompts:
    ```text
    Generating RSA private key, 4096 bit long modulus
    ..........................++
    ...............................................++
    e is 65537 (0x10001)
    Enter country name (2 letter code) [DE]: US
    Enter state or province name (full name) [Some-State]: California
    Enter locality name (e.g., city) []: San Francisco
    Enter organization name (e.g., company) [Internet Widgits Pty Ltd]: MyOrganization
    Enter organizational unit name (e.g., section) []: IT Department
    Enter common name (e.g., your name or your server's hostname) []: openvas.example.com
    Enter email address []: admin@example.com
    ```

3. **Completion**:
    Once you have entered all the required information, `openvas-mkcert` will generate the necessary SSL certificates and store them in the appropriate directories.

#### Files Generated

- **CA Certificate**: Typically stored in `/var/lib/openvas/CA/cacert.pem`.
- **Server Certificate**: Typically stored in `/var/lib/openvas/CA/servercert.pem`.
- **Private Key**: Typically stored in `/var/lib/openvas/private/CA/serverkey.pem`.

These certificates and keys are used by OpenVAS services to secure communications.

#### Example

Here is an example of running the `openvas-mkcert` command and generating certificates:

```bash
$ sudo openvas-mkcert
Generating RSA private key, 4096 bit long modulus
..........................++
...............................................++
e is 65537 (0x10001)
Enter country name (2 letter code) [DE]: US
Enter state or province name (full name) [Some-State]: California
Enter locality name (e.g., city) []: San Francisco
Enter organization name (e.g., company) [Internet Widgits Pty Ltd]: MyOrganization
Enter organizational unit name (e.g., section) []: IT Department
Enter common name (e.g., your name or your server's hostname) []: openvas.example.com
Enter email address []: admin@example.com
```

#### Security Considerations

1. **Secure Storage**: Ensure that the generated private keys are stored securely and only accessible by authorized users.
2. **Certificate Validity**: Regularly check the validity period of your certificates and renew them before they expire.
3. **Configuration**: Verify that OpenVAS services are properly configured to use the generated certificates.

#### Conclusion

The `openvas-mkcert` command is a crucial tool for setting up SSL certificates for OpenVAS. By following the prompts and generating the necessary certificates, you can ensure secure communications between OpenVAS components. Proper management and regular renewal of these certificates are essential for maintaining the security and integrity of your OpenVAS deployment.
