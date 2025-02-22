import Web3 from "web3";
import UserContractABI from "../../../../ethereum/build/contracts/UserContract.json"; // Sesuaikan path ke ABI

let web3;
let userContract;
const contractAddress = "0x585ffc4122eE8156551Ce68118F3C8322E3ef759";

if (window.ethereum) {
  web3 = new Web3(window.ethereum);
  window.ethereum.enable();
} else {
  console.error("Metamask tidak terdeteksi");
}

userContract = new web3.eth.Contract(UserContractABI.abi, contractAddress);

export { web3, userContract };
