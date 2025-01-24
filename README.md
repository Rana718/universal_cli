# 🌌 Universe CLI

A powerful and user-friendly command-line interface tool built in Go for managing universal tasks effortlessly. 🚀

---

## 🛠 Installation

Follow the steps below to install Universe CLI:

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/Rana718/universal_cli.git
```

### 2️⃣ Navigate to the Project Directory
```bash
cd universal_cli
```

### 3️⃣ Build the Project
```bash
go build
```

### 4️⃣ Install the CLI
```bash
go install
```

---

## 📖 Usage

Universe CLI (`gos`) provides a variety of commands to manage paths and execute common tasks.

### 🚀 Getting Started
To display the help menu with available commands and flags:
```bash
gos -h
```

To check the current version of the tool:
```bash
gos version
```

---

### 🔧 Commands and Examples

#### ➕ Add a Path
Add a custom path to the JSON file with a specific name:
```bash
gos addPath backend D:/code/backend
```

#### 📂 Open a Path in Explorer
Open the stored path in your system's file explorer:
```bash
gos ex D:/code/backend
```
Or use the `-p` flag to open the path directly from the JSON file:
```bash
gos ex -p backend
```

#### 💻 Open a Path in Visual Studio Code
Quickly open a stored path in VS Code:
```bash
gos vs D:/code/backend
```
Or use the `-p` flag to open the path directly from the JSON file:
```bash
gos vs -p backend
```

#### 🔍 Get a Path
Retrieve the path associated with a specific name:
```bash
gos get backend
```

#### ❌ Remove a Path
Delete the path associated with a specific name:
```bash
gos rem backend
```

<!-- #### 🔄 Autocompletion
Generate an autocompletion script for your shell:
```bash
gos completion [shell]
``` -->

---

### 🏷 Flags
- `-h, --help`: Display help information.
- `-v, --version`: Show the CLI version.
- `-p, --path`: Open the path directly from the JSON file.

---

### 🖥 Commends
```bash
PS D:\code\Experment\unverse_cli> gos
gos is a CLI tool to manage anything

Usage:
  gos [command]

Available Commands:
  addPath     Add a path to the JSON file with a custom name
  completion  Generate the autocompletion script for the specified shell
  ex          Open the path in Explorer
  get         Get the path associated with the given name
  help        Help about any command
  rem         Remove the path associated with the given name
  vs          Open the path in Visual Studio Code

Flags:
  -h, --help      help for gos
  -p, --path      Open the path from the JSON file
  -v, --version   version for gos

Use "gos [command] --help" for more information about a command.
```

---

## 📜 License

This project is licensed under the MIT License. For more details, see the [LICENSE](LICENSE) file.

---

## 🤝 Contributing

We welcome contributions to improve Universe CLI! Follow these steps to contribute:

1. 🍴 Fork the repository:
   ```bash
   git clone https://github.com/your-username/universal_cli.git
   ```

2. 🌟 Create your feature branch:
   ```bash
   git checkout -b feature/amazing-feature
   ```

3. 📥 Commit your changes:
   ```bash
   git commit -m 'Add some amazing feature'
   ```

4. 🚀 Push to the branch:
   ```bash
   git push origin feature/amazing-feature
   ```

5. 🛠 Open a Pull Request to the main repository.

Let’s build something amazing together! 💡✨
