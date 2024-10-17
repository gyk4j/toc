# Install configuration files

Copy/download files to MaraDNS directory:

1. [mararc.txt](https://raw.githubusercontent.com/gyk4j/toc/main/maradns/mararc.txt)
2. [db.toc.local](https://raw.githubusercontent.com/gyk4j/toc/main/maradns/db.toc.local)
3. [dwood3rc.txt](https://raw.githubusercontent.com/gyk4j/toc/main/maradns/dwood3rc.txt)

# Create secret.txt files

MaraDNS requires a `secret.txt` random seed data file. See 
[MaraDNS reference manual](https://github.com/samboy/MaraDNS/blob/master/REFERENCE.md)

We can either:

1. Run `mkSecretTxt.exe` to generate it *-OR-*
2. Create a 64-bytes (512 bits) random file

# Start MaraDNS authoritative DNS server

In MaraDNS directory:

```bat
maradns.exe -f mararc.txt
```

# Install Deadwood recursive DNS server

In MaraDNS directory (with **Administrator: Command Prompt**):

```bat
install.bat
```

# Start Deadwood recursive DNS server

```bat
net start Deadwood
```

# Stop Deadwood recursive DNS server

```bat
net stop Deadwood
```

# Uninstall Deadwood recursive DNS server

In MaraDNS directory (with **Administrator: Command Prompt**):

```bat
uninstall.bat
```

# Stop authoritative DNS server

Press `CTRL-C` to stop maradns.exe.