default :
	make --no-print-directory all

all :
	make create

# --- Initialize
create :
	ansible-playbook .pool.yml -t create
	make l

# --- Remove
wipe remove:
	ansible-playbook .pool.yml -t remove


l:
	sudo virsh pool-list
	sudo virsh pool-list --all

help:
	cat makefile

