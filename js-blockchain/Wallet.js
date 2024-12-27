const EC = require('elliptic').ec;
const ec = new EC('secp256k1');

class Wallet {
    constructor() {
        this.keyPair = ec.genKeyPair();
        this.publicKey = this.keyPair.getPublic('hex');
        this.privateKey = this.keyPair.getPrivate('hex');
    }

    getPublicKey() {
        return this.publicKey;
    }

    getPrivateKey() {
        return this.privateKey;
    }

    signTransaction(transaction) {
        transaction.signTransaction(this.keyPair);
    }
}

module.exports = Wallet;
