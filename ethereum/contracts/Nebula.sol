pragma solidity >=0.5.16 <=0.6.6;

import "./libs/Queue.sol";
import "./Models.sol";
import "./ISubscription.sol";

contract Nebula {
    uint8 constant oracleCountInEpoch = 5;
    uint256 constant epochInterval = 10;

    QueueLib.Queue public oracleQueue;
    QueueLib.Queue public subscriptionsQueue;
    QueueLib.Queue public pulseQueue;

    string public name;
    uint256 public oraclesCount;
    bytes32 public currentOracleId;
    uint256 public currentEpoch;
    uint256 public endEpochHeight;

    mapping(bytes32 => address) public oracles;
    mapping(address => Models.Oracle) public oracleInfo;
    mapping(bytes32 => Models.Subscription) public subscriptions;
    mapping(uint256 => Models.Pulse) public pulses;
    mapping(uint256 => mapping(bytes32 => bool)) public isPublseSubSent;

    constructor(string memory newName, address[] memory newOracle) public {
        name = newName;
        oraclesCount = newOracle.length;

        for (uint256 i = 0; i < newOracle.length; i++) {
            bytes32 id = keccak256(abi.encode(newOracle[i]));
            QueueLib.push(oracleQueue, id);
            oracles[id] = newOracle[i];

            oracleInfo[newOracle[i]] = Models.Oracle(newOracle[i], true, id);
        }
    }

    function newRound() public {
        require(block.number > endEpochHeight, "round is not ended");

        currentEpoch++;
        if (currentOracleId == oracleQueue.last) {
            currentOracleId = oracleQueue.first;
        } else {
            currentOracleId = QueueLib.next(oracleQueue, currentOracleId);
        }

        endEpochHeight = block.number + epochInterval;
    }

    function confirmData(uint256[] memory data, uint8[] memory v, bytes32[] memory r, bytes32[] memory s) public {
        uint256 count = 0;
        address[oracleCountInEpoch] memory oracleList;
        bytes32 targetId = currentOracleId;
        for(uint i = 0; i < oracleCountInEpoch; i++) {
            oracleList[i] = oracles[targetId];
            targetId = QueueLib.next(oracleQueue, targetId);
        }

        uint256 oracleId = block.number % oracleCountInEpoch;
        require(msg.sender == oracleList[oracleId], "invalid owner");
        for(uint i = 0; i < oracleCountInEpoch; i++) {
            count += ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n33", abi.encodePacked(name, data[i]))),
                v[i], r[i], s[i]) == oracleList[i] ? 1 : 0;
        }

        pulses[block.number] = Models.Pulse(data, oracleList, count);
    }

    function sendData(uint256 blockNumber, bytes32 subscriptionId) public {
        require(blockNumber <= block.number + 1, "invalid block number");
        require(isPublseSubSent[blockNumber][subscriptionId] == false, "sub sent");
        isPublseSubSent[blockNumber][subscriptionId] = true;

        uint256 startBalance = address(this).balance;
        ISubscription(subscriptions[subscriptionId].contractAddress).attachData(pulses[blockNumber].data);

        uint256 endBalance = address(this).balance;
        uint256 profit = endBalance-startBalance;

        require(profit >= subscriptions[subscriptionId].reward, "invalid reward");
    }

    function subscribe(address contractAddress, uint8 minConfirmations, uint256 reward) public {
        bytes32 id = keccak256(abi.encodePacked(abi.encodePacked(msg.sig, msg.sender, contractAddress, minConfirmations)));
        require(subscriptions[id].owner == address(0x00), "rq is exist");
        subscriptions[id] = Models.Subscription(msg.sender, contractAddress, minConfirmations, reward);
        QueueLib.push(subscriptionsQueue, id);
    }

    function unsubscribe(bytes32 id) public {
        require(subscriptions[id].owner == msg.sender, "invalid rq");
        QueueLib.drop(subscriptionsQueue, id);
    }

    function trunOffOracle() public {
        require(oracleInfo[msg.sender].owner == address(0x00), "invalid sender");
        oracleInfo[msg.sender].isOnline = false;
        QueueLib.drop(oracleQueue, oracleInfo[msg.sender].idInQueue);
        oracleInfo[msg.sender].idInQueue = 0x00;
    }

    function trunOnOracle() public {
        require(oracleInfo[msg.sender].owner == address(0x00), "invalid sender");

        bytes32 id = keccak256(abi.encode(oracleInfo[msg.sender].owner));
        QueueLib.push(oracleQueue, id);
        oracles[id] = oracleInfo[msg.sender].owner;

        oracleInfo[msg.sender].isOnline = true;
        oracleInfo[msg.sender].idInQueue = 0x00;
    }

}