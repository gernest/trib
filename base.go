package trib

import (
	"time"

	"golang.org/x/net/context"
)

type BasePlan struct {
	id     string
	stages []Stage
	BaseTime
}

func NewBasePlan(id string) *BasePlan {
	return &BasePlan{id: id}
}

func (b *BasePlan) AddStage(s Stage) {
	b.stages = append(b.stages, s)
}

func (b *BasePlan) ID() string {
	return b.id
}

func (b *BasePlan) Stages() ([]Stage, bool) {
	if len(b.stages) > 0 {
		return b.stages, true
	}
	return nil, false
}

type BaseTime struct {
	isSet    bool
	duration time.Duration
}

func (bt BaseTime) Duration() (time.Duration, bool) {
	if bt.isSet {
		return bt.duration, true
	}
	return bt.duration, false
}
func (bt BaseTime) SetDuration(d time.Duration) {
	ptr := &bt
	ptr.duration = d
	if !ptr.isSet {
		ptr.isSet = true
	}
}

type BaseStage struct {
	id    string
	steps []Step
	BaseTime
}

func NewBaseStage(id string) *BaseStage {
	return &BaseStage{id: id}
}

func (b *BaseStage) AddStep(s Step) {
	b.steps = append(b.steps, s)
}

func (b *BaseStage) ID() string {
	return b.id
}

func (b *BaseStage) Steps() ([]Step, bool) {
	if len(b.steps) > 0 {
		return b.steps, true
	}
	return nil, false
}

type ExecFunc func(context.Context) error

type BaseStep struct {
	id string
	BaseTime
	e          ExecFunc
	isParallel bool
}

func NewBaseStep(id string, isParallel bool, exec ExecFunc) *BaseStep {
	return &BaseStep{
		id:         id,
		isParallel: isParallel,
		e:          exec,
	}
}

func (b *BaseStep) ID() string {
	return b.id
}

func (b *BaseStep) Parallel() bool {
	return b.isParallel
}

func (b *BaseStep) Exec(ctx context.Context) error {
	return b.e(ctx)
}
