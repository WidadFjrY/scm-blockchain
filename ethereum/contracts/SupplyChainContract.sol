// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SupplyChainContract {
    struct User {
        string name;
        address ethAddress;
        string role;
    }

    struct Product {
        uint256 id;
        string name;
        string hashInfo;
        address addedBy;
    }

    mapping(address => User) public users;
    address[] public userAddresses;

    mapping(uint256 => Product) public products;
    uint256 public productCounter;

    event UserRegistered(address indexed ethAddress, string name, string role);
    event ProductAdded(
        uint256 indexed productId,
        string name,
        string hashInfo,
        address indexed addedBy
    );

    modifier onlyWarehouseStaff() {
        require(
            keccak256(abi.encodePacked(users[msg.sender].role)) ==
                keccak256(abi.encodePacked("warehouse_staff")),
            "Only warehouse staff can add products"
        );
        _;
    }

    modifier onlyAdmin() {
        require(
            keccak256(abi.encodePacked(users[msg.sender].role)) ==
                keccak256(abi.encodePacked("admin")),
            "Only admin can add user"
        );
        _;
    }

    function registerUser(string memory _name, string memory _role) public {
        require(
            users[msg.sender].ethAddress == address(0),
            "User already registered"
        );

        users[msg.sender] = User(_name, msg.sender, _role);
        userAddresses.push(msg.sender);

        emit UserRegistered(msg.sender, _name, _role);
    }

    function registerUserWithAdmin(
        string memory _name,
        string memory _role,
        address _ethAddress
    ) public onlyAdmin {
        require(
            users[_ethAddress].ethAddress == address(0),
            "User already registered"
        );

        users[_ethAddress] = User(_name, _ethAddress, _role);
        userAddresses.push(_ethAddress);

        emit UserRegistered(_ethAddress, _name, _role);
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

    // ======================== PRODUCT MANAGEMENT ========================

    function addProduct(
        string memory _name,
        string memory _hashInfo
    ) public onlyWarehouseStaff {
        productCounter++;
        products[productCounter] = Product(
            productCounter,
            _name,
            _hashInfo,
            msg.sender
        );

        emit ProductAdded(productCounter, _name, _hashInfo, msg.sender);
    }

    function getProduct(
        uint256 _productId
    ) public view returns (uint256, string memory, string memory, address) {
        require(products[_productId].id != 0, "Product not found");

        Product memory product = products[_productId];
        return (product.id, product.name, product.hashInfo, product.addedBy);
    }

    function getTotalProducts() public view returns (uint256) {
        return productCounter;
    }
}
