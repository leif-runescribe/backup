const Territory = require('./work/Territory');

class Game {
    constructor(io) {
        this.io = io;
        this.games = new Map();

        this.io.on('connection', (socket) => {
            console.log('New client connected');

            socket.on('createGame', (playerName) => this.createGame(socket, playerName));
            socket.on('joinGame', ({ gameCode, playerName }) => this.joinGame(socket, gameCode, playerName));
            socket.on('claimTerritory', ({ gameCode, playerName, territoryName }) => this.claimTerritory(socket, gameCode, playerName, territoryName));
            socket.on('disconnect', () => this.handleDisconnect(socket));
        });
    }

    createGame(socket, playerName) {
        const gameCode = this.generateGameCode();
        const territories = this.generateMap();
        const playerColor = this.getRandomColor();

        this.games.set(gameCode, {
            players: new Map([[socket.id, { name: playerName, color: playerColor }]]),
            territories: territories,
            currentTurnPlayerId: socket.id // First player's turn
        });

        socket.join(gameCode);
        socket.emit('gameCreated', { gameCode, territories: Array.from(territories.values()) });
        this.io.to(gameCode).emit('newEvent', `${playerName} created the game`);
        this.updateGameStats(gameCode);
    }

    joinGame(socket, gameCode, playerName) {
        const game = this.games.get(gameCode);
        if (game) {
            const playerColor = this.getRandomColor();
            game.players.set(socket.id, { name: playerName, color: playerColor });
            socket.join(gameCode);
            socket.emit('gameJoined', {
                gameCode,
                territories: Array.from(game.territories.values())
            });
            this.io.to(gameCode).emit('newEvent', `${playerName} joined the game`);
            this.updateGameStats(gameCode);
        } else {
            socket.emit('error', 'Game not found');
        }
    }

    claimTerritory(socket, gameCode, playerName, territoryName) {
        const game = this.games.get(gameCode);
        if (game && game.territories.has(territoryName)) {
            const territory = game.territories.get(territoryName);
            // Check if it's the player's turn
            if (game.currentTurnPlayerId !== socket.id) {
                socket.emit('error', 'It is not your turn to claim a territory');
                return;
            }

            if (!territory.owner) {
                territory.owner = socket.id; // Assign the owner
                this.io.to(gameCode).emit('territoryUpdate', {
                    name: territoryName,
                    owner: playerName,
                    color: game.players.get(socket.id).color, // Include player color
                    ...territory
                });
                this.io.to(gameCode).emit('newEvent', `${playerName} claimed ${territoryName}`);
                
                // Change turn to next player
                game.currentTurnPlayerId = this.getNextPlayerId(game, socket.id);
                this.updateGameStats(gameCode);
            } else {
                socket.emit('error', 'Territory already claimed');
            }
        } else {
            socket.emit('error', 'Territory not found');
        }
    }

    handleDisconnect(socket) {
        console.log('Client disconnected');
        for (const [gameCode, game] of this.games.entries()) {
            if (game.players.has(socket.id)) {
                const playerName = game.players.get(socket.id).name;
                game.players.delete(socket.id);
                for (const territory of game.territories.values()) {
                    if (territory.owner === socket.id) {
                        territory.owner = null;
                        this.io.to(gameCode).emit('territoryUpdate', {
                            name: territory.name,
                            owner: null
                        });
                    }
                }
                this.io.to(gameCode).emit('newEvent', `${playerName} left the game`);
                this.updateGameStats(gameCode);
                break;
            }
        }
    }

    generateGameCode() {
        return Math.random().toString(36).substring(2, 8).toUpperCase();
    }

    generateMap() {
        const territories = new Map();
        const usedCoordinates = new Set();
        const minDistance = 50; // Minimum distance between territories

        while (territories.size < 10) {
            let x, y;
            do {
                x = Math.floor(Math.random() * 800);
                y = Math.floor(Math.random() * 600);
            } while (this.isOverlapping(x, y, usedCoordinates, minDistance));

            usedCoordinates.add(`${x},${y}`);
            const name = Territory.getRandomName();
            territories.set(name, new Territory(name, x, y));
        }

        return territories;
    }

    isOverlapping(x, y, usedCoordinates, minDistance) {
        for (const coord of usedCoordinates) {
            const [usedX, usedY] = coord.split(',').map(Number);
            const distance = Math.sqrt(Math.pow(usedX - x, 2) + Math.pow(usedY - y, 2));
            if (distance < minDistance) {
                return true; // Overlapping territory
            }
        }
        return false; // No overlap
    }

    updateGameStats(gameCode) {
        const game = this.games.get(gameCode);
        if (game) {
            const globalStats = {
                players: game.players.size,
                territories: Array.from(game.territories.values()).filter(t => t.owner !== null).length
            };
            this.io.to(gameCode).emit('updateGlobalStats', globalStats);

            for (const [playerId, playerData] of game.players.entries()) {
                const ownedTerritories = Array.from(game.territories.values()).filter(t => t.owner === playerId);
                const playerStats = {
                    name: playerData.name,
                    color: playerData.color, // Include player color
                    territories: ownedTerritories.map(t => ({
                        name: t.name,
                        population: t.population,
                        foodYield: t.foodYield,
                        revenue: t.revenue,
                        militaryStrength: t.militaryStrength,
                        color: playerData.color // Include color for each territory
                    }))
                };
                this.io.to(playerId).emit('updatePlayerStats', playerStats);
            }
        }
    }

    getRandomColor() {
        const colors = ['red', 'blue', 'green', 'yellow', 'purple', 'orange']; // Add more colors as needed
        return colors[Math.floor(Math.random() * colors.length)];
    }

    getNextPlayerId(game, currentPlayerId) {
        const playerIds = Array.from(game.players.keys());
        const currentIndex = playerIds.indexOf(currentPlayerId);
        return playerIds[(currentIndex + 1) % playerIds.length]; // Loop back to the first player
    }
}

module.exports = Game;
