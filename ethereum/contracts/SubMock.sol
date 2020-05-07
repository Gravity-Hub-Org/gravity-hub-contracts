pragma solidity >=0.5.16 <=0.6.6;

import "./ISubscription.sol";

contract SubMock is ISubscription {
    address payable nebulaAddress;
    uint256 reward;
    constructor(address payable newNebulaAddress, uint256 newReward) public {
        nebulaAddress = newNebulaAddress;
        reward = newReward;
    }
    function () external payable {}

    function attachData(uint256[] calldata data) external {
        nebulaAddress.transfer(reward);
    }

}