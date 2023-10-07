package v4tests

import "errors"

const spacexLiveAPIBase = "https://api.spacexdata.com"

const (
	MsgTestingRealAPI = "Testing against the real SpaceX API."

	MsgNotImplemented = "Not implemented"
)

var (
	ErrCollectionLenMismatch = errors.New("collection length mismatch")
	ErrIDMismatch            = errors.New("id mismatch")
)

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

// Dragons dummy data
const (
	dragonsAll = `[
    	{
    	    "heat_shield": {
    	        "material": "PICA-X",
    	        "size_meters": 3.6,
    	        "temp_degrees": 3000,
    	        "dev_partner": "NASA"
    	    },
    	    "launch_payload_mass": {
    	        "kg": 6000,
    	        "lb": 13228
    	    },
    	    "launch_payload_vol": {
    	        "cubic_meters": 25,
    	        "cubic_feet": 883
    	    },
    	    "return_payload_mass": {
    	        "kg": 3000,
    	        "lb": 6614
    	    },
    	    "return_payload_vol": {
    	        "cubic_meters": 11,
    	        "cubic_feet": 388
    	    },
    	    "pressurized_capsule": {
    	        "payload_volume": {
    	            "cubic_meters": 11,
    	            "cubic_feet": 388
    	        }
    	    },
    	    "trunk": {
    	        "trunk_volume": {
    	            "cubic_meters": 14,
    	            "cubic_feet": 494
    	        },
    	        "cargo": {
    	            "solar_array": 2,
    	            "unpressurized_cargo": true
    	        }
    	    },
    	    "height_w_trunk": {
    	        "meters": 7.2,
    	        "feet": 23.6
    	    },
    	    "diameter": {
    	        "meters": 3.7,
    	        "feet": 12
    	    },
    	    "first_flight": "2010-12-08",
    	    "flickr_images": [
    	        "https://i.imgur.com/9fWdwNv.jpg",
    	        "https://live.staticflickr.com/8578/16655995541_078768dea2_b.jpg",
    	        "https://farm3.staticflickr.com/2815/32761844973_4b55b27d3c_b.jpg",
    	        "https://farm9.staticflickr.com/8618/16649075267_d18cbb4342_b.jpg"
    	    ],
    	    "name": "Dragon 1",
    	    "type": "capsule",
    	    "active": true,
    	    "crew_capacity": 0,
    	    "sidewall_angle_deg": 15,
    	    "orbit_duration_yr": 2,
    	    "dry_mass_kg": 4200,
    	    "dry_mass_lb": 9300,
    	    "thrusters": [
    	        {
    	            "type": "Draco",
    	            "amount": 18,
    	            "pods": 4,
    	            "fuel_1": "nitrogen tetroxide",
    	            "fuel_2": "monomethylhydrazine",
    	            "isp": 300,
    	            "thrust": {
    	                "kN": 0.4,
    	                "lbf": 90
    	            }
    	        }
    	    ],
    	    "wikipedia": "https://en.wikipedia.org/wiki/SpaceX_Dragon",
    	    "description": "Dragon is a reusable spacecraft developed by SpaceX, an American private space transportation company based in Hawthorne, California. Dragon is launched into space by the SpaceX Falcon 9 two-stage-to-orbit launch vehicle. The Dragon spacecraft was originally designed for human travel, but so far has only been used to deliver cargo to the International Space Station (ISS).",
    	    "id": "5e9d058759b1ff74a7ad5f8f"
    	},
    	{
    	    "heat_shield": {
    	        "material": "PICA-X",
    	        "size_meters": 3.6,
    	        "temp_degrees": 3000,
    	        "dev_partner": "NASA"
    	    },
    	    "launch_payload_mass": {
    	        "kg": 6000,
    	        "lb": 13228
    	    },
    	    "launch_payload_vol": {
    	        "cubic_meters": 25,
    	        "cubic_feet": 883
    	    },
    	    "return_payload_mass": {
    	        "kg": 3000,
    	        "lb": 6614
    	    },
    	    "return_payload_vol": {
    	        "cubic_meters": 11,
    	        "cubic_feet": 388
    	    },
    	    "pressurized_capsule": {
    	        "payload_volume": {
    	            "cubic_meters": 11,
    	            "cubic_feet": 388
    	        }
    	    },
    	    "trunk": {
    	        "trunk_volume": {
    	            "cubic_meters": 14,
    	            "cubic_feet": 494
    	        },
    	        "cargo": {
    	            "solar_array": 2,
    	            "unpressurized_cargo": true
    	        }
    	    },
    	    "height_w_trunk": {
    	        "meters": 7.2,
    	        "feet": 23.6
    	    },
    	    "diameter": {
    	        "meters": 3.7,
    	        "feet": 12
    	    },
    	    "first_flight": "2019-03-02",
    	    "flickr_images": [
    	        "https://farm8.staticflickr.com/7647/16581815487_6d56cb32e1_b.jpg",
    	        "https://farm1.staticflickr.com/780/21119686299_c88f63e350_b.jpg",
    	        "https://farm9.staticflickr.com/8588/16661791299_a236e2f5dc_b.jpg"
    	    ],
    	    "name": "Dragon 2",
    	    "type": "capsule",
    	    "active": true,
    	    "crew_capacity": 7,
    	    "sidewall_angle_deg": 15,
    	    "orbit_duration_yr": 2,
    	    "dry_mass_kg": 6350,
    	    "dry_mass_lb": 14000,
    	    "thrusters": [
    	        {
    	            "type": "Draco",
    	            "amount": 18,
    	            "pods": 4,
    	            "fuel_1": "nitrogen tetroxide",
    	            "fuel_2": "monomethylhydrazine",
    	            "isp": 300,
    	            "thrust": {
    	                "kN": 0.4,
    	                "lbf": 90
    	            }
    	        },
    	        {
    	            "type": "SuperDraco",
    	            "amount": 8,
    	            "pods": 4,
    	            "fuel_1": "dinitrogen tetroxide",
    	            "fuel_2": "monomethylhydrazine",
    	            "isp": 235,
    	            "thrust": {
    	                "kN": 71,
    	                "lbf": 16000
    	            }
    	        }
    	    ],
    	    "wikipedia": "https://en.wikipedia.org/wiki/Dragon_2",
    	    "description": "Dragon 2 (also Crew Dragon, Dragon V2, or formerly DragonRider) is the second version of the SpaceX Dragon spacecraft, which will be a human-rated vehicle. It includes a set of four side-mounted thruster pods with two SuperDraco engines each, which can serve as a launch escape system or launch abort system (LAS). In addition, it has much larger windows, new flight computers and avionics, and redesigned solar arrays, and a modified outer mold line from the initial cargo Dragon that has been flying for several years.",
    	    "id": "5e9d058859b1ffd8e2ad5f90"
    	}
	]`

	dragonsOne = `{
			"heat_shield": {
			"material": "PICA-X",
			"size_meters": 3.6,
			"temp_degrees": 3000,
			"dev_partner": "NASA"
		},
		"launch_payload_mass": {
			"kg": 6000,
			"lb": 13228
		},
		"launch_payload_vol": {
			"cubic_meters": 25,
			"cubic_feet": 883
		},
		"return_payload_mass": {
			"kg": 3000,
			"lb": 6614
		},
		"return_payload_vol": {
			"cubic_meters": 11,
			"cubic_feet": 388
		},
		"pressurized_capsule": {
			"payload_volume": {
					"cubic_meters": 11,
					"cubic_feet": 388
				}
			},
		"trunk": {
			"trunk_volume": {
				"cubic_meters": 14,
				"cubic_feet": 494
			},
			"cargo": {
				"solar_array": 2,
				"unpressurized_cargo": true
			}
		},
		"height_w_trunk": {
			"meters": 7.2,
			"feet": 23.6
		},
		"diameter": {
			"meters": 3.7,
			"feet": 12
		},
		"first_flight": "2010-12-08",
		"flickr_images": [
			"https://i.imgur.com/9fWdwNv.jpg",
			"https://live.staticflickr.com/8578/16655995541_078768dea2_b.jpg",
			"https://farm3.staticflickr.com/2815/32761844973_4b55b27d3c_b.jpg",
			"https://farm9.staticflickr.com/8618/16649075267_d18cbb4342_b.jpg"
		],
		"name": "Dragon 1",
		"type": "capsule",
		"active": true,
		"crew_capacity": 0,
		"sidewall_angle_deg": 15,
		"orbit_duration_yr": 2,
		"dry_mass_kg": 4200,
		"dry_mass_lb": 9300,
		"thrusters": [
			{
				"type": "Draco",
				"amount": 18,
				"pods": 4,
				"fuel_1": "nitrogen tetroxide",
				"fuel_2": "monomethylhydrazine",
				"isp": 300,
				"thrust": {
					"kN": 0.4,
					"lbf": 90
				}
			}
		],
		"wikipedia": "https://en.wikipedia.org/wiki/SpaceX_Dragon",
		"description": "Dragon is a reusable spacecraft developed by SpaceX, an American private space transportation company based in Hawthorne, California. Dragon is launched into space by the SpaceX Falcon 9 two-stage-to-orbit launch vehicle. The Dragon spacecraft was originally designed for human travel, but so far has only been used to deliver cargo to the International Space Station (ISS).",
		"id": "5e9d058759b1ff74a7ad5f8f"
	}`
)

