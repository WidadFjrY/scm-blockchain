import Web3 from "web3";
import UserContractABI from "../../../../ethereum/build/contracts/SupplyChainContract.json"; // Sesuaikan path ke ABI

let web3;
let userContract;
const contractAddress = "0xF6Cd825d5e918891C7e52cbD3bB4390C19fD1874";

if (window.ethereum) {
  web3 = new Web3(window.ethereum);
  window.ethereum.enable();
} else {
  console.error("Metamask tidak terdeteksi");
}

userContract = new web3.eth.Contract(UserContractABI.abi, contractAddress);

export { web3, userContract };
