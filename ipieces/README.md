# ipieces

ipieces is a Go package used to create Geocaching puzzles such as
[GCB1ZXB](https://coord.info/GCB1ZXB).

## Documentation

Available at https://pkg.go.dev/github.com/bitlux/caches/ipieces.

## Deployment instructions

I deploy on [Google Cloud Run](https://cloud.google.com/run).
https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service is a good
reference on how to us Google Cloud Run.

Create a directory for your project and put your code in a file named `main.go`. Run
`go mod init <module name>` to create a module (because your `main.go` imports packages outside the
standard library, Google Cloud Run requires it to be in its own module).

To deploy, `cd` to the directory with your `main.go` file and run:
```
gcloud run deploy --source . <service> [--allow-unauthenticated]
```
where `<service>` is any name you want.

## Testing

To set the IP address of the request to `foo`, use:
```
curl -H "<backdoor>: foo" localhost:8080/text
```
where `<backdoor>` is the `Puzzle.Backdoor` string you set.

## Contact / Support

I welcome issues and pull requests on GitHub and messages and email on
[geocaching.com](https://www.geocaching.com/profile/?u=bitlux).
