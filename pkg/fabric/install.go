package fabric

import (
	"context"
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
	"gopkg.in/yaml.v2"
)

func (m *Mixin) loadAction_install(ctx context.Context) (*Action, error) {
	var action Action
	err := builder.LoadAction(ctx, m.RuntimeConfig, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &action)
		return &action, err
	})
	return &action, err
}

func (m *Mixin) Install(ctx context.Context) error {
	fmt.Println(m.Out, "Started deployment of fabric artifacts")
	action, err := m.loadAction_install(ctx)
	if err != nil {
		return err
	}
	fmt.Println(action)
	return nil
}

// 	filepath := "/app/ubuntu.16.04-x64"
// 	if _, err := os.Stat(filepath + "/Microsoft.Fabric.Provisioning.Client"); os.IsNotExist(err) {
// 		fmt.Println("File does not exist")
// 		return err
// 	}

// var cmd *exec.Cmd
// print(action)
// 	if action.Steps[0].Type == "Workspace" || action.Steps[0].Type == "" {
// 		fmt.Println("Creating Fabric Workspace")
// 		payload := make(map[string]interface{})
// 		payload["displayName"] = action.Steps[0].DisplayName
// 		payload["description"] = action.Steps[0].Description
// 		payload["capacityId"] = action.Steps[0].CapacityId

// 		jsonString, err := json.Marshal(payload)

// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		cmd = exec.Command(filepath+"/Microsoft.Fabric.Provisioning.Client", "create", "--token", action.Steps[0].Token, "--payload", string(jsonString))
// 		fmt.Println()
// 		fmt.Println()
// 		fmt.Println(cmd)
// 		fmt.Println()
// 		fmt.Println()
// 	} else {
// 		fmt.Println("Creating fabric item")
// 		fmt.Println(action.Steps[0].DisplayName, action.Steps[0].Type, action.Steps[0].Description, action.Steps[0].WorkspaceId)
// 		payload := make(map[string]interface{})
// 		payload["displayName"] = action.Steps[0].DisplayName
// 		payload["type"] = action.Steps[0].Type
// 		payload["description"] = action.Steps[0].Description
// 		payload["workspaceId"] = action.Steps[0].WorkspaceId
// 		jsonString, err := json.Marshal(payload)

// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		cmd = exec.Command(filepath+"/Microsoft.Fabric.Provisioning.Client", "createItem", "--token", action.Steps[0].Token, "--payload", string(jsonString))
// 	}
// 	fmt.Println("Command Executed:")
// 	fmt.Println(cmd)
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return err
// 	}
// 	fmt.Println("Workspace Output")
// 	fmt.Println(string(output))

// 	response := strings.SplitAfter(string(output), "Microsoft.Fabric.Provisioning.Client[0]")

// 	if len(response) >= 3 {
// 		strResp := response[2]
// 		subResp1 := strings.SplitAfter(strResp, "\"id\":\"")
// 		if len(subResp1) >= 2 {
// 			subStr1 := subResp1[1]
// 			str2 := strings.Split(subStr1, "\",")[0]
// 			fmt.Println("Output:", str2)

// 			fmt.Println(action.Steps[0].Outputs)
// 			for _, output := range action.Steps[0].Outputs {
// 				fmt.Println(output)
// 				// ToUpper the key because of the case weirdness with ARM outputs
// 				if "ID" == strings.ToUpper(output.Key) {
// 					err := m.WriteMixinOutputToFile(output.Name, []byte(fmt.Sprintf("%v", str2)))
// 					if err != nil {
// 						return errors.Wrapf(err, "unable to write output '%s'", output.Name)
// 					} else {
// 						fmt.Println(output.Name, []byte(fmt.Sprintf("%v", string(str2))))
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return err
// }
