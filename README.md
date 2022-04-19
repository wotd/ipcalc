# ipcalc

Simple network calculator. Provide IP address and network mask to calculate network address, min and max host.

## Usage:

#### Get help
```
./ipcalc -h
Usage of ./ipcalc:
  -a string
        IP Address (default "192.168.1.1")
  -n string
        Network Mask (default "255.255.255.0")
```

#### Default values
```
./ipcalc
Address:         192.168.1.1     -> 11000000.10101000.00000001.00000001
Netmask:         255.255.255.0   -> 11111111.11111111.11111111.00000000
-----------------------------------------------------------------------
Network:         192.168.1.0     -> 11000000.10101000.00000001.00000000
Broadcast:       192.168.1.255   -> 11000000.10101000.00000001.11111111
Host min:        192.168.1.1     -> 11000000.10101000.00000001.00000001
Host max:        192.168.1.254   -> 11000000.10101000.00000001.11111110
```

### Provided IP address
```
./ipcalc -a=172.168.1.2
Address:         172.168.1.2     -> 10101100.10101000.00000001.00000010
Netmask:         255.255.255.0   -> 11111111.11111111.11111111.00000000
-----------------------------------------------------------------------
Network:         172.168.1.0     -> 10101100.10101000.00000001.00000000
Broadcast:       172.168.1.255   -> 10101100.10101000.00000001.11111111
Host min:        172.168.1.1     -> 10101100.10101000.00000001.00000001
Host max:        172.168.1.254   -> 10101100.10101000.00000001.11111110
```
