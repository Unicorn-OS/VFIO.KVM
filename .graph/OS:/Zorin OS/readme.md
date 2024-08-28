sch: https://www.google.com/search?q=zorin+os+kvm+gpu+passthrough+blacklist+amdgpu, https://www.google.com/search?q=zorin+os+vfio+kvm , https://www.google.com/search?q=zorin+os+blacklist+pci+vfio+kvm

# Discuss:
- https://forum.zorin.com/t/zorin-os-17-doesnt-use-gpu-pass-for-kvm/32050

# Solution:
test: Works!
note: This solved blacklist amdgpu issue.
https://forum.proxmox.com/threads/fix-pci-passthrough-documentation.122521/


src:
```
echo "blacklist amdgpu" >> /etc/modprobe.d/blacklist.conf
echo "blacklist radeon" >> /etc/modprobe.d/blacklist.conf
echo "blacklist nouveau" >> /etc/modprobe.d/blacklist.conf
echo "blacklist nvidia" >> /etc/modprobe.d/blacklist.con
```


## also:
### 1:
>i find solution, and its working now.
>kernel 6.2 vfio gpu pass 14
- https://forum.zorin.com/t/zorin-os-17-doesnt-use-gpu-pass-for-kvm/32050/12

relation: https://forum.level1techs.com/t/linux-kernel-6-seems-to-be-incompatible-with-the-vfio-pci-module-needed-for-pci-passthrough/190039/37

### 2:
>Also for:
>
>    `lsmod | grep kvm vfio vfio_pci`
>
>in order to see if the proper modules are loaded.
>
>    `ls /etc/modprobe.d/`
>
>Do you have a vfio configuration file?
>It should contain your options vfio-pci ids=...
