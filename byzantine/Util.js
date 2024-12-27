const eddsa = require('elliptic').eddsa
const crypt = new eddsa('ed25519')
const uu = require('uuid').v1
const SHA = require('crypto-js').SHA256

class Util{
    static genKeypair(secret){
        console.log(crypt.keyFromSecret(secret))

        return crypt.keyFromSecret(secret)
    }

    static id(){
        console.log(uu());
        return uu();
    }

    static hash(data){
        return SHA(JSON.stringify(data).toString())
    }

    static verifySign(pub, sign, dataHash){
        return crypt.keyFromPublic(pub).verify(dataHash, sign)
    }
}

module.exports = Util;

