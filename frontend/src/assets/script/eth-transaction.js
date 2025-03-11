import Web3 from "web3";
import UserContractABI from "../../../../ethereum/build/contracts/SupplyChainContract.json"; // Sesuaikan path ke ABI

let web3;
let userContract;
const contractAddress = "0x2d17043d2dF2fDF2cE295071DbEc81a6A118d6b0";

if (window.ethereum) {
  web3 = new Web3(window.ethereum);
  window.ethereum.enable();
} else {
  console.error("Metamask tidak terdeteksi");
}

userContract = new web3.eth.Contract(UserContractABI.abi, contractAddress);

export { web3, userContract };
