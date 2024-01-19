package fabric

import (
	"context"

	"get.porter.sh/porter/pkg/exec/builder"
)

// func (m *Mixin) loadAction(ctx context.Context) (*Action, error) {
// 	var action Action
// 	err := builder.LoadAction(ctx, m.RuntimeConfig, "", func(contents []byte) (interface{}, error) {
// 		err := yaml.Unmarshal(contents, &action)
// 		return &action, err
// 	})
// 	return &action, err
// }

func (m *Mixin) Upgrade(ctx context.Context) error {
	action, err := m.loadAction(ctx)
	if err != nil {
		return err
	}

	_, err = builder.ExecuteSingleStepAction(ctx, m.RuntimeConfig, action)
	return err
}
