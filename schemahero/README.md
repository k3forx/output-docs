# SchemaHero

## Install SchemaHero

### Installing the `kubectl` plugin

```bash
kubectl krew install schemahero

kubectl schemahero version
```

### Installing the in-cluster components

```bash
kubectl schemahero install

kubectl get pods -n schemahero-system
NAME           READY   STATUS    RESTARTS   AGE
schemahero-0   1/1     Running   2          12s
```

## Connect a database

### Deploy PostgreSQL

Create a namespace `schemahero-turorial` and PostgreSQL pod in the namespace.

```bash
❯ kubectl create ns schemahero-tutorial
namespace/schemahero-tutorial created

❯ kubectl apply -f demo/postgresql/postgres-11.8.0.yaml
secret/postgresql created
service/postgresql-headless created
service/postgresql created
statefulset.apps/postgresql created

❯ kubectl apply -f demo/postgresql/postgres-11.8.0.yaml -n schemahero-tutorial
secret/postgresql created
service/postgresql-headless created
service/postgresql created
statefulset.apps/postgresql created

❯ kubectl get all -n schemahero-tutorial
NAME               READY   STATUS    RESTARTS   AGE
pod/postgresql-0   1/1     Running   0          12s

NAME                          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/postgresql            ClusterIP   10.106.233.167   <none>        5432/TCP   15s
service/postgresql-headless   ClusterIP   None             <none>        5432/TCP   15s

NAME                          READY   AGE
statefulset.apps/postgresql   1/1     15s
```

### Connect to PostgresSQL

```bash
❯ brew install beekeeper-studio
Updating Homebrew...
==> Auto-updated Homebrew!
Updated 2 taps (homebrew/cask-versions and homebrew/core).
==> Updated Formulae
Updated 5 formulae.
==> Updated Casks
Updated 1 cask.

==> Downloading https://github.com/beekeeper-studio/beekeeper-studio/releases/download/v1.10.2/Beekeeper-Studio-1.10.2
==> Downloading from https://github-releases.githubusercontent.com/198484780/ff682c00-8b25-11eb-9eeb-bb8fbd65a240?X-Am
######################################################################## 100.0%
==> Installing Cask beekeeper-studio
==> Moving App 'Beekeeper Studio.app' to '/Applications/Beekeeper Studio.app'
🍺  beekeeper-studio was successfully installed!

❯ kubectl port-forward -n schemahero-tutorial svc/postgresql 5432:5432
```

With beekeeper, you can check the database `airlinedb`.

![image](https://user-images.githubusercontent.com/45956169/117317659-ab8e8800-aec4-11eb-9bbf-a6d48fc83e71.png)

Click "Connect"

![image](https://user-images.githubusercontent.com/45956169/117317910-ea244280-aec4-11eb-8eb1-0bcf7413c681.png)

### Create SchemaHero Database object

```bash
❯ kubectl apply -f demo/schema/airlinedb.yaml -n schemahero-tutorial
database.databases.schemahero.io/airlinedb created

❯ kubectl get databases -n schemahero-tutorial
NAME        AGE
airlinedb   35s
```

## Create A New Table

### Airports Table

```bash
❯ kubectl apply -f demo/schema/airport-table.yaml
table.schemas.schemahero.io/airport created
```

### Validating the migration

```bash
❯ kubectl schemahero get migrations -n schemahero-tutorial
ID       DATABASE   TABLE    PLANNED  EXECUTED  APPROVED  REJECTED
9e59b6b  airlinedb  airport  45s


```

#### View the migration

```bash
❯ kubectl schemahero describe migration 9e59b6b -n schemahero-tutorial

Migration Name: 9e59b6b

Generated DDL Statement (generated at 2021-05-06T23:53:15+09:00):
  create table "airport" ("code" character (4), "name" character varying (255), primary key ("code"))

To apply this migration:
  kubectl schemahero -n schemahero-tutorial approve migration 9e59b6b

To recalculate this migration against the current schema:
  kubectl schemahero -n schemahero-tutorial recalculate migration 9e59b6b

To deny and cancel this migration:
  kubectl schemahero -n schemahero-tutorial reject migration 9e59b6b
```

### Applying the migration

```bash
❯ kubectl schemahero -n schemahero-tutorial approve migration 9e59b6b
Migration 9e59b6b approved

❯ kubectl schemahero get migrations -n schemahero-tutorial
ID       DATABASE   TABLE    PLANNED  EXECUTED  APPROVED  REJECTED
9e59b6b  airlinedb  airport  4m41s    19s       19s
```

### Verifying the migration

```bash
❯ kubectl exec -it -n schemahero-tutorial postgresql-0 -- bash
I have no name!@postgresql-0:/$ psql -U airlinedb-user  -d airlinedb
Password for user airlinedb-user:
psql (11.8)
Type "help" for help.

airlinedb=> \l
                                      List of databases
   Name    |  Owner   | Encoding |   Collate   |    Ctype    |       Access privileges
-----------+----------+----------+-------------+-------------+-------------------------------
 airlinedb | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 | =Tc/postgres                 +
           |          |          |             |             | postgres=CTc/postgres        +
           |          |          |             |             | "airlinedb-user"=CTc/postgres
 postgres  | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 |
 template0 | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 | =c/postgres                  +
           |          |          |             |             | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 | =c/postgres                  +
           |          |          |             |             | postgres=CTc/postgres
(4 rows)

airlinedb=> \dt
             List of relations
 Schema |  Name   | Type  |     Owner
--------+---------+-------+----------------
 public | airport | table | airlinedb-user
(1 row)

airlinedb=> \d airport
                      Table "public.airport"
 Column |          Type          | Collation | Nullable | Default
--------+------------------------+-----------+----------+---------
 code   | character(4)           |           | not null |
 name   | character varying(255) |           |          |
Indexes:
    "airport_pkey" PRIMARY KEY, btree (code)

airlinedb=> \q
I have no name!@postgresql-0:/$ exit
exit

```

You can check the result of the migration on UI.

![image](https://user-images.githubusercontent.com/45956169/117320248-00cb9900-aec7-11eb-90e2-3de6ac8b28d3.png)