const (
	historyAll = `[
		{
			"links": {
				"article": "http://www.spacex.com/news/2013/02/11/flight-4-launch-update-0"
			},
			"title": "Falcon reaches Earth orbit",
			"event_date_utc": "2008-09-28T23:15:00Z",
			"event_date_unix": 1222643700,
			"details": "Falcon 1 becomes the first privately developed liquid-fuel rocket to reach Earth orbit.",
			"id": "5f6fb2cfdcfdf403df37971e"
		},
		{
			"links": {
				"article": "http://www.spacex.com/news/2013/02/12/falcon-1-flight-5"
			},
			"title": "Falcon delivers payload to orbit",
			"event_date_utc": "2009-07-13T03:35:00Z",
			"event_date_unix": 1247456100,
			"details": "Fifth successful flight of Falcon makes history, becoming the first privately developed liquid-fuel rocket to deliver a commercial satellite to orbit.",
			"id": "5f6fb2efdcfdf403df37971f"
		}
	]`

	historyOne = `{
		"links": {
			"article": "http://www.spacex.com/news/2013/02/12/falcon-1-flight-5"
		},
		"title": "Falcon delivers payload to orbit",
		"event_date_utc": "2009-07-13T03:35:00Z",
		"event_date_unix": 1247456100,
		"details": "Fifth successful flight of Falcon makes history, becoming the first privately developed liquid-fuel rocket to deliver a commercial satellite to orbit.",
		"id": "5f6fb2efdcfdf403df37971f"
	}`
)

