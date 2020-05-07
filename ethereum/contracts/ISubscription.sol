pragma solidity >=0.5.16 <=0.6.6;

interface ISubscription {
    function attachData(uint256[] calldata data) external;
}