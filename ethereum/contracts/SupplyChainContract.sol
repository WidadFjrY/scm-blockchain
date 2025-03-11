// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SupplyChainContract {
    struct User {
        string name;
        address ethAddress;
        string role;
    }

    struct Product {
        string id;
        string name;
        string hashInfo;
        address addedBy;
    }

    struct Transaction {
        uint256 id;
        address buyer;
        string[] productIds;
        uint256[] quantities;
        uint256 totalPrice;
        uint256 timestamp;
        string status;
    }

    mapping(address => User) public users;
    address[] public userAddresses;

    mapping(string => Product) public products;
    string[] public productIds;

    mapping(uint256 => Transaction) public transactions;
    uint256 public transactionCounter;

    address public admin;

    event UserRegistered(address indexed ethAddress, string name, string role);
    event ProductAdded(
        string indexed id,
        string name,
        string hashInfo,
        address indexed addedBy
    );
    event TransactionCreated(
        uint256 indexed id,
        address indexed buyer,
        string[] productIds,
        uint256[] quantities,
        uint256 totalPrice,
        uint256 timestamp,
        string status
    );
    event TransactionStatusUpdated(
        uint256 indexed transactionId,
        string status
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
            "Only admin can perform this action"
        );
        _;
    }

    constructor() {
        admin = msg.sender;
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

    function addProduct(
        string memory _id,
        string memory _name,
        string memory _hashInfo
    ) public onlyWarehouseStaff {
        require(bytes(products[_id].id).length == 0, "Product already exists");

        products[_id] = Product(_id, _name, _hashInfo, msg.sender);
        productIds.push(_id);

        emit ProductAdded(_id, _name, _hashInfo, msg.sender);
    }

    function getProduct(
        string memory _productId
    )
        public
        view
        returns (string memory, string memory, string memory, address)
    {
        require(
            bytes(products[_productId].id).length != 0,
            "Product not found"
        );

        Product memory product = products[_productId];
        return (product.id, product.name, product.hashInfo, product.addedBy);
    }

    function getAllProducts()
        public
        view
        returns (
            string[] memory,
            string[] memory,
            string[] memory,
            address[] memory
        )
    {
        uint256 productCount = productIds.length;

        string[] memory ids = new string[](productCount);
        string[] memory names = new string[](productCount);
        string[] memory hashInfos = new string[](productCount);
        address[] memory addedBys = new address[](productCount);

        for (uint256 i = 0; i < productCount; i++) {
            Product memory product = products[productIds[i]];
            ids[i] = product.id;
            names[i] = product.name;
            hashInfos[i] = product.hashInfo;
            addedBys[i] = product.addedBy;
        }

        return (ids, names, hashInfos, addedBys);
    }

    function getTotalProducts() public view returns (uint256) {
        return productIds.length;
    }

    function createTransaction(
        string[] memory _productIds,
        uint256[] memory _quantities
    ) public payable {
        require(_productIds.length == _quantities.length, "Invalid input");
        require(_productIds.length > 0, "No products specified");
        require(msg.value > 0, "ETH must be sent to make a purchase");

        transactionCounter++;
        transactions[transactionCounter] = Transaction(
            transactionCounter,
            msg.sender,
            _productIds,
            _quantities,
            msg.value,
            block.timestamp,
            "Pending"
        );

        payable(admin).transfer(msg.value);

        emit TransactionCreated(
            transactionCounter,
            msg.sender,
            _productIds,
            _quantities,
            msg.value,
            block.timestamp,
            "Pending"
        );
    }

    function updateTransactionStatus(
        uint256 _transactionId,
        string memory _status
    ) public onlyAdmin {
        require(transactions[_transactionId].id != 0, "Transaction not found");
        transactions[_transactionId].status = _status;

        emit TransactionStatusUpdated(_transactionId, _status);
    }

    function getTransaction(
        uint256 _transactionId
    )
        public
        view
        returns (
            uint256,
            address,
            string[] memory,
            uint256[] memory,
            uint256,
            uint256,
            string memory
        )
    {
        require(transactions[_transactionId].id != 0, "Transaction not found");

        Transaction memory txn = transactions[_transactionId];
        return (
            txn.id,
            txn.buyer,
            txn.productIds,
            txn.quantities,
            txn.totalPrice,
            txn.timestamp,
            txn.status
        );
    }

    function getTransactionsByBuyer(
        address _buyer
    ) public view returns (Transaction[] memory) {
        uint256 count = 0;
        for (uint256 i = 1; i <= transactionCounter; i++) {
            if (
                transactions[i].buyer == _buyer &&
                keccak256(abi.encodePacked(transactions[i].status)) !=
                keccak256(abi.encodePacked("Selesai"))
            ) {
                count++;
            }
        }

        Transaction[] memory result = new Transaction[](count);
        uint256 index = 0;

        for (uint256 i = 1; i <= transactionCounter; i++) {
            if (
                transactions[i].buyer == _buyer &&
                keccak256(abi.encodePacked(transactions[i].status)) !=
                keccak256(abi.encodePacked("Selesai"))
            ) {
                result[index] = transactions[i];
                index++;
            }
        }

        return result;
    }

    function getAllPendingTransactions()
        public
        view
        returns (Transaction[] memory)
    {
        uint256 count = 0;

        for (uint256 i = 1; i <= transactionCounter; i++) {
            if (
                keccak256(abi.encodePacked(transactions[i].status)) !=
                keccak256(abi.encodePacked("Selesai"))
            ) {
                count++;
            }
        }

        Transaction[] memory pendingTransactions = new Transaction[](count);
        uint256 index = 0;

        for (uint256 i = 1; i <= transactionCounter; i++) {
            if (
                keccak256(abi.encodePacked(transactions[i].status)) !=
                keccak256(abi.encodePacked("Selesai"))
            ) {
                pendingTransactions[index] = transactions[i];
                index++;
            }
        }

        return pendingTransactions;
    }
}
