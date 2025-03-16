import Web3 from "web3";
import SupplyChainContractABI from "../../../../ethereum/build/contracts/SupplyChainContract.json"; // Sesuaikan path ke ABI

let web3;
let SupplyChainContract;
const contractAddress = import.meta.env.VITE_CONTRACT_ADDRESS;

if (window.ethereum) {
  web3 = new Web3(window.ethereum);
  window.ethereum.enable();
} else {
  console.error("Metamask tidak terdeteksi");
}

SupplyChainContract = new web3.eth.Contract(
  SupplyChainContractABI.abi,
  contractAddress
);

export { web3, SupplyChainContract };
