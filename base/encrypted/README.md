# Encrypted Package

The `encrypted` package provides a simple and consistent way to encrypt and decrypt verification-related data (e.g., address, credit card, email, phone, ID, face data) using **RSA public/private keys**.  
Encrypted data can be persisted using the [`aria/storage`](https://github.com/ralphferrara/aria/tree/main/storage) backend.

---

## ✨ Features

- RSA encryption using public/private key pairs.
- Typed helpers for common verification data:
  - `MakeADDR` / `FetchADDR`
  - `MakeCRCD` / `FetchCRCD`
  - `MakeMAIL` / `FetchMAIL`
  - `MakePHNE` / `FetchPHNE`
  - `MakeIDEN` / `FetchIDEN`
  - `MakeFACE` / `FetchFACE`
- Generic `MakeCipher` for other payloads.
- Storage integration via `enc.Save(store)` and `FetchX(store, uuid, privateKey)`.

---

## 📦 Installation

```bash
go get github.com/complyage/base/encrypted
```

---

## 🔑 Key Concepts

### Encrypted Struct

```go
type Encrypted struct {
    UUID   string `json:"uuid"`
    Type   string `json:"type"`
    Cipher []byte `json:"cipher"`
}
```

- **UUID** – identifier of the verification record.
- **Type** – data type (ADDR, CRCD, MAIL, PHNE, IDEN, FACE).
- **Cipher** – encrypted JSON payload (RSA-encrypted).

---

## 🚀 Usage

### 1. Generate Keys

```go
priv, _ := rsa.GenerateKey(rand.Reader, 2048)
privBytes, _ := x509.MarshalPKCS8PrivateKey(priv)
privPEM := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})

pubBytes, _ := x509.MarshalPKIXPublicKey(&priv.PublicKey)
pubPEM := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes})
```

---

### 2. Encrypt and Save

```go
store := &storage.Storage{
    Config: storage.StoreConfig{
        Backend:   storage.StorageLocal,
        LocalPath: "/tmp",
    },
}
store.InitStorage()

addr := types.Address{
    Line1: "123 Main St", City: "Metropolis", State: "CA", Postal: "90210", Country: "USA",
}

uuid := "test-uuid-1234"

// Encrypt
enc, err := encrypted.MakeADDR(string(pubPEM), uuid, addr)
if err != nil {
    panic(err)
}

// Save
if err := enc.Save(store); err != nil {
    panic(err)
}
```

---

### 3. Fetch and Decrypt

```go
got, err := encrypted.FetchADDR(store, uuid, string(privPEM))
if err != nil {
    panic(err)
}

fmt.Printf("Decrypted Address: %+v\n", got)
```

---

## 🧪 Testing

Verbose test run:

```bash
go test -v ./...
```

Example output:

```
=== RUN   TestMakeAndFetchADDR
--- PASS: TestMakeAndFetchADDR (0.12s)
=== RUN   TestMakeAndFetchCRCD
--- PASS: TestMakeAndFetchCRCD (0.15s)
PASS
ok   github.com/complyage/base/encrypted   0.42s
```

---

## 📂 Project Structure

```
encrypted/
├── encrypted.go      # Core types and functions
├── addr.go           # Address encrypt/decrypt
├── crcd.go           # Credit card encrypt/decrypt
├── mail.go           # Email encrypt/decrypt
├── phne.go           # Phone encrypt/decrypt
├── iden.go           # Identification encrypt/decrypt
├── face.go           # Face data encrypt/decrypt
└── unit_test.go      # Unit tests
```

---

## ⚖️ License

MIT – see LICENSE file.
