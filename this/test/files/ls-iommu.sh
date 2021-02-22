#!/bin/bash
for d in $(find /sys/kernel/iommu_groups/ -type l | sort -n -k5 -t/); do
    n=${d#*/iommu_groups/*}; n=${n%%/*}
    printf 'IOMMU Group %s ' "$n"
    lspci -nns "${d##*/}"
done;


#prime: https://gist.github.com/flungo/428c374c040de1d0a30fd4a593d39040
