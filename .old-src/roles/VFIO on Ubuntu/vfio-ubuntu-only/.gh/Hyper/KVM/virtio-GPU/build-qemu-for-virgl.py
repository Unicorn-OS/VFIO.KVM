# From:
#https://github.com/openthos/openthos/wiki/virtio-gpu
#https://www.collabora.com/news-and-blog/blog/2018/02/12/virtualizing-gpu-access/

# Not fixed yet:
#https://bugs.launchpad.net/ubuntu/+source/qemu/+bug/1540692


git clone git://git.qemu.org/qemu.git
cd qemu

# config
./configure --target-list=x86_64-softmmu --enable-gtk --with-gtkabi=3.0 --enable-kvm --enable-virglrenderer --disable-werror

# build
make -j7

# install
sudo make install
