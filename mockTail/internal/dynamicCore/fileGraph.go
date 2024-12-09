package dynamicCore

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type NodeDetails struct {
	Name      string
	Path      string
	IsDynamic bool
}

type Node struct {
	Info  NodeDetails
	Nodes []*Node
	Data  []*FileResource
}

func NewDirNode(path string) *Node {
	return &Node{
		Info: NodeDetails{
			Path: path,
		},
	}
}

func (node *Node) BuildTree() error {
	entries, err := os.ReadDir(node.Info.Path)
	if err != nil {
		return nil
	}

	for _, entry := range entries {
		if entry.IsDir() {
			child := NewDirNode(filepath.Join(node.Info.Path, entry.Name()))
			child.BuildMetaData(entry)
			err = child.BuildTree()
			if err != nil {
				return err
			}
			node.Add(child)
		} else {
			node.BuildData(entry)
		}
	}
	return nil
}

func (node *Node) BuildMetaData(entry os.DirEntry) {
	node.Info.Name = entry.Name()
	node.Info.IsDynamic = false
	// TODO: Implement dynamic node detection logic separately
	if strings.HasPrefix(node.Info.Name, "(") && strings.HasSuffix(node.Info.Name, ")") {
		node.Info.IsDynamic = true
	}
}

func (node *Node) FindData(request Request) (*FileResource, error) {
	for _, resource := range node.Data {
		if err := resource.Is(request); err != nil {
			return resource, nil
		}
	}
	return nil, errors.New("resource not found")
}

func (node *Node) BuildData(entry os.DirEntry) {
	p := NewParser(entry)
	res, _ := NewResource(p)
	res.CompletePath = filepath.Join(node.Info.Path, entry.Name())
	node.Data = append(node.Data, res)
}

func (node *Node) Is(query string) bool {
	return node.Info.Path == query
}

func (node *Node) Find(query string) *Node {
	if node.Is(query) {
		return node
	}

	for _, child := range node.Nodes {
		if result := child.Find(query); result != nil {
			return result
		}
	}
	return nil
}
func (node *Node) Add(child *Node) {
	node.Nodes = append(node.Nodes, child)
}

func (node *Node) Remove() error {
	return nil
}

//func watchDirectory(rootPath string) error {
//  watcher, err := fsnotify.NewWatcher()
//  if err != nil {
//    return err
//  }
//  defer watcher.Close()
//
//  done := make(chan bool)
//
//  go func() {
//    for {
//      select {
//      case event, ok := <-watcher.Events:
//        if !ok {
//          return
//        }
//        fmt.Printf("Event: %s, File: %s\n", event.Op, event.Name)
//        if event.Op&fsnotify.Create == fsnotify.Create {
//          fmt.Println("File created:", event.Name)
//        }
//        if event.Op&fsnotify.Write == fsnotify.Write {
//          fmt.Println("File modified:", event.Name)
//        }
//
//      case err, ok := <-watcher.Errors:
//        if !ok {
//          return
//        }
//        log.Println("Error:", err)
//      }
//    }
//  }()
//
//  err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
//    if err != nil {
//      return err
//    }
//    if info.IsDir() {
//      return watcher.Add(path)
//    }
//    return nil
//  })
//  if err != nil {
//    return err
//  }
//
//  <-done
//  return nil
//}
