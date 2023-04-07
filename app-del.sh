kubectl delete -f app/ipvlan-pods-edge1.yaml   --kubeconfig='/root/.kube/config-edge1'
sleep 5
kubectl delete -f app/ipvlan-cni-edge1.yaml   --kubeconfig='/root/.kube/config-edge1'
