https://brainhive.notion.site
https://training.brainhive.nl

# Advent of code 2022
adventofcode.com/2022/day/7
https://github.com/anolivei/advent_of_code_2022
https://nickymeuleman.netlify.app/garden/aoc2022-day07

# vCard specification
https://datatracker.ietf.org/doc/html/rfc6350
https://datatracker.ietf.org/doc/html/rfc6350#section-6.2.2

# Init a go module
cd exercise02
go mod init go-training/exercise02
go mod tidy
go get -u github.com/stretchr/testify/assert


# AWS S3 ListBuckets
https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html


# MinIO
https://min.io/docs/minio/linux/reference/minio-mc-admin.html?ref=docs#installation
https://min.io/docs/minio/linux/reference/minio-mc.html#id3
mc alias set localS3 http://localhost:7000
mc alias list
mc ls localS3
mc ls --recursive --versions localS3/DOC-EXAMPLE-BUCKET

# XML to Go
https://jsonformatter.org/xml-to-go?utm_content=cmp-true


https://github.com/dmeijboom/training-connect-go


cd exercise_grpc_todo
go mod init go-training/exercise_grpc_todo

go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest



# Analyze Docker Images
- Docker dive (https://github.com/wagoodman/dive)
  A tool for exploring a docker image, layer contents, and discovering ways to shrink the size of your Docker/OCI image.

- Docker inspect




cd ./exercise_gin_gonic
docker build -t go-docker-demo .
docker run --rm -p 8080:8080 go-docker-demo

# Can't open a shell on the Multi-staged build since it is not installed as it a minimalistic image
docker exec --rm -it go-docker-demo /bin/bash


# k3d
https://github.com/k3d-io/k3d/releases
https://github.com/k3d-io/k3d
k3d cluster create k3d-training
k3d cluster ls
kubectl get nodes
kubectl get pods -A




# Messaging
- nats and jetstream

# Push image to dockerhub.com
docker tag go-docker-demo:latest miel1980/go-docker-demo:latest
docker push miel1980/go-docker-demo


cd ./k8s
kubectl apply -f deployment.yml
kubectl get deployments
kubectl get pods -l app=go-docker-demo

kubectl apply -f service.yml
kubectl get services


# Install Helm
https://helm.sh/docs/intro/install/


# Install and deploy cloudnative-pg
https://github.com/cloudnative-pg/charts
https://cloudnative-pg.io/documentation/current/quickstart/
https://github.com/cloudnative-pg/cloudnative-pg/blob/main/docs/src/quickstart.md

kubectl apply -f cluster-example.yaml
kubectl get clusters

kubectl get all -A
kubectl api-resources
kubectl get -A deployments
get pods -l cnpg.io/cluster=cluster-example


# Assignment Book review
https://tinyurl.com/book-review-api
https://github.com/dmeijboom/training-book-review/blob/main/openapi.yml
https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/dillendev/training-book-review/main/openapi.yml
https://github.com/dmeijboom/training-book-review/blob/main/books.sql
https://gorm.io/docs/


kubectl get clusters -A
kubectl.exe describe cluster cluster-example
kubectl -n default port-forward cluster-example-1 5432:5432
kubectl port-forward --address localhost cluster-example-1 5432:5432 -n default
kubectl logs cluster-example-1


# Download PostgreSQL
https://www.enterprisedb.com/downloads/postgres-postgresql-downloads
psql -h localhost -p 5432 -U postgres




kubernetes cluster cnpg psql user postgres has no password assigned







# ------------------------------------------------------
# Set password for postgres user
# ------------------------------------------------------


If you have a Kubernetes cluster running and you want to set a password for the `postgres` user in a PostgreSQL database, you can do so by following these steps:

1. **Access the PostgreSQL Pod**: You'll need to access the PostgreSQL pod where your PostgreSQL database is running. You can use `kubectl` to do this:

   ```bash
   kubectl exec -it <pod-name> -- /bin/bash
   ```

   Replace `<pod-name>` with the actual name of your PostgreSQL pod.

2. **Access PostgreSQL Shell**: Once you're inside the pod, you can access the PostgreSQL shell using the `psql` command:

   ```bash
   psql -U postgres
   ```

   This will connect you to the PostgreSQL database as the `postgres` user.

3. **Change Password**: To change the password for the `postgres` user, you can use the following SQL command:

   ```sql
   ALTER USER postgres PASSWORD 'new_password';
   ```

   Replace `'new_password'` with the actual password you want to set.

4. **Exit the PostgreSQL Shell**: You can exit the PostgreSQL shell by typing:

   ```sql
   \q
   ```

5. **Exit the Pod**: Exit the pod shell to return to your local terminal:

   ```bash
   exit
   ```

6. **Restart the Pod**: After changing the password, you may need to restart the PostgreSQL pod for the changes to take effect:

   ```bash
   kubectl delete pod <pod-name>
   ```

   Kubernetes will automatically create a new pod to replace the deleted one.




