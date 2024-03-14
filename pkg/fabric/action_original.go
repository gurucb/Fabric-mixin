package fabric

// import (
// 	"get.porter.sh/porter/pkg/exec/builder"
// )

// var _ builder.ExecutableAction = Action{}
// var _ builder.BuildableAction = Action{}

// type Action struct {
// 	Name  string
// 	Steps []Step // using UnmarshalYAML so that we don't need a custom type per action
// }

// // MarshalYAML converts the action back to a YAML representation
// // install:
// //
// //	fabric:
// //	  ...
// func (a Action) MarshalYAML() (interface{}, error) {
// 	return map[string]interface{}{a.Name: a.Steps}, nil
// }

// // MakeSteps builds a slice of Step for data to be unmarshaled into.
// func (a Action) MakeSteps() interface{} {
// 	return &[]Step{}
// }

// // UnmarshalYAML takes any yaml in this form
// // ACTION:
// // - fabric: ...
// // and puts the steps into the Action.Steps field
// func (a *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	results, err := builder.UnmarshalAction(unmarshal, a)
// 	if err != nil {
// 		return err
// 	}

// 	for actionName, action := range results {
// 		a.Name = actionName
// 		for _, result := range action {
// 			step := result.(*[]Step)
// 			a.Steps = append(a.Steps, *step...)
// 		}
// 		break // There is only 1 action
// 	}
// 	return nil
// }

// func (a Action) GetSteps() []builder.ExecutableStep {
// 	// Go doesn't have generics, nothing to see here...
// 	steps := make([]builder.ExecutableStep, len(a.Steps))
// 	for i := range a.Steps {
// 		steps[i] = a.Steps[i]
// 	}

// 	return steps
// }

// type Step struct {
// 	Instruction `yaml:"fabric"`
// }

// // Actions is a set of actions, and the steps, passed from Porter.
// type Actions []Action

// // UnmarshalYAML takes chunks of a porter.yaml file associated with this mixin
// // and populates it on the current action set.
// // install:
// //
// //	fabric:
// //	  ...
// //	fabric:
// //	  ...
// //
// // upgrade:
// //
// //	fabric:
// //	  ...
// func (a *Actions) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	results, err := builder.UnmarshalAction(unmarshal, Action{})
// 	if err != nil {
// 		return err
// 	}

// 	for actionName, action := range results {
// 		for _, result := range action {
// 			s := result.(*[]Step)
// 			*a = append(*a, Action{
// 				Name:  actionName,
// 				Steps: *s,
// 			})
// 		}
// 	}
// 	return nil
// }

// var _ builder.HasOrderedArguments = Instruction{}
// var _ builder.ExecutableStep = Instruction{}
// var _ builder.StepWithOutputs = Instruction{}

// type Instruction struct {
// 	Outputs        []Output `yaml:"outputs,omitempty"`
// 	SuppressOutput bool     `yaml:"suppress-output,omitempty"`
// 	// https://release-v1.porter.sh/mixins/exec/#ignore-error
// 	builder.IgnoreErrorHandler `yaml:"ignoreError,omitempty"`
// 	FabricFields               `yaml:",inline"`
// }
// type FabricFields struct {
// 	Token       string `yaml:"token"`
// 	DisplayName string `yaml:"displayName,omitempty"`
// 	Type        string `yaml:"type,omitempty"`
// 	Description string `yaml:"description,omitempty"`
// 	CapacityId  string `yaml:"capacityId,omitempty"`
// 	WorkspaceId string `yaml:"workspaceid,omitempty"`
// 	ItemId      string `yaml:"itemid,omitempty"`
// 	Definition  string `yaml:"definition,omitempty"`
// }

// func (s Instruction) GetCommand() string {
// 	return "fabric"
// }

// func (s Instruction) GetWorkingDir() string {
// 	return ""
// }

// func (s Instruction) GetArguments() []string {
// 	return nil
// }

// func (s Instruction) GetSuffixArguments() []string {
// 	return nil
// }

// func (s Instruction) GetFlags() builder.Flags {
// 	return nil
// }

// func (s Instruction) SuppressesOutput() bool {
// 	return s.SuppressOutput
// }

// func (s Instruction) GetOutputs() []builder.Output {
// 	// Go doesn't have generics, nothing to see here...
// 	outputs := make([]builder.Output, len(s.Outputs))
// 	for i := range s.Outputs {
// 		outputs[i] = s.Outputs[i]
// 	}
// 	return outputs
// }

// var _ builder.OutputJsonPath = Output{}
// var _ builder.OutputFile = Output{}
// var _ builder.OutputRegex = Output{}

// type Output struct {
// 	Name string `yaml:"name"`
// 	Key  string `yaml:"key"`

// 	// See https://porter.sh/mixins/exec/#outputs
// 	// TODO: If your mixin doesn't support these output types, you can remove these and the interface assertions above, and from #/definitions/outputs in schema.json
// 	// JsonPath string `yaml:"jsonPath,omitempty"`
// 	// FilePath string `yaml:"path,omitempty"`
// 	// Regex    string `yaml:"regex,omitempty"`

// }

// func (o Output) GetName() string {
// 	return o.Name
// }

// func (o Output) GetJsonPath() string {
// 	return ""
// }

// func (o Output) GetFilePath() string {
// 	return ""
// }

// func (o Output) GetRegex() string {
// 	return ""
// }
