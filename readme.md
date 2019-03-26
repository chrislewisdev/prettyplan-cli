# Prettyplan (CLI Version)

Generate easily-readable HTML versions of your `terraform plan` output right from the comfort of your command line.

## Usage

Simply run `prettyplan` instead of `terraform plan` inside a Terraform project. Prettyplan will capture the output and format it into `prettyplan.html` in the same folder.

## Building from source

If you would like to build prettyplan locally with either `go build` or `go install`, you'll want to first generate the embedded template files using [packr](https://github.com/gobuffalo/packr/tree/master/v2):

```
go get -u github.com/gobuffalo/packr/v2/packr2
go generate
```