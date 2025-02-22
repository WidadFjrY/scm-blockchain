// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract UserContract {
    struct User {
        string name;
        address ethAddress;
        string role;
    }

    mapping(address => User) public users;
    address[] public userAddresses;

    event UserRegistered(address indexed ethAddress, string name, string role);

    function registerUser(string memory _name, string memory _role) public {
        require(
            users[msg.sender].ethAddress == address(0),
            "User already registered"
        );

        users[msg.sender] = User(_name, msg.sender, _role);
        userAddresses.push(msg.sender);

        emit UserRegistered(msg.sender, _name, _role);
    }

    function getUser(
        address _userAddress
    ) public view returns (string memory, address, string memory) {
        require(users[_userAddress].ethAddress != address(0), "User not found");

        User memory user = users[_userAddress];
        return (user.name, user.ethAddress, user.role);
    }

    function getAllUsers()
        public
        view
        returns (address[] memory, string[] memory, string[] memory)
    {
        uint256 userCount = userAddresses.length;
        string[] memory names = new string[](userCount);
        string[] memory roles = new string[](userCount);

        for (uint256 i = 0; i < userCount; i++) {
            address userAddress = userAddresses[i];
            names[i] = users[userAddress].name;
            roles[i] = users[userAddress].role;
        }

        return (userAddresses, names, roles);
    }

    function getTotalUsers() public view returns (uint256) {
        return userAddresses.length;
    }
}
