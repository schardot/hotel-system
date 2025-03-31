# 📌 Hotel Booking System

A **work-in-progress** hotel booking management system built with **Go (Golang)**, **PostgreSQL**, and **Bootstrap**.

This project is designed to manage **customers and bookings**, with plans to include **user authentication, activity logging, and role-based access control**.

---

## 🚀 Features (Current & Planned)
### ✅ Implemented
- 📌 **Customer Management**  
  - Add, edit, and list customers  
  - Search customers by **ID, name, last name, or email**  

- 📌 **Bookings (In Progress)**
  - Create bookings linked to customers  
  - Store **check-in, check-out, and room type**  

### 🔧 Planned Features
- 🛡️ **User Authentication** (Login system with hashed passwords)  
- 👥 **Role-Based Access Control** (Admin, Staff, Guests)  
- 📝 **Activity Logging** (Track user actions)  
- 📊 **Dashboard & Reports**  

---

## 🛠️ Tech Stack
- **Backend**: Go (Golang)  
- **Database**: PostgreSQL  
- **Frontend**: HTML, Bootstrap, jQuery  
- **ORM**: Standard `database/sql`  
- **Migrations**: [Goose](https://github.com/pressly/goose)  

---


## 📂 Project Structure
```
/hotel-booking-system
│── /internal
│   ├── /config          # Configuration files
│   ├── /customers       # Customer management logic
│   ├── /bookings        # Booking management logic (WIP)
│   ├── /db              # Database connection
│   │  ├── /migrations   # Database migrations
│   ├── /entities        # Structs for Customers & Bookings
│── /templates           # HTML templates
│── /migrations          # Database migrations
│── main.go              # Entry point
│── README.md            # You’re here!
```

---

## 📅 Roadmap
🔲 Implement full **booking management**  
🔲 Add **user authentication**  
🔲 Introduce **role-based permissions**  
🔲 Build an **admin dashboard**  
🔲 Deploy to a **cloud server**  

---

## 🤝 Contributing
This project is a work in progress, and contributions are welcome!  

To contribute:
1. **Fork** the repo  
2. **Create a feature branch** (`git checkout -b feature-name`)  
3. **Commit your changes** (`git commit -m "Added new feature"`)  
4. **Push to your fork** and create a **pull request**  

---

## 📜 License
[MIT License](LICENSE) - Feel free to use and modify.

---

### 🔥 Let's Build This Together! 🚀  
Have any ideas or suggestions? Open an **issue** or drop a **comment**! 😊

