# Example GO & Kubernetes REST API

A basic Go REST API application which run on kubernetes. The application will accept full email messages and return a json response having extracted a number of properties.

## Build & Deploy

```
# build the docker image
docker build -t magnuspwhite/validity:1.0.0 .
# push the image to docker hub
docker push magnuspwhite/validity:1.0.0

# in this example we're using mini kube
# start minikube
minikube start
# deploy the kubenetes config
kubectl apply -f k8-deployment.yml
# get the service endpoint URL
minikube service validity-service --url
```

You can now make requests to the service URL (eg. `http://172.17.0.2:32606`) and request the endpoints below.

## Debug locally

```
# build the docker image
docker build -t validity .
# run the docker image forwarding the ports
docker run -p 8080:8080 validity
```

Connect to the docker image on `localhost:8080`

## Endpoints

### /email

| Property      | Detail
| --------------| --------------
| Description   | Parses an email and returns sender metadata
| URL           | `/email`
| Method        | POST
| Parameters    | `text/plain` body containing an email. See `samples` directory for examples
| Returns       | json object as below

Example response

```
{
    "To": "1000mercis@cp.assurance.returnpath.net",
    "From": "\"Darty\" <infos@contact-darty.com>",
    "Date": "01 Apr 2011 16:17:41 +0200",
    "Subject": "Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
    "Message-ID": "<20110401161739.E3786358A9D7B977@contact-darty.com>"
}
```

### /healthcheck

| Property      | Detail
| --------------| --------------
| Description   | Checks health of the system
| URL           | `/healthcheck`
| Method        | GET
| Parameters    | None
| Returns       | 200 HTTP response code if healthy

