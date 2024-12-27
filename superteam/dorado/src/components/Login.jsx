import React, { useState } from "react";
import logo from "./assets/logo.png";
import "./Login.css";
import { Keypair } from "@solana/web3.js";

const Login = ({ setAccount }) => {
  const [privateKeyInput, setPrivateKeyInput] = useState("");

  const createAccount = () => {
    const keypair = Keypair.generate();
    setAccount(keypair);
  };

  const retrieveAccount = () => {
    try {
      const keypair = Keypair.fromSecretKey(Buffer.from(privateKeyInput, "hex"));
      setAccount(keypair);
    } catch (error) {
      console.error("Error retrieving account:", error);
      alert("Invalid private key format");
    }
  };

  return (
    <div className="container">
      <img src={logo} className="logo" alt="Logo" />
      <div className="create">
        <button onClick={createAccount}>Create Account</button>
      </div>
      <div className="login">
        <input
          style={{ marginTop: "32px" }}
          value={privateKeyInput}
          placeholder="Enter your private key"
          onChange={(event) => {
            setPrivateKeyInput(event.target.value);
          }}
        />
        <button style={{ marginTop: "16px" }} onClick={retrieveAccount}>
          Login
        </button>
      </div>
    </div>
  );
};

export default Login;
