#kubectl apply -f kind/multus-daemonset-thick.yml   --kubeconfig='/root/.kube/config-edge1'
kubectl apply -f app/ipvlan-cni-edge1.yaml   --kubeconfig='/root/.kube/config-edge1'
kubectl apply -f app/ipvlan-pods-edge1.yaml   --kubeconfig='/root/.kube/config-edge1'
