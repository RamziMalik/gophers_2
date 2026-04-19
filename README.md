# Gophers 2

A simple CLI-based inventory management application for a small grocery store (**toko kelontong**) built with **Go**.

## Description

This project is the **second task of GDGoC UNJ Mobile Go**.  
The application runs in the terminal and allows users to:

- add items to inventory
- view all available stock
- purchase items
- calculate total price and change automatically

## Features

- Add new items to warehouse inventory
- View current stock list
- Purchase items by ID
- Automatic total price calculation
- Automatic change calculation
- Simple terminal-based interface

## Project Structure

```bash
gophers_2/
├── calculator/
│   └── transaction.go
├── data/
│   └── inventory.go
├── handler/
│   ├── add.go
│   ├── show.go
│   └── update.go
├── input/
│   ├── add.go
│   ├── buy.go
│   └── menu.go
├── model/
│   └── data.go
├── go.mod
└── main.go
