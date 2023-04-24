# Running the server locally

Without Docker:

```console
go mod download
make run
```

With Docker:

```console
make docker-run
```

# Requesting a review

```console
curl -XPOST localhost:9090 -d '{"repositoryOwner": "callrua", "repositoryName": "aye-robot", "pullRequestNumber": 1}'
`
