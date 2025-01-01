package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/go-restful/helper"
)

func createModel(modelName string) {
	filePath := "app/model/" + helper.ToSnakeCase(modelName) + ".go"
	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("Model already exists")
		fmt.Println("===================================")
		return
	}

	fmt.Println("Creating model: ", modelName)
	file, err := os.Create(filePath)

	if err != nil {
		file.Close()
		panic(err)
	}

	defer file.Close()
	tmpl := template.Must(template.New("model"), nil)
	tmpl, err = template.ParseFiles("cmd/templates/model.tmpl")
	if err != nil {
		panic(err)

	}

	tmpl.Execute(file, modelName)
	fmt.Printf("Model created successfully at %s\n", filePath)

}
func createRepository(repoName string) {
	filePath := "app/repository/" + helper.ToSnakeCase(repoName) + "_repository.go"
	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("Repository already exists")
		fmt.Println("===================================")
		return
	}

	fmt.Println("Creating repository: ", repoName)
	file, err := os.Create(filePath)

	if err != nil {
		file.Close()
		panic(err)
	}

	defer file.Close()
	tmpl := template.Must(template.New("repository"), nil)
	tmpl, err = template.ParseFiles("cmd/templates/repository.tmpl")
	if err != nil {
		panic(err)

	}

	tmpl.Execute(file, repoName)
	fmt.Printf("Repository created successfully at %s\n", filePath)
}
func createRequest(requestName string) {
	filePath := "app/request/" + helper.ToSnakeCase(requestName) + "_request.go"
	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("Request already exists")
		fmt.Println("===================================")
		return
	}

	fmt.Println("Creating request: ", requestName)
	file, err := os.Create(filePath)

	if err != nil {
		file.Close()
		panic(err)
	}

	defer file.Close()
	tmpl := template.Must(template.New("request"), nil)
	tmpl, err = template.ParseFiles("cmd/templates/request.tmpl")
	if err != nil {
		panic(err)

	}

	tmpl.Execute(file, requestName)
	fmt.Printf("Request created successfully at %s\n", filePath)
}
func createController(controllerName string) {
	filePath := "app/controller/" + helper.ToSnakeCase(controllerName) + "_controller.go"
	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("Controller already exists")
		fmt.Println("===================================")
		return
	}

	fmt.Println("Creating controller: ", controllerName)
	file, err := os.Create(filePath)

	if err != nil {
		file.Close()
		panic(err)
	}

	defer file.Close()
	tmpl := template.Must(template.New("controller"), nil)
	tmpl, err = template.ParseFiles("cmd/templates/controller.tmpl")
	if err != nil {
		panic(err)

	}

	tmpl.Execute(file, controllerName)
	fmt.Printf("Controller created successfully at %s\n", filePath)
}

func main() {
	var (
		o string
		n string
		h bool
		r bool
	)
	flag.BoolVar(&h, "h", false, "Show help")
	flag.StringVar(&o, "o", "", "Choose the options")
	flag.StringVar(&n, "n", "", "Model name")
	flag.BoolVar(&r, "r", false, "Create model along with request and repository")

	flag.Parse()

	if h && o == "create-model" {
		fmt.Println("Available falgs:")
		fmt.Println("-n: \tModel name")
		fmt.Print("-r: \tDetermines whether to create request and repository along with model\n\n")
		return
	}
	if h && o == "create-repository" {
		fmt.Println("Available falgs:")
		fmt.Println("-n: \tRepository name")
		return
	}
	if h && o == "create-request" {
		fmt.Println("Available falgs:")
		fmt.Println("-n: \tRequest name")
		return
	}
	if h && o == "create-controller" {
		fmt.Println("Available falgs:")
		fmt.Println("-n: \tController name")
		return
	}

	if h {
		fmt.Println("Usage: artics [-o] <operation>")
		fmt.Println("Operations:")
		fmt.Println("create-model: Create a model with given name with -n flag")
		fmt.Println("create-request: Create a request with given name with -n flag")
		fmt.Println("create-repository: Create a repository with given name with -n flag")
		fmt.Println("create-controller: Create a controller with given name with -n flag")
		fmt.Println("-h: show help")
		return
	}

	switch o {
	case "create-model":
		modelName := n
		if n == "" {
			fmt.Println("Model name is required")
			return
		}

		createModel(modelName)
		if r {
			createRepository(modelName)
			createRequest(modelName)
		}
	case "create-repository":
		repoName := n
		if repoName == "" {
			fmt.Println("Repository name is required")
			return

		}

		createRepository(n)
	case "create-request":
		requestName := n
		if requestName == "" {
			fmt.Println("Request name is required")
			return
		}

		createRequest(n)
	case "create-controller":
		controllerName := n
		if controllerName == "" {
			fmt.Println("Controller name is required")
			return
		}

		createController(controllerName)
	default:
		fmt.Println("Invalid command")
	}
}
