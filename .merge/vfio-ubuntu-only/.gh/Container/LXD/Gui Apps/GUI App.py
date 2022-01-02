# from: https://blog.simos.info/how-to-easily-run-graphics-accelerated-gui-apps-in-lxd-containers-on-your-ubuntu-desktop/

src='~/code/.tmp'

mkdir -p $src

cd $src


# Step 2 Creating the gui LXD profile
wget -c https://blog.simos.info/wp-content/uploads/2018/06/lxdguiprofile.txt

lxc profile create gui
cat lxdguiprofile.txt | lxc profile edit gui


# Launching gui containers in LXD
launch_u1604(){
    lxc launch --profile default --profile gui ubuntu:18.04 gui1804
}

launch_u1604(){
    lxc launch --profile default --profile gui ubuntu:16.04 gui1604
}


login(){
    lxc exec gui1804 -- sudo --user ubuntu --login
}
