package queryparser

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FunctionValue represents a function node
type FunctionValue struct {
	Function string `json:"function"`
	Arg1     string `json:"arg1,omitempty"`
	Arg2     string `json:"arg2,omitempty"`
	Arg      string `json:"arg,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value1   string `json:"value1,omitempty"`
	Value2   string `json:"value2,omitempty"`
	Value    string `json:"value,omitempty"`
}

type ParamsTree struct {
	Root Node `json:"paramsTree"`
}

// Node represents a node in the tree
type Node struct {
	Operator string         `json:"operator,omitempty"`
	Func     *FunctionValue `json:"func,omitempty"`
	Children []*Node        `json:"children,omitempty"`
}

func Parse() {
	jsonData := `{
  "paramsTree": {
    "operator": "OR",
    "children": [
      {
        "func": {
          "function": "Gain",
          "arg1": "fc3401",
          "arg2": "ti8910",
          "operator": "between",
          "value1": "-3.0580673",
          "value2": "65.36908"
        }
      },
      {
        "operator": "AND",
        "children": [
          {
            "func": {
              "function": "Gain",
              "arg1": "g2",
              "arg2": "g3",
              "operator": "between",
              "value1": "1",
              "value2": "2"
            }
          },
          {
            "func": {
              "function": "Gain",
              "arg1": "g4",
              "arg2": "g5",
              "operator": "between",
              "value1": "3",
              "value2": "4"
            }
          }
        ]
      },
      {
        "func": {
          "function": "Error",
          "arg": "fc3401",
          "operator": "<",
          "value": "4"
        }
      }
    ]
  }
}`

	var pt ParamsTree
	err := json.Unmarshal([]byte(jsonData), &pt)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	printTree(&pt.Root, 0)
}

func printTree(node *Node, level int) {
	indent := strings.Repeat("\t", level)

	if node.Operator != "" {
		fmt.Printf("%s%s\n", indent, node.Operator)
	}

	if node.Func != nil {
		fmt.Printf("%sFunction: %s\n", indent, node.Func.Function)
		fmt.Printf("%s\tArg1: %s\n", indent, node.Func.Arg1)
		fmt.Printf("%s\tArg2: %s\n", indent, node.Func.Arg2)
		fmt.Printf("%s\tArg: %s\n", indent, node.Func.Arg)
		fmt.Printf("%s\tOperator: %s\n", indent, node.Func.Operator)
		fmt.Printf("%s\tValue1: %s\n", indent, node.Func.Value1)
		fmt.Printf("%s\tValue2: %s\n", indent, node.Func.Value2)
		fmt.Printf("%s\tValue: %s\n", indent, node.Func.Value)
	}

	for _, child := range node.Children {
		printTree(child, level+1)
	}
}
