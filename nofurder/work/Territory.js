class Territory {
    constructor(name, x, y) {
        this.name = name;
        this.x = x;
        this.y = y;
        this.population = Math.floor(Math.random() * 1000000) + 100000;
        this.foodYield = Math.floor(Math.random() * 1000) + 100;
        this.revenue = Math.floor(Math.random() * 10000) + 1000;
        this.militaryStrength = Math.floor(Math.random() * 100) + 10;
        this.owner = null;
    }

    static getRandomName() {
        const territoryNames = [
            "Winterfell", "King's Landing", "The Eyrie", "Riverrun", "Casterly Rock",
            "Highgarden", "Sunspear", "The Twins", "Pyke", "Storm's End",
            "Dragonstone", "Oldtown", "The Dreadfort", "Harrenhal", "The Citadel", "LigmaLand"
        ];
        return territoryNames[Math.floor(Math.random() * territoryNames.length)];
    }
}

module.exports = Territory;
