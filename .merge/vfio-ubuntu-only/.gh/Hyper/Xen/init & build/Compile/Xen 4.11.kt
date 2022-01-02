/*
Supports qemu-xen-traditional
*/

# FIX ./configure errors
sudo apt-get install -y g++-multilib liblzma-dev
sudo apt-get install -y libnl-3-dev libsystemd-dev

# Prereq
sudo apt-get install bcc bin86 gawk bridge-utils iproute libcurl3 libcurl4-openssl-dev bzip2 module-init-tools transfig tgif texinfo texlive-latex-base texlive-latex-recommended texlive-fonts-extra texlive-fonts-recommended pciutils-dev mercurial make gcc libc6-dev zlib1g-dev python python-dev python-twisted libncurses5-dev patch libsdl-dev libjpeg62-dev iasl libbz2-dev e2fslibs-dev git-core uuid-dev ocaml ocaml-findlib libx11-dev bison flex xz-utils libyajl-dev gettext libpixman-1-dev libaio-dev markdown pandoc

# Clone
git clone git://xenbits.xen.org/xen.git
cd xen

# Checkout
git checkout tags/RELEASE-4.10.0 -b my

# Config
./configure --libdir=/usr/lib --enable-systemd

make dist -j16

Package(){
make debball -j16

sudo apt install -y gdebi
cd dist/
sudo gdebi xen*.deb
}

# Install
sudo su

make install -j16

/sbin/ldconfig


# Bootloader
update-grub

# Enable
systemd.kt
