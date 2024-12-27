#!/usr/bin/env node

const Blockchain = require('./Blockchain');
const Transaction = require('./Transaction');
const Wallet = require('./Wallet');
const { saveBlockchain, loadBlockchain } = require('./fsOps');

// Load existing blockchain or create a new one
let blockchain = loadBlockchain();

// Command-line interface
const args = process.argv.slice(2);

if (args.length === 0) {
    console.log("Usage:");
    console.log("  create-wallet");
    console.log("  balance [address]");
    console.log("  send [from_private_key] [to_address] [amount]");
    console.log("  mine [mining_wallet_address]");
    console.log("  reset");
    console.log("  save")
    process.exit(0);
}

const command = args[0];

switch (command) {
    case 'create-wallet':
        const wallet = new Wallet();
        console.log("Wallet created!");
        console.log("Public Key:", wallet.getPublicKey());
        console.log("Private Key:", wallet.getPrivateKey());
        break;

    case 'balance':
        if (args.length < 2) {
            console.log("Usage: balance [wallet_public_key]");
        } else {
            const address = args[1];
            const balance = blockchain.getBalanceOfAddress(address);
            console.log(`Balance of ${address}: $${balance}`);
        }
        break;

    case 'send':
        if (args.length < 4) {
            console.log("Usage: send [from_private_key] [to_address] [amount]");
        } else {
            const fromPrivateKey = args[1];
            const toAddress = args[2];
            const amount = parseFloat(args[3]);

            const wallet = new Wallet();
            wallet.keyPair = wallet.keyPair = wallet.keyPair._importPrivate(fromPrivateKey, 'hex');

            const tx = new Transaction(wallet.getPublicKey(), toAddress, amount);
            wallet.signTransaction(tx);
            blockchain.addTransaction(tx);

            console.log("Transaction added!");
        }
        break;

    case 'mine':
        if (args.length < 2) {
            console.log("Usage: mine [mining_wallet_public_key]");
        } else {
            const minerAddress = args[1];
            blockchain.minePendingTransactions(minerAddress);
            console.log(`Mining rewards sent to ${minerAddress}`);

            // Save the blockchain after mining
            saveBlockchain(blockchain);
        }
        break;

    case 'reset':
        blockchain = new Blockchain();
        saveBlockchain(blockchain);
        console.log("Blockchain reset.");
        break;
    
    
    default:
        console.log("Unknown command");
        break;
}
