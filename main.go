package main

import (
	"gee"
	"net/http"
)

// func main() {
// 	http.HandleFunc("/", indexHandler)
// 	http.HandleFunc("/hello", helloHanlder)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

// func helloHanlder(w http.ResponseWriter, r *http.Request) {
// 	for k, v := range r.Header {
// 		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 	}
// }

// type Engine struct{}

// func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// 	case "/hello":
// 		for k, v := range r.Header {
// 			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 		}
// 	default:
// 		fmt.Fprintf(w, "404 NOT FOUND : %s\n", r.URL)
// 	}
// }

// func main() {
// 	engine := new(Engine)
// 	log.Fatal(http.ListenAndServe(":8080", engine))
// }

// func main() {
// 	r := gee.New()
// 	r.GET("/", func(c *gee.Context) {
// 		c.HTML(http.StatusOK, "<h1>hello Gee</h1>")
// 	})
// 	r.GET("/hello", func(c *gee.Context) {
// 		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Query("name"), c.Path)
// 	})
// 	r.POST("/login", func(c *gee.Context) {
// 		c.JSON(http.StatusOK, gee.H{
// 			"username": c.PostFrom("username"),
// 			"password": c.PostFrom("password"),
// 		})
// 	})

// 	r.Run(":8080")
// }

// func main() {
// 	r := gee.New()

// 	r.GET("/", func(c *gee.Context) {
// 		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
// 	})

// 	r.GET("/hello", func(c *gee.Context) {
// 		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Query("name"), c.Path)
// 	})

// 	r.GET("/hello/:name", func(c *gee.Context) {
// 		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
// 	})

// 	r.GET("/assets/*filepath", func(c *gee.Context) {
// 		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
// 	})

// 	r.Run(":8080")
// }

// func main() {
// 	r := gee.New()

// 	r.GET("/", func(c *gee.Context) {
// 		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
// 	})

// 	v1 := r.Group("/v1")
// 	{
// 		v1.GET("/", func(c *gee.Context) {
// 			c.HTML(http.StatusOK, "<h1>hello gee</h1>")
// 		})

// 		v1.GET("/hello", func(c *gee.Context) {
// 			//expect /hello?name=guagua
// 			c.String(http.StatusOK, "hello %s , you're at %s\n", c.Query("name"), c.Path)
// 		})
// 	}

// 	v2 := r.Group("/v2")
// 	{
// 		v2.GET("/hello/:name", func(c *gee.Context) {
// 			//expect /hello/guagua
// 			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
// 		})

// 		v2.POST("/login", func(c *gee.Context) {
// 			c.JSON(http.StatusOK, gee.H{
// 				"username": c.PostFrom("username"),
// 				"password": c.PostFrom("password"),
// 			})
// 		})
// 	}

// 	r.Run(":8080")
// }

// func onlyForV2() gee.HandleFunc {
// 	return func(c *gee.Context) {
// 		t := time.Now()
// 		c.Fail(500, "Internal Server Error")
// 		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))

// 	}
// }

// func main() {
// 	r := gee.New()
// 	r.Use(gee.Logger()) //global midlleware

// 	r.GET("/", func(c *gee.Context) {
// 		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
// 	})

// 	v2 := r.Group("/v2")
// 	v2.Use(onlyForV2()) //v2 group middleware
// 	{
// 		v2.GET("hello/:name", func(c *gee.Context) {
// 			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
// 		})
// 	}
// 	r.Run(":8080")
// }

// type student struct {
// 	Name string
// 	Age  int8
// }

// func FormatAsDate(t time.Time) string {
// 	year, month, day := t.Date()
// 	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
// }

// func main() {
// 	r := gee.New()
// 	r.Use(gee.Logger())
// 	r.SetFuncMap(template.FuncMap{
// 		"FormaAsDate": FormatAsDate,
// 	})
// 	r.LoadHTMLGlob("templates/*")
// 	r.Static("/assets", "./static")

// 	stu1 := &student{Name: "Geektutu", Age: 20}
// 	stu2 := &student{Name: "Jack", Age: 22}

// 	r.GET("/", func(c *gee.Context) {
// 		c.HTML(http.StatusOK, "css.tmpl", nil)
// 	})

// 	r.GET("/student", func(c *gee.Context) {
// 		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
// 			"title":  "gee",
// 			"stuArr": [2]*student{stu1, stu2},
// 		})
// 	})

// 	r.GET("/date", func(c *gee.Context) {
// 		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
// 			"title": "gee",
// 			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
// 		})
// 	})

// 	r.Run(":8080")

// }

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "hello Geektutu\n")
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":8080")
}
