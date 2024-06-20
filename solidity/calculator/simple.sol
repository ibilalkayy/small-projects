// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract SimpleCalculator {
    uint256 result = 0;

    function add(uint256 num) public {
        result += num;
    }

    function subtract(uint256 num) public {
        result -= num;
    }

    function multiply(uint num) public {
        result *= num;
    }

    function get() public view returns(uint256) {
        return result;
    }
}
