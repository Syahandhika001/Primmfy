# 🎨 PRIMMFY Frontend

Next.js 15 frontend for PRIMMFY learning platform.

## 🚀 Quick Start

See [main README](../README.md) for full setup instructions.

### Development

```bash
npm install
cp .env.example .env.local
npm run dev
```

Open http://localhost:3000

## 📁 Structure

```
src/
├── app/              # Next.js pages (App Router)
├── components/       # React components
└── lib/
    ├── api/         # API client
    ├── contexts/    # React contexts
    ├── hooks/       # Custom hooks
    ├── types/       # TypeScript types
    └── utils/       # Utility functions
```

## 🎨 Design System

- Primary: `rgb(194, 226, 250)` - Brand Blue
- Secondary: `rgb(183, 163, 227)` - Lavender
- Accent: `rgb(255, 241, 203)` - Cream

## 📝 Scripts

- `npm run dev` - Development server
- `npm run build` - Production build
- `npm run start` - Start production
- `npm run lint` - Run linter

For more details, see [main README](../README.md).