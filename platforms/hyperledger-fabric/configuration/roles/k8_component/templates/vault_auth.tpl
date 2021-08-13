apiVersion: v1
kind: ServiceAccount
metadata:
  name: vault-auth
  namespace: {{ component_name }}
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: pod-controller
  namespace: {{ component_name }}
rules:
  - apiGroups:
      - ""
    resources:
      - pods
    verbs:
      - get
      - list
      - watch
      - create
      - delete
  - apiGroups:
      - ""
    resources:
      - events
      - pods/log
      - pods/status
    verbs:
      - get
      - list
      - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: peer-controls-pods
  namespace: {{ component_name }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: pod-controller
subjects:
  - kind: ServiceAccount
    name: vault-auth