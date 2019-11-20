package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"go/types"
	"golang.org/x/tools/go/packages"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

var (
	typeName string
	output string
)

func main() {
	flag.StringVar(&typeName, "type", "", "Name of type to generate mock for")
	flag.StringVar(&output, "output", "", "name of outputed file")
	flag.Parse()
	if strings.TrimSpace(typeName) == "" {
		log.Fatal("type must be provided")
	}
	fmt.Printf("Starting mock generation for %s\n", typeName)
	// load the packing information along with Type and Syntax information.
	// I just want to load the current package, but for  proper version
	// package pattern would be an input field
	pkg, _ := packages.Load(&packages.Config{
		Mode:       packages.NeedTypes | packages.NeedSyntax | packages.NeedDeps | packages.NeedImports,
		Dir:        ".",
	}, ".")
	if len(pkg) == 0{
		log.Fatal("cannot find package main")
	}
	obj := pkg[0].Types.Scope().Lookup(typeName)
	if !types.IsInterface(obj.Type()){
		log.Fatalf("%s is not an Interface", typeName)
	}
	// convert our Object to an interface and call complete to gather all method information.
	intface := obj.Type().Underlying().(*types.Interface).Complete()
	o := Obj{
		PkgName: "main",
		TypeName: typeName,
		Methods:  make([]Method,0),
	}
	fmt.Println("Enumerating methods")
	for i := 0; i < intface.NumMethods(); i++{
		meth := intface.Method(i)
		sig := meth.Type().(*types.Signature)
		method := Method{
			Name:      meth.Name(),
			Params:sig.Params().String(),
			Returns:sig.Results().String(),
		}
		var paramNames []string
		for ii := 0; ii < sig.Params().Len();ii++{
			paramNames = append(paramNames, sig.Params().At(ii).Name())
		}
		method.ParamNames = strings.Join(paramNames,", ")
		o.Methods = append(o.Methods, method)
	}
	fmt.Println("Compiling output")
	bb, _ := ioutil.ReadFile("mock.template")
	tmpl := template.Must(template.New("mock").Parse(fmt.Sprintf("%s", bb)))
	// run the text/template with our args
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, o);err != nil {
		log.Fatalf("failed to exec template %s", err)
	}
	// run go format on the executed template to tidy our source
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("go/format: %s", err)
	}
	// output the file
	var formatBytes bytes.Buffer
	if _, err := formatBytes.Write(formatted); err != nil {
		log.Fatalf("failed to write formated %s", err)
	}
	outputName := strings.ToLower(typeName) + "_mock.go"
	if output != ""{
		outputName = output
		if !strings.HasSuffix(outputName, ".go"){
			outputName = outputName + ".go"
		}
	}
	if err := ioutil.WriteFile( outputName, formatBytes.Bytes(), 0644); err != nil{
		log.Fatalf("failed to write to file %s", err)
	}
	fmt.Println("All done")
}

type Obj struct{
	PkgName string
	TypeName string
	Methods []Method
}

type Method struct{
	Name string
	Params string
	Returns string
	ParamNames string
}
