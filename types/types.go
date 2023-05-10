package types

import (
	"context"
	"reflect"
)

type MaskerName string

// Masker 数据偏移处理方法的抽象
type Masker interface {
	Name() MaskerName
	Mask(context.Context, any) (any, error)
	Args() []reflect.Kind
	Description() string
}

type MaskerInfo struct {
	Name        string
	Args        []reflect.Kind
	Description string
}

// MaskerBuilder Masker的构造函数
type MaskerBuilder func(...any) Masker

// MaskerFactory Masker的构造工厂
type MaskerFactory interface {
	Register(MaskerName, MaskerBuilder) error
	Build(BuilderParams) (Masker, error)
	ListBuilders() map[MaskerName]MaskerBuilder
}

// BuilderParams 数据偏移规则的描述
type BuilderParams struct {
	Name MaskerName
	Args []any
}
