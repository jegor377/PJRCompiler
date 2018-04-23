package main

import "fmt"
import "os"
import "SectionInformations"
import "ScriptsManager"
import "PJRCompiler"

func main() {
	if len(os.Args) != 3 {
		fmt.Println("pjr <input scripts folder> <output executive file name>")
		os.Exit(1)
	}
	sm := ScriptsManager.New()
	err := sm.LoadScripts(os.Args[1])
	if err != nil {
		panic(err)
	}
	sinfo := SectionInformations.New()

	compiler := PJRCompiler.New(sinfo)
	compiler.Load(sm)
	val, err := compiler.Compile()
	if err != nil {
		panic(err)
	}
	sm.SaveScript(os.Args[2], val)
}