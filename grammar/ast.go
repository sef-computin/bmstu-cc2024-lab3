package grammar

import (
	"fmt"
	"os"
	"os/exec"
)

type Node struct {
	Id       int
	Value    string
	Children []*Node
}

var idCounter = 0

func NewNode(val string) *Node {
	idCounter++
	return &Node{Id: idCounter, Value: val, Children: []*Node{}}
}

func (n *Node) AddChild(token string) (ret *Node) {
	ret = NewNode(token)
	n.Children = append(n.Children, ret)
	return
}

func (n *Node) GetName() string {
	return fmt.Sprintf("\"%s-%d\"", n.Value, n.Id)
}

func DrawTree(filename string, root *Node) {
	var buf []byte = []byte("digraph AST{\nforcelabels=true;\n")

	buf = drawNodes(buf, root)

	buf = append(buf, []byte("\n}")...)

	name := fmt.Sprintf("./%s.gv", filename)

	os.WriteFile(name, buf, 0666)

	cmd := exec.Command("dot", "-Tsvg", name)
	stdout, err := cmd.Output()
	if err == nil {
		file, _ := os.Create(name + ".svg")
		file.Write(stdout)
	}
	os.Remove(name)

	cmd = exec.Command("firefox", name+".svg")
	_, _ = cmd.Output()

	// os.Remove(name+"")

	return
}

func drawNodes(buf []byte, node *Node) []byte {
	if node.Children == nil || len(node.Children) == 0 {
		buf = append(buf, []byte(fmt.Sprintf("%s [ shape=ellipse, label=\"%s\" ];\n", node.GetName(), node.Value))...)
	} else {
		buf = append(buf, []byte(fmt.Sprintf("%s [ shape=box, label=\"%s\" ];\n", node.GetName(), node.Value))...)
		for _, child := range node.Children {
			buf = drawNodes(buf, child)
			buf = append(buf, []byte(fmt.Sprintf("%s -> %s;\n", node.GetName(), child.GetName()))...)
		}
	}

	return buf
}