const (
	landpadsAll = `[
		{
			"images": {
				"large": [
					"https://i.imgur.com/KHBk6jO.png"
				]
			},
			"name": "LZ-1",
			"full_name": "Landing Zone 1",
			"status": "active",
			"type": "RTLS",
			"locality": "Cape Canaveral",
			"region": "Florida",
			"latitude": 28.485833,
			"longitude": -80.544444,
			"landing_attempts": 21,
			"landing_successes": 20,
			"wikipedia": "https://en.wikipedia.org/wiki/Landing_Zones_1_and_2",
			"details": "SpaceX's first east coast landing pad is Landing Zone 1, where the historic first Falcon 9 landing occurred in December 2015. LC-13 was originally used as a launch pad for early Atlas missiles and rockets from Lockheed Martin. LC-1 was later expanded to include Landing Zone 2 for side booster RTLS Falcon Heavy missions, and it was first used in February 2018 for that purpose.",
			"launches": [
				"5eb87cefffd86e000604b342",
				"5eb87cf9ffd86e000604b349",
				"5eb87cfeffd86e000604b34d",
				"5eb87d01ffd86e000604b350",
				"5eb87d03ffd86e000604b352",
				"5eb87d07ffd86e000604b356",
				"5eb87d09ffd86e000604b358",
				"5eb87d0effd86e000604b35c",
				"5eb87d10ffd86e000604b35e",
				"5eb87d13ffd86e000604b360",
				"5eb87d26ffd86e000604b371",
				"5eb87d2dffd86e000604b376",
				"5eb87d35ffd86e000604b37a",
				"5eb87d36ffd86e000604b37b",
				"5eb87d42ffd86e000604b384",
				"5eb87d47ffd86e000604b38a",
				"5f8399fb818d8b59f5740d43",
				"600f9b6d8f798e2a4d5f979f",
				"61bf3e31cd5ab50b0d936345",
				"6161d32d6db1a92bfba85359",
				"6243ae24af52800c6e919258"
			],
			"id": "5e9e3032383ecb267a34e7c7"
		},
		{
			"images": {
				"large": [
					"https://i.imgur.com/YE8PMYX.png"
				]
			},
			"name": "LZ-2",
			"full_name": "Landing Zone 2",
			"status": "active",
			"type": "RTLS",
			"locality": "Cape Canaveral",
			"region": "Florida",
			"latitude": 28.485833,
			"longitude": -80.544444,
			"landing_attempts": 3,
			"landing_successes": 3,
			"wikipedia": "https://en.wikipedia.org/wiki/Landing_Zones_1_and_2",
			"details": "SpaceX's first east coast landing pad is Landing Zone 1, where the historic first Falcon 9 landing occurred in December 2015. LC-13 was originally used as a launch pad for early Atlas missiles and rockets from Lockheed Martin. LC-1 was later expanded to include Landing Zone 2 for side booster RTLS Falcon Heavy missions, and it was first used in February 2018 for that purpose.",
			"launches": [
				"5eb87d13ffd86e000604b360",
				"5eb87d2dffd86e000604b376",
				"5eb87d35ffd86e000604b37a"
			],
			"id": "5e9e3032383ecb90a834e7c8"
		}
	]`

	landpadsOne = `{
		"images": {
			"large": [
				"https://i.imgur.com/YE8PMYX.png"
			]
		},
		"name": "LZ-2",
		"full_name": "Landing Zone 2",
		"status": "active",
		"type": "RTLS",
		"locality": "Cape Canaveral",
		"region": "Florida",
		"latitude": 28.485833,
		"longitude": -80.544444,
		"landing_attempts": 3,
		"landing_successes": 3,
		"wikipedia": "https://en.wikipedia.org/wiki/Landing_Zones_1_and_2",
		"details": "SpaceX's first east coast landing pad is Landing Zone 1, where the historic first Falcon 9 landing occurred in December 2015. LC-13 was originally used as a launch pad for early Atlas missiles and rockets from Lockheed Martin. LC-1 was later expanded to include Landing Zone 2 for side booster RTLS Falcon Heavy missions, and it was first used in February 2018 for that purpose.",
		"launches": [
			"5eb87d13ffd86e000604b360",
			"5eb87d2dffd86e000604b376",
			"5eb87d35ffd86e000604b37a"
		],
		"id": "5e9e3032383ecb90a834e7c8"
	}`
)
