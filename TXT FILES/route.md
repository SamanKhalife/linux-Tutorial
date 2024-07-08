# route


The `route` command in Linux is used to view and manipulate the IP routing table. This command displays the current routing table and allows you to add or remove routes. However, like `ifconfig`, the `route` command is considered deprecated in favor of the `ip` command from the `iproute2` package.


1. **View the Routing Table**:
   ```sh
   route
   ```
   or
   ```sh
   route -n
   ```
   The `-n` option prevents the command from attempting to resolve IP addresses to hostnames, making the output quicker to generate and easier to read.

2. **Add a Default Gateway**:
   ```sh
   route add default gw 192.168.1.1
   ```
   This command sets the default gateway to `192.168.1.1`.

3. **Add a Route to a Specific Network**:
   ```sh
   route add -net 192.168.2.0 netmask 255.255.255.0 gw 192.168.1.1
   ```
   This command adds a route to the network `192.168.2.0` with a netmask of `255.255.255.0`, routing through the gateway `192.168.1.1`.

4. **Add a Host Route**:
   ```sh
   route add -host 192.168.1.2 gw 192.168.1.1
   ```
   This command adds a route for a specific host `192.168.1.2` through the gateway `192.168.1.1`.

5. **Delete a Route**:
   ```sh
   route del -net 192.168.2.0 netmask 255.255.255.0
   ```
   This command removes the route to the network `192.168.2.0`.

### Example of `route` Command Output

When you run `route -n`, you might see output similar to this:

```sh
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.1.1     0.0.0.0         UG    100    0        0 eth0
192.168.1.0     0.0.0.0         255.255.255.0   U     100    0        0 eth0
```

### Using `ip` Command for Routing

The `ip` command provides more functionality and is the modern replacement for `route`. Here are equivalent commands using `ip`:

1. **View the Routing Table**:
   ```sh
   ip route show
   ```

2. **Add a Default Gateway**:
   ```sh
   ip route add default via 192.168.1.1
   ```

3. **Add a Route to a Specific Network**:
   ```sh
   ip route add 192.168.2.0/24 via 192.168.1.1
   ```

4. **Add a Host Route**:
   ```sh
   ip route add 192.168.1.2 via 192.168.1.1
   ```

5. **Delete a Route**:
   ```sh
   ip route del 192.168.2.0/24
   ```

### Conclusion

While the `route` command is still available and useful for simple routing table manipulations, it is recommended to use the `ip` command for more advanced and modern network configuration tasks. The `ip` command offers a unified and consistent interface for managing both IP addresses and routes, making it a more powerful and flexible tool for network administrators.



# help 

```
Usage: route [-nNvee] [-FC] [<AF>]           List kernel routing tables
       route [-v] [-FC] {add|del|flush} ...  Modify routing table for AF.

       route {-h|--help} [<AF>]              Detailed usage syntax for specified AF.
       route {-V|--version}                  Display version/author and exit.

        -v, --verbose            be verbose
        -n, --numeric            don't resolve names
        -e, --extend             display other/more information
        -F, --fib                display Forwarding Information Base (default)
        -C, --cache              display routing cache instead of FIB

```



## breakdown

```

```
