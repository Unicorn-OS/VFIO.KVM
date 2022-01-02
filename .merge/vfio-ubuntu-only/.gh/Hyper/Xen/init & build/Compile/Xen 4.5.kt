/*
Supports qemu-xen-traditional
*/

# Prereq
sudo apt-get install bcc bin86 gawk bridge-utils iproute libcurl3 libcurl4-openssl-dev bzip2 module-init-tools transfig tgif texinfo texlive-latex-base texlive-latex-recommended texlive-fonts-extra texlive-fonts-recommended pciutils-dev mercurial make gcc libc6-dev zlib1g-dev python python-dev python-twisted libncurses5-dev patch libsdl-dev libjpeg62-dev iasl libbz2-dev e2fslibs-dev git-core uuid-dev ocaml ocaml-findlib libx11-dev bison flex xz-utils libyajl-dev gettext libpixman-1-dev libaio-dev markdown pandoc

# Clone
git clone git://xenbits.xen.org/xen.git
cd xen

# Checkout
git checkout tags/RELEASE-4.5.0 -b my

# FIX
sudo apt-get install -y g++-multilib

./configure --libdir=/usr/lib

make dist

Package(){
make debball

mkdir -p ~/pkg
cp dist/*.deb ~/pkg/
}

# Install
sudo su

make install

/sbin/ldconfig


# Bootloader
update-grub

# Enable
sudo /etc/init.d/xencommons enable    
sudo /etc/init.d/xendomains enable    
sudo /etc/init.d/xen-watchdog enable    
sudo /etc/init.d/xendriverdomain enable
