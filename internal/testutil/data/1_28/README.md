# Static test data generated by jjaramillo on Tue Oct 24 08:50:35 PDT 2023

### `nri-kubernetes` commit
```
0f6e852cb3c94322d5eebc5b650edc02375ee749
```

`git status --short`

```
 D ../../data/1_28/README.md
 D ../../data/1_28/controlplane/api-server/metrics
 D ../../data/1_28/controlplane/controller-manager/metrics
 D ../../data/1_28/controlplane/etcd/metrics
 D ../../data/1_28/controlplane/scheduler/metrics
 D ../../data/1_28/endpoints.yaml
 D ../../data/1_28/ksm/metrics
 D ../../data/1_28/kubelet/metrics/cadvisor
 D ../../data/1_28/kubelet/pods
 D ../../data/1_28/kubelet/stats/summary
 D ../../data/1_28/namespaces.yaml
 D ../../data/1_28/nodes.yaml
 D ../../data/1_28/pods.yaml
 D ../../data/1_28/services.yaml
?? ./
```

### `$ kubectl version`
```
Client Version: v1.28.3
Kustomize Version: v5.0.4-0.20230601165947-6ce0bf390ce3
Server Version: v1.28.0-rc.1
```

### Kubernetes nodes
```
NAME           STATUS   ROLES           AGE    VERSION        INTERNAL-IP    EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION    CONTAINER-RUNTIME
datagen-1-28   Ready    control-plane   113s   v1.28.0-rc.1   192.168.49.2   <none>        Ubuntu 22.04.2 LTS   6.4.16-linuxkit   containerd://1.6.21
```