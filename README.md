# FinanceApp MVP

A modern personal finance app focused on **expense tracking, budget management, goal setting**, family collaboration, and notifications.

---

## 🚀 Features (MVP)

- Track daily expenses and income with categories and accounts  
- Manage multiple **accounts** (bank, wallet, credit card) with balances  
- Set and monitor budgets with monthly trends  
- Create and track savings goals  
- Receive spend alerts and reminders  
- Collaborate with family members on budgets and expenses  
- Notifications for bills and spend limits

---

## 📱 Tech Stack

- **Frontend:** React Native (cross-platform iOS & Android)  
- **Backend:** Go fiber.go 
- **Database:** PostgreSQL (relational data with UUIDs)  
- **Authentication:** JWT  
- **Push Notifications:** Expo Push  
- **ORM:** GORM (Go) 

---

## 🗂️ Database Schema

The app includes core entities such as:  
- Users & Families  
- Accounts (bank accounts, wallets, credit cards)  
- Transactions (linked to accounts and categories)  
- Categories (global + user-defined)  
- Budgets  
- Goals  
- Notifications  

---

## 📦 Installation & Setup

1. **Clone the repo:**

   ```bash
   git clone https://github.com/yourusername/financeapp.git
   cd financeapp
