# Terraform Watcher

Watcher for terraform.


## Installation

`go get github.com/darkowlzz/terraform-watcher`


## Usage

`terraform-watcher`

Learn more with `terraform-watcher -h`


## Development Setup

1. Clone the repo & `cd` into it.
2. Ensure that [`govendor`](https://github.com/kardianos/govendor) is installed
and install all the vendor dependencies with `govendor sync`.
3. Install with `go install github.com/darkowlzz/terraform-watcher`.

### Dependencies

Add new dependencies with `govendor fetch <packagename>`. This would install
the dependencies under `vendor/` and add them to `vendor/vendor.json`, which
should be checked-in.
