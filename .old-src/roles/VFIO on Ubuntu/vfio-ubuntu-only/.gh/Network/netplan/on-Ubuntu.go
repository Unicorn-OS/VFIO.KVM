/* if fact[dist = 'ubuntu']

https://netplan.io/examples#configuring-network-bridges
*/
echo "network: {config: disabled}" | sudo tee -a /etc/cloud/cloud.cfg.d/99-disable-network-config.cfg
