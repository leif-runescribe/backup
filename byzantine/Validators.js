const Wallet = require('./Wallet')

class Validators{
    constructor(n){
        this.genAddresses(n)
    }

    genAddresses(n){
        let List = []
        for(let i = 0; i < n; i ++){
            List.push(new Wallet("V"+i).getPubKey())
        }
        return List
    }

    isValidatorValid(v){
        return this.List.includes(v)
    }
}

module.exports = Validators