# Nice Token Creation Guide

This guide will help you create a new SPL token on Solana with the Token Extensions program and its token metadata extension. Follow the steps to generate keypairs, fund your account, create a token mint, add metadata, and mint tokens.

## Installation

First, install the Solana CLI by following [this guide](https://docs.solanalabs.com/cli/install).

## Commands and Instructions

### Create Folder and Keypairs

1. **Create a new folder:**
    ```bash
    mkdir nice-token
    cd nice-token
    ```

2. **Generate a keypair for the token owner:**
    ```bash
    solana-keygen grind --starts-with key:1
    ```

3. **Set the Solana CLI to use the generated keypair:**
    ```bash
    solana config set --keypair path/to/keypair.json
    ```

4. **Generate a keypair for the token mint address:**
    ```bash
    solana-keygen grind --starts-with nic:1
    ```

5. **Set Solana CLI to use devnet:**
    ```bash
    solana config set -ud
    ```

### Fund the Account

1. **Airdrop SOL to your account (for devnet):**
    ```bash
    solana airdrop 2
    ```

2. **Check your account address:**
    ```bash
    solana address
    ```

### Create the Token Mint

1. **Create the token mint with metadata extension:**
    ```bash
    spl-token create-token --program-id TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb --enable-metadata path/to/mint-keypair.json
    ```

### Create and Upload Metadata

1. **Create a metadata JSON file:**
    ```json
    {
        "name": "Nice Token",
        "symbol": "NICE",
        "image": "link_to_your_image"
    }
    ```

2. **Upload the metadata file to a storage solution (e.g., Web3 Storage, IPFS).**

### Add Metadata to the Token

1. **Initialize the token metadata:**
    ```bash
    spl-token initialize-metadata path/to/mint-keypair.json 'Nice Token' 'NICE' 'link_to_metadata_file'
    ```

### Mint Tokens

1. **Create a token account:**
    ```bash
    spl-token create-account path/to/mint-keypair.json
    ```

2. **Mint tokens:**
    ```bash
    spl-token mint path/to/mint-keypair.json 100
    ```

3. **Transfer tokens to another account:**
    ```bash
    spl-token transfer path/to/mint-keypair.json 10 recipient_address --fund-recipient
    ```

## Challenges

1. **Generating keypairs with specific prefixes can take time.**
2. **Ensure that metadata and image links are correctly uploaded and accessible.**

## Accomplishments

Successfully created and minted a token with associated metadata on the Solana blockchain using the Token Extensions program.

## What I Learned

- How to set up and configure the Solana CLI.
- How to generate keypairs and configure the Solana environment.
- The process of creating a token mint and initializing metadata.
- How to mint tokens and transfer them between accounts.