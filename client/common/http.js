"use strict";
exports.__esModule = true;
var node_fetch_1 = require("node-fetch");
var URL = 'http://localhost:8080';
exports["default"] = {
    sendData: function (name, battery, currentVideo, deviceTime) { return (node_fetch_1["default"](URL + "/sendData", {
        method: 'POST',
        body: JSON.stringify({ name: name, battery: battery, currentVideo: currentVideo, deviceTime: deviceTime }),
        headers: { 'Content-Type': 'application/json' }
    })); },
    getData: function (name) { return (node_fetch_1["default"](URL + "/getData", {
        method: 'POST',
        body: JSON.stringify({ name: name }),
        headers: { 'Content-Type': 'application/json' }
    })); }
};
