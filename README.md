
# Structure 
/go-tailwind-htmx
|-- /cmd
|   |-- main.go              <-- Main Go file
├── /internal
│   ├── /handlers
│   │   ├── handlers.go
│   │
│   ├── /middleware
│   │   ├── middleware.go
│   │
│   ├── /templates
│   │   ├── template.go
│   │
│   ├── /contextkeys
│   │   ├── contextkeys.go
|-- /views
|   |-- /layout
|       |-- default.templ
|       |-- full-page.templ
|   |-- /pages
|       |-- about.templ
|   |-- /partials
|       |-- footer.templ
|       |-- header.templ
|   |-- index.templ       <-- Main template file (with layout)
|-- /public
|   |-- /dist
|       |-- output.css       <-- Compiled Tailwind CSS
|-- /src
|   |-- /styles
|       |-- tailwind.css     <-- Source Tailwind CSS
|-- go.mod
|-- go.sum
|-- postcss.config.js
|-- tailwind.config.js

# Prerequisite

## 1. INSTALL GO
  ### 1.(a) Download and install Go
  Here: https://go.dev/doc/install
  
  Use below instructions to resolve command not found: go.
  
  ### 1.(b) Add the Go bin Directory to PATH
  Ensure the Go bin directory is added to your system’s PATH, especially for the M1 chip. To do this, add the following to your .zshrc (or .rc): `export PATH=$PATH:$(go env GOPATH)/bin`
  
  ### 1.(c) Then, source your updated .zshrc
  Run: `source ~/.zshrc`
    
  ### 1.(d) Check Go Path Configuration:
  Ensure Go itself is properly set up and you have access to the go env GOPATH. 
  Run: `go env GOPATH` This should return the correct GOPATH where binaries, including Air, are stored.

## 2. INSTALL AIR (Live reload for Go apps)
  In this case we are using this Binary which will be `$(go env GOPATH)/bin/air` https://github.com/air-verse/air?tab=readme-ov-file#installation
  
  ### 2.(a) Install Air
  Run: `curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin`
  
  Run: `air -v` to check if air was installed successfully and running.

  ### 2.(b) Use Air and Tailwind CSS together
  You need to make sure both processes (Go server and Tailwind compilation) are running simultaneously. Here's how to integrate Tailwind CSS watching with Air, so it rebuilds both your Go app and CSS changes automatically.

  ### 2(c). To Watch Tailwind CSS with Air
  Modify the `.air.toml` file by adding `css` ext like this `include_ext = ["go", "tpl", "templ", "html", "css"]`
  
## 3. INSTALL TASK VIA GO
  Ensure you have Go installed by running:
  `go version`

  Install Go Task using go install:
  `go install github.com/go-task/task/v3/cmd/task@latest`
  
  Add Go's binary path to your system's $PATH: Add the following line to your `~/.zshrc` or `~/._profile`, depending on your shell:
  `export PATH=$PATH:$(go env GOPATH)/bin`
  
  Reload your shell configuration:
  `source ~/.zshrc`  # or `~/._profile`

  Verify the installation:
  `task --version`

# Usage
  ## Start the App
  Run `task dev` to start both Air and Tailwind. (For dev environment)
  OR
  Just Run `air`.