const Uril = require('./Util')

class Txn{
    constructor(data, wallet){
        this.id = Util.id
        this.from = wallet.pubKey;
        this.input = {data: data, timestamp: Date.now() }
        this.hash = Util.hash(this.input)
        this.signature = wallet.sign(this.hash)
    }

    static verifyTxn(txn){
        return Util.verifySign(
            txn.from,
            txn.signature,
            Util.hash(txn.input)
        )
    }
}

module.exports = Txn