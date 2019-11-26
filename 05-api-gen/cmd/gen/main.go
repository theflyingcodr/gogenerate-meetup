package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"

	"golang.org/x/tools/go/ast/astutil"
)

var (
	domain      string
	domainTitle string
	templates   = []string{"domain", "http", "mem_store", "service"}
)

func main() {
	log.Println("Starting Generator...")
	flag.StringVar(&domain, "domain", "test", "name of domain to generate")
	flag.Parse()
	if strings.TrimSpace(domain) == "" {
		log.Fatal("failed to generate api domain, no domain passed")
	}
	domain = strings.ToLower(domain)
	domainTitle = strings.Title(domain)
	log.Println("Inputs successfully parsed, starting generation of new domain...")
	generateDomain()
	log.Println("All code generated successfully, updating routes.go with new route constant...")
	updateRoutes()
	updateMain()
	log.Println("Generation completed")
}

func generateDomain() {
	args := map[string]string{
		"Domain":      domain,
		"DomainTitle": domainTitle,
		"Date":        time.Now().Format(time.RFC3339),
	}
	if err := os.Mkdir("../../"+domain+"s", 0755); err != nil {
		log.Fatalf("failed to create %s directory %s", domain, err)
	}
	for _, tmplName := range templates {
		log.Printf("Creating %s.go file\n", tmplName)
		bb, err := ioutil.ReadFile(filepath.Join("templates", tmplName+".template"))
		if err != nil {
			log.Fatalf("failed to read %s: %s", tmplName, err)
		}
		tmpl, err := template.New(tmplName).Parse(string(bb))
		if err != nil {
			log.Fatalf("failed to parse template %s: %s", tmplName, err)
		}
		buf := bytes.Buffer{}
		if err := tmpl.Execute(&buf, args); err != nil {
			log.Fatalf("failed to parse template %s: %s", tmplName, err)
		}
		fn := tmplName + ".go"
		if tmplName == "domain" {
			fn = domain + "s.go"
		}
		formatted, err := format.Source(buf.Bytes())
		if err != nil {
			log.Fatalf("go/format: %s", err)
		}
		if err := ioutil.WriteFile("../../"+domain+"s/"+fn, formatted, 0644); err != nil {
			log.Fatalf("failed to write template %s: %s", tmplName, err)
		}
		log.Printf("Created %s.go file successfully\n", tmplName)
	}
}

func updateRoutes() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "../../routes.go", nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("failed to parse AST from routes.go: %s", err)
	}
	constname := fmt.Sprintf("Route%ssGET", domainTitle)
	d, _ := f.Decls[0].(*ast.GenDecl)
	if d.Tok == token.CONST {
		for _, s := range d.Specs {
			c := s.(*ast.ValueSpec)

			for _, n := range c.Names {
				if n.Name == constname {
					log.Printf("Constant %s already defined %s, nothing more to do\n", domain, n.Name)
					return
				}
			}
		}
	}
	fmt.Printf("Constant for this domain not previously defined, updating routes.go with %s\n", constname)

	con := ast.NewIdent(constname)
	con.Obj = ast.NewObj(ast.Con, constname)
	conVal := &ast.BasicLit{
		Kind:  token.STRING,
		Value: "\"" + "/api/v1/" + domain + "s\"",
	}
	vspec := &ast.ValueSpec{
		Names: []*ast.Ident{con},
		Values: []ast.Expr{
			conVal,
		},
	}

	d.Specs = append(d.Specs, vspec)
	sort.Slice(d.Specs, func(i, j int) bool {
		return d.Specs[i].(*ast.ValueSpec).Names[0].Name < d.Specs[j].(*ast.ValueSpec).Names[0].Name
	})
	fp := "../../routes.go"
	formatAndSave(fs, f, fp)
	fmt.Printf("%s constant successfully added to routes.go\n", constname)
}

func updateMain() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "../api/main.go", nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("failed to parse AST from main.go: %s", err)
	}
	astutil.AddImport(fs, f, fmt.Sprintf("github.com/theflyingcodr/gogenerate-meetup/05-api-gen/%ss", domain))
	ast.Inspect(f, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}
		route := &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent(domain + "s"),
							Sel: ast.NewIdent("NewHttpHandler"),
						},
						Args: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent(domain + "s"),
									Sel: ast.NewIdent("New" + domainTitle + "Svc"),
								},
								Args: []ast.Expr{
									&ast.UnaryExpr{
										Op: token.AND,
										X: &ast.CompositeLit{
											Type: &ast.SelectorExpr{

												X:   ast.NewIdent(domain + "s"),
												Sel: ast.NewIdent("MemoryStore"),
											},
										},
									},
								},
							},
						},
					},
					Sel: ast.NewIdent("Register"),
				},
			},
		}
		fn.Body.List = append([]ast.Stmt{route}, fn.Body.List...)
		return false
	})
	fp := "../api/main.go"
	formatAndSave(fs, f, fp)
	fmt.Printf("main.go wired up with new endpoint")
}

func formatAndSave(fs *token.FileSet, f *ast.File, filePath string) {
	var output []byte
	buffer := bytes.NewBuffer(output)
	if err := printer.Fprint(buffer, fs, f); err != nil {
		log.Fatalf("failed to add constant to resources.go: %s", err)
	}
	bb, err := format.Source(buffer.Bytes())
	if err != nil {
		log.Fatalf("failed to format file %s", err)
	}

	if err := ioutil.WriteFile(filePath, bb, 0755); err != nil {
		log.Fatalf("failed to create %s", err)
	}

}
