# AI Coder

AI Coder is command-line application designed to accelerate software development by generating scaffolding code and/or refactoring existing code. It leverages OpenAI's API to create project structures, boilerplate code, and implementation suggestions based on simple prompts. If you are running ollama, this application can also call the ollama endpoints with local models.

## Features

- Generate complete project scaffolds from simple descriptions
- Create boilerplate code for various programming languages and frameworks
- Interactive mode to refine and expand generated code
- Customizable templates and configurations

## Installation

### Prerequisites

- Go 1.23 or later
- OpenAI API key

### From Source

Clone the repository:

```bash
git clone https://github.com/am8850/ai-scaffolder.git
cd ai-scaffolder
```

Build the executable:

```bash
go build -o aicoder
```

Installation:

```bash
go install
```

### Configuration

Create an `./aicoder.json` file with the following content:

> Note: set the type to either: azure or openai

> Note: to run locally using ollama:<br>set the endpoint to `http://localhost:11434/v1/chat/completions`
<br>Set the model to the desired model: `phi3`
<br>Set the type to: `openai`
<br>Set the key to `ollama`

```json
{
    "endpoint": "https://<NAME>.openai.azure.com/openai/deployments/gpt-4o/chat/completions?api-version=2025-01-01-preview",
    "model": "gpt-4o",
    "key": "<KEY>",
    "type": "azure",
    "code_system_prompt": "You are an AI that can help scaffold code in any programming language.\n\nRules:\n- If the user requests something not related to scaffold code, do not generate any code.\n- Do your best to make the code very usable from the start.\n\nNo prologue or epilogue.\n\nRespond in the following JSON format:\n{\"files\":[\n{ \"filepath\":\"main.py\", \"code\":\"print('Hello World')\" }\n]}",
    "refactor_system_prompt": "You are an AI that can evaluate the programming code for readability and cyclomatic complexity. \n\nRules:\n- Code can be in any programming language.\n- Provide a readability score from 1 to 10 with 10 being very clear.\n- Provide a cyclomatic complexity score from 1 to 10 with 10 being very complex.\n- Provide the reasons for the scores. \n- Generate version of the code that includes the proposed changes to improve readability and cyclomatic compexity. Do your best to provde the best possible version of the code. Add missing comments to the functions.\n- The code should be in ISO-8859-1 encoding.\n- No prologue or epilogue.\n- Output in the following JSON format only:\n\n{\n\"readability_score\":0,\n\"readability_reason\":\"\",\n\"cyclomatic_score\":0,\n\"cyclomatic_reason\":\"\",\n\"improved_code\":\"import os\nmsg=\"Hello World\"\nprint(msg)\",\n}"

}
```

## Usage

### Basic Usage

```bash
aicoder -p "Create a Python hello world application."
```

Generate main.py:

```python
print('Hello, World!')
```

## Examples

Generate a Go microservice:

```bash
aicoder -p "Create a Golang REST API using Gin to manage customers."
```

Generate main.go

```go
package main

import (
        "github.com/gin-gonic/gin"
)

type Customer struct {
        ID    string `json:"id"`
        Name  string `json:"name"`
        Email string `json:"email"`
}

var customers = []Customer{
        {ID: "1", Name: "John Doe", Email: "john.doe@example.com"},
        {ID: "2", Name: "Jane Smith", Email: "jane.smith@example.com"},
}

func getCustomers(c *gin.Context) {
        c.JSON(200, customers)
}

func getCustomerByID(c *gin.Context) {
        id := c.Param("id")
        for _, customer := range customers {
                if customer.ID == id {
                        c.JSON(200, customer)
                        return
                }
        }
        c.JSON(404, gin.H{"message": "Customer not found"})
}

func createCustomer(c *gin.Context) {
        var newCustomer Customer
        if err := c.ShouldBindJSON(&newCustomer); err != nil {
                c.JSON(400, gin.H{"message": "Invalid input"})
                return
        }
        customers = append(customers, newCustomer)
        c.JSON(201, newCustomer)
}

func deleteCustomer(c *gin.Context) {
        id := c.Param("id")
        for i, customer := range customers {
                if customer.ID == id {
                        customers = append(customers[:i], customers[i+1:]...)
                        c.JSON(200, gin.H{"message": "Customer deleted"})
                        return
                }
        }
        c.JSON(404, gin.H{"message": "Customer not found"})
}

func main() {
        r := gin.Default()

        r.GET("/customers", getCustomers)
        r.GET("/customers/:id", getCustomerByID)
        r.POST("/customers", createCustomer)
        r.DELETE("/customers/:id", deleteCustomer)

        r.Run() // Default listens on :8080
}
```

Scaffold a React component:

```bash
INSTRUCTIONS=$(cat ./instructions.md)
PRD=$(cat ./prd.md)
cliai sc -p "Generate the application based on the following intructions and product specifications: $INSTRUCTIONS $PRD"
```

## Project Structure

- `cmd/`: Command-line interface implementation
- `pkg/`: Core packages
  - `config/`: Configuration handling
  - `openai/`: OpenAI API integration
  - `scaffolder/`: Main scaffolding logic
  - `templates/`: Template management
- `examples/`: Example outputs and use cases
- `docs/`: Documentation

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
