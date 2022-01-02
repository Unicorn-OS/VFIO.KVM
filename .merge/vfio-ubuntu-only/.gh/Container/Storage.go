/*
Graph: LXD.Storage
docs . https://lxd.readthedocs.io/en/latest/storage/
guide . https://blog.ubuntu.com/2017/07/12/storage-management-in-lxd-2-15
examples . https://lxd.readthedocs.io/en/latest/storage/#notes-and-examples
*/

func create(){
  /*
  Create a btrfs subvolume named "pool1" on the btrfs filesystem /some/path and use as pool.
  `lxc storage create pool1 btrfs source=/some/path`
  */
}

func works(){
  # Create a separate partiton
  mkfs.btrfs -L container
  
  # Create a pool on that partition
  lxc storage create pool btrfs source=/dev/sdX1
}


func FIX(){
  /* This one is Broken . `sudo` won't work in containers:
  Create a btrfs subvolume named "pool1" on the btrfs filesystem /some/path and use as pool.
  `lxc storage create pool1 btrfs source=/some/path`
  
  But this one Works:
  Create a new pool called "pool1" on /dev/sdX.
  `lxc storage create pool1 btrfs source=/dev/sdX`
}
