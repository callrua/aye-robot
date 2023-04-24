# aye-robot

Peer reviews from ChatGPT :forever-alone:

Leveraging ChatGPT to post peer reviews on GitHub pull requests.

# Requirements

$GH_TOKEN environment variable.
> A GitHub API key, with permission to access and post comments to Pull Requests.

$AI_TOKEN environment variable.
> An API key for ChatGPT, to ask it to review Pull Requests.

# Running the CLI

Without Docker:

```console
go mod download
make run -o $(repositoryOwner) -n $(repositoryName) -pr $(pullRequestNumber)
```
