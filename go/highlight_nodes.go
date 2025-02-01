package main

import "fmt"

// if there are cycles in this data then we have to mark the visited nodes in the graph map

type nodeID string
type nodeType string

type node struct {
	nodeID
	nodeType
	highlighted bool
	links       map[string][]nodeID
}

type highlightGraph struct {
	m    map[nodeID]*node
	head *node
}

func newHighlightGraph(nodes []*node) highlightGraph {
	m := make(map[nodeID]*node)
	var head *node

	for _, n := range nodes {
		if n.nodeType == "input" {
			head = n
		}
		m[n.nodeID] = n
	}

	return highlightGraph{
		m,
		head,
	}
}

func solveHighlightNodes() {

	nodeData := []*node{
		{
			nodeID:      "a",
			nodeType:    "input",
			highlighted: false,
			links: map[string][]nodeID{
				"outgoing": {"b"},
			},
		},
		{
			nodeID:      "b",
			nodeType:    "processor",
			highlighted: false,
			links: map[string][]nodeID{
				"incoming": {"a"},
				"outgoing": {"c"},
			},
		},
		{
			nodeID:      "c",
			nodeType:    "output",
			highlighted: false,
			links: map[string][]nodeID{
				"incoming": {"b"},
			},
		},
	}

	hg := newHighlightGraph(nodeData)
	hg.highlightDFS("c")
	fmt.Println(nodeData)
}

func (hg *highlightGraph) highlightDFS(id nodeID) {
	n := hg.m[id]
	n.highlighted = true

	for _, nl := range n.links {
		for _, nid := range nl {
			hg.highlightDFS(nid)
		}
	}
}
