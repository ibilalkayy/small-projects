// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract AdvancedCalculator {
    uint256 public result;

    function add(uint256 a, uint256 b) public {
        result = a + b;
    }

    function subtract(uint256 a, uint256 b) public {
        result = a - b;
    }

    function multiply(uint256 a, uint256 b) public {
        result = a * b;
    }

    function get(uint256 num1, uint256 num2, uint8 operation) public {
        if (operation == 1) add(num1, num2);
        else if (operation == 2) subtract(num1, num2);
        else if (operation == 3) multiply(num1, num2);
        else revert("invalid operation");
    }
}