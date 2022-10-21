package grule

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// KnowledgeLib ast.KnowledgeLibrary
type KnowledgeLib = ast.KnowledgeLibrary

// Grule extend ast.KnowledgeLibrary
type Grule struct {
	*KnowledgeLib
}

// Init init grule
func Init() (grule *Grule, err error) {
	knowledgeLib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLib)

	err = filepath.WalkDir("./grules", func(s string, file fs.DirEntry, err error) error {
		if filepath.Ext(file.Name()) == ".grl" {
			fileNameWithoutExt := strings.TrimSuffix(file.Name(), ".grl")
			fileNameSplit := strings.Split(fileNameWithoutExt, "_")

			ruleName := fileNameSplit[0]
			ruleVersion := "0.0.1"

			if len(fileNameSplit) == 2 {
				ruleVersion = fileNameSplit[1]
			}

			fmt.Printf("loading %s as %s(%s)\n", file.Name(), ruleName, ruleVersion)

			fileRes := pkg.NewFileResource(fmt.Sprintf("./grules/%s", file.Name()))
			err := ruleBuilder.BuildRuleFromResource(ruleName, ruleVersion, fileRes)
			if err != nil {
				panic(err)
			}

			fmt.Printf("%s loaded as %s(%s)\n", file.Name(), ruleName, ruleVersion)
		}

		return nil
	})

	if err != nil {
		return
	}

	grule = &Grule{KnowledgeLib: knowledgeLib}

	return grule, err
}
