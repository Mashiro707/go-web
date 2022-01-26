package gow

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由
	part     string  // 路由中的一部分
	children []*node // 字节点
	isWild   bool    // 是否精准匹配
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

// 一边匹配一边插入
func (n *node) insert(pattern string, parts []string, height int) {
	// 如果已经匹配完了，则将pattern赋值给该node，表示他是一个完整的url
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	// 如果没有匹配上，那么进行生成，放到n节点的子列表中
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		if child.isWild && len(n.children) > 0 {
			panic(part + "同级已经有路由")
		}
		if part[0] == ' ' && len(parts) > height+1 {
			panic(part + "不能出现在中间")
		}
		if (part[0] == ' ' || part[0] == ':') && len(part) == 1 {
			panic(part + "不能单独出现")
		}
		n.children = append(n.children, child)
	}
	// 接着插入下一个part节点
	child.insert(pattern, parts, height+1)
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.isWild {
			panic(part + "同级已经有" + child.part)
		}
		if child.part == part {
			return child
		}
	}
	return nil
}

func (n *node) search(parts []string, height int) *node {
	// 找到末尾或者通配符
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	// 获取所有可能的子路径
	children := n.matchChildren(part)

	for _, child := range children {
		// 对每个路径接着查找下一个part
		result := child.search(parts, height+1)
		// 如果找到了就立即返回
		if result != nil {
			return result
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 查找所有完整的url，保存到列表中
func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		// 递归终止条件
		*list = append(*list, n)
	}
	for _, child := range n.children {
		// 一层一层的递归找pattern是非空的节点
		child.travel(list)
	}
}
