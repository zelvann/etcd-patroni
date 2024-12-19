*This is a submission for the Distributed System's Final Assigment*

## What I Built
Simple case which uses patroni that supported by etcd for maximum accessibility and citus for distributed database

### Architecture
<img width="2245" alt="architecture" src="https://github.com/user-attachments/assets/c40b6400-d71a-4820-b4f7-e93c437f917b" />

### How to use
```bash
git clone https://github.com/zelvann/etcd-patroni.git
cd etcd-patroni && cp .env.example .env && cp patroni.env.example patroni.env
```
Then edit your environment files, make sure that's a corrent environment. After that,
```bash
docker compose up -d
```
> TIP: if you are a nonpriveleged user, you can up docker container as daemon by using sudo