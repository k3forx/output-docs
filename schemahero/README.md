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

## Modify Table

In this step, we are going to deploy a table and then make a change to it.

### Create the schedule table

```bash
❯ cat demo/schema/schedule.yaml
apiVersion: schemas.schemahero.io/v1alpha4
kind: Table
metadata:
  name: schedule
  namespace: schemahero-tutorial
spec:
  database: airlinedb
  name: schedule
  schema:
    postgres:
      primaryKey: [flight_num]
      columns:
        - name: flight_num
          type: int
        - name: origin
          type: char(4)
          constraints:
            notNull: true
        - name: destination
          type: char(4)
          constraints:
            notNull: true
        - name: departure_time
          type: time
          constraints:
            notNull: true
        - name: arrival_time
          type: time
          constraints:
            notNull: true

❯ kubectl apply -f demo/schema/schedule.yaml -n schemahero-tutorial
table.schemas.schemahero.io/schedule created

❯ kubectl schemahero get migrations -n schemahero-tutorial
ID       DATABASE   TABLE     PLANNED  EXECUTED  APPROVED  REJECTED
9e59b6b  airlinedb  airport   13h      13h       13h
a9626a8  airlinedb  schedule  16s

❯ kubectl schemahero -n schemahero-tutorial approve migration a9626a8
Migration a9626a8 approved

❯ kubectl exec -it -n schemahero-tutorial postgresql-0 -- bash
I have no name!@postgresql-0:/$ psql -U airlinedb-user  -d airlinedb
Password for user airlinedb-user:
psql (11.8)
Type "help" for help.

airlinedb=> \dt
             List of relations
 Schema |   Name   | Type  |     Owner
--------+----------+-------+----------------
 public | airport  | table | airlinedb-user
 public | schedule | table | airlinedb-user
(2 rows)

airlinedb=> exit
I have no name!@postgresql-0:/$ exit
exit
```

### Change columns

Let's make a few changes to this table schema now: - Make the `departure_time` and `arrival_time` columns nullable - Add a new column named `duration`

```bash
❯ git diff demo/schema/schedule.yaml
diff --git a/schemahero/demo/schema/schedule.yaml b/schemahero/demo/schema/schedule.yaml
index 2541e70..88d3e48 100644
--- a/schemahero/demo/schema/schedule.yaml
+++ b/schemahero/demo/schema/schedule.yaml
@@ -22,9 +22,7 @@ spec:
             notNull: true
         - name: departure_time
           type: time
-          constraints:
-            notNull: true
         - name: arrival_time
           type: time
-          constraints:
-            notNull: true
+        - name: duration
+          type: int

❯ kubectl apply -f demo/schema/schedule.yaml
table.schemas.schemahero.io/schedule configured

❯ kubectl schemahero get migrations -n schemahero-tutorial
ID       DATABASE   TABLE     PLANNED  EXECUTED  APPROVED  REJECTED
9e59b6b  airlinedb  airport   16h      16h       16h
a9626a8  airlinedb  schedule  3h       3h        3h
fa32022  airlinedb  schedule  14s
```

### View the migration

```bash
❯ kubectl schemahero -n schemahero-tutorial describe migration fa32022

Migration Name: fa32022

Generated DDL Statement (generated at 2021-05-07T16:47:48+09:00):
  alter table "schedule" alter column "departure_time" type time, alter column "departure_time" drop not null;
alter table "schedule" alter column "arrival_time" type time, alter column "arrival_time" drop not null;
alter table "schedule" add column "duration" integer

To apply this migration:
  kubectl schemahero -n schemahero-tutorial approve migration fa32022

To recalculate this migration against the current schema:
  kubectl schemahero -n schemahero-tutorial recalculate migration fa32022

To deny and cancel this migration:
  kubectl schemahero -n schemahero-tutorial reject migration fa32022
```

### Approve the migration

```bash
❯ kubectl schemahero -n schemahero-tutorial approve migration fa32022
Migration fa32022 approved

❯ kubectl schemahero get migrations -n schemahero-tutorial
ID DATABASE TABLE PLANNED EXECUTED APPROVED REJECTED
9e59b6b airlinedb airport 16h 16h 16h
a9626a8 airlinedb schedule 3h 3h 3h
fa32022 airlinedb schedule 1m44s 12s 12s
```

### Verify in the database management utility

