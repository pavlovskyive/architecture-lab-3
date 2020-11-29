const http = require('../tools/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listMenuItems: () => client.get('/menu'),
        createMenuItem: (name, price) => client.post('/menu', { name, price }),
        listOrders: () => client.get('/orders'),
        createOrder: (menu_item_id, table_number) => client.post('/orders', { menu_item_id, table_number })
    }

};

module.exports = { Client };