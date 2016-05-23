/*
Package version provides a basic structure for setting
the application version during build using the build
-ldflags parameter.

Example:

  go build -ldflags \
      "-X github.com/dgarlitt/urban-gopher/version.Version=1.0.0 \
      -X github.com/dgarlitt/urban-gopher/version.Commit=ba43d3f9 \
      -X github.com/dgarlitt/urban-gopher/version.Branch=my-feature-branch"

*/
package version