```bash
❯ kubectl exec -it -n schemahero-tutorial postgresql-0 -- bash
I have no name!@postgresql-0:/$ psql -U airlinedb-user  -d airlinedb
Password for user airlinedb-user:
psql (11.8)
Type "help" for help.

airlinedb=> \dt
             List of relations
 Schema |   Name   | Type  |     Owner
--------+----------+-------+----------------
 public | airport  | table | airlinedb-user
 public | schedule | table | airlinedb-user
(2 rows)

airlinedb=> \d schedule
                         Table "public.schedule"
     Column     |          Type          | Collation | Nullable | Default
----------------+------------------------+-----------+----------+---------
 flight_num     | integer                |           | not null |
 origin         | character(4)           |           | not null |
 destination    | character(4)           |           | not null |
 departure_time | time without time zone |           |          |
 arrival_time   | time without time zone |           |          |
 duration       | integer                |           |          |
Indexes:
    "schedule_pkey" PRIMARY KEY, btree (flight_num)

airlinedb=> exit
I have no name!@postgresql-0:/$ exit
exit
```

### Adding a foreign key

```bash
❯ git diff demo/schema/schedule.yaml
diff --git a/schemahero/demo/schema/schedule.yaml b/schemahero/demo/schema/schedule.yaml
index 2541e70..e36fbe3 100644
--- a/schemahero/demo/schema/schedule.yaml
+++ b/schemahero/demo/schema/schedule.yaml
@@ -9,6 +9,19 @@ spec:
   schema:
     postgres:
       primaryKey: [flight_num]
+      foreignKeys:
+        - columns:
+          - origin
+          references:
+            table: airport
+            columns:
+              - code
+        - columns:
+          - destination
+          references:
+            table: airport
+            columns:
+              - code
       columns:
         - name: flight_num
           type: int
@@ -22,9 +35,7 @@ spec:
             notNull: true
         - name: departure_time
           type: time
-          constraints:
-            notNull: true
         - name: arrival_time
           type: time
-          constraints:
-            notNull: true
+        - name: duration
+          type: int

❯ kubectl apply -f demo/schema/schedule.yaml -n schemahero-tutorial
table.schemas.schemahero.io/schedule configured

❯ kubectl schemahero get migrations -n schemahero-tutorial
ID       DATABASE   TABLE     PLANNED  EXECUTED  APPROVED  REJECTED
9e59b6b  airlinedb  airport   18h      18h       18h
a9626a8  airlinedb  schedule  5h       5h        5h
b12d3fd  airlinedb  schedule  22s
fa32022  airlinedb  schedule  1h       1h        1h

❯ kubectl schemahero -n schemahero-tutorial describe migration b12d3fd

Migration Name: b12d3fd

Generated DDL Statement (generated at 2021-05-07T18:14:20+09:00):
  alter table schedule add constraint schedule_origin_fkey foreign key (origin) references airport (code);
alter table schedule add constraint schedule_destination_fkey foreign key (destination) references airport (code)

To apply this migration:
  kubectl schemahero -n schemahero-tutorial approve migration b12d3fd

To recalculate this migration against the current schema:
  kubectl schemahero -n schemahero-tutorial recalculate migration b12d3fd

To deny and cancel this migration:
  kubectl schemahero -n schemahero-tutorial reject migration b12d3fd
```

### Approve and verify

```bash
❯ kubectl schemahero -n schemahero-tutorial approve migration b12d3fd
Migration b12d3fd approved

❯ kubectl exec -it -n schemahero-tutorial postgresql-0 -- bash
I have no name!@postgresql-0:/$ psql -U airlinedb-user  -d airlinedb
Password for user airlinedb-user:
psql (11.8)
Type "help" for help.

airlinedb=> \dt
             List of relations
 Schema |   Name   | Type  |     Owner
--------+----------+-------+----------------
 public | airport  | table | airlinedb-user
 public | schedule | table | airlinedb-user
(2 rows)

airlinedb=> \d schedule
                         Table "public.schedule"
     Column     |          Type          | Collation | Nullable | Default
----------------+------------------------+-----------+----------+---------
 flight_num     | integer                |           | not null |
 origin         | character(4)           |           | not null |
 destination    | character(4)           |           | not null |
 departure_time | time without time zone |           |          |
 arrival_time   | time without time zone |           |          |
 duration       | integer                |           |          |
Indexes:
    "schedule_pkey" PRIMARY KEY, btree (flight_num)
Foreign-key constraints:
    "schedule_destination_fkey" FOREIGN KEY (destination) REFERENCES airport(code)
    "schedule_origin_fkey" FOREIGN KEY (origin) REFERENCES airport(code)

airlinedb=> \d airport
                      Table "public.airport"
 Column |          Type          | Collation | Nullable | Default
--------+------------------------+-----------+----------+---------
 code   | character(4)           |           | not null |
 name   | character varying(255) |           |          |
Indexes:
    "airport_pkey" PRIMARY KEY, btree (code)
Referenced by:
    TABLE "schedule" CONSTRAINT "schedule_destination_fkey" FOREIGN KEY (destination) REFERENCES airport(code)
    TABLE "schedule" CONSTRAINT "schedule_origin_fkey" FOREIGN KEY (origin) REFERENCES airport(code)

airlinedb=> exit
I have no name!@postgresql-0:/$ exit
exit
```
