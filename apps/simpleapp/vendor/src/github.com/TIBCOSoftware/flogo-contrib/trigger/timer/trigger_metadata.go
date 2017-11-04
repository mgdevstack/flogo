package timer

import (
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

var jsonMetadata = `{
  "name": "tibco-timer",
  "type": "flogo:trigger",
  "ref": "github.com/TIBCOSoftware/flogo-contrib/trigger/timer",
  "version": "0.0.1",
  "title": "Timer",
  "description": "Simple Timer trigger",
  "homepage": "https://github.com/TIBCOSoftware/flogo-contrib/tree/master/trigger/timer",
  "settings": [
  ],
  "outputs": [
    {
      "name": "params",
      "type": "params"
    },
    {
      "name": "content",
      "type": "object"
    }
  ],
  "handler": {
    "settings": [
      {
        "name": "repeating",
        "type": "string",
        "required" : true
      },
      {
        "name": "notImmediate",
        "type": "string",
        "required" : false
      },
      {
        "name": "startDate",
        "type": "string",
        "required" : false
      },
      {
        "name": "hours",
        "type": "string",
        "required" : false
      },
      {
        "name": "minutes",
        "type": "string",
        "required" : false
      },
      {
        "name": "seconds",
        "type": "string",
        "required" : false
      }
    ]
  }
}
`

// init create & register trigger factory
func init() {
	md := trigger.NewMetadata(jsonMetadata)
	trigger.RegisterFactory(md.ID, NewFactory(md))
}
