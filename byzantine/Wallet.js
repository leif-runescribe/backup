const Util = require('./Util')
const Txn = require('./Txn')

class Wallet{
    constructor(secret){
        this.keyPair = Util.genKeypair(secret)
        this.pubKey = this.keyPair.getPublic('hex')

    }

    toString(){
        return `Wallet -
        pubKey: ${this.pubKey.toString()}`
    
    }
sign(dataHash){
    return this.keyPair.sign(dataHash).toHex()
}

createTxn(data){
    return new Txn(data,this)
}
getPubKey(){
    return this.pubKey


}
}

module.exports = Wallet;