package units

import "time"

type Unit struct {
	unitID           int
	name             string
	description      string
	price            float32
	timeOfRecruiting string
	timeOfDisbanding string
}

var unitsList = []*Unit{
	&Unit{
		unitID:        10001,
		name:             "Stark bowmen",
		description:      "Band of 14 Stark bowmen, from a variety of Nothern banners of House Stark. Good mobility, ranged strength.",
		price:            "200 Talents",
		timeOfRecruiting:   time.Now().UTC().String(),
		timeOfDisbanding: time.Now().UTC().String(),
	},
	&Unit{
		unitID:        10002,
		name:             "Greyjoy sailors",
		description:      "Crew of 32 Greyjoy sailors, bloodthirsty and ruthless, plunder units.",
		price:            "170",
		timeOfRecruiting:   time.Now().UTC().String(),
		timeOfDisbanding: time.Now().UTC().String(),
	},
	&Unit{
		unitID:        10003,
		name:             "Bolton Riders",
		description:      "Half a dozen Bolton hunters, armed with spears, on horseback, wreak havoc in the field.",
		price:            "400",
		timeOfRecruiting:   time.Now().UTC().String(),
		timeOfDisbanding: time.Now().UTC().String(),
	},
	&Unit{
		unitID:        10004,
		name:             "Tully Militiamen",
		description:      "Company of 40 Tully village militamen, armed with an assortment of weapons. stout defenders of their homes.",
		price:            "150",
		timeOfRecruiting:   time.Now().UTC().String(),
		timeOfDisbanding: time.Now().UTC().String(),
	},
	&Unit{
		unitID:        10006,
		name:             "Gold Cloaks",
		description:      "12 men of King's Landing's city guards. Off duty hunters",
		price:            "240",
		timeOfRecruiting:   time.Now().UTC().String(),
		timeOfDisbanding: time.Now().UTC().String(),
	},
}

func getUnits() []*Unit{
	return unitsList
}

func getUnitbyId(id int) *Unit{
	for _, x:= range unitsList{
		if x.unitID == id
		return x
		
	}
	return nil
}
