const request = require('request')

const Client = (baseUrl) => {
    const responceHandler = (res) => {
        if (res.ok) {
            return res.json()
        }

        throw new Error(`unexpected responce from server: ${res.status}, ${res.statusText}`)
    }

    return {       
        get: (path) => {
            return new Promise((resolve, reject) => {
                request(`${baseUrl}${path}`, {json: true}, (err, res, body) => {
                    if (err) {
                        reject(err)
                        return;
                    }
                    resolve(body)
                })
            })
        },

        post: async (path, data) => {
            return new Promise((resolve, reject) => {
                request(`${baseUrl}${path}`, {json: true, method: 'POST', body: data}, (err, res, body) => {
                    if (err) {
                        reject(err)
                        return
                    }
                    resolve(body)
                })
            })
        }
    }
}

module.exports = { Client }