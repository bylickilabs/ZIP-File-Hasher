|[![Go](https://github.com/bylickilabs/ZIP-File-Hasher/actions/workflows/go.yml/badge.svg)](https://github.com/bylickilabs/ZIP-File-Hasher/actions/workflows/go.yml)|
|---|

# 🔐 ZIP File Hasher

|<img width="1280" height="640" alt="logo" src="https://github.com/user-attachments/assets/1cea97c3-c965-49ec-af0c-fa880643aeff" />|
|---|

A powerful Go-based CLI tool to securely hash the contents of ZIP files – 100% local, GDPR-compliant, with zero external dependencies.

<details>
  <summary>✅ Go Developer</summary>

| ![certificate_of_completion_go](https://github.com/user-attachments/assets/b8eca14b-a2fe-4419-abf0-a7d110a07dd2) |
|---|

</details>

---

## 🚀 Features

- 🗜️ Handles `.zip` files only
- 📁 Extracts ZIP content into temporary directories
- 🔐 Supported hash algorithms: `md5`, `sha1`, `sha256`, `sha512`
- 🧾 Optional JSON output for automation and integration
- 🛡️ Fully local processing – no upload, no cloud, no tracking

---

## 🖥️ Requirements

- [Install Go (version 1.18 or higher)](https://go.dev/dl/)
- Recommended: Windows 10 / 11

---

## 📦 Project Structure

```plaintext
ziphasher/
├── main.go                 # Application entry point
├── hasher/
│   └── filehash.go        # File hashing functionality
└── zip/
    └── unzip.go           # ZIP extraction logic
```

---

## ⚙️ Initialization

```yarn
go mod init ziphasher
```

> ⚠️ The module name **must match** the import paths in `main.go`.

---

## ▶️ Run

```powershell
go run main.go -zip="your_zip_file.zip" -algo=sha512 -json
```

### 🔧 Parameters

| Parameter | Description |
|-----------|-------------|
| `-zip`    | Path to the ZIP file |
| `-algo`   | Hash algorithm: `md5`, `sha1`, `sha256`, `sha512` |
| `-json`   | Enables JSON output format |

---

## 🔨 Optional: Build

```powershell
go build -o ziphasher.exe
```

Then run via:

```powershell
.\ziphasher.exe -zip="test.zip" -algo=sha256 -json
```

---

## 📑 Example Output (JSON)

```json
{
  "test.zip_extracted/readme.txt": "abc123def456...",
  "test.zip_extracted/docs/manual.pdf": "f7b9aa22bb77..."
}
```

---

## 👨‍💻 Author & License

BYLICKILABS – 2025  

License: MIT  
[LICENSE](LICENSE)
