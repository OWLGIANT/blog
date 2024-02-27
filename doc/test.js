const axios = require('axios')

async function auth(ip){
    const resp = await axios.get(`http://auth.thousandquant.com:8500/add?ip=${ip}`)
    return resp
}

const ipList = [
    '43.207.105.77',
    '57.180.58.204',
]

async function sleep(delay) {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve()
        }, delay)
    })
}

async function init(){
    for (let i = 0; i < ipList.length; i++) {
        const resp = await auth(ipList[i])
        console.log(`ip ${ipList[i]} ${resp.status}`)
        await sleep(200)
    }
}

init()