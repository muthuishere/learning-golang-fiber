### Learning Golang with Fiber


- [ x ] Printing Hello World

```go

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
```

- [ x ] Get Path Variables
    
    ```go
        app.Get("/:name", func(c *fiber.Ctx) error {
            return c.SendString("Hello, " + c.Params("name"))
        })
    ```
  
- How to Convert Data Types from String

    [strconv](https://pkg.go.dev/strconv)


- How to remove unused Dependencies

    ```go
        go mod tidy
    ```
  

- How to have a live reload for Fiber application 
  - Install air in your system
    ```bash
        go install github.com/cosmtrek/air@latest 
    ```
    - add a config file for air
      ```bash
          touch .air.toml
      ```
      - add the following code to the config file
          ```toml
            # .air.toml
          root = "."
          tmp_dir = "tmp"
          [build]
          cmd = "go build -o ./tmp/main ."
          bin = "./tmp/main"
          delay = 1000 # ms
          exclude_dir = ["apiscripts", "tmp", "vendor"]
          include_ext = ["go", "tpl", "tmpl", "html"]
          exclude_regex = ["_test\\.go"]
          ``` 
    

- What should be the status code , if the parameters are invalid
    - 400 Bad Request ( I am going with it)

- How to Send Invalid Status
    ```go
        c.SendStatus(fiber.StatusBadRequest)
    ```
  


