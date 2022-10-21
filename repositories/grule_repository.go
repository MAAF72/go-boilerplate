package repositories

import (
	"github.com/MAAF72/go-boilerplate/adapters/grule"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
)

// GruleRepository grule repository
type GruleRepository struct {
	grule *grule.Grule
}

// GruleRepositoryImpl grule repository implementations
type GruleRepositoryImpl interface {
	NewGruleKnowledgeBaseInstance(ruleSlug string, ruleVersion string) *ast.KnowledgeBase
	NewGruleEngine() *engine.GruleEngine
	NewGruleDataContext() ast.IDataContext
}

// NewGruleKnowledgeBaseInstance create new grule knowledge base instance
func (repo GruleRepository) NewGruleKnowledgeBaseInstance(ruleSlug string, ruleVersion string) *ast.KnowledgeBase {
	return repo.grule.NewKnowledgeBaseInstance(ruleSlug, ruleVersion)
}

// NewGruleEngine create new grule engine
func (repo GruleRepository) NewGruleEngine() *engine.GruleEngine {
	return engine.NewGruleEngine()
}

// NewGruleDataContext create new grule data context
func (repo GruleRepository) NewGruleDataContext() ast.IDataContext {
	return ast.NewDataContext()
}
