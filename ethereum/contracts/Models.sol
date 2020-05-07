pragma solidity >=0.5.16 <=0.6.6;

library Models {
    uint8 constant oracleCountInEpoch = 5;

    struct Subscription {
        address owner;
        address contractAddress;
        uint8 minConfirmations;
        uint256 reward;
    }

    struct Pulse {
        uint256[] data;
        address[oracleCountInEpoch] oracles;
        uint256 confirmationCount;
    }

    struct Oracle {
        address owner;
        bool isOnline;
        bytes32 idInQueue;
    }
}