### Cluster wide example
1. Create namespace
```
kubectl create namespace contrail
```
2. Create a profile config map in the namespace
```
cat << EOF > cluster.profile
KEY1=value1
KEY2=value2
EOF
kubectl create configmap cluster.profile --from-env-file=./cluster.profile -n contrail
```
3. Label ALL nodes with the profile
```
for node in `kubectl get node -o=custom-columns=NAME:.metadata.name |grep -v NAME`
do
  kubectl label node $node opencontrail.org/profile_cluster="cluster.profile"
done
```
4. Apply OpenContrail manifest
```
kubectl apply -f https://raw.githubusercontent.com/michaelhenkel/contrail-kubeadm/master/contrail.yaml
```
### Per Node example
1. Create namespace
```
kubectl create namespace contrail
```
2. Create a profile config map in the namespace
```
cat << EOF > leaf1.profile
KEY1=value1
KEY2=value2
EOF
kubectl create configmap leaf1.profile --from-env-file=./leaf1.profile -n contrail
```
3. Label only specific nodes with the profile
```
kubectl label node kvm4-eth2.local opencontrail.org/profile_leaf1=leaf1.profile
```
4. Apply OpenContrail manifest
```
kubectl apply -f https://raw.githubusercontent.com/michaelhenkel/contrail-kubeadm/master/contrail.yaml
```
### Multiple profiles example
1. Create namespace
```
kubectl create namespace contrail
```
2. Create 1st profile config map
```
cat << EOF > cluster.profile
KEY1=value1
KEY2=value2
EOF
kubectl create configmap cluster.profile --from-env-file=./cluster.profile -n contrail
```
3. Create 2nd profile config map
```
cat << EOF > leaf1.profile
KEY3=value3
KEY4=value4
EOF
kubectl label node kvm4-eth2.local opencontrail.org/profile_leaf1=leaf1.profile
```
4. Label node with both profiles
```
kubectl label node kvm4-eth2.local opencontrail.org/profile_cluster="cluster.profile"
kubectl label node kvm4-eth2.local opencontrail.org/profile_leaf1="leaf1.profile"
```
5. Apply OpenContrail manifest
```
kubectl apply -f https://raw.githubusercontent.com/michaelhenkel/contrail-kubeadm/master/contrail.yaml
```
