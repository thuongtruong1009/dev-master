const express = require('express')
const bodyParser = require('body-parser')
const app = express()

const Producer = require('./producer')
const producer = new Producer()

app.use(bodyParser.json("application/json"))

app.post("/sendLog", async(req, res, next) => {
    await producer.publishMessage(req.body.logType, req.body.message)
    res.send("Message sent to RabbitMQ")
})

app.listen(3000, () => {
    console.log("Server is listening on port 3000")
})