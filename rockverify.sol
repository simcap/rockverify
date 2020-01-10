pragma solidity ^0.6.1;

contract RockVerify {

    mapping(bytes32 => bytes32) public downloadables;

    event Registered(address from, bytes32 urlShasum, bytes32 fileShasum);

    function register(bytes32 urlShasum, bytes32 fileShasum) public {
        require(downloadables[urlShasum] == 0x00000000000000000000000000000000, "url shasum already registered");
        downloadables[urlShasum] = fileShasum;
        emit Registered(msg.sender, urlShasum, fileShasum);
    }

    function lookup(bytes32 urlShasum) public view returns (bytes32) {
        return downloadables[urlShasum];
    }
}