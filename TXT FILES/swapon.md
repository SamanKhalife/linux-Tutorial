



# help 

```
Usage:
 swapon [options] [<spec>]

Enable devices and files for paging and swapping.

Options:
 -a, --all                enable all swaps from /etc/fstab
 -d, --discard[=<policy>] enable swap discards, if supported by device
 -e, --ifexists           silently skip devices that do not exist
 -f, --fixpgsz            reinitialize the swap space if necessary
 -o, --options <list>     comma-separated list of swap options
 -p, --priority <prio>    specify the priority of the swap device
 -s, --summary            display summary about used swap devices (DEPRECATED)
     --show[=<columns>]   display summary in definable table
     --noheadings         don't print table heading (with --show)
     --raw                use the raw output format (with --show)
     --bytes              display swap size in bytes in --show output
 -v, --verbose            verbose mode

 -h, --help               display this help
 -V, --version            display version

The <spec> parameter:
 -L <label>             synonym for LABEL=<label>
 -U <uuid>              synonym for UUID=<uuid>
 LABEL=<label>          specifies device by swap area label
 UUID=<uuid>            specifies device by swap area UUID
 PARTLABEL=<label>      specifies device by partition label
 PARTUUID=<uuid>        specifies device by partition UUID
 <device>               name of device to be used
 <file>                 name of file to be used

Available discard policy types (for --discard):
 once    : only single-time area discards are issued
 pages   : freed pages are discarded before they are reused
If no policy is selected, both discard types are enabled (default).

Available output columns:
 NAME   device file or partition path
 TYPE   type of the device
 SIZE   size of the swap area
 USED   bytes in use
 PRIO   swap priority
 UUID   swap uuid
 LABEL  swap label

```



## breakdown

```

```