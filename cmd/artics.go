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

	if h {
		fmt.Println("Usage: artics [-o] <option>")
		fmt.Println("Options:")
		fmt.Println("create-model [-n] <model-name> [args]")
		fmt.Println("Args:")
		fmt.Println("-r : Create model along with request and repository")
		fmt.Println("")

		fmt.Println("create-request [-n] <request-name>")
		fmt.Println("create-repository [-n] <request-name>")
		fmt.Println("create-controller [-n] <controller-name>")
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
