The ethtool command is used to control and monitor network interfaces. It is a powerful tool that can be used to configure a wide range of network settings.

For example, to display the current settings of the interface eth0, you would use the following command:

`ethtool -s eth0`

To display the statistics of the interface eth0, you would use the following command:

`ethtool -g eth0`

To set the interface eth0's speed and duplex mode to 1000Mbps full duplex, you would use the following command:

`ethtool -i eth0 speed 1000 duplex full`

To set the interface eth0's link speed to 100Mbps, you would use the following command:

`ethtool -l eth0 speed 100`

To display all of the available options for the interface eth0, you would use the following command:

`ethtool -a eth0`


# help 

```
Usage:
        ethtool [ FLAGS ]  DEVNAME      Display standard information about device
        ethtool [ FLAGS ] -s|--change DEVNAME   Change generic options
                [ speed %d ]
                [ lanes %d ]
                [ duplex half|full ]
                [ port tp|aui|bnc|mii|fibre|da ]
                [ mdix auto|on|off ]
                [ autoneg on|off ]
                [ advertise %x[/%x] | mode on|off ... [--] ]
                [ phyad %d ]
                [ xcvr internal|external ]
                [ wol %d[/%d] | p|u|m|b|a|g|s|f|d... ]
                [ sopass %x:%x:%x:%x:%x:%x ]
                [ msglvl %d[/%d] | type on|off ... [--] ]
                [ master-slave preferred-master|preferred-slave|forced-master|forced-slave ]
        ethtool [ FLAGS ] -a|--show-pause DEVNAME       Show pause options
        ethtool [ FLAGS ] -A|--pause DEVNAME    Set pause options
                [ autoneg on|off ]
                [ rx on|off ]
                [ tx on|off ]
        ethtool [ FLAGS ] -c|--show-coalesce DEVNAME    Show coalesce options
        ethtool [ FLAGS ] -C|--coalesce DEVNAME Set coalesce options
                [adaptive-rx on|off]
                [adaptive-tx on|off]
                [rx-usecs N]
                [rx-frames N]
                [rx-usecs-irq N]
                [rx-frames-irq N]
                [tx-usecs N]
                [tx-frames N]
                [tx-usecs-irq N]
                [tx-frames-irq N]
                [stats-block-usecs N]
                [pkt-rate-low N]
                [rx-usecs-low N]
                [rx-frames-low N]
                [tx-usecs-low N]
                [tx-frames-low N]
                [pkt-rate-high N]
                [rx-usecs-high N]
                [rx-frames-high N]
                [tx-usecs-high N]
                [tx-frames-high N]
                [sample-interval N]
                [cqe-mode-rx on|off]
                [cqe-mode-tx on|off]
        ethtool [ FLAGS ] -g|--show-ring DEVNAME        Query RX/TX ring parameters
        ethtool [ FLAGS ] -G|--set-ring DEVNAME Set RX/TX ring parameters
                [ rx N ]
                [ rx-mini N ]
                [ rx-jumbo N ]
                [ tx N ]
                [ rx-buf-len N]
             [ cqe-size N]
                [ tx-push on|off]
        ethtool [ FLAGS ] -k|--show-features|--show-offload DEVNAME     Get state of protocol offload and other features
        ethtool [ FLAGS ] -K|--features|--offload DEVNAME       Set protocol offload and other features
                FEATURE on|off ...
        ethtool [ FLAGS ] -i|--driver DEVNAME   Show driver information
        ethtool [ FLAGS ] -d|--register-dump DEVNAME    Do a register dump
                [ raw on|off ]
                [ file FILENAME ]
        ethtool [ FLAGS ] -e|--eeprom-dump DEVNAME      Do a EEPROM dump
                [ raw on|off ]
                [ offset N ]
                [ length N ]
        ethtool [ FLAGS ] -E|--change-eeprom DEVNAME    Change bytes in device EEPROM
                [ magic N ]
                [ offset N ]
                [ length N ]
                [ value N ]
        ethtool [ FLAGS ] -r|--negotiate DEVNAME        Restart N-WAY negotiation
        ethtool [ FLAGS ] -p|--identify DEVNAME Show visible port identification (e.g. blinking)
               [ TIME-IN-SECONDS ]
        ethtool [ FLAGS ] -t|--test DEVNAME     Execute adapter self test
               [ online | offline | external_lb ]
        ethtool [ FLAGS ] -S|--statistics DEVNAME       Show adapter statistics
               [ --all-groups | --groups [eth-phy] [eth-mac] [eth-ctrl] [rmon] ]
        ethtool [ FLAGS ] --phy-statistics DEVNAME      Show phy statistics
        ethtool [ FLAGS ] -n|-u|--show-nfc|--show-ntuple DEVNAME        Show Rx network flow classification options or rules
                [ rx-flow-hash tcp4|udp4|ah4|esp4|sctp4|tcp6|udp6|ah6|esp6|sctp6 [context %d] |
                  rule %d ]
        ethtool [ FLAGS ] -N|-U|--config-nfc|--config-ntuple DEVNAME    Configure Rx network flow classification options or rules
                rx-flow-hash tcp4|udp4|ah4|esp4|sctp4|tcp6|udp6|ah6|esp6|sctp6 m|v|t|s|d|f|n|r... [context %d] |
                flow-type ether|ip4|tcp4|udp4|sctp4|ah4|esp4|ip6|tcp6|udp6|ah6|esp6|sctp6
                        [ src %x:%x:%x:%x:%x:%x [m %x:%x:%x:%x:%x:%x] ]
                        [ dst %x:%x:%x:%x:%x:%x [m %x:%x:%x:%x:%x:%x] ]
                        [ proto %d [m %x] ]
                        [ src-ip IP-ADDRESS [m IP-ADDRESS] ]
                        [ dst-ip IP-ADDRESS [m IP-ADDRESS] ]
                        [ tos %d [m %x] ]
                        [ tclass %d [m %x] ]
                        [ l4proto %d [m %x] ]
                        [ src-port %d [m %x] ]
                        [ dst-port %d [m %x] ]
                        [ spi %d [m %x] ]
                        [ vlan-etype %x [m %x] ]
                        [ vlan %x [m %x] ]
                        [ user-def %x [m %x] ]
                        [ dst-mac %x:%x:%x:%x:%x:%x [m %x:%x:%x:%x:%x:%x] ]
                        [ action %d ] | [ vf %d queue %d ]
                        [ context %d ]
                        [ loc %d]] |
                delete %d
        ethtool [ FLAGS ] -T|--show-time-stamping DEVNAME       Show time stamping capabilities
        ethtool [ FLAGS ] -x|--show-rxfh-indir|--show-rxfh DEVNAME      Show Rx flow hash indirection table and/or RSS hash key
                [ context %d ]
        ethtool [ FLAGS ] -X|--set-rxfh-indir|--rxfh DEVNAME    Set Rx flow hash indirection table and/or RSS hash key
                [ context %d|new ]
                [ equal N | weight W0 W1 ... | default ]
                [ hkey %x:%x:%x:%x:%x:.... ]
                [ hfunc FUNC ]
                [ delete ]
        ethtool [ FLAGS ] -f|--flash DEVNAME    Flash firmware image from the specified file to a region on the device
               FILENAME [ REGION-NUMBER-TO-FLASH ]
        ethtool [ FLAGS ] -P|--show-permaddr DEVNAME    Show permanent hardware address
        ethtool [ FLAGS ] -w|--get-dump DEVNAME Get dump flag, data
                [ data FILENAME ]
        ethtool [ FLAGS ] -W|--set-dump DEVNAME Set dump flag of the device
                N
        ethtool [ FLAGS ] -l|--show-channels DEVNAME    Query Channels
        ethtool [ FLAGS ] -L|--set-channels DEVNAME     Set Channels
               [ rx N ]
               [ tx N ]
               [ other N ]
               [ combined N ]
        ethtool [ FLAGS ] --show-priv-flags DEVNAME     Query private flags
        ethtool [ FLAGS ] --set-priv-flags DEVNAME      Set private flags
                FLAG on|off ...
        ethtool [ FLAGS ] -m|--dump-module-eeprom|--module-info DEVNAME Query/Decode Module EEPROM information and optical diagnostics if available
                [ raw on|off ]
                [ hex on|off ]
                [ offset N ]
                [ length N ]
                [ page N ]
                [ bank N ]
                [ i2c N ]
        ethtool [ FLAGS ] --show-eee DEVNAME    Show EEE settings
        ethtool [ FLAGS ] --set-eee DEVNAME     Set EEE settings
                [ eee on|off ]
                [ advertise %x ]
                [ tx-lpi on|off ]
                [ tx-timer %d ]
        ethtool [ FLAGS ] --set-phy-tunable DEVNAME     Set PHY tunable
                [ downshift on|off [count N] ]
                [ fast-link-down on|off [msecs N] ]
                [ energy-detect-power-down on|off [msecs N] ]
        ethtool [ FLAGS ] --get-phy-tunable DEVNAME     Get PHY tunable
                [ downshift ]
                [ fast-link-down ]
                [ energy-detect-power-down ]
        ethtool [ FLAGS ] --get-tunable DEVNAME Get tunable
                [ rx-copybreak ]
                [ tx-copybreak ]
                [ tx-buf-size ]
                [ pfc-precention-tout ]
        ethtool [ FLAGS ] --set-tunable DEVNAME Set tunable
                [ rx-copybreak N]
                [ tx-copybreak N]
                [ tx-buf-size N]
                [ pfc-precention-tout N]
        ethtool [ FLAGS ] --reset DEVNAME       Reset components
                [ flags %x ]
                [ mgmt ]
                [ mgmt-shared ]
                [ irq ]
                [ irq-shared ]
                [ dma ]
                [ dma-shared ]
                [ filter ]
                [ filter-shared ]
                [ offload ]
                [ offload-shared ]
                [ mac ]
                [ mac-shared ]
                [ phy ]
                [ phy-shared ]
                [ ram ]
                [ ram-shared ]
                [ ap ]
                [ ap-shared ]
                [ dedicated ]
                [ all ]
        ethtool [ FLAGS ] --show-fec DEVNAME    Show FEC settings
        ethtool [ FLAGS ] --set-fec DEVNAME     Set FEC settings
                [ encoding auto|off|rs|baser|llrs [...]]
        ethtool [ FLAGS ] -Q|--per-queue DEVNAME        Apply per-queue command. 
The supported sub commands include --show-coalesce, --coalesce             [queue_mask %x] SUB_COMMAND
        ethtool [ FLAGS ] --cable-test DEVNAME  Perform a cable test
        ethtool [ FLAGS ] --cable-test-tdr DEVNAME      Print cable test time domain reflectrometery data
                [ first N ]
                [ last N ]
                [ step N ]
                [ pair N ]
        ethtool [ FLAGS ] --show-tunnels DEVNAME        Show NIC tunnel offload information
        ethtool [ FLAGS ] --show-module DEVNAME Show transceiver module settings
        ethtool [ FLAGS ] --set-module DEVNAME  Set transceiver module settings
                [ power-mode-policy high|auto ]
        ethtool [ FLAGS ] -h|--help             Show this help
        ethtool [ FLAGS ] --version             Show version number
        ethtool --monitor               Show kernel notifications
                ( [ --all ]
                  | -s | --change
                  | -k | --show-features | --show-offload | -K | --features | --offload
                  | --show-priv-flags | --set-priv-flags
                  | -g | --show-ring | -G | --set-ring
                  | -l | --show-channels | -L | --set-channels
                  | -c | --show-coalesce | -C | --coalesce
                  | -a | --show-pause | -A | --pause
                  | --show-eee | --set-eee
                  | --cable-test
                  | --cable-test-tdr
                  | --show-module | --set-module )
                [ DEVNAME | * ]

FLAGS:
        --debug MASK    turn on debugging messages
        --json          enable JSON output format (not supported by all commands)
        -I|--include-statistics         request device statistics related to the command (not supported by all commands)
```


## breakdown

```
-s: This option displays the current settings of the interface.
-g: This option displays the statistics of the interface.
-i: This option sets the interface's speed and duplex mode.
-l: This option sets the interface's link speed.
-a: This option displays all of the available options for the interface.
-c: This option sets the interface's collision count.
-d: This option disables the interface.
-e: This option enables the interface.
-f: This option performs a reset on the interface.
-r: This option resets the interface's statistics.
```
