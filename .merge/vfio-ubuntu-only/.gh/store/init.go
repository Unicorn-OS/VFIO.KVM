/*
~/.uni/disk is a high-performance FS for Running instances!
*/

func init(){
  mkdir -p ~/.uni/disk
  mount --bind $FS ~/.uni/disk
}
