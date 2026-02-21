# 🎯 Gopher Guess

A minimalist, high-efficiency terminal-based number guessing game. Designed for those who want a quick mental challenge without leaving their keyboard-centric workflow.

Built with **Go** and the **Bubble Tea** (Charm.sh) framework.

---

## ✨ Features

- **Dynamic Difficulty**: Choose between **Easy** (more attempts) or **Hard** (fewer attempts) to test your intuition.
- **Range-Based Logic**: Guess any secret number between **1 and 100**.
- **Real-time Feedback**: Instant UI updates indicating if your guess is "Too High" or "Too Low."
- **Visual Clarity**: Modern, color-coded terminal interface powered by **Lip Gloss**.
- **Quick Reset**: Start a fresh round instantly with a single keystroke.

---

## ⌨️ Shortcuts & Controls

### Gameplay Controls

| Key       | Action                |
| :-------- | :-------------------- |
| `j` / `↓` | Move selection down   |
| `k` / `↑` | Move selection up     |
| `Enter`   | Submit current guess  |
| `r`       | **Reset** / New Round |
| `q`       | **Quit** application  |

## 🚀 Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (1.19 or higher)

### Build from Source

1. **Clone the repository:**

   ```bash
   git clone https://github.com/DevSatyamCollab/number-guessing-game.git
   ```

2. Build the binary

   ```bash
   go build -o guess-game .
   ```

3. (Optional) Move to your path:
   ```bash
   sudo mv guess-game /usr/local/bin/
   ```
