// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

contract MySmartContract {
    uint256 balance = 0;
    address public admin;

    constructor() payable  {
        admin = msg.sender;
        updateBalance();
    }

    receive() external payable {
        updateBalance();
    }

    function updateBalance() internal {
        balance += msg.value;
    }

    function Withdrawl(uint256 _amt) public{
        require(msg.sender == admin);
        balance = balance - _amt;
    }

    function Deposit(uint256 amount) payable public {
        require(msg.value == amount);
             updateBalance();
  
    }


    function Balance() public view returns (uint256) {
        return balance;
    }

}
