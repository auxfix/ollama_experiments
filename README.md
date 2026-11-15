# 🤖 Mii AI Assistant

A lightweight and elegant AI companion powered by **Ollama** and **Next.js**. 🚀

## ✨ Features

- 💬 **Real-time Chat**: Seamless conversation with an intuitive interface.
- 🧠 **Gemma 4 Powered**: Utilizing the powerful `gemma4:12b` model.
- ⚡ **Fast & Responsive**: Built with Go for high-performance backend processing.
- 🎨 **Modern UI**: A clean, user-friendly frontend built with Next.js and Tailwind CSS.

## 🛠️ Tech Stack

- **Backend**: [Go](https://go.dev/) (Gin Framework) 🐹
- **Frontend**: [Next.js](https://nextjs.org/) + Tailwind CSS 🎨
- **AI Engine**: [Ollama](https://ollama.com/) 🧠

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Node.js 18+
- [Ollama](https://ollama.com/) installed and running.

### Installation
1. **Clone the repo**:
   ```bash
   git clone https://github.com/your-repo/ollama_experiments.git
   cd ollama_experiments
   ```

2. **Setup Backend**:
   ```bash
   cd server
   go build -o server_bin main.go
   ./server_bin
   ```

3. **Setup Frontend**:
   ```bash
   cd ../client
   npm install
   npm run dev
   ```

## 📖 Usage
Once both the server and frontend are running, open your browser to `http://localhost:3000` and start chatting with Mii! 💬✨

---
*Made with ❤️ by the Mii Team*
