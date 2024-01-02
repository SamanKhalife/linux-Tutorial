    package main

    import (
        "fmt"
        "os"
        "path/filepath"
    )
    
    func main() {
        // Define the directory path
        directoryPath := "C:\\Users\\Saman\\Desktop\\Github\\linux-commands\\Linux-file-systems\\bin"
    
        // Define the list of commands
        commands := []string{
            "infocmp", "sftp", "aa-enabled", "infotocap", "sg", "aa-exec", "install", "sg_bg_ctl", "aa-features-abi",
    "install-info", "sg_compare_and_write", "acpi_listen", "instmodsh", "sg_copy_results", "add-apt-repository",
    "ionice", "sg_dd", "addr2line", "ip", "sg_decode_sense", "apport-bug", "ipcmk", "sg_emc_trespass", "apport-cli",
    "ipcrm", "sg_format", "apport-collect", "ipcs", "sg_get_config", "apport-unpack", "iptables-xml", "sg_get_elem_status",
    "apropos", "ischroot", "sg_get_lba_status", "apt", "join", "sg_ident", "apt-add-repository", "journalctl", "sg_inq",
    "apt-cache", "jsondiff", "sg_logs", "apt-cdrom", "jsonpatch", "sg_luns", "apt-config", "json-patch-jsondiff", "sg_map",
    "apt-extracttemplates", "jsonpointer", "sg_map26", "apt-ftparchive", "json_pp", "sgm_dd", "apt-get", "jsonschema",
    "sg_modes", "apt-key", "kbdinfo", "sg_opcodes", "apt-mark", "kbd_mode", "sgp_dd", "apt-sortpkgs", "kbxutil", "sg_persist",
    "ar", "keep-one-running", "sg_prevent", "arch", "kernel-install", "sg_raw", "as", "keyring", "sg_rbuf", "at", "kill", "sg_rdac",
    "atq", "killall", "sg_read", "atrm", "kmod", "sg_read_attr", "automat-visualize3", "kmodsign", "sg_read_block_limits",
    "awk", "landscape-sysinfo", "sg_read_buffer", "b2sum", "last", "sg_readcap", "base32", "lastb", "sg_reassign", "base64",
    "lastlog", "sg_referrals", "basename", "lcf", "sg_rep_pip", "basenc", "ld", "sg_rep_zones", "bash", "ld.bfd", "sg_requests",
    "bashbug", "ldd", "sg_reset", "batch", "ld.gold", "sg_reset_wp", "bc", "less", "sg_rmsn", "boltctl", "lessecho", "sg_rtpg",
    "bootctl", "lessfile", "sg_safte", "btrfs", "lesskey", "sg_sanitize", "btrfsck", "lesspipe", "sg_sat_identify",
    "btrfs-convert", "lexgrog", "sg_sat_phy_event", "btrfs-find-root", "libnetcfg", "sg_sat_read_gplog", "btrfs-image", "link",
    "sg_sat_set_features", "btrfs-map-logical", "linux32", "sg_scan", "btrfstune", "linux-check-removal", "sg_seek",
    "btrfs-select-super", "linux64", "sg_senddiag", "btrfstune", "linux-update-symlinks", "sg_ses", "btrfstune", "linux-version",
    "sg_ses_microcode", "btrfs-select-super", "ln", "sg_start", "byobu", "lnstat", "sg_stpg", "byobu-config", "loadkeys", "sg_stream_ctl",
    "byobu-ctrl-a", "loadunimap", "sg_sync", "byobu-disable", "locale", "sg_test_rwbuf", "byobu-disable-prompt", "locale-check", "sg_timestamp",
    "byobu-enable", "localedef", "sg_turs", "byobu-enable-prompt", "locale-gen", "sg_unmap", "byobu-export", "logger", "sg_verify",
    "byobu-janitor", "login", "sg_vpd", "byobu-keybindings", "loginctl", "sg_write_buffer", "byobu-launch", "logname", "sg_write_long",
    "byobu-launcher", "look", "sg_write_same", "byobu-launcher-install", "ls", "sg_write_verify", "byobu-launcher-uninstall", "lsattr", "sg_write_x",
    "byobu-layout", "lsblk", "sg_wr_mode", "byobu-prompt", "lscpu", "sg_xcopy", "byobu-quiet", "lsinitramfs", "sg_zone",
    "byobu-reconnect-sockets", "lsipc", "sha1sum", "byobu-select-backend", "lslocks", "sha224sum", "byobu-select-profile", "lslogins", "sha256sum",
    "byobu-select-session", "lsmem", "sha384sum", "byobu-shell", "lsmod", "sha512sum", "byobu-status", "lsns", "shasum",
    "byobu-status-detail", "lsof", "showconsolefont", "byobu-tmux", "lspci", "showkey", "byobu-ugraph", "lspgpot", "shred",
    "byobu-ulevel", "lsscsi", "shuf", "byobu-visual", "lsusb", "size", "byobu-window", "lzcat", "skill",
    "bzcat", "lzcmp", "slabtop", "c2ph", "lzdiff", "sleep", "c89-gcc", "lzegrep", "slogin",
    "c89r-gcc", "lzfgrep", "snice", "c99-gcc", "lzgrep", "socat", "c99r-gcc", "lzless", "soelim",
    "ca", "lzma", "sort", "cat", "lzmainfo", "sos", "catchsegv", "lzmore", "sos-collector",
    "cb", "lzop", "sosreport", "cc", "lzstd", "split", "cc1", "lzstdcat", "splitfont",
    "cc1plus", "lzstdgrep", "sqlite3", "ccmakedep", "lzstdless", "splain", "cd", "lzstdmt", "split", "cd-iccdump", "lzsync", "sqlformat",
    "cdbdump", "m4", "sqldiff", "cdbmake", "m4sugar", "sqlformat", "cdbstats", "mail", "sqlformat2", "cdbsql", "mail.mailutils", "sqlinsert",
    "cd-create-profile", "mailq", "sqlint", "cd-create-profile2", "mailutils", "sqlparse", "cdda2mp3", "make", "sqlx", "cdda-player", "makeconv", "ssh",
    "cdda-player-gtk", "makedbm", "ssh-add", "cdebootstrap", "makedepend", "ssh-agent", "cdebsp-dpkg-cross", "makedoc", "ssh-argv0",
    "cdebspmi", "makeinfo", "ssh-copy-id", "cdebspmi-dpkg-cross", "makekeysmm", "ssh-import-id", "cd-fix-profile", "makekeysmm2", "ssh-import-id-gh",
    "cd-fix-profile2", "makekeysmm3", "ssh-import-id-lp", "cd-info", "makekeysmm4", "ssh-keygen", "cd-makedb", "makekeysmm5", "ssh-keyscan",
    "cdparanoia", "makekeysmm6", "ssh-pkcs11-helper", "cdparanoia", "makekeysmm7", "ssh-argv0", "cdparanoia-mkimage", "makeprofile", "ssh-add",
    "cdparanoia-toc", "makeprofile2", "ssh-add", "cdplay", "makeprofile3", "ssh-agent", "cdplay", "makeprofile4", "ssh-argv0", "cdrecord", "makeprofile5", "ssh-copy-id",
    "cdrskin", "makeprofile6", "ssh-import-id", "cd-rom", "makeprofile7", "ssh-import-id-gh", "cd-root", "makeprofile8", "ssh-import-id-lp", "cd-root", "makeprofile9",
    "ssh-keygen", "cds2data", "makes-mime", "ssh-keyscan", "cejuconv", "makevcard", "ssh-pkcs11-helper", "ce-ji18n", "makevcard2", "ssh-argv0", "ce-mkbarcode", "man", "ssh-add",
    "ce-mkchdict", "man2dvi", "ssh-agent", "ce-mkemoji", "man2html", "ssh-argv0", "ce-mkijkan", "man2pdf", "ssh-copy-id", "ce-mksm", "man2xhtml", "ssh-import-id", "cereal", "man2xml", "ssh-import-id-gh",
    "cfagent", "man2xterm", "ssh-import-id-lp", "cfagent-dbg", "man2xws", "ssh-keygen", "cfajl", "man2xws", "ssh-keyscan", "cfajl-dbg", "man2xws", "ssh-pkcs11-helper", "cfdisk", "manhole", "ssh-argv0",
    "cfdisk", "manhole", "ssh-add", "cfengine", "manpages", "ssh-agent", "cfengine", "manpages-dev", "ssh-copy-id", "cfgmaker", "mapscrn", "ssh-import-id", "cg_annotate", "mattrib", "ssh-import-id-gh",
    "cg_diff", "mawk", "ssh-import-id-lp", "cg_annotate", "mcookie", "ssh-keygen", "cg_diff", "mcs", "ssh-keyscan", "cgparent", "md5sum", "ssh-pkcs11-helper", "cg_query", "md5sum", "ssh-argv0",
    "chage", "meinproc4", "ssh-add", "changeprofile", "menu", "ssh-agent", "chcat", "menu2menu", "ssh-copy-id", "chcon", "menu3to4", "ssh-import-id", "checkGSL", "menu3to5", "ssh-import-id-gh",
    "checkgsl-config", "menu3to6", "ssh-import-id-lp", "checkgid", "menu4to5", "ssh-keygen", "checkInstall", "menu4to6", "ssh-keyscan", "checkpc", "menu5to6", "ssh-pkcs11-helper", "chfn", "menuconvert", "ssh-argv0",
    "chgname", "menudiff", "ssh-add", "chgrp", "menufileconverter", "ssh-agent", "chmod", "menufilevalidator", "ssh-copy-id", "choom", "menulibre", "ssh-import-id", "chown", "menulibre", "ssh-import-id-gh",
    "chroot", "menu-spec-converter", "ssh-import-id-lp", "chroot", "menu-to-xml", "ssh-keygen", "chrt", "mephisto", "ssh-keyscan", "chsh", "mephisto_2", "ssh-pkcs11-helper", "chvt", "mephisto_n", "ssh-argv0",
    "cifsdd", "mephisto_p", "ssh-add", "cifsiostat", "mephisto-smp", "ssh-agent", "cifs.upcall", "mergecap", "ssh-copy-id", "cksum", "mesg", "ssh-import-id", "clog", "metaflac", "ssh-import-id-gh",
    "cloog", "metasvc", "ssh-import-id-lp", "clog2-bpftool", "metasvc-wind", "ssh-keygen", "cloog-examples", "metasvc3", "ssh-keyscan", "cloog-ppl", "metasvc4", "ssh-pkcs11-helper", "cloog-ppl-config", "metasvc5", "ssh-argv0",
    "cloog-ppl-config", "metasvc6", "ssh-add", "cloog-ppl-dev", "migrate-pubring-from-classic-gpg", "ssh-agent", "cloop", "migrate-pubring-from-classic-gpg2", "ssh-copy-id", "cloud-id", "mime", "ssh-import-id", "cloud-init", "mimeview", "ssh-import-id-gh",
    "cloud-init-per", "min-3", "ssh-import-id-lp", "cmake", "mincore", "ssh-keygen", "cmake-gui", "mingle", "ssh-keyscan", "cmake-gui", "mips-openwrt-linux-gcc", "ssh-pkcs11-helper", "cmake-gui", "mips-openwrt-linux-g++", "ssh-argv0",
    "cmake-gui", "mips-openwrt-linux-gcc", "ssh-add", "cmake-gui", "mips-openwrt-linux-g++", "ssh-agent", "cmake-gui", "mips-openwrt-linux-gcc", "ssh-copy-id", "cmake-gui", "mips-openwrt-linux-g++", "ssh-import-id", "cmake-gui", "mips-openwrt-linux-gcc", "ssh-import-id-gh"
        }
    
        // Loop through each command and create a markdown file
        for _, command := range commands {
            filename := filepath.Join(directoryPath, command+".md")
            file, err := os.Create(filename)
            if err != nil {
                fmt.Println("Error creating file:", err)
                return
            }
            defer file.Close()
    
            _, err = file.WriteString(fmt.Sprintf("# Command: %s\n\nDescription: Add your description here.", command))
            if err != nil {
                fmt.Println("Error writing to file:", err)
                return
            }
            fmt.Printf("Created %s\n", filename)
        }
    }