# 🛍️ E-Commerce Full Stack Application

A modern, full-stack e-commerce platform built with **Go** (backend) and **React Native/Expo** (mobile frontend). This project features a RESTful API with JWT authentication, integrated payment processing with Stripe, and a beautiful mobile-first shopping experience.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
- [Mobile App](#mobile-app)
- [Environment Variables](#environment-variables)
- [Development](#development)
- [Architecture](#architecture)

---

## 🎯 Overview

This is a production-ready e-commerce application with two main components:

1. **Backend API** (`go-ecommerce-api/`) - A RESTful API built with Go, Gin, PostgreSQL, and Redis
2. **Mobile App** (`native-ecommerce/`) - A cross-platform mobile app built with React Native, Expo, and TypeScript

The backend and frontend communicate via a type-safe API client automatically generated from OpenAPI specifications.

---

## ✨ Features

### Backend Features
- 🔐 **JWT Authentication** - Secure access/refresh token authentication with HTTP-only cookies
- 👤 **User Management** - Registration, login, profile management with role-based access (User/Admin)
- 🛍️ **Product Management** - CRUD operations, search, filtering by category, featured products
- 🛒 **Shopping Cart** - Add/update/remove items, automatic price calculations with tax and shipping
- 📦 **Order Management** - Order creation, tracking, status updates, order history
- 💳 **Payment Processing** - Full Stripe integration with payment intents, webhooks, and payment history
- ❤️ **Wishlist** - Save and manage favorite products
- ⚡ **Redis Caching** - Improved performance with intelligent caching layer
- 📊 **Admin Panel** - User management, product CRUD, order management, bulk operations
- 📚 **Swagger/OpenAPI** - Interactive API documentation with OpenAPI 3.0

### Mobile App Features
- 📱 **Cross-Platform** - Runs on iOS, Android, and Web
- 🎨 **Modern UI** - Built with Tailwind CSS (NativeWind) and shadcn-inspired components
- 🔒 **Secure Authentication** - Token-based auth with persistent sessions
- 🛍️ **Product Browsing** - Browse, search, filter products with smooth animations
- 🛒 **Cart Management** - Real-time cart updates with optimistic UI
- 💳 **Stripe Checkout** - Integrated payment flow with Stripe SDK
- ❤️ **Wishlist** - Save favorites with instant feedback
- 📦 **Order Tracking** - View order history and status
- 🌙 **Dark Mode** - Full dark mode support
- 🔄 **Offline Support** - Works with React Query for data synchronization
- 📱 **Tab Navigation** - Intuitive bottom tab navigation (Home, Search, Cart, Profile)

---

## 🛠️ Tech Stack

### Backend
| Technology | Purpose |
|------------|---------|
| **Go 1.25.1** | Backend language |
| **Gin** | Web framework |
| **PostgreSQL** | Primary database |
| **GORM** | ORM for database operations |
| **Redis** | Caching layer |
| **JWT** | Authentication tokens |
| **bcrypt** | Password hashing |
| **Stripe** | Payment processing |
| **Swag** | OpenAPI/Swagger documentation |

### Frontend
| Technology | Purpose |
|------------|---------|
| **React Native 0.81** | Mobile framework |
| **Expo 54** | Development platform |
| **TypeScript 5.9** | Type safety |
| **Expo Router 6** | File-based routing |
| **TanStack Query 5** | Data fetching & caching |
| **NativeWind 4** | Tailwind CSS for React Native |
| **React Hook Form** | Form management |
| **Zod** | Schema validation |
| **Stripe React Native** | Payment integration |
| **Axios** | HTTP client |
| **AsyncStorage** | Local data persistence |

---

## 📁 Project Structure

```
go-native/
├── go-ecommerce-api/          # Backend API (Go)
│   ├── database/              # Database & Redis connections
│   ├── handlers/              # HTTP request handlers
│   ├── middleware/            # Auth, logging, CORS middleware
│   ├── models/                # Data models (User, Product, Cart, Order, etc.)
│   ├── repositories/          # Data access layer
│   ├── router/                # Route definitions
│   ├── services/              # Business logic layer
│   ├── utils/                 # Utilities (JWT, bcrypt, cart calculations)
│   ├── docs/                  # Auto-generated Swagger/OpenAPI docs
│   ├── scripts/               # Build & documentation scripts
│   ├── main.go                # Application entry point
│   ├── go.mod                 # Go dependencies
│   ├── Makefile               # Build commands
│   └── README.md              # Backend documentation
│
└── native-ecommerce/          # Mobile App (React Native + Expo)
    ├── app/                   # Expo Router file-based routing
    │   ├── (tabs)/           # Tab navigation screens
    │   │   ├── (home)/       # Home tab stack
    │   │   ├── profile/      # Profile tab stack (orders, wishlist)
    │   │   ├── cart.tsx      # Cart screen
    │   │   └── search.tsx    # Search screen
    │   ├── index.tsx         # Landing/Login screen
    │   ├── register.tsx      # Registration screen
    │   └── sign-in.tsx       # Sign in screen
    ├── components/            # Reusable UI components
    │   ├── auth/             # Authentication components
    │   ├── custom/           # Custom app components
    │   ├── products/         # Product-related components
    │   └── ui/               # Base UI components (shadcn-style)
    ├── api/                   # API hooks and mutations
    ├── client/                # Auto-generated API client
    ├── context/               # React Context (Auth)
    ├── hooks/                 # Custom React hooks
    ├── lib/                   # Utilities and helpers
    ├── assets/                # Images and static assets
    ├── package.json           # Node dependencies
    ├── app.json               # Expo configuration
    └── README.md              # Frontend documentation
```

---

## 🚀 Getting Started

### Prerequisites

- **Go 1.25.1+** - [Install Go](https://golang.org/doc/install)
- **Node.js 18+** - [Install Node.js](https://nodejs.org/)
- **PostgreSQL** - [Install PostgreSQL](https://www.postgresql.org/download/)
- **Redis** - [Install Redis](https://redis.io/download)
- **pnpm** (optional) - `npm install -g pnpm`
- **Expo CLI** - `npm install -g expo-cli`

### Backend Setup

1. **Navigate to backend directory**
   ```bash
   cd go-ecommerce-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Create `.env` file**
   ```bash
   cp .env.example .env
   ```
   
   Edit `.env` with your configuration:
   ```env
   DB_URL=postgres://username:password@localhost:5432/ecommerce_db?sslmode=disable
   REDIS_URL=localhost:6379
   JWT_SECRET=your-super-secret-jwt-key
   STRIPE_SECRET_KEY=sk_test_your_stripe_secret_key
   STRIPE_WEBHOOK_SECRET=whsec_your_webhook_secret
   ```

4. **Run database migrations**
   ```bash
   go run main.go
   ```
   GORM will automatically migrate your database schemas.

5. **Generate API documentation**
   ```bash
   make docs
   # or
   swag init
   ```

6. **Start the server**
   ```bash
   go run main.go
   # Server will start on http://localhost:8080
   ```

7. **Access Swagger UI**
   Open http://localhost:8080/swagger/index.html

### Frontend Setup

1. **Navigate to frontend directory**
   ```bash
   cd native-ecommerce
   ```

2. **Install dependencies**
   ```bash
   pnpm install
   # or
   npm install
   ```

3. **Generate API client** (after backend is running)
   ```bash
   pnpm generate:sdk
   ```
   This generates a type-safe API client from the OpenAPI specification.

4. **Start the development server**
   ```bash
   pnpm dev
   # or
   npm run dev
   ```

5. **Run on platform**
   - Press `i` for iOS simulator (Mac only)
   - Press `a` for Android emulator
   - Press `w` for web browser
   - Scan QR code with Expo Go app on physical device

---

## 📚 API Documentation

### Interactive Documentation
Once the backend is running, access the interactive Swagger UI at:
```
http://localhost:8080/swagger/index.html
```

### OpenAPI Specifications
- **JSON**: `http://localhost:8080/openapi.json`
- **YAML**: `http://localhost:8080/openapi.yaml`

### Main API Endpoints

#### 🔐 Authentication
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/auth/register` | User registration |
| POST | `/auth/login` | User login |
| GET | `/auth/refresh` | Refresh access token |

#### 🛍️ Products (Public)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/products` | Get all products |
| GET | `/products/featured` | Get featured products |
| GET | `/products/category/:category` | Get products by category |
| GET | `/products/search?query=` | Search products |
| GET | `/products/:id` | Get product by ID |

#### 🛒 Cart (Protected)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/cart` | Get user's cart |
| POST | `/cart/items` | Add item to cart |
| PUT | `/cart/items/:id` | Update cart item |
| DELETE | `/cart/items/:id` | Remove item |
| DELETE | `/cart` | Clear cart |

#### 📦 Orders (Protected)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/orders` | Get user's orders |
| POST | `/orders` | Create order |
| POST | `/orders/checkout` | Integrated checkout with payment |
| GET | `/orders/:id` | Get order by ID |

#### 💳 Payments (Protected)
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/payments/create-intent` | Create payment intent |
| POST | `/payments/confirm/:id` | Confirm payment |
| POST | `/payments/cancel/:id` | Cancel payment |
| GET | `/payments/status/:id` | Get payment status |
| GET | `/payments/history` | Get payment history |

#### ❤️ Wishlist (Protected)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/wishlist` | Get user's wishlist |
| POST | `/wishlist` | Add to wishlist |
| DELETE | `/wishlist/:product_id` | Remove from wishlist |
| GET | `/wishlist/:product_id` | Check if in wishlist |

#### 👑 Admin (Protected - Admin Only)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/admin/users` | Get all users |
| POST | `/admin/products` | Create product |
| POST | `/admin/products/bulk` | Bulk create products |
| PUT | `/admin/products/:id` | Update product |
| DELETE | `/admin/products/:id` | Delete product |
| GET | `/admin/orders` | Get all orders |
| PUT | `/admin/orders/:id/status` | Update order status |

---

## 📱 Mobile App

### Features by Screen

#### 🏠 Home Tab
- Featured products carousel
- Browse all products
- Add to cart/wishlist
- Product details

#### 🔍 Search Tab
- Search by keyword
- Filter by category
- View search results

#### 🛒 Cart Tab
- View cart items
- Update quantities
- Remove items
- See price summary (subtotal, tax, shipping)
- Proceed to checkout

#### 👤 Profile Tab
- User profile information
- Order history
- Wishlist management
- Dark mode toggle
- Logout

### Type-Safe API Integration

The mobile app uses an auto-generated API client:

```bash
# Generate after backend changes
pnpm generate:sdk
```

This creates:
- Type-safe API functions
- TanStack Query hooks
- Zod validation schemas
- Request/response types

Example usage:
```typescript
import { useQuery } from '@tanstack/react-query';
import { getProductsOptions } from '@/client/@tanstack/react-query.gen';

function ProductList() {
  const { data, isLoading } = useQuery(getProductsOptions());
  // data is fully typed!
}
```

---

## 🔐 Environment Variables

### Backend (.env)
```env
# Database
DB_URL=postgres://user:password@localhost:5432/ecommerce_db?sslmode=disable

# Redis
REDIS_URL=localhost:6379

# JWT
JWT_SECRET=your-jwt-secret-key-min-32-chars

# Stripe
STRIPE_SECRET_KEY=sk_test_your_stripe_secret_key
STRIPE_PUBLISHABLE_KEY=pk_test_your_stripe_publishable_key
STRIPE_WEBHOOK_SECRET=whsec_your_webhook_secret

# Server
PORT=8080
```

### Frontend (.env - optional)
The API URL is configured in the generated client. Update `client/index.ts` if needed:
```typescript
export const client = createClient({
  baseURL: 'http://localhost:8080'
});
```

---

## 🏗️ Architecture

### Backend Architecture

```
┌─────────────┐
│   Client    │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  Middleware │ ◄── Authentication, Logging, CORS
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  Handlers   │ ◄── HTTP Request Handling
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  Services   │ ◄── Business Logic
└──────┬──────┘
       │
       ▼
┌─────────────┐
│Repositories │ ◄── Data Access Layer
└──────┬──────┘
       │
       ▼
┌─────────────────┐
│ PostgreSQL/Redis│ ◄── Data Storage
└─────────────────┘
```

**Design Patterns:**
- **Clean Architecture** - Separation of concerns with clear boundaries
- **Repository Pattern** - Abstract data access for testability
- **Dependency Injection** - Loose coupling between components
- **Middleware Pattern** - Request processing pipeline

### Frontend Architecture

```
┌──────────────┐
│   UI Layer   │ ◄── Components (React Native)
└──────┬───────┘
       │
       ▼
┌──────────────┐
│  API Hooks   │ ◄── TanStack Query (Data fetching)
└──────┬───────┘
       │
       ▼
┌──────────────┐
│ API Client   │ ◄── Auto-generated from OpenAPI
└──────┬───────┘
       │
       ▼
┌──────────────┐
│ Backend API  │ ◄── Go Backend
└──────────────┘
```

**Key Patterns:**
- **File-based Routing** - Expo Router for navigation
- **Component Composition** - Reusable UI components
- **Custom Hooks** - Encapsulated business logic
- **Context API** - Global state (Auth)
- **Optimistic Updates** - Better UX with TanStack Query

---

## 🔧 Development

### Backend Commands

```bash
# Run server
go run main.go

# Run with hot reload (using air)
air

# Generate Swagger docs
make docs
swag init

# Build binary
make build
go build -o tmp/main .

# Run tests
go test ./...

# Format code
go fmt ./...

# Lint code
golangci-lint run
```

### Frontend Commands

```bash
# Start dev server
pnpm dev

# Generate API client
pnpm generate:sdk

# Run on iOS
pnpm ios

# Run on Android
pnpm android

# Run on Web
pnpm web

# Build for production
eas build --platform all

# Type check
tsc --noEmit

# Lint
eslint .
```

### SDK Generation Workflow

1. Make changes to backend API
2. Run backend server
3. Generate OpenAPI spec: `make docs`
4. Go to frontend: `cd ../native-ecommerce`
5. Generate client: `pnpm generate:sdk`
6. Types and hooks are now updated!

---

## 📝 Key Features Explained

### 1. **Integrated Checkout Flow**
The app supports a complete checkout with Stripe:
- Create payment intent
- Process payment via Stripe SDK
- Confirm payment on backend
- Create order
- Clear cart
- Show order confirmation

### 2. **Redis Caching Strategy**
- Product listings cached for 5 minutes
- User cart cached for 30 minutes
- Wishlist cached for 10 minutes
- Automatic cache invalidation on mutations

### 3. **JWT Authentication**
- Access token (1 hour) for API requests
- Refresh token (24 hours) for getting new access tokens
- Tokens stored in HTTP-only cookies
- Mobile app uses AsyncStorage for persistence

### 4. **Role-Based Access Control**
- User role: Browse, cart, orders, wishlist
- Admin role: All user features + product/order management

### 5. **Optimistic UI Updates**
The mobile app uses optimistic updates for better UX:
- Adding to cart shows immediately
- Wishlist updates are instant
- Real data syncs in background

---


## 🙏 Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/)
- [Expo](https://expo.dev/)
- [React Native Reusables](https://reactnativereusables.com/)
- [TanStack Query](https://tanstack.com/query)
- [Stripe](https://stripe.com/)
- [Hey API](https://heyapi.dev/)

---

**Built with ❤️ using Go and React Native**
