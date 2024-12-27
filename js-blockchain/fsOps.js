const fs = require('fs');
const path = require('path');
const Blockchain = require('./Blockchain');

const blockchainFile = path.join(__dirname, 'blockchain.json');

function saveBlockchain(blockchain) {
    fs.writeFileSync(blockchainFile, JSON.stringify(blockchain, null, 2), 'utf8');
}

function loadBlockchain() {
    if (fs.existsSync(blockchainFile)) {
        const data = fs.readFileSync(blockchainFile, 'utf8');
        const chainData = JSON.parse(data);

        const blockchain = new Blockchain();
        blockchain.chain = chainData.chain;
        blockchain.pendingTransactions = chainData.pendingTransactions;
        return blockchain;
    }
    return new Blockchain();
}

module.exports = { saveBlockchain, loadBlockchain };
