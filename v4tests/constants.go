package v4tests

const spacexLiveAPIBase = "https://api.spacexdata.com/v4"

// Capsules dummy data
const (
	capsulesAll = `[
		{
			"reuse_count":0,
			"water_landings":1,
			"land_landings":0,
			"last_update":"Hanging in atrium at SpaceX HQ in Hawthorne ",
			"launches":["5eb87cdeffd86e000604b330"],
			"serial":"C101",
			"status":"retired",
			"type":"Dragon 1.0",
			"id":"5e9e2c5bf35918ed873b2664"
		},
		{
			"reuse_count":0,
			"water_landings":1,
			"land_landings":0,
			"last_update":"On display at KSC Visitor's Center ",
			"launches":["5eb87cdfffd86e000604b331"],
			"serial":"C102",
			"status":"retired",
			"type":"Dragon 1.0",
			"id":"5e9e2c5bf3591882af3b2665"
		}
	]`

	capsulesOne = `{
		"reuse_count":0,
		"water_landings":1,
		"land_landings":0,
		"last_update":"Hanging in atrium at SpaceX HQ in Hawthorne ",
		"launches":["5eb87cdeffd86e000604b330"],
		"serial":"C101",
		"status":"retired",
		"type":"Dragon 1.0",
		"id":"5e9e2c5bf35918ed873b2664"
	}`
)

// Core dummy data
const (
	coresAll = `[
		{
			"block":null,
			"reuse_count":0,
			"rtls_attempts":0,
			"rtls_landings":0,
			"asds_attempts":0,
			"asds_landings":0,
			"last_update":"Engine failure at T+33 seconds resulted in loss of vehicle",
			"launches":["5eb87cd9ffd86e000604b32a"],
			"serial":"Merlin1A",
			"status":"lost",
			"id":"5e9e289df35918033d3b2623"
		},
		{
			"block":null,
			"reuse_count":0,
			"rtls_attempts":0,
			"rtls_landings":0,
			"asds_attempts":0,
			"asds_landings":0,
			"last_update":"Successful first-stage burn and transition to second stage, maximal altitude 289 km. Harmonic oscillation at T+5 minutes Premature engine shutdown at T+7 min 30 s. Failed to reach orbit.",
			"launches":["5eb87cdaffd86e000604b32b"],
			"serial":"Merlin2A",
			"status":"lost",
			"id":"5e9e289ef35918416a3b2624"
		}
	]`

	coresOne = `{
		"block":null,
		"reuse_count":0,
		"rtls_attempts":0,
		"rtls_landings":0,
		"asds_attempts":0,
		"asds_landings":0,
		"last_update":"Engine failure at T+33 seconds resulted in loss of vehicle",
		"launches":["5eb87cd9ffd86e000604b32a"],
		"serial":"Merlin1A",
		"status":"lost",
		"id":"5e9e289df35918033d3b2623"
	}`
)

// Crew dummy data
const (
	crewAll = `[
		{
			"name":"Robert Behnken",
			"agency":"NASA",
			"image":"https://imgur.com/0smMgMH.png",
			"wikipedia":"https://en.wikipedia.org/wiki/Robert_L._Behnken",
			"launches":["5eb87d46ffd86e000604b388"],
			"status":"active",
			"id":"5ebf1a6e23a9a60006e03a7a"
		},
		{
			"name":"Douglas Hurley",
			"agency":"NASA",
			"image":"https://i.imgur.com/ooaayWf.png",
			"wikipedia":"https://en.wikipedia.org/wiki/Douglas_G._Hurley",
			"launches":["5eb87d46ffd86e000604b388"],
			"status":"active",
			"id":"5ebf1b7323a9a60006e03a7b"
		}
	]`

	crewOne = `{
		"name":"Robert Behnken",
		"agency":"NASA",
		"image":"https://imgur.com/0smMgMH.png",
		"wikipedia":"https://en.wikipedia.org/wiki/Robert_L._Behnken",
		"launches":["5eb87d46ffd86e000604b388"],
		"status":"active",
		"id":"5ebf1a6e23a9a60006e03a7a"
	}`
)
