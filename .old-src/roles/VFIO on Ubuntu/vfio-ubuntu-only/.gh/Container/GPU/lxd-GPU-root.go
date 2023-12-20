/* enables `sudo app` to work in GPU containers:
- Only changes is we use `raw.idmap: both 1000 0` instead

Graph: LXD.GPU.root
. guide: https://github.com/lxc/lxd/blob/master/doc/userns-idmap.md
. guide: https://blog.simos.info/how-to-easily-run-graphics-accelerated-gui-apps-in-lxd-containers-on-your-ubuntu-desktop/

Graph: LXD.set
. multiple values: https://superuser.com/questions/1174344/syntax-for-setting-lxd-container-raw-idmap

*/

func init(){
lxc profile create rootgui
cat rootguiprofile.txt | lxc profile edit rootgui

}
