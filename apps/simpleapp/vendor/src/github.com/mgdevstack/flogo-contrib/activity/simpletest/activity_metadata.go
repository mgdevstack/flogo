package simpletest

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "simpletest",
  "version": "0.0.1",
  "type": "flogo:activity",
  "description": "Simple Activity to test",
  "author": "Mayank Gaikwad <mg.devzone@gmail.com>",
  "inputs":[
    {
      "name": "input",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "output",
      "type": "string"
    }
  ]
}`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
