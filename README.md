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
![Screenshot from 2023-01-09 04-21-20](https://user-images.githubusercontent.com/63949797/211236434-7617a53c-93b4-4e54-b85d-a78721542c80.png)
![Screenshot from 2023-01-09 04-21-28](https://user-images.githubusercontent.com/63949797/211236437-26130ea6-d21b-443b-816e-18b547ee9eea.png)
![Screenshot from 2023-01-09 04-21-33](https://user-images.githubusercontent.com/63949797/211236440-e40ca21c-8784-45cf-8c00-819894e26d1b.png)



- Play around with the store methods in `main.go` to simulate how an auto store 
manager will use this project to manage sales and purchases.











