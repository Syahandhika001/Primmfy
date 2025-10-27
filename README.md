# 🎓 PRIMMFY

Interactive programming learning platform using **PRIMM** (Predict, Run, Investigate, Modify, Make) methodology.

## 📦 Project Structure

This is a **monorepo** containing both frontend and backend:

```
PRIMMFY/
├── Frontend/       # Next.js 15 + TypeScript
└── Backend/        # Go + Gin framework
```

---

## 🚀 Quick Start

### Prerequisites
- Node.js 18+
- Go 1.21+
- PostgreSQL 14+ (or your database)

### 1. Clone Repository
```bash
git clone https://github.com/YOUR_USERNAME/PRIMMFY.git
cd PRIMMFY
```

### 2. Setup Backend
```bash
cd Backend
cp .env.example .env
# Edit .env with your configuration
go mod download
go run main.go
```

Backend will run on: `http://localhost:8080`

### 3. Setup Frontend
```bash
cd Frontend
npm install
cp .env.example .env.local
# Edit .env.local:
# NEXT_PUBLIC_API_URL=http://localhost:8080/api
npm run dev
```

Frontend will run on: `http://localhost:3000`

---

## 🏗️ Tech Stack

### Frontend
- **Framework:** Next.js 15 (App Router)
- **Language:** TypeScript
- **Styling:** TailwindCSS v4
- **State:** React Context API
- **Forms:** React Hook Form + Zod
- **HTTP:** Axios

### Backend
- **Language:** Go 1.21
- **Framework:** Gin
- **Database:** PostgreSQL (with GORM)
- **Auth:** JWT
- **Validation:** go-playground/validator

---

## 📚 Documentation

- [Frontend Documentation](./Frontend/README.md)
- [Backend Documentation](./Backend/README.md)
- [API Documentation](./Backend/docs/API.md) _(if exists)_

---

## 🔐 Environment Variables

### Frontend `.env.local`
```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```

### Backend `.env`
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=primmfy
JWT_SECRET=your_secret_key
```

---

## 🎯 Features

### Phase 1: Authentication ✅
- [x] User registration (Teacher/Student)
- [x] Login with JWT
- [x] Role-based authentication
- [x] Form validation

### Phase 2: Dashboard (In Progress)
- [ ] Protected routes
- [ ] Teacher dashboard
- [ ] Student dashboard
- [ ] Profile management

### Phase 3: Content Management
- [ ] Create lessons (Teacher)
- [ ] View lessons (Student)
- [ ] PRIMM steps implementation

### Phase 4: Interactive Learning
- [ ] Code editor
- [ ] Code execution
- [ ] Real-time feedback
- [ ] Progress tracking

---

## 🤝 Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m '✨ Add AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Open Pull Request

---

## 📝 Commit Convention

- `✨ Feature:` New feature
- `🐛 Fix:` Bug fix
- `♻️ Refactor:` Code refactoring
- `📝 Docs:` Documentation
- `🎨 Style:` Formatting, styling
- `🧪 Test:` Add tests
- `🚀 Deploy:` Deployment changes

---

## 👨‍💻 Team

- **Developer:** [Your Name]
- **University:** [Your University]
- **Project Type:** Final Year Project / Thesis

---

## 📄 License

MIT License - See [LICENSE](LICENSE) file for details

---

**Made with ❤️ for better programming education**