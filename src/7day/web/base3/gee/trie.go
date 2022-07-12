package gee

import (
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || n.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查询
func (n *node) matchChildren(part string) []*node {
	matchNodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || n.isWild {
			matchNodes = append(matchNodes, child)
		}
	}
	return matchNodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil { // 说明当前节点不存在该子节点，则在下一级添加该节点
		child = &node{part: part, isWild: isWild(part)}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" { // 末尾节点的 pattern 是完整的，在 insert() 中设置过
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}

	for _, child := range n.children {
		child.travel(list)
	}
}

func isWild(part string) bool {
	if part == "" {
		return false
	}
	return part[0] == ':' || part[0] == '*'
}
