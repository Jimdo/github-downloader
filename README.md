# github-downloader

Download releases from private repositories with ease!

## Usage

```bash
$ docker run -it --rm \
    -v <out_dir>:/opt \
    -e GITHUB_TOKEN=<github_token> \
    quay.io/jimdo/github-downloader \
    <repo> <version> <file>
```

This command downloads 
* the file `<file>` 
* from the repository `<repo>` 
* released with the tag `<version>` 
* and stores it in `<out_dir>`.
* For authentication, it uses the `<github_token>` that needs `repo` scope. 