# aye-robot

Peer reviews from ChatGPT :forever-alone:

Leveraging ChatGPT to post peer reviews on GitHub pull requests.

# Requirements

$GH_TOKEN environment variable.
> A GitHub API key, with permission to access and post comments to Pull Requests.

$AI_TOKEN environment variable.
> An API key for ChatGPT, to ask it to review Pull Requests.

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
```
