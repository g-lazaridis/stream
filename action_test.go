package stream

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/app/resource"
	"github.com/project-flogo/core/engine/channels"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
	"gitlab.f-in.io/project-flogo/stream/pipeline"
)

const testConfig string = `{
  "id": "flogo-stream",
  "ref": "gitlab.f-in.io/project-flogo/stream",
  "settings": {
    "streamURI": "res://stream:test",
    "outputChannel": "testChan"
  }
}
`
const resData string = `{
        "metadata": {
          "input": [
            {
              "name": "input",
              "type": "integer"
            }
          ]
        },
        "stages": [
        ]
      }`

func TestActionFactory_New(t *testing.T) {

	cfg := &action.Config{}
	err := json.Unmarshal([]byte(testConfig), cfg)

	if err != nil {
		t.Error(err)
		return
	}

	af := ActionFactory{}
	ctx := test.NewActionInitCtx()

	err = af.Initialize(ctx)
	assert.Nil(t, err)

	resourceCfg := &resource.Config{ID: "stream:test"}
	resourceCfg.Data = []byte(resData)
	err = ctx.AddResource(pipeline.ResType, resourceCfg)
	assert.Nil(t, err)

	_, err = channels.New("testChan", 5)
	assert.Nil(t, err)

	act, err := af.New(cfg)
	assert.Nil(t, err)
	assert.NotNil(t, act)
}
