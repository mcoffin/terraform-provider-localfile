# terraform-provider-localfile
Simple, lightweight terraform provider to provide easy creation of local files.

# Caveats

## API Version

`terraform-provider-localfile` must be compiled against a version of terraform that uses the same version of the plugin api as the terraform version where the plugin will be deployed. If you get an error similar to this one:

```
Error configuring: 1 error(s) occurred:

* Incompatible API version with plugin. Plugin version: 3, Ours: 2
```

Then you'll need to make sure to checkout the right version of terraform prior to compiling `terraform-provider-localfile`.

# License
`terraform-provider-localfile` is licensed under the MIT License. See `LICENSE` file.
