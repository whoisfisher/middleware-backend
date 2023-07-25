INSERT INTO `rdev_user`(`created_at`, `updated_at`, `id`, `name`, `password`, `is_active`, `is_admin`, `type`, `role`)
VALUES (date_add(now(), interval 8 HOUR), date_add(now(), interval 8 HOUR), '5e81095f-3c0c-4cb2-8033-bde03d60135c',
        'admin', 'Pb4BAQEBAQHH2XmOtIOlsIViX4E8vnrkYHoWEi6lUFo=', 1, 1, 'LOCAL', 0);


INSERT INTO rdev_cluster (created_at, updated_at, id, name, api_server, version, token, type, status)
VALUES ('2023-07-12 13:45:00', '2023-07-12 13:45:00', '5e810d5f-1c0c-dcb2-8031-bde03d62135d', 'prod',
        '172.30.1.84:6443', 'v1.24.9',
        'eyJhbGciOiJSUzI1NiIsImtpZCI6IkRXYUdjWGtpMGswaTJuUHRUeWtCSWZydmNkbGEzYVctd1dOQkg4NEFNMDQifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJrdWJlcm5ldGVzLWFkbWluIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Imt1YmVybmV0ZXMtYWRtaW4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiIwNDg1ZGM1YS04OWUzLTQ2M2MtOTM5Zi0zOTNmNjdmOGJiOTEiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06a3ViZXJuZXRlcy1hZG1pbiJ9.AAk-TMVwlMUhpvwQHPVVhitfYOix-aY2FNxfj8bKmAjT2cD2DpjHurmWdeIJUmENLGcyi2_ZxKtKDiWvXgxY7X_5eFsXaEAxXaXUgzqlHWldq8j0dBx4KMBqk22055F41KTn1ICIpTzrV6ITass5a_ggObJ5h64TlmSZehHSwM6XuzcFjohBl3NXBEUXYjbqB8Z-RGklmmYq8l8L2H8568S6l7V0PJGBnHicaeabZzdxLyJeCN_4_URQh1eZsjJ1bxCKssKsJx3hMVQxvuPcRYRmg2ObKyjw9QYUY4tWCGSVBks31V9Ro3l5C8SmRGTcUYKDiC1ja8USy6GJBhRqcQ',
        'local', 'running');

INSERT INTO rdev_user_cluster (user_id, cluster_id)
VALUES ('5e81095f-3c0c-4cb2-8033-bde03d60135c', '5e810d5f-1c0c-dcb2-8031-bde03d62135d');

INSERT INTO rdev_templates (created_at, updated_at, id, name, base_template, advance_template)
VALUES ('2023-07-12 14:29:55', '2023-07-12 14:29:57', '5281095f-3c0c-4352-8033-bde03d60125c', 'postgres', 'apiVersion: "acid.zalan.do/v1"
kind: postgresql
metadata:
  name: acid-minimal-cluster
  labels:
    team: acid
spec:
  teamId: "acid"
  postgresql:
    version: "12"
  numberOfInstances: 1
  users:
    zalando:
      - superuser
      - createdb
  volume:
    size: "1Gi"
  resources:
    requests:
      cpu: 100m
      memory: 100Mi
    limits:
      cpu: 500m
      memory: 500Mi', 'apiVersion: "acid.zalan.do/v1"
kind: postgresql
metadata:
  name: acid-minimal-cluster
  labels:
    team: acid
spec:
  teamId: "acid"
  postgresql:
    version: "12"
  numberOfInstances: 3
  users:
    zalando:
      - superuser
      - createdb
  volume:
    size: "3Gi"
  resources:
    requests:
      cpu: 300m
      memory: 300Mi
    limits:
      cpu: 1500m
      memory: 1500Mi');

