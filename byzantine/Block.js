class Block{
    constructor(tiemstamp, prevHash, hash, data, proposer, sign, seq){
        this.tiemstamp = tiemstamp
        this.hash = hash
        this.prevHash = prevHash
        this.data = data
        this.proposer = proposer
        this.sign = sign
        this.seq = seq
    }

    toString(){
        return `Block -
        Timestamp : ${this.timestamp}
        Previous Hash : ${this.prevHash}
        Hash : ${this.hash}
        Data : ${this.data}
        Proposer : ${this.proposer}
        Signature : ${this.sign}
        Sequence Number : ${this.seq}
        `
    }

    static genesis() {
        return new this(
          `genesis time`,
          "----",
          "genesis-hash",
          [],
          "P4@P@53R",
          "SIGN",
          0
        );
      }
    
      // creates a block using the passed lastblock, transactions and wallet instance
      static createBlock(lastBlock, data, wallet) {
        let hash;
        let timestamp = Date.now();
        const lastHash = lastBlock.hash;
        hash = Block.hash(timestamp, lastHash, data);
        let proposer = wallet.getPublicKey();
        let signature = Block.signBlockHash(hash, wallet);
        return new this(
          timestamp,
          lastHash,
          hash,
          data,
          proposer,
          signature,
          1 + lastBlock.sequenceNo
        );
      }
    
      // hashes the passed values
      static hash(timestamp, lastHash, data) {
        return SHA256(JSON.stringify(`${timestamp}${lastHash}${data}`)).toString();
      }
    
      // returns the hash of a block
      static blockHash(block) {
        const { timestamp, lastHash, data } = block;
        return Block.hash(timestamp, lastHash, data);
      }
    
      // signs the passed block using the passed wallet instance
      static signBlockHash(hash, wallet) {
        return wallet.sign(hash);
      }
    
      // checks if the block is valid
      static verifyBlock(block) {
        return ChainUtil.verifySignature(
          block.proposer,
          block.signature,
          Block.hash(block.timestamp, block.lastHash, block.data)
        );
      }
    
      // verifies the proposer of the block with the passed public key
      static verifyProposer(block, proposer) {
        return block.proposer == proposer ? true : false;
      }
    }
    
    module.exports = Block;