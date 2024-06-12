const express = require("express");
const app = express();
const axios = require("axios");
const cors = require("cors");
const port = 3001;
require("dotenv").config();
app.use(cors());
app.use(express.json());

const HOST = process.env.HOST || "localhost"; // 192.168.1.2:4444/api

app.get("/api/room", async (req, res) => {
  try {
    const {store} = req.query;

    const token = req.headers.authorization.split(" ")[1];

    let url = `${HOST}/room`;

    if (store) {
      url = `${url}?store=${store}`;
    }

    const room = await axios.get(url, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    res.send(room.data);
  } catch (err) {
    response(res);
  }
});

app.get("/api/stores", async(req, res) => {
  try {
    const token = req.headers.authorization.split(" ")[1];

    let url = `${HOST}/stores`;

    const store = await axios.get(url, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })

    res.send(store.data);
  } catch (err) {
    response(res);
  }
})

app.post("/api/auth/login", async (req, res) => {
  const { username, password } = req.body;
  try {
    const loginResponse = await axios.post(`${HOST}/auth/login`, {
      username,
      password,
    });

    res.send(loginResponse.data);
  } catch (error) {
    console.error(error);
    res.status(401).send("invalid token format");
  }
});

app.get('/api/receipts', async (req, res) => {
  try {
    const token = req.headers.authorization.split(" ")[1];

    let url = `${HOST}/receipts`;

    const {room, store} = req.query

    if (room && store) {
      url = `${url}?room=${room}&store=${store}`;
    }

    const receipts = await axios.get(url, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })

    res.send(receipts.data);
  } catch (err) {
    response(res);
  } 
})

const response = (res) => 
  res.status(401).json({
    "data": null,
    "errorCode": 0,
    "errors": [
        "signature is invalid"
    ],
    "message": ""
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
