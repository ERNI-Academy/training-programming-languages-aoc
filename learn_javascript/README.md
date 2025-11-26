# 📘 How to Run JavaScript on Your PC (From Scratch)

This guide explains how to run JavaScript **outside the browser** using **Node.js** and **Visual Studio Code**.  
If you’ve never executed JS on your machine before, this README will walk you through every step.

---

## 🚀 1. Install Node.js

To run JavaScript on your PC, you need **Node.js**, which includes the V8 engine and the npm package manager.

Download the installer here:

👉 https://nodejs.org/

Choose the **LTS version** (recommended).

After installing, open a terminal and verify:

```bash
node -v
npm -v
```

If you see version numbers, Node is correctly installed.

---

## 📝 2. Create a JavaScript File

Create a new folder, for example:

```
test/
```

Inside, create a file named:

```
index.js
```

Add some test code:

```js
console.log("Hello from Node!");
```

---

## 🖥 3. Open the Terminal in Visual Studio Code

1. Open the folder in **VS Code**.
2. Go to the top menu → **Terminal → New Terminal**  
   (shortcut: **Ctrl + Shift + `** on Windows).

This will open the integrated terminal inside VS Code.

---

## ▶️ 4. Run Your JavaScript File

In the terminal, run:

```bash
node index.js
```

Output:

```
Hello from Node!
```

Congrats — you’re executing JavaScript natively on your machine 🎉

---

## 📂 5. Running More Advanced Scripts

You can import modules, read files, etc. Example:

```js
const fs = require("fs");

const text = fs.readFileSync("input.txt", "utf8");
console.log(text);
```

Run again with:

```bash
node index.js
```

---

## 🧩 6. Passing Arguments to Your Script

You can provide parameters directly from the terminal:

```bash
node index.js input.txt
```

Access them inside your script:

```js
const file = process.argv[2];
console.log("Received file:", file);
```

---

## 📦 7. (Optional) Initialize an npm Project

If you want to install third-party packages:

```bash
npm init -y
```

This creates a `package.json`.

Install a package:

```bash
npm install lodash
```

Use it in your script:

```js
const _ = require("lodash");
console.log(_.random(1, 100));
```