# auto-shop
A micro tool for managing materials in an auto-mechanics shop.

This project was written for [Alt-school's](https://www.altschoolafrica.com/) second semester exam.

## Description

auto-shop uses Object-Oriented Programming (OOP) paradigms in Go to create a mini shop
management system for an auto-mechanics distributor. 

## Features

The store manager will be able to perform the following: 

- Check the number of products in the store that are still up for sale
- Add an Item to the store
- List all product items in the store
- Sell an item
- Show a list of sold items and the total price

## Added Features

- Validate each added product on it's product type field
- Create a set containing all default valid types
- Include a function that allows a store manager add a new type
- Display all store products in JSON format
- Create two generic types - Vehicle and Accessory - which can accomodate all possible
products that could be added to autoshop

## Setup

To use this project, you must have `golang 1.18` or higher installed. If installed, go through
the following processes:

- Clone this project using the following command: 

```shell
git clone https://github.com/Emmrys-Jay/auto-shop.git
```

- Change into project directory using:

```shell
cd auto-shop
```

- Run project demo using:
```shell
go run main.go
```

- Demo output should be in form specified in the screenshot below.

![Screenshot from 2022-11-04 13-34-10](https://user-images.githubusercontent.com/63949797/200057365-2798a6c7-8ee6-47c0-9ca3-d6293e511b32.png)
![Screenshot from 2022-11-04 13-34-14](https://user-images.githubusercontent.com/63949797/200057372-d8991411-7e9a-43d5-8929-acd69df4f5e7.png)


- Play around with the store methods in `main.go` to simulate how an auto store 
manager will use this project to manage sales and purchases.











