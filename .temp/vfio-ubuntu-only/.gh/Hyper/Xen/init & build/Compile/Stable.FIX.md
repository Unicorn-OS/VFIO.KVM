# FIX
## FIX 1 pkgs
https://libvmi.wordpress.com/2015/01/23/libvmi-xen-setup/

## FIX 2
xl: error while loading shared libraries: libxlutil.so.4.12
https://stackoverflow.com/questions/16914243/unable-to-find-libxlutil-so-1-0-on-linux

" had a similar problem when installing xen 4.3 on Fedora 18. When configuring I figured out that: ./configure --libdir=/usr/lib did the trick for me."

## FIX 3 build Xen OVMF, but it failed as well
http://xen.1045712.n5.nabble.com/Xen-4-4-1-OVMF-on-Debian-compile-failed-td5729019.html

'install:
nasm
xz-utils
python-lzma
'


# Future
https://wiki.xen.org/wiki/Xen_VGA_Passthrough

'qemu-xen-traditional'
https://wiki.xenproject.org/wiki/QEMU_Upstream

Search for it on:
http://manpages.ubuntu.com/manpages/trusty/man5/xl.cfg.5.html
