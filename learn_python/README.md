## Intro to Python

### 1. Setup:

1.1. Option A) (Basic) Just install Python:  
    - Install Python from the official website: https://www.python.org/downloads/.  
    - Verify installation by running `python --version` in your terminal or command prompt.  

1.1. Option B) (Advanced) Install `uv` and it will install and manage Python and its dependencies:

- Windows:  

```ps
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex" 
```

- Linux/Mac:

```bash
curl -LsSf https://astral.sh/uv/install.sh | sh
```

- Then run 
```bash
uv install python
uv run python --version
```

- If you create a new repo for Python projects, you can run:
```bash
uv init .
uv install python
uv run python <your_script>.py
```

1.2. Install a code editor like VSCode or PyCharm for writing Python code:  
- [VSCode](https://code.visualstudio.com/) (recommended)
- [PyCharm](https://www.jetbrains.com/pycharm/) (alternative)


### 2. Free resources to learn Python:

- [Kaggle Python Course](https://www.kaggle.com/learn/python)
- [Automate the Boring Stuff with Python](https://automatetheboringstuff.com)
- [YouTube - Learn Python](https://www.youtube.com/watch?v=K5KVEU3aaeQ)


### 3. Advent of Code 2024 day 1 solution in Python:

- See the `aoc_2024_day01.ipynb` Jupyter Notebook file in this folder for the complete solution.


### 4. Your turn:

- Solve the 2nd, 3rd and following days of Advent of Code 2024 using Python, and after the 1st of December of 2025, you can also start solving the 2025 edition!



  