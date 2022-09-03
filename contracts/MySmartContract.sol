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
        require(msg.sender == admin, "Only contract owner can withdraw funds");
        balance = balance - _amt;
    }

    function Deposit() payable public {
        require(msg.value >=1 ether, "Deposit should be greater than 1 ether");
             updateBalance();
    }


    function Balance() public view returns (uint256) {
        return balance;
    }

    function ContractBalance() public view returns (uint256) {
        return address(this).balance;
    }

}