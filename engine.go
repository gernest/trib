package trib

import (
	"errors"
	"sync"

	"golang.org/x/net/context"
)

type Engine struct {
}

func (e Engine) ExecPlan(ctx context.Context, p Plan) error {
	stags, ok := p.Stages()
	if !ok {
		return errors.New("empty plan")
	}
	for _, v := range stags {
		err := e.ExecStage(ctx, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e Engine) ExecStage(ctx context.Context, s Stage) error {
	var (
		parallel    []int
		normal      []int
		wg          sync.WaitGroup
		hasParallel bool
	)

	steps, ok := s.Steps()
	if !ok {
		return errors.New("empty stage")
	}
	for i := 0; i < len(steps); i++ {
		if steps[i].Parallel() {
			parallel = append(parallel, i)
			continue
		}
		normal = append(normal, i)
	}

	if len(parallel) > 0 {
		hasParallel = true
	}
	if hasParallel {
		for _, v := range parallel {
			go e.StepParallel(ctx, steps[v], &wg)
		}
	}
	for _, v := range normal {
		err := steps[v].Exec(ctx)
		if err != nil {
			return err
		}
	}
	if hasParallel {
		wg.Wait()
	}
	return nil
}

func (e Engine) StepParallel(ctx context.Context, s Step, wait *sync.WaitGroup) {
	wait.Add(1)
	defer wait.Done()
	s.Exec(ctx)
}
