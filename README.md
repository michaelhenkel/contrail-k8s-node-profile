#Per Node profile
1. Create namespace    
```
kubectl create namespace contrail
```
2. Create a profile config map in the namespace    
```
cat << EOF > leaf1.conf
KEY1=value1
KEY2=value2
EOF
kubectl create configmap leaf1 --from-env-file=./leaf1.conf -n contrail
```
3. label a node with the profile
```
kubectl label node kvm4-eth2.local opencontrail.org/profile=leaf1
```
4. run contrail-k8s-node-profile in the entrypoint    
```
for i in `/contrail-k8s-node-profile`
do
  export $i
done
```
