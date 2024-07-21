<h1 align="center">
HTML.tf
</h1>

<p align="center">
<b>
<i>
HTML × Terraform
</i>
</b>
</p>

<p align="center">
The Next Generation HTML Builder.
</p>

<p align="center">
<a href="https://github.com/koki-develop/terraform-provider-html/releases/latest"><img src="https://img.shields.io/github/v/release/koki-develop/terraform-provider-html" alt="GitHub release (latest by date)"></a>
<a href="https://github.com/koki-develop/terraform-provider-html/actions/workflows/test.yml"><img src="https://img.shields.io/github/actions/workflow/status/koki-develop/terraform-provider-html/test.yml?logo=github" alt="GitHub Workflow Status"></a>
<a href="https://goreportcard.com/report/github.com/koki-develop/terraform-provider-html"><img src="https://goreportcard.com/badge/github.com/koki-develop/terraform-provider-html" alt="Go Report Card"></a>
<a href="./LICENSE"><img src="https://img.shields.io/github/license/koki-develop/terraform-provider-html" alt="LICENSE"></a>
</p>

```hcl
# TODO: example
```

```console
# TODO: example
```

# Getting Started

HTML.tf is a Terraform provider that allows you to write HTML in Terraform configuration files.
To use it, add the following provider configuration.

```hcl
terraform {
  required_providers {
    js = {
      source = "koki-develop/html"
    }
  }
}

provider "html" {}
```

Next, run `terraform init`.

```console
$ terraform init
```

That's it. You are ready to use HTML.tf.

# Examples

- [`examples/`](./examples)

# Documentation

- [Terraform Registry](https://registry.terraform.io/providers/koki-develop/html/latest/docs)

# LICENSE

[MIT](./LICENSE)
