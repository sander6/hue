package rule

import (
	"testing"
)

var singleRuleResponse = []byte(`{
	"name": "Wall Switch Rule",
	"owner": "ruleOwner",
	"created": "2014-07-23T15:02:56",
	"lasttriggered": "none",
	"timestriggered": 0,
	"status": "enabled",
	"conditions": [
	{
		"address": "/sensors/2/state/buttonevent",
		"operator": "eq",
		"value": "16"
	},
	{
		"address": "/sensors/2/state/lastupdated",
		"operator": "dx"
	}
	],
	"actions": [
	{
		"address": "/groups/0/action",
		"method": "PUT",
		"body": {
			"scene": "S3"
		}
	}
	]
}`)

var allRulesResponse = []byte(`{
	"1": {
		"name": "Wall Switch Rule",
		"lasttriggered": "2013-10-17T01:23:20",
		"creationtime": "2013-10-10T21:11:45",
		"timestriggered": 27,
		"owner": "78H56B12BA",
		"status": "enabled",
		"conditions": [
		{
			"address": "/sensors/2/state/buttonevent",
			"operator": "eq",
			"value": "16"
		},
		{
			"address": "/sensors/2/state/lastupdated",
			"operator": "dx"
		}
		],
		"actions": [
		{
			"address": "/groups/0/action",
			"method": "PUT",
			"body": {
				"scene": "S3"
			}
		}
		]
	} 
}`)
