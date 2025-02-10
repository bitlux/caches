# ipieces

ipieces is a Go package used to create Geocaching puzzles such as [GCB1ZXB](https://coord.info/GCB1ZXB).

## Documentation

Available at https://pkg.go.dev/github.com/bitlux/caches/ipieces.

## Deployment instructions

I deploy on [Google Cloud Run](https://cloud.google.com/run). To do that, you must first sign
into the [Google Cloud console](https://console.cloud.google.com/),
[create a project](https://cloud.google.com/resource-manager/docs/creating-managing-projects), and have
[`gcloud`](https://cloud.google.com/sdk) installed.

To deploy, `cd` to the directory with your `main.go` file and run:
```
gcloud run deploy --source . <project> [--allow-unauthenticated]
```

## Testing

To set the IP address of the request to `foo`, use:
```
curl -H "<backdoor>: foo" localhost:8080/text
```
where `<backdoor>` is the `Puzzle.Backdoor` string you set.

## Contact / Support

I welcome issues and pull requests on GitHub and messages and email on
[geocaching.com](https://www.geocaching.com/profile/?u=bitlux).