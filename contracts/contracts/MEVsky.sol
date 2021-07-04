// SPDX-License-Identifier: MIT

pragma solidity =0.8.6;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";


contract MEVsky is Ownable {
    using Address for address payable;

    error WrongState(bool);
    error BountyTooSmall(uint256, uint256);

    event TurnedOn(
        uint256 bounty
    );
    event TurnedOff(
        address receiver
    );
    event MinBountySet(
        uint256 minBounty
    );

    bool public on;
    uint256 public minBounty;

    function turnOn() public payable {
        if (on) {
            revert WrongState(on);
        }
        if (msg.value < minBounty) {
            revert BountyTooSmall(msg.value, minBounty);
        }

        on = true;

        emit TurnedOn({
            bounty: msg.value
        });
    }

    function turnOff(address payable receiver) public {
        if (!on) {
            revert WrongState(on);
        }

        on = false;

        emit TurnedOff({
            receiver: receiver
        });

        receiver.sendValue(address(this).balance);
    }

    function setMinBounty(uint256 newMinBounty) public onlyOwner {
        minBounty = newMinBounty;

        emit MinBountySet({
            minBounty: minBounty
        });
    }
}