package main

import (
	"flag"

	"github.com/fatih/color"
)

const (
	ks = `
	    ■■                  
	    ■■        ■■■■■■■■  
	■■■■■■■■■■           ■  
	    ■    ■          ■   
	    ■    ■          ■   
	    ■    ■         ■    
	   ■■    ■        ■■    
	   ■     ■       ■■■■   
	  ■■     ■      ■■  ■■  
	 ■■     ■■    ■■■    ■■ 
	 ■    ■■■    ■■       ■ 
`
)

var (
	red     = flag.Bool("red", false, "color: red")
	blue    = flag.Bool("blue", false, "color: blue")
	green   = flag.Bool("green", false, "color: green")
	magenta = flag.Bool("magenta", false, "color: magenta")
	yellow  = flag.Bool("yellow", false, "color: yellow")
	cyan    = flag.Bool("cyan", false, "color: cyan")
)

func init() {
	flag.Parse()
}

func main() {
	var co *color.Color
	switch {
	case *red:
		co = color.New(color.FgHiRed)
	case *blue:
		co = color.New(color.FgHiBlue)
	case *green:
		co = color.New(color.FgHiGreen)
	case *magenta:
		co = color.New(color.FgHiMagenta)
	case *yellow:
		co = color.New(color.FgHiYellow)
	case *cyan:
		co = color.New(color.FgHiCyan)
	default:
		co = color.New()
	}
	co.Add(color.Bold)

	co.Println(ks)
}
