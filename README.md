# Network Manager

# Methods

# Common issues

### Wired device is strictly unmanaged

Make sure netplan configuration exists

```
cd /etc/netplan
```

```
sudo nano 50-cloud-init.yaml
```


If not, run the following

```
sudo netplan generate && sudo netplan apply
```

Navigate to `/etc/netplan`, edit `50-cloud-init.yaml` and add NetWorkManager as the renderer

```
network:
    renderer: NetworkManager
    ethernets:
        eth0:
            dhcp4: true
    version: 2
```

# Compatability

This software can be installed on any Linux system that uses NetWorkManager.

NetworkManager is bundled with most Debian-based Linux distributions. However, it can be installed on any Linux distribution.

This software, unfortunately, is not compatible with Windows and MacOS which do not use NetworkManager.

This software can be run in a Docker container, but a mapping to the host's DBUS is required in the Dockerfile, thus the same limitation applies if the Docker image is not Linux-based.

Successfully tested on Ubuntu Desktop, Ubuntu Server, & BalenaOS.

## Author

### [Sean Velasco](https://seanvelasco.com)