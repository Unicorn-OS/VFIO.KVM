/*
https://wiki.xenproject.org/wiki/OVMF

enable OVMF at build time in 4.8 as it does not require non-free anymore
https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=858962

https://github.com/intel/gvt-linux/wiki/GVTd_Setup_Guide
*/

# FIX 3
sudo apt install -y nasm xz-utils python-lzma

# clone
cd xen

# Checkout
// stable-4.11 works perfect!
git checkout stable-4.11

# Dependencies
sudo apt-get install libaio1 libaio-dev

# Compile
./configure --libdir=/usr/lib --enable-ovmf --with-system-ovmf=/usr/share/ovmf/OVMF.fd

make xen -j4
make tools -j4

make install-xen
make install-tools
