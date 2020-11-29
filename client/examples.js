// This file contains examples of implementation scenarios using
// the SDK for menu management

const restaurant = require('./restaurant/client');

const client = restaurant.Client('http://localhost:8080');

const printMenuItems = (menu) => {
    console.log('Menu:')
    menu.forEach((menuItem) => console.log(`${menuItem.name}: ${menuItem.price} UAH`))
}

const printOrders = (orders, menu) => {
    console.log('Orders:')
    orders.forEach((order) => {
        let menuItem = menu.filter((item) => item.id == order.menu_item_id)

        console.log(`${menuItem[0].name} for table number ${order.table_number}`)
    })
}

const printExample = (number) => {
    console.log()
    console.log('==========')
    console.log(`Example #${number}`)
    console.log('==========')
}

// 1st example: Showing all menu items
client.listMenuItems()
    .then((menu) => {
        printExample(1)
        printMenuItems(menu)    
    })
    .catch((err) => {
        console.log(`Problem listing menu items: ${err.message}`);
    });

// 2nd example: Creation of new item
client.createMenuItem('Milkshake', 34)
    .then((res) => {
        client.listMenuItems()
            .then((menu) => {
                printExample(2)
        
                console.log('Server response for creating a new menu item:', res)

                console.log("Menu after inserting a new item: ")
                printMenuItems(menu)
            })
            .catch((err) => {
                console.log(`Problem listing menu items: ${err.message}`);
            });
    })
    .catch((err) => {
        console.log(`Problem creating a new menu item: ${err.message}`);
    })

// 3rd example: Showing all orders
client.listOrders()
    .then((orders) => {
        client.listMenuItems()
            .then((menu) => {
                printExample(3)

                printOrders(orders, menu)
            })
            .catch((err) => {
                console.log(`Problem listing menu items: ${err.message}`);
            });
    })
    .catch((err) => {
        console.log(`Problem listing orders: ${err.message}`);
    });

// 4th example: Creation of new order
client.createOrder(2, 4)
    .then((res) => {
        client.listOrders()
            .then((orders) =>
                client.listMenuItems()
                    .then((menu) => {
                        printExample(4)
        
                        console.log('Server response for creating a new order:', res)

                        console.log("Orders after inserting a new one: ")
                        printOrders(orders, menu)
                    })
                    .catch((err) => {
                        console.log(`Problem listing menu items: ${err.message}`);
                    })
            )
            .catch((err) => {
                console.log(`Problem listing orders: ${err.message}`);
            });
    })
    .catch((err) => {
        console.log(`Problem creating a new order: ${err.message}`);
    })