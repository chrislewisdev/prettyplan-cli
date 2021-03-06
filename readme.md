# Prettyplan (CLI Version)

Generate easily-readable HTML versions of your `terraform plan` output right from the comfort of your command line. (a web version is available at https://prettyplan.chrislewisdev.com/)

![prettyplan report](https://raw.githubusercontent.com/chrislewisdev/prettyplan-cli/master/screenshot.png)

## Terraform Version Compatibility

Prettyplan was written to work on Terraform plans from 0.11 and earlier. In 0.12, the plan output was significantly changed, addressing many of the pain points that Prettyplan addresses; for this reason, there are no current plans to update Prettyplan to work with 0.12. In my case, Prettyplan was made unnecessary by Terraform's improvements.

Contributions are still welcome if anyone would like to upgrade the code to handle plans from 0.12 onwards.

## Installation

Head on over to the [Releases page](https://github.com/chrislewisdev/prettyplan-cli/releases) and download the latest release executable for your platform. Place the executable somewhere your command line will be able to find it (i.e. your PATH), and you should be good to go!

(Note: the Linux/MacOS versions have not yet been tested. Feedback on these versions will be much appreciated!)

## Usage

In a terraform project (where you would normally run `terraform plan`), simply run:
```
prettyplan
```
Prettyplan will run `terraform plan`, capture its output and write your prettified report into a `prettyplan.html` file in the same folder.

To open the generated report as soon as it is ready, use the `-open` flag.

## Building from source

If you would like to build prettyplan locally with either `go build` or `go install`, you'll want to first generate the embedded template files using [packr](https://github.com/gobuffalo/packr/tree/master/v2):

```
go get -u github.com/gobuffalo/packr/v2/packr2
go generate
go build
```
