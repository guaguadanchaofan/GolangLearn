package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandleFunc
}

func newRouter() *router {
	//return &router{handlers: make(map[string]HandleFunc)}
	return &router{roots: make(map[string]*node),
		handlers: make(map[string]HandleFunc)}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	// log.Printf("Route %4s - %s", method, pattern)
	// key := method + "-" + pattern
	// r.handlers[key] = handler
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	serachParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}
	n := root.search(serachParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = serachParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(serachParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	// key := c.Method + "-" + c.Path
	// if handler, ok := r.handlers[key]; ok {
	// 	handler(c)
	// } else {
	// 	c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	// }

	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.params = params
		key := c.Method + "-" + n.pattern
		//r.handlers[key](c)
		c.handlders = append(c.handlders, r.handlers[key])
	} else {
		c.handlders = append(c.handlders, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
